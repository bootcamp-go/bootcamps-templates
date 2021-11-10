package handler

import (
	"github.com/gin-gonic/gin"
)

type Product struct {
	// productService product.Service
}

func NewProduct() *Product {
	return &Product{
		// productService: w,
	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (p *Product) Get() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (p *Product) Create() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (p *Product) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (p *Product) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
