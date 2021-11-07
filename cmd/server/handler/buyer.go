package handler

import (
	"github.com/gin-gonic/gin"
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
	return func(c *gin.Context) {}
}

func (b *Buyer) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (b *Buyer) Store() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (b *Buyer) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (b *Buyer) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
