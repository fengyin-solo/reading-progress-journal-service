# Bug

阅读摘录批次在解析器复用缓冲区、延迟回执、缓存回查和排队导出四个边界上共享可变切片，后续处理会改写已经交付的内容。

# Trigger

依次解析 alpha 与 bravo；创建 cedar 延迟回执后修改来源；缓存 delta 后修改一次查询结果；将 echo 排队后修改生产者对象，最后读取四个结果。

# Error

`first parsed batch was overwritten by buffer reuse: "bravo"`

`delayed receipt observed caller mutation: "Xedar"`

`cached batch was mutated through returned slice: "Xelta"`

`queued export observed producer buffer reuse: "Xcho"`
