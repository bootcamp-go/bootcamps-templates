package handler

import (
	"github.com/gin-gonic/gin"
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
	return func(c *gin.Context) {}
}

func (e *Employee) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (e *Employee) Create() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (e *Employee) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (e *Employee) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
