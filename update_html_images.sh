#!/bin/bash

# 批量更新HTML文件中的图片引用为WebP格式

echo "开始更新HTML文件中的图片引用..."

# 定义要处理的HTML文件
html_files=(
    "templates/ServicePages/semiconductor_automation.html"
    "templates/ServicePages/medical_equipment_automation.html"
    "templates/ServicePages/chemical_automation.html"
    "templates/ServicePages/ServicePages3C.html"
    "templates/ServicePages/New_Energy_Services.html"
)

# 备份原文件
backup_dir="templates/backup_$(date +%Y%m%d_%H%M%S)"
mkdir -p "$backup_dir"

for file in "${html_files[@]}"; do
    if [ -f "$file" ]; then
        echo "处理文件: $file"
        # 备份原文件
        cp "$file" "$backup_dir/$(basename "$file")"
        
        # 替换PNG为WebP
        sed -i '' 's/\.png"/.webp"/g' "$file"
        sed -i '' 's/\.png"/.webp"/g' "$file"
        sed -i '' 's/\.png"/.webp"/g' "$file"
        
        # 为所有图片添加loading="lazy"属性（除了logo）
        sed -i '' 's/<img src="\.\.\/\.\.\/static\/picture\/[^l]/<img src="..\/..\/static\/picture\//g; s/\(<img [^>]*\) class="/\1 loading="lazy" class="/g' "$file"
        
        echo "  ✓ 已更新: $file"
    fi
done

echo ""
echo "=========================================="
echo "HTML文件更新完成！"
echo "=========================================="
echo "备份文件保存在: $backup_dir"
echo "=========================================="
