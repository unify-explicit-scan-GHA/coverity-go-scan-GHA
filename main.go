package main

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/calculi-corp/template-go-testing/calculator"
	"github.com/labstack/echo"
)

func main() {
	cal := calculator.Calculator{}
	fmt.Println("Add returns: ", cal.Add(10, 10))
	fmt.Println("Sub returns: ", cal.Sub(10, 5))
	fmt.Println("Mul returns: ", cal.Mul(10, 10))
	fmt.Println("Div returns: ", cal.Div(10, 2))

	data := []byte("This is a test")
	hash := md5.Sum(data)
	fmt.Printf("MD5 hash: %x\n", hash)

	// Echo instance
	e := echo.New()
	e.GET("/", hello)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

// go test resulr in .xml format use below command
// go test -v 2>&1 ./... | go-junit-report -set-exit-code > report.xml
