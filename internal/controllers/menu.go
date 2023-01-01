package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manot40/gomenu/internal/models"
	"github.com/manot40/gomenu/internal/utils"
	"github.com/manot40/gomenu/internal/utils/rekuest"
)

type menu struct {
	Description *string  `json:"description"`
	Tags_       []string `json:"tags"`
	Like        *uint    `json:"like"`
}

type createMenu struct {
	menu
	Name  string `json:"name" binding:"required"`
	Price uint   `json:"price" binding:"required"`
}

type updateMenu struct {
	menu
	Name  string `json:"name"`
	Price *uint  `json:"price"`
	Tags  string `json:"-" binding:"-"` // ignore this field
}

func GetMenu(ctx *gin.Context) {
	var menu models.Menu

	if err := models.DB.First(&menu, ctx.Param("id")).Error; err != nil {
		utils.SendJson(ctx, http.StatusNotFound, err.Error(), "Menu not found")
		return
	}

	utils.SendJson(ctx, http.StatusOK, &menu, "Menu found")
}

func GetAllMenu(ctx *gin.Context) {
	var menus []models.Menu
	query := utils.Pagination(ctx, &menus)

	if err := query.Find(&menus).Error; err != nil {
		utils.SendJson(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if len(menus) == 0 {
		utils.SendJson(ctx, http.StatusNotFound, nil, "No menu found, please check again later")
		return
	}

	utils.SendJson(ctx, http.StatusOK, &menus, "List of menu(s) found")
}

func CreateMenu(ctx *gin.Context) {
	var input createMenu

	if err := ctx.ShouldBindJSON(&input); err != nil {
		utils.SendJson(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var tags string
	if err := rekuest.TransformTags(&tags, input.Tags_); err != nil {
		utils.SendJson(ctx, http.StatusBadRequest, err.Error(), "Tag not found")
		return
	}

	menu := models.Menu{
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
		Tags:        tags,
	}

	if err := menu.Create(); err != nil {
		utils.SendJson(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJson(ctx, http.StatusCreated, &menu, "Menu created")
}

func UpdateMenu(ctx *gin.Context) {
	var reqBody updateMenu
	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		utils.SendJson(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var menu models.Menu
	if err := models.DB.First(&menu, ctx.Param("id")).Error; err != nil {
		utils.SendJson(ctx, http.StatusNotFound, err.Error(), "Menu not found")
		return
	}

	if err := rekuest.TransformTags(&reqBody.Tags, reqBody.Tags_); err != nil {
		utils.SendJson(ctx, http.StatusBadRequest, err.Error(), "Tag not found")
		return
	}

	data := utils.StructToMap(menu, reqBody)
	data["Tags"] = reqBody.Tags
	if err := models.DB.Model(&menu).Updates(&data).Error; err != nil {
		utils.SendJson(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJson(ctx, http.StatusCreated, &menu, "Menu updated")
}

func DeleteMenu(ctx *gin.Context) {
	var menu models.Menu

	if err := models.DB.Where("id = ?", ctx.Param("id")).First(&menu).Error; err != nil {
		utils.SendJson(ctx, http.StatusNotFound, err.Error(), "Menu not found")
		return
	}

	if err := menu.Delete(); err != nil {
		utils.SendJson(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendJson(ctx, http.StatusOK, nil, "Menu deleted")
}
