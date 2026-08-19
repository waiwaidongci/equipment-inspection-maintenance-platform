# Bug 是什么

维修事件缺失错误在仓储格式化时断链，分类、HTTP 状态和重试策略随后全部误判。

# 如何触发

执行 `go test ./internal/repairevents -run '^TestArchivedEventReturnsMissingWithoutRetry$' -count=1`。

# 错误信息

测试报告 repository broke missing chain、classifier returned system、handler returned 500、missing event was retried。
