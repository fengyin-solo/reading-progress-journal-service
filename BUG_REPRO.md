# Bug

取消首个阅读元数据补全请求后，下游仍被调用两次；随后发起的全新请求继承首请求的取消状态并在调用下游前失败。

# Trigger

运行 context 场景：绑定 reader-A 的请求上下文并取消它，随后执行补全；紧接着用新的上下文执行 reader-B 的补全。

# Error

`canceled request continued metadata calls: error="" calls=2`

`fresh request inherited old cancellation or retried: error="context canceled" calls=0`
