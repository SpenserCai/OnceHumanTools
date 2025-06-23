package services

import (
	"sort"
)

// AffixProbabilityService 词条概率计算服务
type AffixProbabilityService struct{}

// NewAffixProbabilityService 创建词条概率计算服务
func NewAffixProbabilityService() *AffixProbabilityService {
	return &AffixProbabilityService{}
}

// CalculateProbability 计算词条出现概率
func (s *AffixProbabilityService) CalculateProbability(slotCount int, targetAffixIDs []int, showCombinations bool) *AffixProbabilityResult {
	const totalAffixes = 10

	// 参数验证
	if slotCount <= 0 || slotCount > totalAffixes {
		return &AffixProbabilityResult{
			Error: "词条数量必须在1-10之间",
		}
	}

	// 去重目标词条
	targetSet := make(map[int]bool)
	for _, id := range targetAffixIDs {
		if id >= 1 && id <= totalAffixes {
			targetSet[id] = true
		}
	}

	rangeSize := len(targetSet)
	if rangeSize == 0 {
		return &AffixProbabilityResult{
			Error: "目标范围中没有有效的词条编号",
		}
	}

	// 如果需要的词条数量大于范围大小，概率为0
	if slotCount > rangeSize {
		return &AffixProbabilityResult{
			Probability:        0,
			ProbabilityPercent: 0,
			TotalCombinations:  combination(totalAffixes, slotCount),
			ValidCombinations:  0,
			SlotCount:          slotCount,
			TargetRange:        getSortedKeys(targetSet),
		}
	}

	// 计算总的可能组合数
	totalCombinations := combination(totalAffixes, slotCount)

	// 计算满足条件的组合数
	validCombinations := combination(rangeSize, slotCount)

	// 计算概率
	probability := float64(validCombinations) / float64(totalCombinations)

	result := &AffixProbabilityResult{
		Probability:        probability,
		ProbabilityPercent: probability * 100,
		TotalCombinations:  totalCombinations,
		ValidCombinations:  validCombinations,
		SlotCount:          slotCount,
		TargetRange:        getSortedKeys(targetSet),
	}

	// 如果需要显示组合
	if showCombinations && validCombinations > 0 && validCombinations <= 1000 {
		result.Combinations = generateCombinations(getSortedKeys(targetSet), slotCount)
	}

	return result
}

// AffixProbabilityResult 词条概率计算结果
type AffixProbabilityResult struct {
	Probability        float64 `json:"probability"`
	ProbabilityPercent float64 `json:"probabilityPercent"`
	TotalCombinations  int64   `json:"totalCombinations"`
	ValidCombinations  int64   `json:"validCombinations"`
	SlotCount          int     `json:"slotCount"`
	TargetRange        []int   `json:"targetRange"`
	Combinations       [][]int `json:"combinations,omitempty"`
	Error              string  `json:"error,omitempty"`
}

// combination 计算组合数 C(n,r)
func combination(n, r int) int64 {
	if r > n || r < 0 {
		return 0
	}
	if r == 0 || r == n {
		return 1
	}
	if r > n-r {
		r = n - r
	}

	result := int64(1)
	for i := 0; i < r; i++ {
		result = result * int64(n-i) / int64(i+1)
	}
	return result
}

// getSortedKeys 获取map的有序键
func getSortedKeys(m map[int]bool) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

// generateCombinations 生成所有组合
func generateCombinations(items []int, r int) [][]int {
	n := len(items)
	if r > n {
		return nil
	}

	var result [][]int
	combo := make([]int, r)

	var generate func(start, depth int)
	generate = func(start, depth int) {
		if depth == r {
			comboCopy := make([]int, r)
			copy(comboCopy, combo)
			result = append(result, comboCopy)
			return
		}

		for i := start; i <= n-(r-depth); i++ {
			combo[depth] = items[i]
			generate(i+1, depth+1)
		}
	}

	generate(0, 0)
	return result
}
