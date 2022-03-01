package main

import (
	gc "apigogame/src/game/controller"
	gr "apigogame/src/game/repo"
	gs "apigogame/src/game/service"
	mc "apigogame/src/masterdata/controller"
	mdr "apigogame/src/masterdata/repo"
	mds "apigogame/src/masterdata/service"

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
	viper.SetConfigFile("config/Config.json")
	err := viper.ReadInConfig()
	if err != nil {
		logger.Fatal(err)
	}
}

// @title GOGAME
// @version 1.0
// @description This page is API documentation for a little game like predict age, gender by name, get random fact, image, joke
// @schemes https
// @host apigogame.herokuapp.com
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
	// https://apigogame.herokuapp.com
	// viper.GetString("host") + ":" + viper.GetString("port")
	url := ginSwagger.URL("https://apigogame.herokuapp.com" + "/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Initiate Repo
	mdRepo := mdr.CreateMasterDataRepoImpl()
	gRepo := gr.CreateGameRepoImpl()

	// Initiate Service
	mdService := mds.CreateMasterDataServiceImpl(mdRepo)
	gService := gs.CreateGameServiceImpl(gRepo, mdService)

	// Initiate Controller
	mc.CreateMasterDataController(r, mdService)
	gc.CreateGameController(r, gService)

	// r.Run(":" + viper.GetString("port"))
	r.Run() // Heroku will supply automatically
}
