package routes

import (
	"epm/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Rota para criar um livro
	r.POST("/employees", controllers.CreateEmployee)

	// Rota para listar todos os livros
	r.GET("/employees", controllers.GetEmployees)

	// Rota para obter um livro pelo ID
	r.GET("/employees/:id", controllers.GetEmployee)

	// Rota para atualizar um livro
	r.PUT("/employees/:id", controllers.UpdateEmployee)

	// Rota para deletar um livro
	r.DELETE("/employees/:id", controllers.DeleteEmployee)

	return r
}
