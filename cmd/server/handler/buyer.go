package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/usuario/repositorio/internal/domain"
)

type Buyer struct {
	// buyerService buyer.Service
}

func NewBuyer() *Buyer {
	return &Buyer{
		// buyerService: b,
	}
}

func (b *Buyer) Get() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
	}
}

func (b *Buyer) GetAll() gin.HandlerFunc {
	type response struct {
		Data []domain.Buyer `json:"data"`
	}

	return func(c *gin.Context) {

	}
}

func (b *Buyer) Store() gin.HandlerFunc {
	type request struct {
		CardNumberID string `json:"card_number_id"`
		FirstName    string `json:"first_name"`
		LastName     string `json:"last_name"`
	}

	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {
	}
}

func (b *Buyer) Update() gin.HandlerFunc {
	type request struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}

	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {

	}
}

func (b *Buyer) Delete() gin.HandlerFunc {
	type response struct {
		Data interface{} `json:"data"`
	}

	return func(c *gin.Context) {

	}
}
