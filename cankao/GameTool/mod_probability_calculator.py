#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
七日世界模组词条强化概率计算工具
"""

import itertools
from typing import List, Tuple, Dict
from collections import defaultdict


class ModProbabilityCalculator:
    def __init__(self):
        self.max_level = 5
        self.max_enhancements = 5
        self.total_outcomes = 0
        self.successful_outcomes = 0
        self.all_paths = []
    
    def reset(self):
        """重置计算器状态"""
        self.total_outcomes = 0
        self.successful_outcomes = 0
        self.all_paths = []
    
    def is_valid_levels(self, levels: List[int]) -> bool:
        """验证词条等级是否有效"""
        if len(levels) != 4:
            return False
        return all(1 <= level <= self.max_level for level in levels)
    
    def get_available_slots(self, levels: List[int]) -> List[int]:
        """获取可以强化的词条索引（未满级的词条）"""
        return [i for i, level in enumerate(levels) if level < self.max_level]
    
    def _check_success(self, current_levels: List[int], target_levels: List[int], order_independent: bool) -> bool:
        """
        检查是否达到目标
        
        Args:
            current_levels: 当前词条等级
            target_levels: 目标词条等级
            order_independent: 是否顺序无关
            
        Returns:
            是否成功
        """
        if order_independent:
            # 顺序无关：对两个列表排序后比较
            sorted_current = sorted(current_levels, reverse=True)
            sorted_target = sorted(target_levels, reverse=True)
            return all(sorted_current[i] >= sorted_target[i] for i in range(4))
        else:
            # 位置对应：按位置严格比较
            return all(current_levels[i] >= target_levels[i] for i in range(4))
    
    def calculate_probability(self, initial_levels: List[int], target_levels: List[int], 
                            show_all_paths: bool = False, order_independent: bool = True) -> Dict:
        """
        计算达到目标等级的概率
        
        Args:
            initial_levels: 初始4个词条等级
            target_levels: 目标4个词条等级
            show_all_paths: 是否显示所有路径
            order_independent: 是否顺序无关（True: 只要总体等级达到即可，False: 位置必须对应）
            
        Returns:
            包含概率和路径信息的字典
        """
        self.reset()
        
        # 验证输入
        if not self.is_valid_levels(initial_levels):
            raise ValueError("初始等级必须是4个1-5之间的整数")
        if not self.is_valid_levels(target_levels):
            raise ValueError("目标等级必须是4个1-5之间的整数")
        
        # 检查目标是否可达
        for i in range(4):
            if target_levels[i] < initial_levels[i]:
                raise ValueError(f"目标等级不能低于初始等级：词条{i+1}")
        
        # 递归计算所有可能的强化路径
        self._calculate_recursive(initial_levels, target_levels, 0, [], show_all_paths, order_independent)
        
        # 计算概率
        probability = self.successful_outcomes / self.total_outcomes if self.total_outcomes > 0 else 0
        
        result = {
            'probability': probability,
            'probability_percent': probability * 100,
            'successful_outcomes': self.successful_outcomes,
            'total_outcomes': self.total_outcomes,
            'paths': self.all_paths if show_all_paths else []
        }
        
        return result
    
    def _calculate_recursive(self, current_levels: List[int], target_levels: List[int], 
                           enhancement_count: int, path: List[Tuple], show_all_paths: bool, order_independent: bool):
        """
        递归计算所有可能的强化路径
        
        Args:
            current_levels: 当前词条等级
            target_levels: 目标词条等级
            enhancement_count: 已使用的强化次数
            path: 当前路径（用于记录强化历史）
            show_all_paths: 是否记录所有路径
            order_independent: 是否顺序无关
        """
        # 如果强化次数用完
        if enhancement_count >= self.max_enhancements:
            self.total_outcomes += 1
            
            # 检查是否达到目标
            is_success = self._check_success(current_levels, target_levels, order_independent)
            if is_success:
                self.successful_outcomes += 1
                if show_all_paths:
                    self.all_paths.append({
                        'path': path.copy(),
                        'final_levels': current_levels.copy(),
                        'success': True
                    })
            elif show_all_paths:
                self.all_paths.append({
                    'path': path.copy(),
                    'final_levels': current_levels.copy(),
                    'success': False
                })
            return
        
        # 获取可以强化的词条
        available_slots = self.get_available_slots(current_levels)
        
        # 如果没有可强化的词条，结束
        if not available_slots:
            # 使用剩余的强化次数，但不会改变结果
            remaining_enhancements = self.max_enhancements - enhancement_count
            for _ in range(remaining_enhancements):
                self.total_outcomes += 1
                is_success = self._check_success(current_levels, target_levels, order_independent)
                if is_success:
                    self.successful_outcomes += 1
            return
        
        # 对每个可能的强化选择进行递归
        for slot in available_slots:
            new_levels = current_levels.copy()
            new_levels[slot] += 1
            new_path = path + [(enhancement_count + 1, slot, new_levels[slot])]
            
            self._calculate_recursive(new_levels, target_levels, enhancement_count + 1, 
                                    new_path, show_all_paths, order_independent)
    
    def print_result(self, result: Dict, initial_levels: List[int], target_levels: List[int]):
        """打印计算结果"""
        print("\n" + "="*60)
        print("七日世界模组词条强化概率计算结果")
        print("="*60)
        
        print(f"初始等级: {initial_levels}")
        print(f"目标等级: {target_levels}")
        print(f"强化次数: {self.max_enhancements}")
        print("-"*60)
        
        print(f"成功次数: {result['successful_outcomes']}")
        print(f"总次数: {result['total_outcomes']}")
        print(f"成功概率: {result['probability']:.6f}")
        print(f"成功概率: {result['probability_percent']:.4f}%")
        
        if result['paths']:
            print(f"\n所有路径详情 (共{len(result['paths'])}条):")
            print("-"*60)
            
            successful_paths = [p for p in result['paths'] if p['success']]
            failed_paths = [p for p in result['paths'] if not p['success']]
            
            if successful_paths:
                print(f"\n成功路径 (共{len(successful_paths)}条):")
                for i, path_info in enumerate(successful_paths[:10], 1):  # 只显示前10条
                    print(f"路径{i}: ", end="")
                    for step, slot, level in path_info['path']:
                        print(f"第{step}次强化词条{slot+1}到{level}级 ", end="")
                    print(f"-> 最终: {path_info['final_levels']}")
                
                if len(successful_paths) > 10:
                    print(f"... 还有{len(successful_paths) - 10}条成功路径")
            
            if failed_paths and len(failed_paths) <= 20:  # 只在失败路径不太多时显示
                print(f"\n失败路径 (共{len(failed_paths)}条):")
                for i, path_info in enumerate(failed_paths[:5], 1):  # 只显示前5条
                    print(f"路径{i}: ", end="")
                    for step, slot, level in path_info['path']:
                        print(f"第{step}次强化词条{slot+1}到{level}级 ", end="")
                    print(f"-> 最终: {path_info['final_levels']}")
                
                if len(failed_paths) > 5:
                    print(f"... 还有{len(failed_paths) - 5}条失败路径")


def main():
    """主函数"""
    calculator = ModProbabilityCalculator()
    
    print("七日世界模组词条强化概率计算工具")
    print("="*60)
    
    try:
        # 输入初始等级
        print("请输入初始4个词条的等级 (1-5):")
        initial_input = input("格式: 等级1,等级2,等级3,等级4 (例如: 1,2,3,1): ").strip()
        initial_levels = [int(x.strip()) for x in initial_input.split(',')]
        
        # 输入目标等级
        print("\n请输入目标4个词条的等级 (1-5):")
        target_input = input("格式: 等级1,等级2,等级3,等级4 (例如: 3,4,5,2): ").strip()
        target_levels = [int(x.strip()) for x in target_input.split(',')]
        
        # 询问是否显示所有路径
        show_paths_input = input("\n是否显示所有强化路径？(y/n): ").strip().lower()
        show_all_paths = show_paths_input in ['y', 'yes', '是', '1']
        
        # 询问判断模式
        print("\n请选择判断模式:")
        print("1. 顺序无关模式（默认）：只要总体等级达到即可，不考虑位置")
        print("2. 位置对应模式：词条位置必须严格对应")
        mode_input = input("请选择 (1/2): ").strip()
        order_independent = mode_input != '2'
        
        # 计算概率
        print(f"\n正在计算...（使用{'顺序无关' if order_independent else '位置对应'}模式）")
        result = calculator.calculate_probability(initial_levels, target_levels, show_all_paths, order_independent)
        
        # 打印结果
        calculator.print_result(result, initial_levels, target_levels)
        
    except ValueError as e:
        print(f"输入错误: {e}")
    except KeyboardInterrupt:
        print("\n计算被中断")
    except Exception as e:
        print(f"发生错误: {e}")


if __name__ == "__main__":
    main() 