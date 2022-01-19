package main

import (
	gc "apigogame/src/game/controller"
	gr "apigogame/src/game/repo"
	gs "apigogame/src/game/service"
	"apigogame/src/util"
	"net/http"

	"github.com/gin-gonic/contrib/static"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	_ "apigogame/src/docs"
)

func init() {
	viper.SetConfigFile("src/config/Config.json")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal(err)
	}
}

// @title GOGAME
// @version 1.0
// @description This page is API documentation for a little game like predict age by name
// @schemes http
// @host localhost:1111
// @BasePath /game
// @contact.name Developer
// @contact.email intanmarsjaf@outlook.com
func main() {
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})
	r.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Handle wrong method
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		util.HandleError(c, http.StatusMethodNotAllowed, "045", "Method Not Allowed", "Method Not Allowed", "Method Not Allowed")
	})
	// Handle no route
	r.NoRoute(func(c *gin.Context) {
		util.HandleError(c, http.StatusNotFound, "044", "Page Not Found", "Page Not Found", "Page Not Found")
	})

	// Swagger
	url := ginSwagger.URL(viper.GetString("host") + ":" + viper.GetString("port") + "/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))                            // http://localhost:1111/swagger/index.html

	// Initiate Repo
	gRepo := gr.CreateGameRepoImpl()

	// Initiate Service
	gService := gs.CreateGameServiceImpl(gRepo)

	// Initiate Controller
	gc.CreateGameController(r, gService)

	// r.Run(":" + viper.GetString("port"))
	r.Run() // Heroku will supply automatically
}
