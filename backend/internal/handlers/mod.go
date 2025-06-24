package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	internalModels "github.com/SpenserCai/OnceHumanTools/backend/internal/models"
	"github.com/SpenserCai/OnceHumanTools/backend/internal/services"
	"github.com/SpenserCai/OnceHumanTools/backend/models"
	"github.com/SpenserCai/OnceHumanTools/backend/restapi/operations/mod"
)

// ModHandler 模组处理器
type ModHandler struct {
	affixService      *services.AffixProbabilityService
	strengthenService *services.StrengthenProbabilityService
}

// NewModHandler 创建模组处理器
func NewModHandler() *ModHandler {
	return &ModHandler{
		affixService:      services.NewAffixProbabilityService(),
		strengthenService: services.NewStrengthenProbabilityService(),
	}
}

// ListAffixes 获取词条列表
func (h *ModHandler) ListAffixes(params mod.ListAffixesParams) middleware.Responder {
	affixes := internalModels.GetAllAffixes()

	// 转换为API模型
	affixList := make([]*models.Affix, 0, len(affixes))
	for _, affix := range affixes {
		id := int32(affix.ID)
		name := affix.Name
		category := affix.Category
		affixList = append(affixList, &models.Affix{
			ID:          &id,
			Name:        &name,
			Description: affix.Description,
			Category:    category,
		})
	}

	total := int32(len(affixList))
	response := &models.AffixListResponse{
		Affixes: affixList,
		Total:   total,
	}

	return mod.NewListAffixesOK().WithPayload(response)
}

// CalculateAffixProbability 计算词条概率
func (h *ModHandler) CalculateAffixProbability(params mod.CalculateAffixProbabilityParams) middleware.Responder {
	// 转换参数
	slotCount := int(*params.Body.SlotCount)
	targetAffixIDs := make([]int, len(params.Body.TargetAffixIds))
	for i, id := range params.Body.TargetAffixIds {
		targetAffixIDs[i] = int(id)
	}

	showCombinations := false
	if params.Body.ShowCombinations != nil {
		showCombinations = *params.Body.ShowCombinations
	}

	// 调用服务计算
	result := h.affixService.CalculateProbability(slotCount, targetAffixIDs, showCombinations)

	// 检查错误
	if result.Error != "" {
		errorMsg := result.Error
		error := "bad_request"
		return mod.NewCalculateAffixProbabilityBadRequest().WithPayload(&models.ErrorResponse{
			Error:   &error,
			Message: &errorMsg,
		})
	}

	// 转换结果
	slotCount32 := int32(result.SlotCount)
	targetRange := make([]int32, len(result.TargetRange))
	for i, id := range result.TargetRange {
		targetRange[i] = int32(id)
	}

	response := &models.AffixProbabilityResponse{
		Probability:        &result.Probability,
		ProbabilityPercent: &result.ProbabilityPercent,
		TotalCombinations:  &result.TotalCombinations,
		ValidCombinations:  &result.ValidCombinations,
		SlotCount:          slotCount32,
		TargetRange:        targetRange,
	}

	// 添加组合数据
	if showCombinations && len(result.Combinations) > 0 {
		combinations := make([][]int32, len(result.Combinations))
		for i, combo := range result.Combinations {
			combinations[i] = make([]int32, len(combo))
			for j, id := range combo {
				combinations[i][j] = int32(id)
			}
		}
		response.Combinations = combinations
	}

	return mod.NewCalculateAffixProbabilityOK().WithPayload(response)
}

// CalculateStrengthenProbability 计算强化概率
func (h *ModHandler) CalculateStrengthenProbability(params mod.CalculateStrengthenProbabilityParams) middleware.Responder {
	// 转换参数
	initialLevels := make([]int, len(params.Body.InitialLevels))
	for i, level := range params.Body.InitialLevels {
		initialLevels[i] = int(level)
	}

	targetLevels := make([]int, len(params.Body.TargetLevels))
	for i, level := range params.Body.TargetLevels {
		targetLevels[i] = int(level)
	}

	orderIndependent := true
	if params.Body.OrderIndependent != nil {
		orderIndependent = *params.Body.OrderIndependent
	}

	showPaths := false
	if params.Body.ShowPaths != nil {
		showPaths = *params.Body.ShowPaths
	}

	// 调用服务计算
	result := h.strengthenService.CalculateProbability(initialLevels, targetLevels, orderIndependent, showPaths)

	// 检查错误
	if result.Error != "" {
		errorMsg := result.Error
		error := "bad_request"
		return mod.NewCalculateStrengthenProbabilityBadRequest().WithPayload(&models.ErrorResponse{
			Error:   &error,
			Message: &errorMsg,
		})
	}

	// 转换结果
	response := &models.StrengthenProbabilityResponse{
		Probability:        &result.Probability,
		ProbabilityPercent: &result.ProbabilityPercent,
		SuccessfulOutcomes: &result.SuccessfulOutcomes,
		TotalOutcomes:      &result.TotalOutcomes,
	}

	// 添加路径数据
	if showPaths && len(result.Paths) > 0 {
		paths := make([]*models.StrengthenPath, 0, len(result.Paths))
		for _, path := range result.Paths {
			// 转换等级
			finalLevels := make([]int32, len(path.FinalLevels))
			for i, level := range path.FinalLevels {
				finalLevels[i] = int32(level)
			}

			// 转换步骤
			steps := make([]*models.StrengthenStep, 0, len(path.Steps))
			for _, step := range path.Steps {
				step32 := int32(step.Step)
				slot32 := int32(step.Slot)
				level32 := int32(step.NewLevel)
				steps = append(steps, &models.StrengthenStep{
					Step:     step32,
					Slot:     slot32,
					NewLevel: level32,
				})
			}

			paths = append(paths, &models.StrengthenPath{
				Success:     path.Success,
				FinalLevels: finalLevels,
				Steps:       steps,
			})
		}
		response.Paths = paths
	}

	return mod.NewCalculateStrengthenProbabilityOK().WithPayload(response)
}
