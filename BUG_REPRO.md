# Bug 是什么

复核状态活跃判定、转换图、重试成功回写和查询过滤四处不同步，成功记录无法关闭且中间态不可见。

# 如何触发

执行 `go test ./internal/reviewflow -run '^TestReviewRetryClosesAndRemainsQueryable$' -count=1`。

# 错误信息

测试报告 state activity contract inconsistent、transition graph incomplete、retry success remained retrying、active query lost retrying。
