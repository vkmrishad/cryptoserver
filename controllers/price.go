package controllers

import (
	"net/http"
	"io/ioutil"
	"encoding/json"

	"github.com/labstack/echo"
	"github.com/vipindasvg/cryptoserver/common"
	"github.com/vipindasvg/cryptoserver/models"
)

const (
	 btstickerUrl = "https://api.hitbtc.com/api/2/public/ticker" 
	 btscurrencyUrl = "https://api.hitbtc.com/api/2/public/currency"
	 btssymbolUrl = "https://api.hitbtc.com/api/2/public/symbol"
)	 

// request get courses /api/v1/getcourses?limit=2
func GetCryptoPrice(c echo.Context) error {
	symbol := c.Param("symbol")
	
	client := &http.Client{}
	//To get the real time crypto value by symbol
	req, err := http.NewRequest("GET", btstickerUrl + "/" + symbol, nil)
	if err != nil {
		common.Log.WithField("handler", "get-currency-by-symbol").WithField("issue", "request").
			Errorln("bad request to bts:", err)
	}
	resp, err := client.Do(req)
	if resp.Status != "200 OK" {
		common.Log.WithField("handler", "get-currency-by-symbol").WithField("issue", "request").
			Errorln("can not read response from bts:", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		common.Log.WithField("handler", "get-currency-by-symbol").WithField("issue", "request").
			Errorln("can not read response from bts:", err)
	}
	price := new(models.Price)
	if err := json.Unmarshal(body, price); err != nil {
		common.Log.WithField("handler", "get-currency-by-symbol").WithField("issue", "request").
			Errorln("can not unmarshall response json from bts:", err)
	}
	defer resp.Body.Close()
	//To get the symbol informations mainly feecurrency and basecurrency
	req1, err := http.NewRequest("GET", btssymbolUrl + "/" + symbol, nil)
	if err != nil {
		common.Log.WithField("handler", "get-currency-by-symbol").WithField("issue", "request").
			Errorln("bad request to bts:", err)
	}
	resp1, err := client.Do(req1)
	if resp1.Status != "200 OK" {
		common.Log.WithField("handler", "get-currency-by-symbol").WithField("issue", "request").
			Errorln("can not read response from bts:", err)
	}
	body1, err := ioutil.ReadAll(resp1.Body)
	if err != nil {
		common.Log.WithField("handler", "get-currency-by-symbol").WithField("issue", "request").
			Errorln("can not read response from bts:", err)
	}
	sym := new(models.Symbol)
	if err := json.Unmarshal(body1, sym); err != nil {
		common.Log.WithField("handler", "get-currency-by-symbol").WithField("issue", "request").
			Errorln("can not unmarshall response json from bts:", err)
	}
	defer resp1.Body.Close()

	//To get the currency information mainly the fullname
	req2, err := http.NewRequest("GET", btscurrencyUrl + "/" + sym.BaseCurrency, nil)
	if err != nil {
		common.Log.WithField("handler", "get-currency-by-symbol").WithField("issue", "request").
			Errorln("bad request to bts:", err)
	}
	resp2, err := client.Do(req2)
	if resp2.Status != "200 OK" {
		common.Log.WithField("handler", "get-currency-by-symbol").WithField("issue", "request").
			Errorln("can not read response from bts:", err)
	}
	body2, err := ioutil.ReadAll(resp2.Body)
	if err != nil {
		common.Log.WithField("handler", "get-currency-by-symbol").WithField("issue", "request").
			Errorln("can not read response from bts:", err)
	}
	cur := new(models.Currency)
	if err := json.Unmarshal(body2, cur); err != nil {
		common.Log.WithField("handler", "get-currency-by-symbol").WithField("issue", "request").
			Errorln("can not unmarshall response json from bts:", err)
	}
	defer resp2.Body.Close()

	//Get complete information
	total := new(models.TotalPrice)
	total.Id = sym.BaseCurrency
	total.Price = price
	total.FullName = cur.FullName
	total.FeeCurrency = sym.FeeCurrency
	return c.JSON(http.StatusOK, total)
}

// request get courses /api/v1/getcourses?limit=2
func GetCryptoPrices(c echo.Context) error {
	client := &http.Client{}
	//To get the real time crypto value by symbol
	req, err := http.NewRequest("GET", btstickerUrl, nil)
	if err != nil {
		common.Log.WithField("handler", "get-currency").WithField("issue", "request").
			Errorln("bad request to bts:", err)
	}
	resp, err := client.Do(req)
	if resp.Status != "200 OK" {
		common.Log.WithField("handler", "get-currency").WithField("issue", "request").
			Errorln("can not read response from bts:", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		common.Log.WithField("handler", "freshdesk-list").WithField("issue", "request").
			Errorln("can not read response from bts:", err)
	}
	var prices []*models.Price
	if err := json.Unmarshal(body, &prices); err != nil {
		common.Log.WithField("handler", "freshdesk-list").WithField("issue", "request").
			Errorln("can not unmarshall response json from bts:", err)
	}
	defer resp.Body.Close()
	var totalprices []*models.TotalPrice
	for _, price := range prices {
		//To get the symbol informations mainly feecurrency and basecurrency
		req1, err := http.NewRequest("GET", btssymbolUrl + "/" + price.Symbol, nil)
		if err != nil {
			common.Log.WithField("handler", "get-currency").WithField("issue", "request").
			Errorln("bad request to bts:", err)
		}
		resp1, err := client.Do(req1)
		if resp1.Status != "200 OK" {
			common.Log.WithField("handler", "get-currency").WithField("issue", "request").
			Errorln("can not read response from bts:", err)
		}
		body1, err := ioutil.ReadAll(resp1.Body)
		if err != nil {
			common.Log.WithField("handler", "freshdesk-list").WithField("issue", "request").
			Errorln("can not read response from bts:", err)
		}
		sym := new(models.Symbol)
		if err := json.Unmarshal(body1, sym); err != nil {
			common.Log.WithField("handler", "freshdesk-list").WithField("issue", "request").
			Errorln("can not unmarshall response json from bts:", err)
		}
		defer resp1.Body.Close()

		//To get the currency information mainly the fullname
		req2, err := http.NewRequest("GET", btscurrencyUrl + "/" + sym.BaseCurrency, nil)
		if err != nil {
			common.Log.WithField("handler", "get-currency").WithField("issue", "request").
			Errorln("bad request to bts:", err)
		}
		resp2, err := client.Do(req2)
		if resp2.Status != "200 OK" {
			common.Log.WithField("handler", "get-currency").WithField("issue", "request").
			Errorln("can not read response from bts:", err)
		}
		body2, err := ioutil.ReadAll(resp2.Body)
		if err != nil {
			common.Log.WithField("handler", "freshdesk-list").WithField("issue", "request").
			Errorln("can not read response from bts:", err)
		}
		cur := new(models.Currency)
		if err := json.Unmarshal(body2, cur); err != nil {
			common.Log.WithField("handler", "freshdesk-list").WithField("issue", "request").
			Errorln("can not unmarshall response json from bts:", err)
		}
		defer resp2.Body.Close()

		//Get complete information
		total := new(models.TotalPrice)
		total.Id = sym.BaseCurrency
		total.Price = price
		total.FullName = cur.FullName
		total.FeeCurrency = sym.FeeCurrency
		totalprices = append(totalprices, total)
	}	
	return c.JSON(http.StatusOK, totalprices)
}

