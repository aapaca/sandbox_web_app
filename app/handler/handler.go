package handler

import (
	"net/http"
	"github.com/labstack/echo"
)

func MainPage() echo.HandlerFunc {
	return func(c echo.Context) error {     //c をいじって Request, Responseを色々する
		return c.String(http.StatusOK, "Hello. Aapaca is the most excited musician's library in the world. you can explore musicians, albums, and songs here.")
	}
}