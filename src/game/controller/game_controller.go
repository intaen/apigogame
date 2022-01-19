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
	var input domain.InputGameName
	err := c.ShouldBind(&input)
	if err != nil {
		util.HandleError(c, http.StatusBadRequest, "004", "Name can't be empty", err, "Error in BindJSON")
		return
	}

	result, err := g.gService.PredictionByName(input.Name)
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in Homepage")
		return
	}
	resp, log := util.ConvertResponse(result)
	util.HandleSuccess(c, http.StatusOK, "000", "Hello "+input.Name+"!", resp, log, "Success")
}

// Random Activity godoc
// @Tags Game
// @Summary Random Activity Game
// @Description This is API to get get random activity
// @Produce json
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /v1/dare [get]
func (g *GameController) Gamev1Dare(c *gin.Context) {
	result, err := g.gService.Punishment()
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in Homepage")
		return
	}

	resp, log := util.ConvertResponse(result)
	util.HandleSuccess(c, http.StatusOK, "000", "Here's your dare!", resp, log, "Success")
}

// Today Fact godoc
// @Tags Game
// @Summary Fact Game
// @Description This is API to get get fact about math
// @Produce json
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /v1/fact [get]
func (g *GameController) Gamev1Fact(c *gin.Context) {
	result, err := g.gService.CheckFact()
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in Homepage")
		return
	}

	resp, log := util.ConvertResponse(result)
	util.HandleSuccess(c, http.StatusOK, "000", "Here's today fact!", resp, log, "Success")
}
