package handler

import (
	"admin-app/Playlist/business"
	"admin-app/Playlist/commons/constants"
	"admin-app/Playlist/models"
	"encoding/json"
	"errors"
	"net/http"
	genericModels "playlist-app/src/models"
	"playlist-app/src/utils/validations"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, genericModels.ErrorAPIResponse{ErrorMessage: constants.SongOrPlaylistNotFoundError})
			return
		} else if strings.Contains(strings.ToLower(err.Error()), constants.EmptySongIdsError) {
			ctx.JSON(http.StatusBadRequest, genericModels.ErrorAPIResponse{ErrorMessage: constants.NoSongIDsProvidedError})
		} else if strings.Contains(strings.ToLower(err.Error()), constants.InvalidAction) {
			ctx.JSON(http.StatusBadRequest, genericModels.ErrorAPIResponse{ErrorMessage: constants.InvalidActionsError})
		} else if strings.Contains(strings.ToLower(err.Error()), constants.DuplicateKeyError) {
			ctx.JSON(http.StatusConflict, genericModels.ErrorAPIResponse{ErrorMessage: constants.SongAlreadyExistsInPlaylistError})
		} else if strings.Contains(strings.ToLower(err.Error()), constants.ForeignKeyError) {
			ctx.JSON(http.StatusBadRequest, genericModels.ErrorAPIResponse{ErrorMessage: constants.InvalidPlaylistOrSongId})
		} else {
			ctx.JSON(http.StatusInternalServerError, genericModels.ErrorAPIResponse{ErrorMessage: constants.UnexpectedError})
		}
		return
	}

	msg := constants.SongsAddedToPlaylistSuccess
	if bffAdSongsRequest.Action == constants.Delete {
		msg = constants.SongsDeletedFromPlaylistSuccess
	}

	ctx.JSON(http.StatusOK, models.BFFAdSongsFromPlaylistResponse{
		Message:  msg,
		Playlist: *playlist,
	})
}
