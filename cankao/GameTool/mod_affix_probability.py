#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
七日世界模组词条出现概率计算工具
"""

import math
from typing import List, Set, Dict
from itertools import combinations


class ModAffixProbabilityCalculator:
    def __init__(self):
        # 定义所有可能的词条
        self.affixes = {
            1: "异常伤害",
            2: "弹匣容量",
            3: "换弹速度加成",
            4: "对普通敌人伤害",
            5: "对精英敌人伤害",
            6: "对上位者伤害",
            7: "最大生命值",
            8: "头部受伤减免",
            9: "枪械伤害减免",
            10: "异常伤害减免"
        }
        self.total_affixes = len(self.affixes)
    
    def calculate_combination(self, n: int, r: int) -> int:
        """
        计算组合数 C(n, r)
        
        Args:
            n: 总数量
            r: 选择数量
            
        Returns:
            组合数
        """
        if r > n or r < 0:
            return 0
        return math.comb(n, r)
    
    def calculate_probability(self, n_slots: int, target_range: List[int]) -> Dict:
        """
        计算指定词条范围内的概率
        
        Args:
            n_slots: 参与随机的词条数量
            target_range: 目标词条范围（词条编号列表）
            
        Returns:
            包含概率信息的字典
        """
        # 验证输入
        if n_slots <= 0 or n_slots > self.total_affixes:
            raise ValueError(f"词条数量必须在1-{self.total_affixes}之间")
        
        if not target_range:
            raise ValueError("目标范围不能为空")
        
        # 验证目标范围中的词条编号
        valid_range = []
        for affix_id in target_range:
            if affix_id in self.affixes:
                valid_range.append(affix_id)
            else:
                print(f"警告：词条编号{affix_id}无效，已忽略")
        
        if not valid_range:
            raise ValueError("目标范围中没有有效的词条编号")
        
        # 去重
        valid_range = list(set(valid_range))
        range_size = len(valid_range)
        
        if n_slots > range_size:
            # 如果需要的词条数量大于范围大小，概率为0
            return {
                'probability': 0.0,
                'probability_percent': 0.0,
                'total_combinations': self.calculate_combination(self.total_affixes, n_slots),
                'valid_combinations': 0,
                'n_slots': n_slots,
                'target_range': valid_range,
                'range_size': range_size,
                'all_combinations': []
            }
        
        # 计算总的可能组合数
        total_combinations = self.calculate_combination(self.total_affixes, n_slots)
        
        # 计算满足条件的组合数（从目标范围中选择n_slots个）
        valid_combinations = self.calculate_combination(range_size, n_slots)
        
        # 计算概率
        probability = valid_combinations / total_combinations if total_combinations > 0 else 0
        
        return {
            'probability': probability,
            'probability_percent': probability * 100,
            'total_combinations': total_combinations,
            'valid_combinations': valid_combinations,
            'n_slots': n_slots,
            'target_range': valid_range,
            'range_size': range_size,
            'all_combinations': []  # 这里可以添加具体的组合列表
        }
    
    def get_all_combinations(self, n_slots: int, target_range: List[int] = None) -> List[List[int]]:
        """
        获取所有可能的组合
        
        Args:
            n_slots: 词条数量
            target_range: 目标范围，如果为None则使用所有词条
            
        Returns:
            所有组合的列表
        """
        if target_range is None:
            source_range = list(self.affixes.keys())
        else:
            source_range = [x for x in target_range if x in self.affixes]
        
        return [list(combo) for combo in combinations(source_range, n_slots)]
    
    def calculate_detailed_probability(self, n_slots: int, target_range: List[int], 
                                     show_combinations: bool = False) -> Dict:
        """
        计算详细的概率信息，包括具体组合
        
        Args:
            n_slots: 参与随机的词条数量
            target_range: 目标词条范围
            show_combinations: 是否显示具体组合
            
        Returns:
            详细的概率信息
        """
        result = self.calculate_probability(n_slots, target_range)
        
        if show_combinations:
            # 获取所有总组合
            all_combos = self.get_all_combinations(n_slots)
            result['all_combinations'] = all_combos
            
            # 获取满足条件的组合
            valid_combos = self.get_all_combinations(n_slots, target_range)
            result['valid_combinations_list'] = valid_combos
        
        return result
    
    def print_result(self, result: Dict, show_combinations: bool = False):
        """
        打印计算结果
        
        Args:
            result: 计算结果字典
            show_combinations: 是否显示具体组合
        """
        print("\n" + "="*60)
        print("七日世界模组词条出现概率计算结果")
        print("="*60)
        
        print(f"词条数量: {result['n_slots']}")
        print(f"目标范围: {result['target_range']}")
        print(f"目标范围词条: {[self.affixes[i] for i in result['target_range']]}")
        print(f"目标范围大小: {result['range_size']}")
        print("-"*60)
        
        print(f"满足条件的组合数: {result['valid_combinations']}")
        print(f"总组合数: {result['total_combinations']}")
        print(f"概率: {result['probability']:.6f}")
        print(f"概率: {result['probability_percent']:.4f}%")
        
        if show_combinations and 'valid_combinations_list' in result:
            print(f"\n满足条件的组合详情 (共{len(result['valid_combinations_list'])}种):")
            print("-"*60)
            
            for i, combo in enumerate(result['valid_combinations_list'], 1):
                combo_names = [f"{affix_id}.{self.affixes[affix_id]}" for affix_id in combo]
                print(f"组合{i}: {combo} -> {combo_names}")
                
                if i >= 20:  # 只显示前20个
                    remaining = len(result['valid_combinations_list']) - 20
                    if remaining > 0:
                        print(f"... 还有{remaining}种组合")
                    break
    
    def print_all_affixes(self):
        """打印所有可用的词条"""
        print("\n所有可用词条:")
        print("-"*40)
        for affix_id, affix_name in self.affixes.items():
            print(f"{affix_id}. {affix_name}")


def main():
    """主函数"""
    calculator = ModAffixProbabilityCalculator()
    
    print("七日世界模组词条出现概率计算工具")
    print("="*60)
    
    # 显示所有词条
    calculator.print_all_affixes()
    
    try:
        # 输入词条数量
        print("\n请输入参与随机的词条数量:")
        n_slots = int(input("词条数量 (1-10): ").strip())
        
        # 输入目标范围
        print("\n请输入目标词条范围:")
        print("格式: 词条编号,词条编号,... (例如: 1,4,5,6)")
        range_input = input("目标范围: ").strip()
        target_range = [int(x.strip()) for x in range_input.split(',')]
        
        # 询问是否显示具体组合
        show_combos_input = input("\n是否显示具体组合？(y/n): ").strip().lower()
        show_combinations = show_combos_input in ['y', 'yes', '是', '1']
        
        # 计算概率
        print("\n正在计算...")
        result = calculator.calculate_detailed_probability(n_slots, target_range, show_combinations)
        
        # 打印结果
        calculator.print_result(result, show_combinations)
        
    except ValueError as e:
        print(f"输入错误: {e}")
    except KeyboardInterrupt:
        print("\n计算被中断")
    except Exception as e:
        print(f"发生错误: {e}")


if __name__ == "__main__":
    main() 