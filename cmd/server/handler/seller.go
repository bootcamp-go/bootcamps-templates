package handler

import (
	"github.com/usuario/repositorio/internal/domain"

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

	type response struct {
		Data []domain.Seller `json:"data"`
	}

	return func(c *gin.Context) {
	}
}

func (s *Seller) Get() gin.HandlerFunc {

	type response struct {
		Data domain.Seller `json:"data"`
	}

	return func(c *gin.Context) {
	}
}

func (s *Seller) Store() gin.HandlerFunc {
	type request struct {
		CID         int    `json:"cid"`
		CompanyName string `json:"company_name"`
		Address     string `json:"address"`
		Telephone   string `json:"telephone"`
		LocalityID  int    `json:"locality_id"`
	}

	type response struct {
		Data domain.Seller `json:"data"`
	}

	return func(c *gin.Context) {
	}
}

func (s *Seller) Update() gin.HandlerFunc {

	type request struct {
		SellerID    int    `json:"seller_id"`
		CID         int    `json:"cid"`
		CompanyName string `json:"company_name"`
		Address     string `json:"address"`
		Telephone   string `json:"telephone"`
		LocalityID  int    `json:"locality_id"`
	}

	type response struct {
		Data string `json:"data"`
	}

	return func(c *gin.Context) {
	}
}

func (s *Seller) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {
	}
}
