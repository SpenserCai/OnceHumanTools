package services

import (
	"sort"
)

// StrengthenProbabilityService 强化概率计算服务
type StrengthenProbabilityService struct{}

// NewStrengthenProbabilityService 创建强化概率计算服务
func NewStrengthenProbabilityService() *StrengthenProbabilityService {
	return &StrengthenProbabilityService{}
}

// CalculateProbability 计算强化成功概率
func (s *StrengthenProbabilityService) CalculateProbability(initialLevels, targetLevels []int, orderIndependent bool, showPaths bool) *StrengthenProbabilityResult {
	// 参数验证
	if len(initialLevels) != 4 || len(targetLevels) != 4 {
		return &StrengthenProbabilityResult{
			Error: "必须提供4个词条的等级",
		}
	}

	// 验证等级范围
	for i := 0; i < 4; i++ {
		if initialLevels[i] < 1 || initialLevels[i] > 5 {
			return &StrengthenProbabilityResult{
				Error: "初始等级必须在1-5之间",
			}
		}
		if targetLevels[i] < 1 || targetLevels[i] > 5 {
			return &StrengthenProbabilityResult{
				Error: "目标等级必须在1-5之间",
			}
		}
		if targetLevels[i] < initialLevels[i] {
			return &StrengthenProbabilityResult{
				Error: "目标等级不能低于初始等级",
			}
		}
	}

	calculator := &strengthenCalculator{
		maxLevel:        5,
		maxEnhancements: 5,
		orderIndependent: orderIndependent,
		showPaths:       showPaths,
	}

	return calculator.calculate(initialLevels, targetLevels)
}

// StrengthenProbabilityResult 强化概率计算结果
type StrengthenProbabilityResult struct {
	Probability         float64            `json:"probability"`
	ProbabilityPercent  float64            `json:"probabilityPercent"`
	SuccessfulOutcomes  int64              `json:"successfulOutcomes"`
	TotalOutcomes       int64              `json:"totalOutcomes"`
	Paths               []StrengthenPath   `json:"paths,omitempty"`
	Error               string             `json:"error,omitempty"`
}

// StrengthenPath 强化路径
type StrengthenPath struct {
	Success     bool             `json:"success"`
	FinalLevels []int            `json:"finalLevels"`
	Steps       []StrengthenStep `json:"steps"`
}

// StrengthenStep 强化步骤
type StrengthenStep struct {
	Step     int `json:"step"`
	Slot     int `json:"slot"`
	NewLevel int `json:"newLevel"`
}

// strengthenCalculator 强化计算器
type strengthenCalculator struct {
	maxLevel           int
	maxEnhancements    int
	orderIndependent   bool
	showPaths          bool
	totalOutcomes      int64
	successfulOutcomes int64
	paths              []StrengthenPath
}

func (c *strengthenCalculator) calculate(initialLevels, targetLevels []int) *StrengthenProbabilityResult {
	c.totalOutcomes = 0
	c.successfulOutcomes = 0
	c.paths = nil

	// 递归计算所有可能的强化路径
	c.calculateRecursive(copyIntSlice(initialLevels), targetLevels, 0, nil)

	probability := float64(c.successfulOutcomes) / float64(c.totalOutcomes)
	if c.totalOutcomes == 0 {
		probability = 0
	}

	return &StrengthenProbabilityResult{
		Probability:        probability,
		ProbabilityPercent: probability * 100,
		SuccessfulOutcomes: c.successfulOutcomes,
		TotalOutcomes:      c.totalOutcomes,
		Paths:              c.paths,
	}
}

func (c *strengthenCalculator) calculateRecursive(currentLevels, targetLevels []int, enhancementCount int, path []StrengthenStep) {
	// 如果强化次数用完
	if enhancementCount >= c.maxEnhancements {
		c.totalOutcomes++
		
		isSuccess := c.checkSuccess(currentLevels, targetLevels)
		if isSuccess {
			c.successfulOutcomes++
		}
		
		if c.showPaths && len(c.paths) < 100 { // 限制路径数量
			c.paths = append(c.paths, StrengthenPath{
				Success:     isSuccess,
				FinalLevels: copyIntSlice(currentLevels),
				Steps:       copySteps(path),
			})
		}
		return
	}

	// 获取可以强化的词条索引
	availableSlots := c.getAvailableSlots(currentLevels)
	
	// 如果没有可强化的词条
	if len(availableSlots) == 0 {
		// 直接结束，不再继续强化
		c.totalOutcomes++
		isSuccess := c.checkSuccess(currentLevels, targetLevels)
		if isSuccess {
			c.successfulOutcomes++
		}
		return
	}

	// 对每个可能的强化选择进行递归
	for _, slot := range availableSlots {
		newLevels := copyIntSlice(currentLevels)
		newLevels[slot]++
		
		newPath := append(path, StrengthenStep{
			Step:     enhancementCount + 1,
			Slot:     slot,
			NewLevel: newLevels[slot],
		})
		
		c.calculateRecursive(newLevels, targetLevels, enhancementCount+1, newPath)
	}
}

func (c *strengthenCalculator) getAvailableSlots(levels []int) []int {
	var slots []int
	for i, level := range levels {
		if level < c.maxLevel {
			slots = append(slots, i)
		}
	}
	return slots
}

func (c *strengthenCalculator) checkSuccess(currentLevels, targetLevels []int) bool {
	if c.orderIndependent {
		// 顺序无关：对两个列表排序后比较
		sortedCurrent := copyIntSlice(currentLevels)
		sortedTarget := copyIntSlice(targetLevels)
		sort.Sort(sort.Reverse(sort.IntSlice(sortedCurrent)))
		sort.Sort(sort.Reverse(sort.IntSlice(sortedTarget)))
		
		for i := 0; i < 4; i++ {
			if sortedCurrent[i] < sortedTarget[i] {
				return false
			}
		}
		return true
	} else {
		// 位置对应：按位置严格比较
		for i := 0; i < 4; i++ {
			if currentLevels[i] < targetLevels[i] {
				return false
			}
		}
		return true
	}
}

func copyIntSlice(slice []int) []int {
	result := make([]int, len(slice))
	copy(result, slice)
	return result
}

func copySteps(steps []StrengthenStep) []StrengthenStep {
	result := make([]StrengthenStep, len(steps))
	copy(result, steps)
	return result
}