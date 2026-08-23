# 阅读记录系统（readinglog）

纯 Go 标准库实现的阅读记录管理后端服务，零第三方依赖。用于管理书籍、书架、读书笔记、书摘与阅读进度。

## 运行

```bash
# 启动服务（默认监听 :8080）
go run ./cmd/server

# 自定义端口与配置
PORT=9090 MAX_PAGE_SIZE=100 go run ./cmd/server
```

## 环境变量

| 变量 | 默认值 | 说明 |
|------|--------|------|
| PORT | 8080 | 监听端口 |
| ADDR | 空（覆盖 PORT） | 完整监听地址 |
| MAX_PAGE_SIZE | 100 | 分页单页最大条数 |
| LOG_LEVEL | info | 日志级别（debug/info/warn/error） |

## API 一览

统一响应结构：`{"code":0,"message":"ok","data":...}`；错误时 `code` 非 0，HTTP 状态码对应 400/404/409/500。

### 书籍 Book

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/books | 新增书籍 |
| GET | /api/books?category=&status=&keyword=&page=&size= | 分页列表 |
| GET | /api/books/{id} | 查询书籍 |
| PUT | /api/books/{id} | 更新书籍 |
| DELETE | /api/books/{id} | 删除书籍 |
| POST | /api/books/{id}/start | 开始阅读（wishlist→reading） |
| POST | /api/books/{id}/finish | 标记读完（reading→finished） |
| POST | /api/books/{id}/drop | 弃读（wishlist/reading→dropped） |

书籍状态枚举：`wishlist`（想读）/ `reading`（在读）/ `finished`（读完）/ `dropped`（弃读）

### 书架 Shelf

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/shelves | 新增书架（名称唯一） |
| GET | /api/shelves?keyword=&page=&size= | 分页列表 |
| GET | /api/shelves/{id} | 查询书架 |
| PUT | /api/shelves/{id} | 更新书架 |
| DELETE | /api/shelves/{id} | 删除书架 |

### 笔记 Note

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/notes | 新增笔记（校验 book_id 存在） |
| GET | /api/notes?book_id=&keyword=&page=&size= | 分页列表 |
| GET | /api/notes/{id} | 查询笔记 |
| PUT | /api/notes/{id} | 更新笔记 |
| DELETE | /api/notes/{id} | 删除笔记 |

### 书摘 Highlight

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/highlights | 新增书摘（校验 book_id 存在） |
| GET | /api/highlights?book_id=&chapter=&page=&size= | 分页列表 |
| GET | /api/highlights/{id} | 查询书摘 |
| PUT | /api/highlights/{id} | 更新书摘 |
| DELETE | /api/highlights/{id} | 删除书摘 |
| POST | /api/highlights/batch-delete | 批量删除，body `{"ids":[...]}` |

### 进度 Progress

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/progresses | 新增进度（校验 book_id 存在，页数不能超过总页数） |
| GET | /api/progresses?book_id=&page=&size= | 分页列表 |
| GET | /api/progresses/{id} | 查询进度 |
| PUT | /api/progresses/{id} | 更新进度 |
| DELETE | /api/progresses/{id} | 删除进度 |

进度字段说明：`current_page` 为当前页码，`percentage` 为进度百分比（未显式指定时按 `current_page/total_pages*100` 自动计算）。

### 统计 Stats

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/stats/overview | 整体概览统计 |
| GET | /healthz | 健康检查 |

## 项目结构

```
origin/
├── cmd/server/main.go       # 入口：配置加载、依赖装配、优雅关闭
├── internal/
│   ├── app/app.go           # 依赖装配 store -> service -> handler
│   ├── config/config.go     # 环境变量配置
│   ├── model/               # 领域模型 + 校验 + 状态机
│   ├── store/               # Store 接口 + 内存实现
│   ├── service/             # 业务逻辑
│   └── handler/             # HTTP 路由 + 处理器
└── pkg/
    ├── httpx/               # 统一响应、分页、JSON 解析
    ├── idgen/               # ID 生成
    └── logger/              # 分级日志
```

## 测试

```bash
go test ./...
```
