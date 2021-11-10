package handler

import (
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
	return func(c *gin.Context) {}
}

func (s *Section) Get() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (s *Section) Create() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (s *Section) Update() gin.HandlerFunc {
	return func(c *gin.Context) {}
}

func (s *Section) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {}
}
