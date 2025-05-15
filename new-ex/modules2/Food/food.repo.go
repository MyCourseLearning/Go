package food

import (
	"db-excercise/model"

	"gorm.io/gorm"
)

type FoodRepository interface {
	GetAll() ([]interface{}, error)
}

type foodRepository struct {
	db *gorm.DB
}

func NewFoodRepository(db *gorm.DB) FoodRepository {
	return &foodRepository{
		db: db,
	}
}

func (f *foodRepository) GetAll() ([]interface{}, error) {
	var foods []model.Food
	if err := f.db.Find(&foods).Error; err != nil {
		return nil, err
	}
	var result []interface{}
	for _, food := range foods {
		result = append(result, food)
	}
	return result, nil
}
