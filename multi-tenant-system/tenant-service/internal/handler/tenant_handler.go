package handler

import (
	"fmt"
	"net/http"
	"tenant-service/internal/model"
	"tenant-service/internal/operator"

	"github.com/gin-gonic/gin"
)

type TenantHandler struct {
	Op *operator.TenantOperator
}

func NewTenantHandler(op *operator.TenantOperator) *TenantHandler {
	return &TenantHandler{Op: op}
}

func (h *TenantHandler) CreateTenant(c *gin.Context) {
	var tenant model.Tenant
	if err := c.ShouldBindJSON(&tenant); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("invalid payload - %v", err)})
		return
	}

	if err := h.Op.CreateTenant(tenant); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to create tenant - %v", err)})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Tenant created"})
}

func (h *TenantHandler) GetTenantByID(c *gin.Context) {
	tenantID := c.Param("id")
	tenant, err := h.Op.GetTenantByID(tenantID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to fetch tenant - %v", err)})
		return
	}

	c.JSON(http.StatusOK, tenant)
}

func (h *TenantHandler) GetAllTenants(c *gin.Context) {
	tenants, err := h.Op.GetAllTenants()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to fetch tenants - %v", err)})
		return
	}

	c.JSON(http.StatusOK, tenants)
}
