package food_domain

import (
	"golang-kubernetes/foods/db"
	"golang-kubernetes/foods/utils/error_formats"
	"golang-kubernetes/foods/utils/error_utils"
)

var FoodRepo foodDomain = &foodRepo{}

const (
	queryGetFoodById     = `SELECT id, name, price, stock FROM foods WHERE id = $1`
	queryCreateFood      = `INSERT INTO foods(name, price, stock) VALUES($1, $2, $3) RETURNING id, name, price, stock, created_at`
	queryReduceFoodStock = `UPDATE foods SET stock = stock - $2 WHERE id = $1`
	queryGetAllFoods     = `SELECT id, name, price, stock, created_at FROM foods`
)

type foodDomain interface {
	GetFoodById(id int32) (*Food, error_utils.MessageErr)
	CreateFood(*Food) (*Food, error_utils.MessageErr)
	ReduceFoodStock(*Food) error_utils.MessageErr
	GetAllFoods() ([]*Food, error_utils.MessageErr)
}

type foodRepo struct{}

func (f *foodRepo) GetFoodById(id int32) (*Food, error_utils.MessageErr) {
	db := db.GetDB()

	row := db.QueryRow(queryGetFoodById, id)

	var food Food

	if err := row.Scan(&food.Id, &food.Name, &food.Price, &food.Stock); err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &food, nil
}

func (f *foodRepo) CreateFood(foodReq *Food) (*Food, error_utils.MessageErr) {
	db := db.GetDB()

	row := db.QueryRow(queryCreateFood, foodReq.Name, foodReq.Price, foodReq.Stock)

	var food Food

	if err := row.Scan(&food.Id, &food.Name, &food.Price, &food.Stock, &food.CreatedAt); err != nil {
		return nil, error_formats.ParseError(err)
	}

	return &food, nil
}

func (f *foodRepo) ReduceFoodStock(foodReq *Food) error_utils.MessageErr {
	db := db.GetDB()

	_, err := db.Exec(queryReduceFoodStock, foodReq.Id, foodReq.Stock)

	if err != nil {
		return error_formats.ParseError(err)
	}

	return nil
}

func (f *foodRepo) GetAllFoods() ([]*Food, error_utils.MessageErr) {
	db := db.GetDB()

	rows, err := db.Query(queryGetAllFoods)
	if err != nil {
		return nil, error_formats.ParseError(err)
	}

	var foods []*Food

	for rows.Next() {
		var food Food
		err = rows.Scan(&food.Id, &food.Name, &food.Price, &food.Stock, &food.CreatedAt)

		if err != nil {
			return nil, error_formats.ParseError(err)
		}

		foods = append(foods, &food)
	}

	return foods, nil
}
