# Bug

章节批量导入遇到损坏内容时吞掉解码错误，HTTP 返回成功，审计记录为 success；循环中的读取器直到函数收尾仍未释放。

# Trigger

依次导入 intro、details 和内容为 MALFORMED 的 broken 章节，再检查 HTTP 响应、审计状态以及读取器的当前数量和峰值。

# Error

`chapter import returned 200: {"code":0,"message":"ok","data":{"imported":["intro","details"],"error":"","audit_status":"success","open_readers":3,"peak_readers":3}}`
