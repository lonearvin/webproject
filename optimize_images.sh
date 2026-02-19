#!/bin/bash

# 图片优化脚本 - 将PNG/JPG转换为WebP格式以减小文件大小

echo "开始优化图片..."

# 定义要优化的目录
directories=(
    "static/picture/car"
    "static/picture/energy"
    "static/picture/phone"
    "static/picture/machine"
    "static/picture/semiconductor"
    "static/picture/medical"
    "static/picture/chemicalAutomation"
    "static/case/casePicture"
)

# WebP质量设置 (80-85通常在质量和大小之间有很好的平衡)
QUALITY=85

# 统计信息
total_files=0
optimized_files=0
total_original_size=0
total_optimized_size=0

# 遍历目录
for dir in "${directories[@]}"; do
    if [ -d "$dir" ]; then
        echo "处理目录: $dir"
        
        # 处理PNG文件
        for file in "$dir"/*.png; do
            if [ -f "$file" ]; then
                total_files=$((total_files + 1))
                
                # 获取原始文件大小
                original_size=$(stat -f%z "$file" 2>/dev/null || stat -c%s "$file" 2>/dev/null)
                total_original_size=$((total_original_size + original_size))
                
                # 生成WebP文件名
                webp_file="${file%.png}.webp"
                
                # 检查是否已存在WebP文件
                if [ -f "$webp_file" ]; then
                    echo "  ⊘ $(basename "$file") - WebP文件已存在，跳过"
                    optimized_files=$((optimized_files + 1))
                    optimized_size=$(stat -f%z "$webp_file" 2>/dev/null || stat -c%s "$webp_file" 2>/dev/null)
                    total_optimized_size=$((total_optimized_size + optimized_size))
                    continue
                fi
                
                # 转换为WebP
                cwebp -q $QUALITY "$file" -o "$webp_file" 2>/dev/null
                
                if [ -f "$webp_file" ]; then
                    # 获取优化后文件大小
                    optimized_size=$(stat -f%z "$webp_file" 2>/dev/null || stat -c%s "$webp_file" 2>/dev/null)
                    total_optimized_size=$((total_optimized_size + optimized_size))
                    optimized_files=$((optimized_files + 1))
                    
                    # 计算压缩率
                    reduction=$(( (original_size - optimized_size) * 100 / original_size ))
                    
                    echo "  ✓ $(basename "$file") -> $(basename "$webp_file") (减少 ${reduction}%)"
                fi
            fi
        done
        
        # 处理JPG文件
        for file in "$dir"/*.jpg "$dir"/*.jpeg; do
            if [ -f "$file" ]; then
                total_files=$((total_files + 1))
                
                # 获取原始文件大小
                original_size=$(stat -f%z "$file" 2>/dev/null || stat -c%s "$file" 2>/dev/null)
                total_original_size=$((total_original_size + original_size))
                
                # 生成WebP文件名
                webp_file="${file%.*}.webp"
                
                # 检查是否已存在WebP文件
                if [ -f "$webp_file" ]; then
                    echo "  ⊘ $(basename "$file") - WebP文件已存在，跳过"
                    optimized_files=$((optimized_files + 1))
                    optimized_size=$(stat -f%z "$webp_file" 2>/dev/null || stat -c%s "$webp_file" 2>/dev/null)
                    total_optimized_size=$((total_optimized_size + optimized_size))
                    continue
                fi
                
                # 转换为WebP
                cwebp -q $QUALITY "$file" -o "$webp_file" 2>/dev/null
                
                if [ -f "$webp_file" ]; then
                    # 获取优化后文件大小
                    optimized_size=$(stat -f%z "$webp_file" 2>/dev/null || stat -c%s "$webp_file" 2>/dev/null)
                    total_optimized_size=$((total_optimized_size + optimized_size))
                    optimized_files=$((optimized_files + 1))
                    
                    # 计算压缩率
                    reduction=$(( (original_size - optimized_size) * 100 / original_size ))
                    
                    echo "  ✓ $(basename "$file") -> $(basename "$webp_file") (减少 ${reduction}%)"
                fi
            fi
        done
    fi
done

# 处理根目录下的图片
echo "处理根目录图片..."
for file in static/picture/*.png static/picture/*.jpg static/picture/*.jpeg; do
    if [ -f "$file" ]; then
        total_files=$((total_files + 1))
        
        original_size=$(stat -f%z "$file" 2>/dev/null || stat -c%s "$file" 2>/dev/null)
        total_original_size=$((total_original_size + original_size))
        
        webp_file="${file%.*}.webp"
        
        # 检查是否已存在WebP文件
        if [ -f "$webp_file" ]; then
            echo "  ⊘ $(basename "$file") - WebP文件已存在，跳过"
            optimized_files=$((optimized_files + 1))
            optimized_size=$(stat -f%z "$webp_file" 2>/dev/null || stat -c%s "$webp_file" 2>/dev/null)
            total_optimized_size=$((total_optimized_size + optimized_size))
            continue
        fi
        
        cwebp -q $QUALITY "$file" -o "$webp_file" 2>/dev/null
        
        if [ -f "$webp_file" ]; then
            optimized_size=$(stat -f%z "$webp_file" 2>/dev/null || stat -c%s "$webp_file" 2>/dev/null)
            total_optimized_size=$((total_optimized_size + optimized_size))
            optimized_files=$((optimized_files + 1))
            
            reduction=$(( (original_size - optimized_size) * 100 / original_size ))
            echo "  ✓ $(basename "$file") -> $(basename "$webp_file") (减少 ${reduction}%)"
        fi
    fi
done

# 处理companyLogo目录
if [ -d "static/picture/companyLogo" ]; then
    echo "处理目录: static/picture/companyLogo"
    for file in static/picture/companyLogo/*.{png,jpg}; do
        if [ -f "$file" ]; then
            total_files=$((total_files + 1))
            
            original_size=$(stat -f%z "$file" 2>/dev/null || stat -c%s "$file" 2>/dev/null)
            total_original_size=$((total_original_size + original_size))
            
            webp_file="${file%.*}.webp"
            
            # 检查是否已存在WebP文件
            if [ -f "$webp_file" ]; then
                echo "  ⊘ $(basename "$file") - WebP文件已存在，跳过"
                optimized_files=$((optimized_files + 1))
                optimized_size=$(stat -f%z "$webp_file" 2>/dev/null || stat -c%s "$webp_file" 2>/dev/null)
                total_optimized_size=$((total_optimized_size + optimized_size))
                continue
            fi
            
            cwebp -q $QUALITY "$file" -o "$webp_file" 2>/dev/null
            
            if [ -f "$webp_file" ]; then
                optimized_size=$(stat -f%z "$webp_file" 2>/dev/null || stat -c%s "$webp_file" 2>/dev/null)
                total_optimized_size=$((total_optimized_size + optimized_size))
                optimized_files=$((optimized_files + 1))
                
                reduction=$(( (original_size - optimized_size) * 100 / original_size ))
                echo "  ✓ $(basename "$file") -> $(basename "$webp_file") (减少 ${reduction}%)"
            fi
        fi
    done
fi

# 输出统计信息
echo ""
echo "=========================================="
echo "图片优化完成！"
echo "=========================================="
echo "处理文件总数: $total_files"
echo "成功优化文件: $optimized_files"
echo "原始总大小: $(echo "scale=2; $total_original_size / 1024 / 1024" | bc) MB"
echo "优化后总大小: $(echo "scale=2; $total_optimized_size / 1024 / 1024" | bc) MB"
echo "节省空间: $(echo "scale=2; ($total_original_size - $total_optimized_size) / 1024 / 1024" | bc) MB"
echo "平均压缩率: $(echo "scale=2; ($total_original_size - $total_optimized_size) * 100 / $total_original_size" | bc)%"
echo "=========================================="
