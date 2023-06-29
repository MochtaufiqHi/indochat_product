package handler

import (
	dto "diskon/dto/result"
	"diskon/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type cusHandlers struct {
	CusRepo repository.CusRepo
}

type catHandlers struct {
	CatRepo repository.CateRepo
}

type proHandlers struct {
	ProRepo repository.ProRepo
}

type orHandlers struct {
	OrRepo repository.OrRepo
}

func HandlerCus(CusRepo repository.CusRepo) *cusHandlers {
	return &cusHandlers{CusRepo}
}

func HandlerCat(CatRepo repository.CateRepo) *catHandlers {
	return &catHandlers{CatRepo}
}

func HandlerPro(ProRepo repository.ProRepo) *proHandlers {
	return &proHandlers{ProRepo}
}

func HandlerOr(OrRepo repository.OrRepo) *orHandlers {
	return &orHandlers{OrRepo}
}

func (h *cusHandlers) GetCus(c echo.Context) error {
	cus, err := h.CusRepo.GetCus()

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: "salahhhhhh"})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: cus})
}

func (h *catHandlers) GetCat(c echo.Context) error {
	cat, err := h.CatRepo.GetCat()

	if err != nil {
		return c.JSON(http.StatusOK, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Errorrr boy"})
	}

	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: cat})
}

func (h *proHandlers) GetPro(c echo.Context) error {
	pro, err := h.ProRepo.GetPro()

	if err != nil {
		return c.JSON(http.StatusOK, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Errorrr boy"})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: pro})
}

func (h *orHandlers) GetOr(c echo.Context) error {
	or, err := h.OrRepo.GetOr()
	if err != nil {
		return c.JSON(http.StatusOK, dto.ErrorResult{Code: http.StatusBadRequest, Message: "Errorrr boy"})
	}
	return c.JSON(http.StatusOK, dto.SuccessResult{Code: http.StatusOK, Data: or})
}
