package companies

import (
	"net/http"
	"xmgo/pkg/usecase/companies"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	usecase companies.ICompanies
}

type CompanyGet struct {
	ID int `uri:"id" binding:"required"`
}

func NewHandler(uc companies.ICompanies) *Handler {
	return &Handler{usecase: uc}
}

func (h *Handler) InitRoutes(r gin.IRouter) {
	r.GET("/", h.GetAll)
	r.GET("/:id", h.Get)
}

func (h *Handler) GetAll(c *gin.Context) {
	result := h.usecase.GetAll()
	c.String(http.StatusOK, result)
}

func (h *Handler) Get(c *gin.Context) {
	var err error
	//id, err := strconv.Atoi(c.Param("id"))

	var req = CompanyGet{}
	if err = c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	company, err := h.usecase.Get(c, req.ID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, company)
}
