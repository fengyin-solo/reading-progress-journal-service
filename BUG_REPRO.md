# Bug 说明

阅读摘要构建发生 panic 时，半成品在恢复之前已发布到缓存，后续请求会读到未就绪对象。

## 如何触发

以标题 `broken` 构建 `digest-3`，收到构建错误后再读取同一 ID，然后查询摘要列表。

## 错误信息

第二次读取触发 `digest digest-3 is not ready` 并返回 500，列表同时暴露 `id=digest-3, ready=false, title=broken`。
