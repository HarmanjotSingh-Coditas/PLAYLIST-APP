package handler

import (
	"admin-app/Playlist/business"
	"admin-app/Playlist/commons/constants"
	"admin-app/Playlist/models"
	"encoding/json"
	"net/http"
	genericModels "playlist-app/src/models"
	"playlist-app/src/utils/validations"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AdSongsFromPlaylistController struct {
	service business.AdSongsFromPlaylistService
}

func NewADSongsFromPlaylistController(service business.AdSongsFromPlaylistService) *AdSongsFromPlaylistController {
	return &AdSongsFromPlaylistController{
		service: service,
	}
}

// HandleAdSongsFromPlaylist adds or deletes songs from playlist
// @Summary Modify playlist songs
// @Description Add new songs to or delete existing songs from a playlist
// @Tags Playlists
// @Accept json
// @Produce json
// @Param request body models.BFFAdSongsFromPlaylistRequest true "Add/Delete songs request"
// @Success 200 {object} models.BFFAdSongsFromPlaylistResponse "Operation successful"
// @Failure 400 {object} models.ErrorAPIResponse "Invalid input or action"
// @Failure 404 {object} models.ErrorAPIResponse "Playlist or songs not found"
// @Failure 409 {object} models.ErrorAPIResponse "Songs already exist in playlist"
// @Failure 500 {object} models.ErrorAPIResponse "Internal server error"
// @Router /v1/api/playlists/ad [put]s
func (controller *AdSongsFromPlaylistController) HandleAdSongsFromPlaylist(ctx *gin.Context) {
	var bffAdSongsRequest models.BFFAdSongsFromPlaylistRequest

	if err := ctx.ShouldBindJSON(&bffAdSongsRequest); err != nil {
		errorMsgs := genericModels.ErrorMessage{
			Key:          err.(*json.UnmarshalTypeError).Field,
			ErrorMessage: constants.JsonBindingFieldError,
		}
		errorResponse := genericModels.ErrorAPIResponse{
			Message: []genericModels.ErrorMessage{errorMsgs},
		}
		ctx.JSON(http.StatusBadRequest, errorResponse)
		return
	}

	if err := validations.GetBFFValidator(ctx).Struct(&bffAdSongsRequest); err != nil {
		validationErrors, _ := validations.FormatValidationErrors(err.(validator.ValidationErrors))
		ctx.JSON(http.StatusBadRequest, validationErrors)
		return
	}

	playlist, err := controller.service.AdSongsPlaylistService(ctx, bffAdSongsRequest)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), constants.PlaylistDoesNotExistsError):
			ctx.JSON(http.StatusNotFound, genericModels.ErrorAPIResponse{
				ErrorMessage: constants.PlaylistNotFoundError,
			})
		case strings.Contains(err.Error(), constants.NoValidSongsToAddError):
			ctx.JSON(http.StatusNotFound, genericModels.ErrorAPIResponse{
				ErrorMessage: err.Error(),
			})
		case strings.Contains(err.Error(), constants.NoValidSongsToBeDeletedError):
			ctx.JSON(http.StatusBadRequest, genericModels.ErrorAPIResponse{
				ErrorMessage: err.Error(),
			})
		case strings.Contains(err.Error(), constants.InvalidAction):
			ctx.JSON(http.StatusBadRequest, genericModels.ErrorAPIResponse{
				ErrorMessage: constants.InvalidActionsError,
			})
		case strings.Contains(err.Error(), constants.SongsWithIds):
			ctx.JSON(http.StatusConflict, genericModels.ErrorAPIResponse{
				ErrorMessage: err.Error(),
			})
		default:
			ctx.JSON(http.StatusInternalServerError, genericModels.ErrorAPIResponse{
				ErrorMessage: constants.UnexpectedError,
			})
		}
		return
	}

	response := models.BFFAdSongsFromPlaylistResponse{
		Playlist: *playlist,
	}

	switch bffAdSongsRequest.Action {
	case "ADD":
		response.Message = constants.SongsAddedToPlaylistSuccess
	case "DELETE":
		response.Message = constants.SongsDeletedFromPlaylistSuccess
	}

	ctx.JSON(http.StatusOK, response)
}
