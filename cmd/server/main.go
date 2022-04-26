package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/usuario/repositorio/cmd/server/routes"
)

func main() {
	// NO MODIFICAR
	r := gin.Default()

	router := routes.NewRouter(r)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}
