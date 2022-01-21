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

	// JSON
	v1 := r.Group("/game/v1")
	{
		v1.POST("/name", GameController.Gamev1Name)
		v1.GET("/dare", GameController.Gamev1Dare)
		v1.GET("/fact", GameController.Gamev1Fact)
		v1.GET("/img", GameController.Gamev1Img)
		v1.POST("/joke", GameController.Gamev1Joke)
	}
}

// Name Prediction godoc
// @Tags Game
// @Summary Prediction Game
// @Description This is API to get prediction from name input by user
// @Accept json
// @Produce json
// @Param Gamev1 body domain.ExampleGameName true "Name Prediction"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /v1/name [post]
func (g *GameController) Gamev1Name(c *gin.Context) {
	var input domain.InputPredictName
	err := c.ShouldBind(&input)
	if err != nil {
		util.HandleError(c, http.StatusBadRequest, "004", err.Error(), err, "Error in BindJSON")
		return
	}

	result, err := g.gService.PredictionByName(input.Name)
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in PredictionByName")
		return
	}
	resp, log := util.ConvertResponse(result)
	util.HandleSuccess(c, http.StatusOK, "000", "Hello "+input.Name+"!", resp, log, "Success")
}

// Random Activity godoc
// @Tags Game
// @Summary Random Activity Game
// @Description This is API to get random activity
// @Produce json
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /v1/dare [get]
func (g *GameController) Gamev1Dare(c *gin.Context) {
	result, err := g.gService.DoYouDare()
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in DoYouDare")
		return
	}

	resp, log := util.ConvertResponse(result)
	util.HandleSuccess(c, http.StatusOK, "000", "Here's your dare!", resp, log, "Success")
}

// Fact Check godoc
// @Tags Game
// @Summary Fact Game
// @Description This is API to get random fact
// @Produce json
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /v1/fact [get]
func (g *GameController) Gamev1Fact(c *gin.Context) {
	result, err := g.gService.CheckFact()
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in CheckFact")
		return
	}

	resp, log := util.ConvertResponse(result)
	util.HandleSuccess(c, http.StatusOK, "000", "Here's today fact!", resp, log, "Success")
}

// Random Image godoc
// @Tags Game
// @Summary Random Image Game
// @Description This is API to get random image
// @Produce json
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /v1/img [get]
func (g *GameController) Gamev1Img(c *gin.Context) {
	result, err := g.gService.CheckImg()
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in CheckImg")
		return
	}

	resp, log := util.ConvertResponse(result)
	util.HandleSuccess(c, http.StatusOK, "000", "Be happy!", resp, log, "Success")
}

// Random Joke godoc
// @Tags Game
// @Summary Random Joke Game
// @Description This is API to get random joke
// @Accept json
// @Produce json
// @Param Gamev1 body domain.InputCheckJoke true "Random Joke"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /v1/name [post]
func (g *GameController) Gamev1Joke(c *gin.Context) {
	var input domain.InputCheckJoke
	err := c.ShouldBind(&input)
	if err != nil {
		util.HandleError(c, http.StatusBadRequest, "004", err.Error(), err, "Error in BindJSON")
		return
	}

	mode := ""
	if input.SafeMode {
		mode = "safe-mode"
	}

	result, err := g.gService.CheckJoke(mode)
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in CheckJoke")
		return
	}
	resp, log := util.ConvertResponse(result)
	util.HandleSuccess(c, http.StatusOK, "000", "Here's the joke!", resp, log, "Success")
}
