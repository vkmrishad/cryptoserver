package routers

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/vipindasvg/cryptoserver/controllers"
	"os"
)

const (
	versionpref = "/api/v1"
)

func InitRoutes() *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.Logger.SetLevel(log.DEBUG)
	e.Logger.SetOutput(os.Stdout)
	if l, ok := e.Logger.(*log.Logger); ok {
		l.SetHeader("${time_rfc3339} ${level}")
	}
	// liveness probe for k8s
	e.GET(versionpref+"cryptoprice/currency/:symbol", controllers.GetCryptoPrice)
	e.GET(versionpref+"cryptoprice/currency/all", controllers.GetCryptoPrices)
	return e
}	
