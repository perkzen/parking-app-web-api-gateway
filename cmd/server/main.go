package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"parking-app-web-api-gateway/internal/routes"
	"parking-app-web-api-gateway/internal/utils"
)

func main() {

	_ = godotenv.Load(".env")

	app := gin.Default()

	routes.InitRouter(app)

	port, _ := utils.GetEnvVar("PORT")

	log.Printf("Swagger UI is available on http://localhost:%s/swagger/index.html\n\n", port)
	log.Fatal(app.Run(":" + port))
}
