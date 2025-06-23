#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
七日世界模组词条概率计算使用示例
"""

from mod_affix_probability import ModAffixProbabilityCalculator


def example_1():
    """示例1: 你提到的例子"""
    print("示例1: N=3，目标范围=[1,4,5,6]")
    print("-" * 40)
    
    calculator = ModAffixProbabilityCalculator()
    
    n_slots = 3
    target_range = [1, 4, 5, 6]
    
    result = calculator.calculate_detailed_probability(n_slots, target_range, True)
    calculator.print_result(result, True)


def example_2():
    """示例2: 4个词条，全伤害类型"""
    print("\n\n示例2: N=4，全伤害类型词条")
    print("-" * 40)
    
    calculator = ModAffixProbabilityCalculator()
    
    n_slots = 4
    # 伤害类词条：1.异常伤害, 4.对普通敌人伤害, 5.对精英敌人伤害, 6.对上位者伤害
    target_range = [1, 4, 5, 6]
    
    result = calculator.calculate_detailed_probability(n_slots, target_range, True)
    calculator.print_result(result, True)


def example_3():
    """示例3: 2个词条，防御类型"""
    print("\n\n示例3: N=2，防御类型词条")
    print("-" * 40)
    
    calculator = ModAffixProbabilityCalculator()
    
    n_slots = 2
    # 防御类词条：7.最大生命值, 8.头部受伤减免, 9.枪械伤害减免, 10.异常伤害减免
    target_range = [7, 8, 9, 10]
    
    result = calculator.calculate_detailed_probability(n_slots, target_range, True)
    calculator.print_result(result, True)


def example_4():
    """示例4: 对比不同词条数量的概率"""
    print("\n\n示例4: 对比不同词条数量（目标范围=[1,2,3,4,5]）")
    print("-" * 40)
    
    calculator = ModAffixProbabilityCalculator()
    
    target_range = [1, 2, 3, 4, 5]
    
    for n_slots in range(1, 6):
        result = calculator.calculate_probability(n_slots, target_range)
        print(f"N={n_slots}: 概率={result['probability_percent']:.2f}% ({result['valid_combinations']}/{result['total_combinations']})")


def example_5():
    """示例5: 极端情况测试"""
    print("\n\n示例5: 极端情况测试")
    print("-" * 40)
    
    calculator = ModAffixProbabilityCalculator()
    
    # 情况1: 只有1个词条在范围内
    print("情况1: N=3，目标范围=[1] (只有1个词条)")
    result1 = calculator.calculate_probability(3, [1])
    print(f"概率: {result1['probability_percent']:.2f}%")
    
    # 情况2: 所有词条都在范围内
    print("\n情况2: N=3，目标范围=[1,2,3,4,5,6,7,8,9,10] (所有词条)")
    result2 = calculator.calculate_probability(3, list(range(1, 11)))
    print(f"概率: {result2['probability_percent']:.2f}%")
    
    # 情况3: 需要的词条数量等于范围大小
    print("\n情况3: N=4，目标范围=[1,2,3,4] (范围大小等于需要数量)")
    result3 = calculator.calculate_probability(4, [1, 2, 3, 4])
    print(f"概率: {result3['probability_percent']:.2f}%")


def comparison_analysis():
    """对比分析：不同目标范围对概率的影响"""
    print("\n\n对比分析：不同目标范围对概率的影响 (N=3)")
    print("=" * 60)
    
    calculator = ModAffixProbabilityCalculator()
    
    test_cases = [
        ([1, 2, 3], "前3个词条"),
        ([1, 4, 5, 6], "你的例子"),
        ([1, 2, 3, 4, 5], "前5个词条"),
        ([6, 7, 8, 9, 10], "后5个词条"),
        (list(range(1, 8)), "前7个词条"),
        (list(range(1, 11)), "所有词条"),
    ]
    
    for target_range, description in test_cases:
        result = calculator.calculate_probability(3, target_range)
        range_names = [calculator.affixes[i] for i in target_range]
        print(f"{description} {target_range}:")
        print(f"  词条: {range_names}")
        print(f"  概率: {result['probability_percent']:.2f}% ({result['valid_combinations']}/{result['total_combinations']})")
        print()


if __name__ == "__main__":
    print("七日世界模组词条概率计算工具 - 使用示例")
    print("=" * 60)
    
    # 运行所有示例
    example_1()
    example_2() 
    example_3()
    example_4()
    example_5()
    comparison_analysis()
    
    print("\n" + "=" * 60)
    print("所有示例运行完成！")
    print("你可以运行 'python3 mod_affix_probability.py' 进行交互式计算") 