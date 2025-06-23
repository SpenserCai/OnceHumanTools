package models

// Affix 词条
type Affix struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Category    string `json:"category,omitempty"`
}

// AffixCategory 词条分类
type AffixCategory string

const (
	AffixCategoryDamage  AffixCategory = "damage"  // 伤害类
	AffixCategoryDefense AffixCategory = "defense" // 防御类
	AffixCategoryUtility AffixCategory = "utility" // 功能类
)

// GetAllAffixes 获取所有词条定义
func GetAllAffixes() []Affix {
	return []Affix{
		{ID: 1, Name: "异常伤害", Description: "提升异常状态伤害", Category: string(AffixCategoryDamage)},
		{ID: 2, Name: "弹匣容量", Description: "增加武器弹匣容量", Category: string(AffixCategoryUtility)},
		{ID: 3, Name: "换弹速度加成", Description: "提升换弹速度", Category: string(AffixCategoryUtility)},
		{ID: 4, Name: "对普通敌人伤害", Description: "对普通敌人造成额外伤害", Category: string(AffixCategoryDamage)},
		{ID: 5, Name: "对精英敌人伤害", Description: "对精英敌人造成额外伤害", Category: string(AffixCategoryDamage)},
		{ID: 6, Name: "对上位者伤害", Description: "对上位者敌人造成额外伤害", Category: string(AffixCategoryDamage)},
		{ID: 7, Name: "最大生命值", Description: "增加角色最大生命值", Category: string(AffixCategoryDefense)},
		{ID: 8, Name: "头部受伤减免", Description: "减少头部受到的伤害", Category: string(AffixCategoryDefense)},
		{ID: 9, Name: "枪械伤害减免", Description: "减少枪械造成的伤害", Category: string(AffixCategoryDefense)},
		{ID: 10, Name: "异常伤害减免", Description: "减少异常状态伤害", Category: string(AffixCategoryDefense)},
	}
}

// GetAffixByID 根据ID获取词条
func GetAffixByID(id int) *Affix {
	affixes := GetAllAffixes()
	for _, affix := range affixes {
		if affix.ID == id {
			return &affix
		}
	}
	return nil
}

// GetAffixesByCategory 根据分类获取词条
func GetAffixesByCategory(category AffixCategory) []Affix {
	var result []Affix
	affixes := GetAllAffixes()
	for _, affix := range affixes {
		if affix.Category == string(category) {
			result = append(result, affix)
		}
	}
	return result
}