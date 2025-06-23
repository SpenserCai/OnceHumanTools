package handlers

import (
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"

	"github.com/SpenserCai/OnceHumanTools/backend/models"
	"github.com/SpenserCai/OnceHumanTools/backend/restapi/operations/system"
)

// SystemHandler 系统处理器
type SystemHandler struct{}

// NewSystemHandler 创建系统处理器
func NewSystemHandler() *SystemHandler {
	return &SystemHandler{}
}

// HealthCheck 健康检查
func (h *SystemHandler) HealthCheck(params system.HealthCheckParams) middleware.Responder {
	status := "ok"
	version := "1.0.0"
	timestamp := strfmt.DateTime(time.Now())

	response := &models.HealthResponse{
		Status:    &status,
		Timestamp: &timestamp,
		Version:   version,
	}

	return system.NewHealthCheckOK().WithPayload(response)
}
