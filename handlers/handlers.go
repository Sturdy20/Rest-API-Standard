package handlers

import (
	"company/models"
	"company/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IHandler interface {
	AddRegisteHandler(c *gin.Context)
	GetAllEmployeesHandler(c *gin.Context)
	GetCompanyByIDHandler(c *gin.Context)
	UpdateEmployeeHandler(c *gin.Context)
	DeleteEmployeeHandler(c *gin.Context)

}

type handler struct {
	s services.IService
}

func NewHandler(s services.IService) IHandler {
	return &handler{s: s}
}

func (h *handler) AddRegisteHandler(c *gin.Context) {
	var register models.RequestRegister

	if err := c.ShouldBindJSON(&register); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	err := h.s.AddRegisterService(register)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "OK", "message": "Register Successfully"})
}

func (h *handler) GetAllEmployeesHandler(c *gin.Context) {
	employees, err := h.s.GetAllEmployeesService()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": "Failed to get employees"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "GetAll Successfully", "employees": employees})
}

func (h *handler) GetCompanyByIDHandler(c *gin.Context) {
	cpnID := c.Param("cpn_id")
	company, err := h.s.GetCompanyByIDService(cpnID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Get company by ID successfully", "company": company})
}

func (h *handler) UpdateEmployeeHandler(c *gin.Context) {
	var employee models.RequestUpdateEmployee

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "Error", "message": err.Error()})
		return
	}

	err := h.s.UpdateEmployeeService(employee)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Employee updated successfully"})
}

func (h *handler) DeleteEmployeeHandler(c *gin.Context) {
	epyID := c.Param("epy_id")
	err := h.s.DeleteEmployeeService(epyID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK", "message": "Employee deleted successfully"})
}
