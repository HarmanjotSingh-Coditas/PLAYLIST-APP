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

type CreateUserPlaylistController struct {
	service business.CreateUserPlaylistService
}

func NewCreateUserPlaylistController(service business.CreateUserPlaylistService) *CreateUserPlaylistController {
	return &CreateUserPlaylistController{
		service: service,
	}
}

func (controller *CreateUserPlaylistController) HandleCreateUserPlaylist(ctx *gin.Context) {
	var bffCreateUserPlaylist models.BFFCreateUserPlaylistRequest

	if err := ctx.ShouldBindJSON(&bffCreateUserPlaylist); err != nil {
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

	if err := validations.GetBFFValidator(ctx).Struct(bffCreateUserPlaylist); err != nil {
		validationErrors, _ := validations.FormatValidationErrors(err.(validator.ValidationErrors))
		ctx.JSON(http.StatusBadRequest, validationErrors)
		return
	}

	created, err := controller.service.CreateUserPlaylistService(ctx, bffCreateUserPlaylist)
	if err != nil {
		if strings.Contains(err.Error(), "one or more song IDs do not exist") {
			ctx.JSON(http.StatusNotFound, genericModels.ErrorAPIResponse{
				ErrorMessage: "One or more song IDs do not exist",
			})
		} else if strings.Contains(err.Error(), "playlist already exists") {
			ctx.JSON(http.StatusConflict, genericModels.ErrorAPIResponse{
				ErrorMessage: "Playlist already exists",
			})
		} else if strings.Contains(err.Error(), "playlist could not be created") {
			ctx.JSON(http.StatusInternalServerError, genericModels.ErrorAPIResponse{
				ErrorMessage: "Failed to create playlist",
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, genericModels.ErrorAPIResponse{
				ErrorMessage: "An unexpected error occurred",
			})
		}
		return
	}

	if created {
		ctx.JSON(http.StatusOK, models.BFFCreateUserPlaylistResponse{
			Message: "Playlist created successfully",
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, genericModels.ErrorAPIResponse{
			ErrorMessage: "Failed to create playlist",
		})
	}
}
