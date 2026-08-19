# Bug 是什么

证据批处理吞掉业务错误、资源 Close 不生效、空证据校验遗漏，非法持久化路径还会 panic。

# 如何触发

执行 `go test ./internal/evidencebatch -run '^TestArchiveClosesEvidenceHandlesAndJoinsFailures$' -count=1`，再执行 `go test ./internal/evidencebatch -run '^TestBadEvidencePersistenceReturnsError$' -count=1`。

# 错误信息

首条测试报告 batch lost an error、三条 resource remains open、service accepted empty evidence；第二条 panic：`save evidence: invalid inspection evidence`。
