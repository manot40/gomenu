package utils

import (
	"math"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/manot40/gomenu/internal/models"
	"gorm.io/gorm"
)

type pagination struct {
	CurrentPage  int     `json:"currentPage"`
	TotalPage    float64 `json:"totalPage"`
	TotalRecords int     `json:"totalRecords"`
}

func Pagination(ctx *gin.Context, model interface{}) *gorm.DB {
	query := models.DB.Model(&model)

	// Pagination
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))

	// Optional Queries
	sort := ctx.DefaultQuery("sortby", "id")
	order := ctx.DefaultQuery("order", "DESC")
	search := ctx.DefaultQuery("search", "")

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	var count int64
	query.Count(&count)

	// Set Meta
	meta := pagination{
		CurrentPage:  page,
		TotalPage:    math.Ceil(float64(count) / float64(limit)),
		TotalRecords: int(count),
	}

	ctx.Set("meta", meta)

	return query.Order(sort + " " + strings.ToUpper(order)).Offset((page - 1) * limit).Limit(limit)
}
