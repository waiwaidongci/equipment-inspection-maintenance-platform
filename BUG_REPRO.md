# Bug
空巡检计划留下 nil map，添加检查项会 panic；可选校验器以 typed nil 形式返回并被误判为可用。

# 触发
在项目根目录运行 `go test ./internal/checklist -run TestEmptyPlanDoesNotPanicOrBypassNilChecks`。

# 错误信息
`add panicked: assignment to entry in nil map`、`optional validator is a typed nil`
