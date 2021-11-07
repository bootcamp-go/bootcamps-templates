package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/usuario/repositorio/cmd/server/routes"
)

func main() {

	db, _ := sql.Open("mysql", "meli_sprint_user:Meli_Sprint#123/melisprint")
	r := gin.Default()

	router := routes.NewRouter(r, db)
	router.MapRoutes()

	if err := r.Run(); err != nil {
		panic(err)
	}
}
