package controllers

import (
	"net/http"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"strings"

	"github.com/labstack/echo"
	"github.com/vipindasvg/cryptoserver/common"
)

const (
	 btsUrl = "https://api.hitbtc.com/api/2/public/currency" 
)	 
// request get courses /api/v1/getcourses?limit=2
func GetCurrency(c echo.Context) error {
	symbol := c.Param("symbol")
	client := &http.Client{}
	var url string
	if symbol != "all" {
		url = btsUrl + "/" + symbol
	} else {
		url = btsUrl
	}
	req, err := http.NewRequest("GET", btsUrl, nil)
	if err != nil {
		common.Log.WithField("handler", "get-currency").WithField("issue", "request").
			Errorln("bad request to bts:", err)
	}
	resp, err := client.Do(req)
	if resp.Status != "200 OK" {
		common.Log.WithField("handler", "get-currency").WithField("issue", "request").
			Errorln("can not read response from bts:", err)
	}
	return c.JSON(http.StatusOK, resp)
}