# Bug
已删除设备的 missing 错误在传递中丢失，接口返回 500 并继续重试查询。

# 触发
在项目根目录运行 `go test ./internal/equipmentlookup -run TestUnknownEquipmentReturnsOneNotFoundWithoutRetry`。

# 错误信息
`missing identity was discarded`
