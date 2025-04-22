#!/bin/bash

echo "运行所有示例代码..."

# 查找并运行所有示例目录中的main.go文件
for example_dir in examples/*/; do
  if [ -f "${example_dir}main.go" ]; then
    echo "========================================"
    echo "运行示例: ${example_dir}"
    echo "========================================"
    go run ${example_dir}main.go
    
    # 检查运行结果
    if [ $? -ne 0 ]; then
      echo "❌ 示例运行失败: ${example_dir}"
      exit 1
    else
      echo "✅ 示例运行成功: ${example_dir}"
    fi
    echo ""
  fi
done

echo "所有示例运行完成！" 