package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/SpenserCai/OnceHumanTools/backend/models"
	"github.com/SpenserCai/OnceHumanTools/backend/restapi/operations/tools"
)

// ToolsHandler 工具处理器
type ToolsHandler struct{}

// NewToolsHandler 创建工具处理器
func NewToolsHandler() *ToolsHandler {
	return &ToolsHandler{}
}

// ListTools 获取工具列表
func (h *ToolsHandler) ListTools(params tools.ListToolsParams) middleware.Responder {
	// 定义可用工具
	toolList := []*models.Tool{
		{
			ID:          stringPtr("affix-probability"),
			Name:        stringPtr("模组词条概率计算器"),
			Description: "计算特定词条组合出现的概率",
			Category:    stringPtr("mod"),
			Icon:        "dice",
		},
		{
			ID:          stringPtr("strengthen-probability"),
			Name:        stringPtr("模组强化概率计算器"),
			Description: "计算模组词条强化到目标等级的概率",
			Category:    stringPtr("mod"),
			Icon:        "trending-up",
		},
	}

	// 获取分类
	categories := []string{"mod", "weapon", "character"}

	response := &models.ToolsListResponse{
		Tools:      toolList,
		Categories: categories,
	}

	return tools.NewListToolsOK().WithPayload(response)
}

// 辅助函数：创建字符串指针
func stringPtr(s string) *string {
	return &s
}
