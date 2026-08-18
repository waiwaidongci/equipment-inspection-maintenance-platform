# Bug
维修批处理在循环末尾才释放租约，提交结果覆盖校验错误，失败路径和重复释放使资源计数失真。

# 触发
在项目根目录运行 `go test ./internal/repairbatch -run TestRepairBatchReleasesLeasesAndKeepsErrors`。

# 错误信息
`peak=7 open=0`、`validation error lost`、`open=-1`
