package rekuest

import (
	"strings"

	"github.com/manot40/gomenu/internal/models"
)

func TransformTags(bind *string, input []string) error {
	for _, tag := range input {
		var tagModel models.Tag

		if err := models.DB.Where("name = ?", tag).First(&tagModel).Error; err != nil {
			return err
		}
	}

	*bind = strings.Join(input, ",")

	return nil
}
