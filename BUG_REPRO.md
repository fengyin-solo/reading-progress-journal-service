# Bug 说明

缺省阅读规则同时持有 typed-nil 校验器和 nil 规则表，一条路径静默跳过校验，另一条路径写 map 时 panic。

## 如何触发

在不提供额外阅读规则配置的情况下，先检查 0 页内容，再写入名为 daily、值为 12 的规则。

## 错误信息

第一次检查返回 `allowed=true`，随后规则写入返回 500，日志为 `assignment to entry in nil map`。
