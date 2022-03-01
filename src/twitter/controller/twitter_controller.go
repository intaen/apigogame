package controller

import (
	"apigogame/src/domain"
	"apigogame/src/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TwitterController struct {
	tService domain.TwitterService
	gService domain.GameService
}

func CreateTwitterController(r *gin.Engine, tService domain.TwitterService, gService domain.GameService) {
	TwitterController := TwitterController{tService: tService, gService: gService}

	// Twitter
	tweet := r.Group("/twitter")
	{
		tweet.POST("/timeline", TwitterController.GetTweets)
		tweet.GET("/tweet/fact", TwitterController.TweetRandomFact)
	}
}

// Get Tweets godoc
// @Tags Twitter
// @Summary Get Tweets in Timeline
// @Description This is API to see tweets in timeline
// @Accept json
// @Produce json
// @Param domain.ExamplePredictName body domain.InputTimeline true "Body"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /twitter/timeline [post]
func (t *TwitterController) GetTweets(c *gin.Context) {
	var input domain.InputTimeline
	err := c.ShouldBind(&input)
	if err != nil {
		util.HandleError(c, http.StatusBadRequest, "004", err.Error(), err, "Error in BindJSON")
		return
	}

	twts, err := t.tService.FindAllTweet(input.TotalDataShow)
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in Timeline")
		return
	}

	util.HandleSuccess(c, http.StatusOK, "000", "Check out your timeline!", gin.H{
		"Home": twts,
	}, twts, "Success")
}

// Tweet Random Fact godoc
// @Tags Twitter
// @Summary Random Fact
// @Description This is API to tweet random fact
// @Produce json
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /twitter/tweet/fact [get]
func (t *TwitterController) TweetRandomFact(c *gin.Context) {
	result, err := t.gService.FindRandomFact()
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in FindRandomFact")
		return
	}

	url, err := t.tService.CreateTweet(result.Fact)
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in CreateTweet")
		return
	}

	util.HandleSuccess(c, http.StatusOK, "000", "Check out the tweet!", gin.H{
		"URL": url,
	}, url, "Success")
}
