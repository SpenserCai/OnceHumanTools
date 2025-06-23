package services

import (
	"github.com/SpenserCai/OnceHumanTools/backend/internal/services"
)

// StrengthenProbabilityResult 强化概率计算结果
type StrengthenProbabilityResult = services.StrengthenProbabilityResult

// StrengthenPath 强化路径
type StrengthenPath = services.StrengthenPath

// StrengthenStep 强化步骤
type StrengthenStep = services.StrengthenStep

// NewStrengthenProbabilityService 创建强化概率服务
func NewStrengthenProbabilityService() *services.StrengthenProbabilityService {
	return services.NewStrengthenProbabilityService()
}
