package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manot40/gomenu/internal/models"
	"github.com/manot40/gomenu/internal/utils"
)

type createTags struct {
	Name string `json:"name" binding:"required"`
}

func GetAllTag(ctx *gin.Context) {
	var tags []models.Tag
	query := utils.Pagination(ctx, &tags)

	if err := query.Find(&tags).Error; err != nil {
		utils.SendJson(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if len(tags) == 0 {
		utils.SendJson(ctx, http.StatusNotFound, nil, "No tags found, please check again later")
		return
	}

	utils.SendJson(ctx, http.StatusOK, &tags, "List of tags(s) found")
}

func CreateTag(ctx *gin.Context) {
	var input createTags

	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.SendJson(ctx, http.StatusBadRequest, err.Error())
		return
	}

	tag := models.Tag{Name: input.Name}

	if err := tag.Create(); err != nil {
		utils.SendJson(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJson(ctx, http.StatusCreated, &tag, "Tags created")
}

func DeleteTag(ctx *gin.Context) {
	var tag models.Tag

	if err := models.DB.First(&tag, ctx.Param("id")).Error; err != nil {
		utils.SendJson(ctx, http.StatusNotFound, err.Error(), "Tags not found")
		return
	}

	if err := tag.Delete(); err != nil {
		utils.SendJson(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJson(ctx, http.StatusOK, nil, "Tags deleted")
}
