# Bug 是什么

巡检缓存构造、快照读取、汇总和 worker 合并都共享内部 map，导致历史污染和 data race。

# 如何触发

执行 `go test -race ./internal/inspectioncache -run '^TestCacheReadWriteKeepsHistoricalInspectionView$' -count=1`。

# 错误信息

race detector 报 Worker.Merge 写入与 Service.Summarize 读取冲突，同时历史快照和 view copy 断言失败。
