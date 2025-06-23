# 远程二进制文件调试指南

## 前提条件

### 1. 编译带调试信息的二进制文件

二进制文件必须包含调试信息，编译时使用：

```bash
go build -gcflags="all=-N -l" -o foo-debug ./cmd/foo
```

### 2. 在远程服务器上启动 dlv

```bash
# 方法1: 直接调试二进制文件
dlv exec ./foo-debug --headless --listen=:12345 --api-version=2 --accept-multiclient -- test-arg

# 方法2: 附加到运行中的进程
./foo-debug test-arg &  # 后台启动程序
dlv attach $(pgrep foo-debug) --headless --listen=:12345 --api-version=2 --accept-multiclient
```

## substitutePath 配置说明

### 为什么需要 substitutePath？

- 本地源码路径：`/home/kira/workspace/debug/go/delve/vscode-remote-headless`
- 二进制文件编译时的路径：可能是 `/go/src/demo` 或其他路径
- 调试器需要知道如何将二进制文件中的路径映射到本地源码

### 如何确定正确的路径映射？

1. 查看二进制文件的调试信息：

```bash
go tool objdump -s "main.main" foo-debug | head -20
```

2. 或者查看符号表：

```bash
go tool nm foo-debug | grep main.main
```

3. 使用 dlv 直接查看源码路径：

```bash
dlv exec ./foo-debug
(dlv) sources
```

### 常见的路径映射情况

#### 情况 1: 使用 Go modules（推荐）

```json
"substitutePath": [
  {
    "from": "${workspaceFolder}",
    "to": "/go/src/demo"  // demo 是您的模块名
  }
]
```

#### 情况 2: 在容器中编译

```json
"substitutePath": [
  {
    "from": "${workspaceFolder}",
    "to": "/app"  // 容器中的工作目录
  }
]
```

#### 情况 3: 相对路径编译

```json
"substitutePath": [
  {
    "from": "${workspaceFolder}",
    "to": "/path/to/project"  // 编译时的完整路径
  }
]
```

## 调试步骤

1. **启动远程 dlv 服务器**
2. **修改配置中的 host 地址**
3. **根据实际情况调整 substitutePath**
4. **在 VS Code 中启动调试配置**
5. **设置断点并开始调试**

## 常见问题

### Q: 无法命中断点

A: 检查 substitutePath 配置是否正确，路径映射必须准确

### Q: 看不到源码

A: 确保二进制文件包含调试信息，且路径映射正确

### Q: 调试信息不准确

A: 确保本地源码版本与二进制文件版本一致

## 验证配置

启动调试后，在 Debug Console 中执行：

```
-exec sources
```

查看调试器能否正确找到源文件路径。
