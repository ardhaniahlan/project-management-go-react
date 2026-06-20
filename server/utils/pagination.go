package utils

import (
	"strings"

	"gorm.io/gorm"
)

func Paginate[T any](db *gorm.DB, limit, offset int, sort string, dest *[]T) (int64, error) {
	var count int64

	if err := db.Count(&count).Error; err != nil {
		return 0, err
	}

	if sort != "" {
		if strings.HasPrefix(sort, "-") {
			sort = strings.TrimPrefix(sort, "-") + " DESC"
		} else {
			sort += " ASC"
		}
		db = db.Order(sort)
	}

	err := db.Limit(limit).Offset(offset).Find(dest).Error

	return count, err
}