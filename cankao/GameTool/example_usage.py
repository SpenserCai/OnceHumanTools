'''
Author: SpenserCai
Date: 2025-06-11 11:41:15
version: 
LastEditors: SpenserCai
LastEditTime: 2025-06-11 11:41:42
Description: file content
'''
#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
七日世界模组词条强化概率计算工具使用示例
"""

from mod_probability_calculator import ModProbabilityCalculator


def example_1():
    """示例1: 基本概率计算"""
    print("示例1: 基本概率计算")
    print("-" * 40)
    
    calculator = ModProbabilityCalculator()
    
    # 初始等级: [1, 1, 1, 1]，目标等级: [2, 2, 2, 2]
    initial_levels = [1, 1, 1, 1]
    target_levels = [2, 2, 2, 2]
    
    result = calculator.calculate_probability(initial_levels, target_levels, show_all_paths=False)
    calculator.print_result(result, initial_levels, target_levels)


def example_2():
    """示例2: 显示所有路径的概率计算"""
    print("\n\n示例2: 显示所有路径的概率计算")
    print("-" * 40)
    
    calculator = ModProbabilityCalculator()
    
    # 初始等级: [2, 3, 1, 2]，目标等级: [3, 4, 2, 3]
    initial_levels = [2, 3, 1, 2]
    target_levels = [3, 4, 2, 3]
    
    result = calculator.calculate_probability(initial_levels, target_levels, show_all_paths=True)
    calculator.print_result(result, initial_levels, target_levels)


def example_3():
    """示例3: 高难度目标"""
    print("\n\n示例3: 高难度目标")
    print("-" * 40)
    
    calculator = ModProbabilityCalculator()
    
    # 初始等级: [1, 1, 1, 1]，目标等级: [5, 5, 1, 1]
    initial_levels = [1, 1, 1, 1]
    target_levels = [5, 5, 1, 1]
    
    result = calculator.calculate_probability(initial_levels, target_levels, show_all_paths=False)
    calculator.print_result(result, initial_levels, target_levels)


def example_4():
    """示例4: 不可能的目标"""
    print("\n\n示例4: 不可能的目标")
    print("-" * 40)
    
    calculator = ModProbabilityCalculator()
    
    # 初始等级: [1, 1, 1, 1]，目标等级: [5, 5, 5, 5]
    initial_levels = [1, 1, 1, 1]
    target_levels = [5, 5, 5, 5]
    
    result = calculator.calculate_probability(initial_levels, target_levels, show_all_paths=False)
    calculator.print_result(result, initial_levels, target_levels)


def example_5():
    """示例5: 对比两种判断模式"""
    print("\n\n示例5: 对比两种判断模式")
    print("-" * 40)
    
    calculator = ModProbabilityCalculator()
    
    # 使用一个能看出区别的例子
    initial_levels = [1, 1, 1, 1]
    target_levels = [2, 2, 2, 1]
    
    print(f"初始等级: {initial_levels}")
    print(f"目标等级: {target_levels}")
    print()
    
    # 位置对应模式
    result1 = calculator.calculate_probability(initial_levels, target_levels, False, False)
    print(f"位置对应模式: {result1['probability_percent']:.2f}%")
    
    # 顺序无关模式  
    result2 = calculator.calculate_probability(initial_levels, target_levels, False, True)
    print(f"顺序无关模式: {result2['probability_percent']:.2f}%")
    
    print(f"概率提升: {result2['probability_percent'] - result1['probability_percent']:.2f}%")


def custom_calculation():
    """自定义计算"""
    print("\n\n自定义计算")
    print("-" * 40)
    
    calculator = ModProbabilityCalculator()
    
    # 你可以在这里修改初始等级和目标等级
    initial_levels = [3, 2, 4, 1]  # 修改这里
    target_levels = [4, 4, 5, 3]   # 修改这里
    show_all_paths = True          # 是否显示所有路径
    order_independent = True       # 是否顺序无关
    
    result = calculator.calculate_probability(initial_levels, target_levels, show_all_paths, order_independent)
    calculator.print_result(result, initial_levels, target_levels)


if __name__ == "__main__":
    print("七日世界模组词条强化概率计算工具 - 使用示例")
    print("=" * 60)
    
    # 运行所有示例
    example_1()
    example_2()
    example_3()
    example_4()
    example_5()
    custom_calculation()
    
    print("\n" + "=" * 60)
    print("所有示例运行完成！")
    print("你可以运行 'python3 mod_probability_calculator.py' 进行交互式计算") 