# Bug

阅读日志首次写入失败并由服务重试后，未提交的第一次尝试也发布成功通知，审计同时保留成功、失败和重试成功状态，且第一次写入错误被回滚清理错误覆盖。

# Trigger

调用保存场景处理器保存 `entry-42`。第一次写入固定失败且回滚清理也失败，服务随后执行第二次尝试并成功提交。

# Error

`first failure lost write or rollback error: "rollback cleanup failed"`

`notifications include an uncommitted or duplicate success: [{EntryID:entry-42 Attempt:1 State:saved} {EntryID:entry-42 Attempt:2 State:saved}]`

`audit trail did not converge on the committed retry: [{EntryID:entry-42 Attempt:1 State:saved} {EntryID:entry-42 Attempt:1 State:failed} {EntryID:entry-42 Attempt:2 State:saved}]`
