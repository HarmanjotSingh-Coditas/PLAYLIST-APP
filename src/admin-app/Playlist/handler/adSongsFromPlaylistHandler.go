package handler

import (
	"admin-app/Playlist/business"
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

func (controller *AdSongsFromPlaylistController) HandleAdSongsFromPlaylist(ctx *gin.Context) {
	var bffAdSongsRequest models.BFFAdSongsFromPlaylistRequest

	if err := ctx.ShouldBindJSON(&bffAdSongsRequest); err != nil {
		errorMsgs := genericModels.ErrorMessage{
			Key:          err.(*json.UnmarshalTypeError).Field,
			ErrorMessage: "JsonBindingFieldError",
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
		case strings.Contains(err.Error(), "playlist does not exist"):
			ctx.JSON(http.StatusNotFound, genericModels.ErrorAPIResponse{
				ErrorMessage: "Playlist not found",
			})
		case strings.Contains(err.Error(), "no valid songs to add"):
			ctx.JSON(http.StatusBadRequest, genericModels.ErrorAPIResponse{
				ErrorMessage: err.Error(),
			})
		case strings.Contains(err.Error(), "no valid songs to delete"):
			ctx.JSON(http.StatusBadRequest, genericModels.ErrorAPIResponse{
				ErrorMessage: err.Error(),
			})
		case strings.Contains(err.Error(), "invalid action"):
			ctx.JSON(http.StatusBadRequest, genericModels.ErrorAPIResponse{
				ErrorMessage: "Invalid action. Must be either 'ADD' or 'DELETE'",
			})
		case strings.Contains(err.Error(), "songs with IDs"):
			ctx.JSON(http.StatusConflict, genericModels.ErrorAPIResponse{
				ErrorMessage: err.Error(),
			})
		default:
			ctx.JSON(http.StatusInternalServerError, genericModels.ErrorAPIResponse{
				ErrorMessage: "An unexpected error occurred",
			})
		}
		return
	}

	// Handle success response
	response := models.BFFAdSongsFromPlaylistResponse{
		Playlist: playlist,
	}

	switch bffAdSongsRequest.Action {
	case "ADD":
		response.Message = "Songs added to playlist successfully"
	case "DELETE":
		response.Message = "Songs deleted from playlist successfully"
	}

	ctx.JSON(http.StatusOK, response)
}
