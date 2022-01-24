package controller

import (
	"apigogame/src/domain"
	"apigogame/src/util"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type MasterDataController struct {
	mdService domain.MasterDataService
}

func CreateMasterDataController(r *gin.Engine, mdService domain.MasterDataService) {
	MasterDataController := MasterDataController{mdService: mdService}
	v1 := r.Group("/master")
	{
		v1.GET("/api", MasterDataController.GetListAPIByAuth)
		v1.GET("/country", MasterDataController.GetListCountry)
	}
}

// Get List Country godoc
// @Tags Master Data
// @Summary List Country
// @Description This is API to get list country by country id or region
// @Accept json
// @Produce json
// @Param id query string false "Params 1"
// @Param region query string false "Params 2"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /master/country [get]
func (md *MasterDataController) GetListCountry(c *gin.Context) {
	// Set params
	cid := c.Query("id")
	reg := c.Query("region")

	// Get Country
	result, err := md.mdService.FindCountry(strings.ToUpper(cid), strings.ToUpper(reg))
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in FindCountry")
		return
	}

	resp, logResp := util.ConvertResponse(result)
	util.HandleSuccess(c, http.StatusOK, "000", "Data found", resp, logResp, "Success")
}

// Get List API godoc
// @Tags Master Data
// @Summary List Public API
// @Description This is API to get list of public API by auth or category
// @Accept json
// @Produce json
// @Param auth query boolean false "Params 1"
// @Param category query string false "Params 2"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.BadRequestResponse
// @Router /master/api [get]
func (md *MasterDataController) GetListAPIByAuth(c *gin.Context) {
	// Set params
	auth := c.Query("auth")
	cat := c.Query("category")

	// Check isAuth and add kind of auth
	isAuth, _ := strconv.ParseBool(auth)
	var auths []string
	if !isAuth && auth != "" {
		auths = append(auths, "null")
	} else if isAuth && auth != "" {
		auths = append(auths, "apikey", "oauth")
	}

	// Get API
	result, err := md.mdService.FindAPI(auths, cat)
	if err != nil {
		util.HandleSuccess(c, http.StatusOK, "001", err.Error(), nil, err, "Error in FindAPI")
		return
	}

	resp, _ := util.ConvertResponse(result)
	util.HandleSuccess(c, http.StatusOK, "000", "Data found", resp, "TotalData: "+strconv.Itoa(result.Count), "Success")
}
