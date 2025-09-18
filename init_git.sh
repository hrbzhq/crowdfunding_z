#!/bin/bash

# 设置邮箱和密钥路径
EMAIL="hrbzhq@163.com"
KEY_PATH="$HOME/.ssh/id_ed25519"

# 1. 生成 SSH 密钥
ssh-keygen -t ed25519 -C "$EMAIL" -f "$KEY_PATH" -N ""

# 2. 显示公钥（用于复制到 GitHub）
echo "🔑 请将以下公钥添加到 GitHub SSH 设置页面：https://github.com/settings/keys"
cat "${KEY_PATH}.pub"

# 3. 测试 SSH 连接
ssh -T git@github.com

# 4. 初始化 Git 仓库（如果尚未初始化）
if [ ! -d .git ]; then
  git init
  git remote add origin git@github.com:hrbzhq/crowdfunding_z.git
  git add .
  git commit -m "Initial commit"
fi

# 5. 推送到 GitHub
git push -u origin master
