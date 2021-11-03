package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/usuario/repositorio/internal/domain"
)

type Employee struct {
	// employeeService employee.Service
}

func NewEmployee() *Employee {
	return &Employee{
		// employeeService: e,
	}
}

func (e *Employee) Get() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {

	}
}

func (e *Employee) GetAll() gin.HandlerFunc {
	type response struct {
		Data []domain.Employee `json:"data"`
	}

	return func(c *gin.Context) {

	}
}

func (e *Employee) Store() gin.HandlerFunc {
	type request struct {
		CardNumberID string `json:"card_number_id"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
		WarehouseID  int    `json:"warehouse_id"`
	}

	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {

	}
}

func (e *Employee) Update() gin.HandlerFunc {
	type request struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		WarehouseID int    `json:"warehouse_id"`
	}

	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {

	}
}

func (e *Employee) Delete() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {

	}
}
