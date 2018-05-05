
package main

import(
	"net/http"
	"github.com/sirupsen/logrus"
	"github.com/labstack/echo"

)

func init(){
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})

}

func main(){
	e := echo.New()	
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Hello")
	
	})

	e.logrus.Fatal(e.Start(":5000"))

}
