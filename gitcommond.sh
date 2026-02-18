#!/bin/bash

# Git 常用命令脚本

# 提交并推送代码
commit_and_push() {
    local message="$1"
    
    if [ -z "$message" ]; then
        echo "错误：请提供提交描述"
        echo "使用方法：./gitcommond.sh '提交描述'"
        exit 1
    fi
    
    echo "=========================================="
    echo "Git 提交和推送"
    echo "=========================================="
    echo "提交描述: ${message}"
    echo "=========================================="
    
    # 添加所有更改
    git add .
    
    # 提交
    git commit -m "${message}"
    
    if [ $? -ne 0 ]; then
        echo "提交失败"
        exit 1
    fi
    
    echo "✓ 提交成功"
    
    # 推送到远程仓库
    git push origin main
    
    if [ $? -ne 0 ]; then
        echo "推送失败"
        exit 1
    fi
    
    echo "✓ 推送成功"
    echo "=========================================="
}

# 查看状态
git_status() {
    echo "=========================================="
    echo "Git 状态"
    echo "=========================================="
    git status
    echo "=========================================="
}

# 查看日志
git_log() {
    local count="${1:-10}"
    echo "=========================================="
    echo "Git 日志 (最近 ${count} 条)"
    echo "=========================================="
    git log -n "${count}" --oneline --decorate
    echo "=========================================="
}

# 拉取最新代码
git_pull() {
    echo "=========================================="
    echo "拉取最新代码"
    echo "=========================================="
    git pull origin main
    echo "=========================================="
}

# 查看分支
git_branch() {
    echo "=========================================="
    echo "Git 分支"
    echo "=========================================="
    git branch -a
    echo "=========================================="
}

# 显示帮助信息
show_help() {
    echo "=========================================="
    echo "Git 常用命令脚本"
    echo "=========================================="
    echo "使用方法："
    echo "  ./gitcommond.sh commit '提交描述'    - 提交并推送代码"
    echo "  ./gitcommond.sh status              - 查看状态"
    echo "  ./gitcommond.sh log [数量]          - 查看日志（默认10条）"
    echo "  ./gitcommond.sh pull               - 拉取最新代码"
    echo "  ./gitcommond.sh branch             - 查看分支"
    echo "  ./gitcommond.sh help               - 显示帮助"
    echo "=========================================="
    echo "示例："
    echo "  ./gitcommond.sh commit '优化图片加载速度'"
    echo "  ./gitcommond.sh log 5"
    echo "=========================================="
}

# 主函数
main() {
    local command="$1"
    local param="$2"
    
    case "$command" in
        commit)
            commit_and_push "$param"
            ;;
        status)
            git_status
            ;;
        log)
            git_log "$param"
            ;;
        pull)
            git_pull
            ;;
        branch)
            git_branch
            ;;
        help|--help|-h)
            show_help
            ;;
        "")
            show_help
            ;;
        *)
            echo "错误：未知命令 '$command'"
            echo ""
            show_help
            exit 1
            ;;
    esac
}

# 执行主函数
main "$@"
