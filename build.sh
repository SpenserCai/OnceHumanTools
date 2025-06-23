#!/bin/bash
###
 # @Author: SpenserCai
 # @Date: 2025-06-23 16:38:14
 # @version: 
 # @LastEditors: SpenserCai
 # @LastEditTime: 2025-06-23 16:53:32
 # @Description: file content
### 

# 快速构建脚本

# 设置 Go 环境变量
export GOPATH=$(go env GOPATH)
export PATH="$GOPATH/bin:$PATH"

echo "OnceHuman工具集 - 快速构建"
echo ""
echo "使用方法:"
echo "  ./build.sh              # 构建集成版本"
echo "  ./build.sh release      # 构建发布版本"
echo "  ./build.sh integrated   # 构建集成版本（默认）"
echo "  ./build.sh cross linux amd64  # 交叉编译"
echo ""

# 检查是否安装了必要的工具
command -v go >/dev/null 2>&1 || { echo "错误: 需要安装 Go"; exit 1; }
command -v npm >/dev/null 2>&1 || { echo "错误: 需要安装 Node.js/npm"; exit 1; }
command -v make >/dev/null 2>&1 || { echo "错误: 需要安装 make"; exit 1; }

# 使用 make 或 scripts/build.sh
if [ -f "Makefile" ]; then
    case "$1" in
        release)
            make build-release
            ;;
        cross)
            bash scripts/build.sh cross $2 $3
            ;;
        *)
            make build-integrated
            ;;
    esac
else
    bash scripts/build.sh $@
fi