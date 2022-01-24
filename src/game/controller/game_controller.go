package controller

import (
	"apigogame/src/domain"
	"apigogame/src/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GameController struct {
	gService domain.GameService
}

func CreateGameController(r *gin.Engine, gService domain.GameService) {
	GameController := GameController{gService: gService}
	v1 := r.Group("/game/v1")
	{
		v1.POST("/predict", GameController.GetPredictionByName)
		v1.GET("/dare", GameController.GetRandomActivity)
		v1.GET("/fact", GameController.GetRandomFact)
		v1.GET("/img", GameController.GetRandomImage)
		v1.POST("/joke", GameController.GetRandomJoke)
	}
}

// Get Prediction By Name godoc
// @Tags Game
// @Summary Prediction By Name
// @Description This is API to get prediction from name input by user
// @Accept json
// @Produce json
// @Param domain.ExamplePredictName body domain.ExamplePredictName true "Body"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /game/v1/predict [post]
func (g *GameController) GetPredictionByName(c *gin.Context) {
	var input domain.InputPredictName
	err := c.ShouldBind(&input)
	if err != nil {
		util.HandleError(c, http.StatusBadRequest, "004", err.Error(), err, "Error in BindJSON")
		return
	}

	result, err := g.gService.FindPredictionByName(input.Name)
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in FindPredictionByName")
		return
	}
	resp, log := util.ConvertResponse(result)
	util.HandleSuccess(c, http.StatusOK, "000", "Hello "+input.Name+"!", resp, log, "Success")
}

// Get Random Activity godoc
// @Tags Game
// @Summary Random Activity
// @Description This is API to get random activity
// @Produce json
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /game/v1/dare [get]
func (g *GameController) GetRandomActivity(c *gin.Context) {
	result, err := g.gService.FindRandomActivity()
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in FindRandomActivity")
		return
	}

	resp, log := util.ConvertResponse(result)
	util.HandleSuccess(c, http.StatusOK, "000", "Here's your dare!", resp, log, "Success")
}

// Get Random Fact godoc
// @Tags Game
// @Summary Random Fact
// @Description This is API to get random fact
// @Produce json
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /game/v1/fact [get]
func (g *GameController) GetRandomFact(c *gin.Context) {
	result, err := g.gService.FindRandomFact()
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in FindRandomFact")
		return
	}

	resp, log := util.ConvertResponse(result)
	util.HandleSuccess(c, http.StatusOK, "000", "Here's today fact!", resp, log, "Success")
}

// Get Random Image godoc
// @Tags Game
// @Summary Random Image
// @Description This is API to get random image
// @Produce json
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /game/v1/img [get]
func (g *GameController) GetRandomImage(c *gin.Context) {
	result, err := g.gService.FindRandomImage()
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in FindRandomImage")
		return
	}

	resp, log := util.ConvertResponse(result)
	util.HandleSuccess(c, http.StatusOK, "000", "Be happy!", resp, log, "Success")
}

// Get Random Joke godoc
// @Tags Game
// @Summary Random Joke
// @Description This is API to get random joke
// @Accept json
// @Produce json
// @Param domain.InputRandomJoke body domain.InputRandomJoke true "Body"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /game/v1/joke [post]
func (g *GameController) GetRandomJoke(c *gin.Context) {
	var input domain.InputRandomJoke
	err := c.ShouldBind(&input)
	if err != nil {
		util.HandleError(c, http.StatusBadRequest, "004", err.Error(), err, "Error in BindJSON")
		return
	}

	result, err := g.gService.FindRandomJoke(input.IsSafe)
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in FindRandomJoke")
		return
	}
	resp, log := util.ConvertResponse(result)
	util.HandleSuccess(c, http.StatusOK, "000", "Here's the joke!", resp, log, "Success")
}
