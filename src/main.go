package main

import (
	"apigogame/src/driver"
	gc "apigogame/src/game/controller"
	gr "apigogame/src/game/repo"
	gs "apigogame/src/game/service"
	mc "apigogame/src/masterdata/controller"
	mdr "apigogame/src/masterdata/repo"
	mds "apigogame/src/masterdata/service"

	tc "apigogame/src/twitter/controller"
	tr "apigogame/src/twitter/repo"
	ts "apigogame/src/twitter/service"
	"apigogame/src/util"
	"log"
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
// @schemes http
// @host localhost:1111
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

	// Initiate Twitter Client
	client, err := driver.InitiateClient()
	if err != nil {
		log.Printf("Error %v", err)
	}

	// Initiate Repo
	mdRepo := mdr.CreateMasterDataRepoImpl()
	gRepo := gr.CreateGameRepoImpl()
	tRepo := tr.CreateTwitterRepoImpl(client)

	// Initiate Service
	mdService := mds.CreateMasterDataServiceImpl(mdRepo)
	gService := gs.CreateGameServiceImpl(gRepo, mdService)
	tService := ts.CreateTwitterServiceImpl(tRepo, gRepo)

	// Initiate Controller
	mc.CreateMasterDataController(r, mdService)
	gc.CreateGameController(r, gService)
	tc.CreateTwitterController(r, tService, gService)

	r.Run(":" + viper.GetString("port"))
	// r.Run() // Heroku will supply automatically
}
