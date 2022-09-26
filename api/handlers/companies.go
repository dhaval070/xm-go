package handlers

import (
	"errors"
	"log"
	"net/http"
	"xmgo/pkg/apperror"
	"xmgo/pkg/usecase/companies"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase companies.ICompanies
}

type CompanyGetReq struct {
	ID uint `uri:"id" binding:"required"`
}

// NewHandler returns new Handler instance
func NewHandler(uc companies.ICompanies) *Handler {
	return &Handler{usecase: uc}
}

// InitRoutes binds handler functions to urls
func (h *Handler) InitRoutes(r gin.IRouter) {
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
}

// GetAll returns all companies records
func (h *Handler) GetAll(c *gin.Context) {
	result := h.usecase.GetAll()
	c.String(http.StatusOK, result)
}

// Get returns company record by ID
func (h *Handler) Get(c *gin.Context) {
	var err error

	var req = CompanyGetReq{}
	if err = c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	company, err := h.usecase.Get(c, req.ID)

	if err != nil {
		if errors.Is(err, apperror.ErrNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, company)
}
