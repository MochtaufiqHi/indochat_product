package routes

import (
	"diskon/handler"
	"diskon/pkg/mysql"
	"diskon/repository"

	"github.com/labstack/echo/v4"
)

func RouteInit(e *echo.Group) {
	CusRoute(e)
	CatRoute(e)
	ProRepo(e)
	OrRepo(e)
}

func CusRoute(e *echo.Group) {
	CusRepo := repository.RepoConnCus(mysql.DB)
	h := handler.HandlerCus(CusRepo)

	e.GET("/customer", h.GetCus)
}

func CatRoute(e *echo.Group) {
	CatRepo := repository.RepoConnCus(mysql.DB)
	h := handler.HandlerCat(CatRepo)

	e.GET("/category", h.GetCat)
}

func ProRepo(e *echo.Group) {
	ProRepo := repository.RepoConnCus(mysql.DB)
	h := handler.HandlerPro(ProRepo)

	e.GET("/product", h.GetPro)
}

func OrRepo(e *echo.Group) {
	OrRepo := repository.RepoConnCus(mysql.DB)
	h := handler.HandlerOr(OrRepo)

	e.GET("/order", h.GetOr)
}
