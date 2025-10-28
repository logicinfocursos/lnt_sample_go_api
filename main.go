// main.go
package main

// 0. Importações de dependências
import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/learning_new_techs/go/database"
	"github.com/learning_new_techs/go/structs"
)

func main() {
	// 1. Carrega variáveis do .env
	// Equivalente ao dotenv.config() no Node.js e load_dotenv() no Python
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	// 2. Define a porta da API
	// Equivalente ao uso de process.env.API_PORT ou os.getenv("API_PORT")
	API_PORT := os.Getenv("API_PORT")
	if API_PORT == "" {
		API_PORT = "3000"
	}

	// 3. Inicialização do framework (gin) e ORM (gorm)
	// Equivalente a const app = express() e const prisma = new PrismaClient() no Node.js
	r := gin.Default()
	database.ConnectDatabase()
	database.SeedMovies("database/movies.json")

	// 4. Middleware CORS (equivalente ao cors() do Express/FastAPI)
	// Permite requisições de qualquer origem e métodos
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Next()
	})

	// 5. Definição das rotas da API
	// Equivalente a app.get, app.post, etc. no Node.js e @app.get, @app.post no FastAPI, cada "verbo HTTP + endpoint" mapeia para uma função handler
	r.GET("/movies", getMovies)
	r.GET("/movies/:id", getMovie)
	r.POST("/movies", createMovie)
	r.PUT("/movies/:id", updateMovie)
	r.DELETE("/movies/:id", deleteMovie)

	// 6. Inicialização do servidor na porta definida (API_PORT)
	// Equivalente a app.listen(API_PORT) no Node.js e uvicorn.run(port=API_PORT) no Python
	log.Printf("API Go rodando na porta %s", API_PORT)
	r.Run(":" + API_PORT)
}

// 4.1 Rota GET /movies - Listar todos os filmes
func getMovies(c *gin.Context) {
	var movies []structs.Movie
	database.DB.Find(&movies)
	c.JSON(http.StatusOK, movies)
}

// 4.2 Rota GET /movies/{movie_id} - Obter um filme por ID
func getMovie(c *gin.Context) {
	id := c.Param("id")
	var movie structs.Movie
	if err := database.DB.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Filme não encontrado"})
		return
	}
	c.JSON(http.StatusOK, movie)
}

// 4.3 Rota POST /movies - Criar um novo filme
func createMovie(c *gin.Context) {
	var movie structs.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	database.DB.Create(&movie)
	c.JSON(http.StatusCreated, movie)
}

// 4.4 Rota PUT /movies/{movie_id} - Atualizar um filme
func updateMovie(c *gin.Context) {
	id := c.Param("id")
	var movie structs.Movie
	if err := database.DB.First(&movie, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Filme não encontrado"})
		return
	}
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	database.DB.Save(&movie)
	c.JSON(http.StatusOK, movie)
}

// 4.5 Rota DELETE /movies/{movie_id} - Deletar um filme
func deleteMovie(c *gin.Context) {
	id := c.Param("id")
	if database.DB.Delete(&structs.Movie{}, id).RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Filme não encontrado"})
		return
	}
	c.Status(http.StatusNoContent)
}
