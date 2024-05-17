# [Delve](https://github.com/go-delve/delve)

Delve 提供了几种主要的调试方式：

1. **[命令行调试](https://github.com/go-delve/delve/blob/master/Documentation/cli/README.md)**：在命令行中使用 `dlv debug` 命令启动 Delve 并加载你的程序。在此模式下，你可以使用各种命令来控制程序的执行，查看变量的值，设置断点等。
2. **API 服务器模式**：在此模式下，你可以使用任何支持 Delve API 的客户端来连接服务器并进行调试。
3. **[集成开发环境（IDE）调试](https://github.com/go-delve/delve/blob/master/Documentation/EditorIntegration.md)**：许多 IDE，如 [Visual Studio Code](https://github.com/golang/vscode-go/wiki/debugging)，GoLand 等，都集成了 Delve，可以在 IDE 中直接进行调试。

## [远程调式模式对比](https://github.com/golang/vscode-go/wiki/debugging#remote-debugging)

| 模式                                                               | 优点                                                                                                                                                                        | 缺点                                                                                                                                                                                                                 |
| ------------------------------------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Connect to headless delve with target specified at server start-up | 1. 支持多客户端连接，可以同时从多个客户端进行调试。<br>2. 支持 `--continue` 标志，允许在调试会话结束后继续运行程序。<br>3. 服务器启动时就指定了目标，简化了客户端的配置。   | 1. 服务器启动时就需要指定目标，这可能限制了一些灵活性。<br>2. 需要手动启动和管理服务器。                                                                                                                             |
| Connect to delve dap with target specified at client start-up      | 1. 提供了更大的灵活性，可以在客户端启动时指定目标。<br>2. 可以轻松地将本地配置适应到连接到外部服务器。<br>3. 不需要手动启动和管理服务器，服务器会在调试会话结束后自动退出。 | 1. 不支持多客户端连接。<br>2. 不支持 `--continue` 标志，调试会话结束后程序将停止运行。<br>3. 安全性问题，任何能够连接到服务器的人都可以让它运行任意程序。<br>4. 需要在客户端配置中指定目标，可能增加了配置的复杂性。 |
