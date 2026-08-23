# Bug 说明

请求身份对象归还 sync.Pool 后仍被延迟回执和审计闭包持有，下一请求复用对象时覆盖了上一请求身份。

## 如何触发

以 reader-A 取得池对象并创建延迟回执与审计，随后归还对象；让 reader-B 立即复用同一对象后，再读取 A 的回执和审计。

## 错误信息

`first_receipt=reader-B` 且 `audit_reader=reader-B`，而这两个字段均属于 reader-A 的请求。
