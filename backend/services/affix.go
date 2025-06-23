package services

import (
	"github.com/SpenserCai/OnceHumanTools/backend/internal/services"
)

// AffixProbabilityResult 词条概率计算结果
type AffixProbabilityResult = services.AffixProbabilityResult

// NewAffixProbabilityService 创建词条概率服务
func NewAffixProbabilityService() *services.AffixProbabilityService {
	return services.NewAffixProbabilityService()
}
