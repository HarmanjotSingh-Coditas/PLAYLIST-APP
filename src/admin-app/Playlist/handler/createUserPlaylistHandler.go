package handler

import (
	"admin-app/Playlist/business"
	"admin-app/Playlist/commons/constants"
	"admin-app/Playlist/models"
	"encoding/json"
	"net/http"
	genericConstants "playlist-app/src/constants"
	genericModels "playlist-app/src/models"

	"playlist-app/src/utils/validations"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type CreateUserPlaylistController struct {
	service business.CreateUserPlaylistService
}

func NewCreateUserPlaylistController(service business.CreateUserPlaylistService) *CreateUserPlaylistController {
	return &CreateUserPlaylistController{
		service: service,
	}
}

// HandleCreateUserPlaylist creates a new playlist
// @Summary Create new playlist
// @Description Creates a new playlist with specified songs
// @Tags Playlists
// @Accept json
// @Produce json
// @Param request body models.BFFCreateUserPlaylistRequest true "Create playlist request"
// @Success 200 {object} models.BFFCreateUserPlaylistResponse "Playlist created successfully"
// @Failure 400 {object} models.ErrorAPIResponse "Invalid input: Validation failed"
// @Failure 404 {object} models.ErrorAPIResponse "Songs not found"
// @Failure 409 {object} models.ErrorAPIResponse "Playlist already exists"
// @Failure 500 {object} models.ErrorAPIResponse "Internal server error"
// @Router /v1/api/playlists/create [post]
func (controller *CreateUserPlaylistController) HandleCreateUserPlaylist(ctx *gin.Context) {
	var bffCreateUserPlaylist models.BFFCreateUserPlaylistRequest

	if err := ctx.ShouldBindJSON(&bffCreateUserPlaylist); err != nil {
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

	if err := validations.GetBFFValidator(ctx).Struct(bffCreateUserPlaylist); err != nil {
		validationErrors, _ := validations.FormatValidationErrors(err.(validator.ValidationErrors))
		ctx.JSON(http.StatusBadRequest, validationErrors)
		return
	}

	_, err := controller.service.CreateUserPlaylistService(ctx, bffCreateUserPlaylist)
	if err != nil {
		if strings.Contains(err.Error(), genericConstants.DuplicateKeyError) {
			ctx.JSON(http.StatusConflict, genericModels.ErrorAPIResponse{
				ErrorMessage: constants.PlaylistAlreadyExistsError,
			})
			return
		}
		if strings.Contains(err.Error(), genericConstants.ForeignKeyError) {
			ctx.JSON(http.StatusNotFound, genericModels.ErrorAPIResponse{
				ErrorMessage: constants.SongIdsDoesNotExistsError,
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, genericModels.ErrorAPIResponse{
			ErrorMessage: constants.UnexpectedError,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.BFFCreateUserPlaylistResponse{
		Message: constants.PlaylistCreationSuccess,
	})
}
