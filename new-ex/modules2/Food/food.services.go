package food

type FoodService interface {
	GetAll() ([]interface{}, error)
}

type foodService struct {
	foodRepo FoodRepository
}

func NewFoodService(foodRepo FoodRepository) FoodService {
	return &foodService{
		foodRepo: foodRepo,
	}
}

func (f *foodService) GetAll() ([]interface{}, error) {
	foods, err := f.foodRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return foods, nil
}
