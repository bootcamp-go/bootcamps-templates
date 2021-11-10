package handler

import (
	"github.com/gin-gonic/gin"
)

type Seller struct {
	// sellerService seller.Service
}

func NewSeller() *Seller {
	return &Seller{
		// sellerService: s,
	}
}

func (s *Seller) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (s *Seller) Get() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (s *Seller) Create() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (s *Seller) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (s *Seller) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
