# Bug
任务生产遇到坏数据不关闭通道，工作计数注册过晚，消费者不响应取消且错误通道没有缓冲。

# 触发
在项目根目录运行 `go test ./internal/taskpipeline -run TestPipelineLifecycleContracts`。

# 错误信息
`producer did not close output`、`wait returned before worker registration`、`consumer ignored cancellation`
