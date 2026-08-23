# Bug

book-A 的阅读次数快照缓存为 1 后，后台并发更新会把已经返回的快照与缓存汇总都改成 100001，并产生数据竞争。

# Trigger

创建包含 book-A=1 的阅读次数索引，将索引快照写入缓存；随后后台执行 100000 次递增，同时持续读取缓存视图，后台完成后再次读取快照和汇总。

# Error

`WARNING: DATA RACE`

`cached snapshot was rewritten by background updates: 100001`

`cached aggregate drifted with live store state: 100001`
