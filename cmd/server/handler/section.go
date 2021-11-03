package handler

import (
	"github.com/usuario/repositorio/internal/domain"

	"github.com/gin-gonic/gin"
)

type Section struct {
	// sectionService section.Service
}

func NewSection() *Section {
	return &Section{
		// sectionService: s,
	}
}

func (s *Section) GetAll() gin.HandlerFunc {

	type response struct {
		Data []domain.Section `json:"data"`
	}

	return func(c *gin.Context) {

	}
}

func (s *Section) Get() gin.HandlerFunc {

	type response struct {
		Data domain.Section `json:"data"`
	}

	return func(c *gin.Context) {

	}
}

func (s *Section) Store() gin.HandlerFunc {
	type request struct {
		SectionNumber      int `json:"section_number"`
		CurrentTemperature int `json:"current_temperature"`
		MinTemperature     int `json:"minimum_temperature"`
		CurrentCapacity    int `json:"current_capacity"`
		MinCapacity        int `json:"minimum_capacity"`
		MaxCapacity        int `json:"maximum_capacity"`
		WarehouseID        int `json:"warehouse_id"`
		ProductTypeID      int `json:"product_type_id"`
	}

	type response struct {
		Data domain.Section `json:"data"`
	}

	return func(c *gin.Context) {

	}
}

func (s *Section) Update() gin.HandlerFunc {

	type request struct {
		SectionNumber      int `json:"section_number"`
		CurrentTemperature int `json:"current_temperature"`
		MinTemperature     int `json:"minimum_temperature"`
		CurrentCapacity    int `json:"current_capacity"`
		MinCapacity        int `json:"minimum_capacity"`
		MaxCapacity        int `json:"maximum_capacity"`
		WarehouseID        int `json:"warehouse_id"`
		ProductTypeID      int `json:"product_type_id"`
	}

	type response struct {
		Data string `json:"data"`
	}

	return func(c *gin.Context) {

	}
}

func (s *Section) Delete() gin.HandlerFunc {

	return func(c *gin.Context) {

	}
}
