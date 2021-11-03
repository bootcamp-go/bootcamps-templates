package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/usuario/repositorio/internal/domain"
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

	type response struct {
		Data []domain.Product `json:"data"`
	}

	return func(c *gin.Context) {

	}
}

func (p *Product) Get() gin.HandlerFunc {

	type response struct {
		Data domain.Product `json:"data"`
	}

	return func(c *gin.Context) {

	}
}

func (p *Product) Store() gin.HandlerFunc {
	type request struct {
		Description    string  `json:"description"`
		ExpirationRate int     `json:"expiration_rate"`
		FreezingRate   int     `json:"freezing_rate"`
		Height         float32 `json:"height"`
		Length         float32 `json:"length"`
		Netweight      float32 `json:"netweight"`
		ProductCode    string  `json:"product_code"`
		RecomFreezTemp float32 `json:"recommended_freezing_temperature"`
		Width          float32 `json:"width"`
		ProductTypeID  int     `json:"product_type_id"`
		SellerID       int     `json:"seller_id"`
	}

	type response struct {
		Data domain.Product `json:"data"`
	}

	return func(c *gin.Context) {

	}
}

func (p *Product) Update() gin.HandlerFunc {

	type request struct {
		Description    string  `json:"description"`
		ExpirationRate int     `json:"expiration_rate"`
		FreezingRate   int     `json:"freezing_rate"`
		Height         float32 `json:"height"`
		Length         float32 `json:"length"`
		Netweight      float32 `json:"netweight"`
		ProductCode    string  `json:"product_code"`
		RecomFreezTemp float32 `json:"recommended_freezing_temperature"`
		Width          float32 `json:"width"`
		ProductTypeID  int     `json:"product_type_id"`
		SellerID       int     `json:"seller_id"`
	}

	type response struct {
		Data string `json:"data"`
	}

	return func(c *gin.Context) {

	}
}

func (p *Product) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {

	}
}
