# Bug 是什么

巡检上传请求 context 在 Request、Gateway、Worker 和 Service 四层被替换或忽略，取消无法生效且请求值丢失。

# 如何触发

执行 `go test ./internal/inspectioncontext -run '^TestUploadContextCancellationIsRequestScoped$' -count=1`。

# 错误信息

测试报告 request discarded cancellation、gateway did not stop、worker retried cancelled upload、service changed context。
