# Bug

阅读摘要第一次已送达但确认超时后，重试会再次产生外部发送；第二次完成后，第一轮陈旧回调又把任务与缓存状态覆盖为 retrying。

# Trigger

发布 reading-weekly 摘要，让第一次发送返回确认超时，再执行一次重试并完成，最后送达第一轮的 retrying 回调。

# Error

`retry duplicated summary delivery: 2`

`stale callback regressed stored job to retrying@1`

`stale callback regressed cached job to retrying@1`
