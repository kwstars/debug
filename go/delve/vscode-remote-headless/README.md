# 远程调试指南

## 场景：将构建好的二进制文件部署到远程服务器并进行调试

### 步骤 1：本地构建调试版本

```bash
# 构建带调试信息的二进制文件
make build-debug
```

这会生成 `bin/foo-debug` 文件，包含调试符号且禁用了优化。

### 步骤 2：传输文件到远程服务器

```bash
# 使用 scp 传输二进制文件到远程服务器
scp bin/foo-debug user@remote-server:/path/to/remote/project/

# 或者传输整个项目目录（如果需要源码对应）
rsync -av --exclude='.git' ./ user@remote-server:/path/to/remote/project/
```

### 步骤 3：在远程服务器上启动调试服务器

登录到远程服务器后，使用以下命令之一：

#### 选项 A：程序自动开始运行（推荐）

```bash
# 程序会立即开始执行，适合调试运行中的程序
dlv exec ./foo-debug --headless --listen=0.0.0.0:12345 --api-version 2 --accept-multiclient --continue --log --log-output=debugger,rpc,dap --log-dest=./dlv.log -- arg1
```

#### 选项 B：程序启动时暂停

```bash
# 程序会在 main 函数入口暂停，等待调试器连接
dlv exec ./foo-debug --headless --listen=0.0.0.0:12345 --api-version 2 --accept-multiclient --log --log-output=debugger,rpc,dap --log-dest=./dlv.log -- arg1
```

#### 选项 C：使用 Makefile（如果传输了完整项目）

```bash
# 自动运行模式
make remote-debug-running

# 暂停模式
make remote-debug-paused
```

### 步骤 4：配置 VSCode

1. 打开 `.vscode/launch.json` 文件
2. 修改以下配置项：
   - `host`: 将 `"your-remote-server-ip"` 替换为实际的远程服务器 IP
   - `substitutePath`: 设置正确的路径映射

```json
{
  "name": "远程调试 - 连接到运行中的程序",
  "host": "192.168.1.100", // 替换为实际IP
  "substitutePath": [
    {
      "from": "${workspaceFolder}",
      "to": "/home/user/myproject" // 远程服务器上的实际路径
    }
  ]
}
```

### 步骤 5：开始调试

1. 在 VSCode 中按 `F5` 或进入调试面板
2. 选择 "远程调试 - 连接到运行中的程序" 配置
3. 点击绿色播放按钮开始调试

## 常用命令参考

### 查看调试服务器状态

```bash
# 查看 dlv 进程
ps aux | grep dlv

# 查看端口占用
netstat -tlnp | grep 12345

# 查看调试日志
tail -f dlv.log
```

### 停止调试服务器

```bash
# 停止所有 dlv 进程
pkill dlv

# 或者使用 Makefile
make clean
```

### 防火墙配置

如果远程服务器有防火墙，需要开放端口：

```bash
# Ubuntu/Debian
sudo ufw allow 12345

# CentOS/RHEL
sudo firewall-cmd --permanent --add-port=12345/tcp
sudo firewall-cmd --reload
```

## 故障排除

### 常见问题

1. **连接超时**

   - 检查远程服务器防火墙设置
   - 确认 dlv 服务器正在运行
   - 验证网络连通性：`telnet remote-ip 12345`

2. **源码路径不匹配**

   - 检查 `substitutePath` 配置
   - 确保本地和远程的源码版本一致

3. **版本兼容性问题**

   - 确保使用 Delve v1.7.3 或更高版本
   - 检查 Go 版本兼容性

4. **权限问题**
   - 确保二进制文件有执行权限：`chmod +x foo-debug`
   - 检查端口绑定权限

### 调试日志

查看详细的调试日志：

```bash
tail -f dlv.log
```

## 注意事项

1. **安全性**：在生产环境中，建议使用 VPN 或 SSH 隧道而不是直接暴露调试端口
2. **性能**：调试版本会影响程序性能，仅在调试时使用
3. **端口冲突**：确保端口 12345 未被其他服务占用
4. **源码同步**：保持本地和远程的源码版本一致，以确保断点和变量显示正确
