# Bug
维修进入复核后无法转为关闭，工作器写回旧状态，复核记录也从活动列表消失。

# 触发
在项目根目录运行 `go test ./internal/maintenanceflow -run TestRepairVerificationCanCloseAndRemainVisible`。

# 错误信息
`retry state="in_progress"`
