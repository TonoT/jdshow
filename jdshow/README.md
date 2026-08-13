# JDShow

JDShow 是面向中国互联网、AI 和科技岗位的招聘雷达。当前版本是基于 `plan.md` 建立的可运行网站 MVP，使用内置演示数据验证首页、搜索、岗位详情和公司详情流程，尚未接入真实招聘采集和 PostgreSQL。

## 本地运行

需要 Go 1.22 或更高版本：

```bash
go run ./cmd/web
```

默认访问：http://localhost:8080

可通过环境变量修改监听地址：

```bash
ADDRESS=:8090 go run ./cmd/web
```

## 测试

```bash
go test ./...
go build ./...
```

## 项目结构

```text
cmd/web/                     网站启动入口
internal/app/                应用依赖装配
internal/company/            公司领域模型
internal/job/                岗位领域模型
internal/demo/               MVP 演示数据和候选目录
internal/source/shanghai/    上海公司 JSON 数据加载
internal/httpserver/         路由、Handler、中间件和模板
internal/platform/           环境配置
web/static/                  静态资源
data/company_jobs/           每家公司独立 JSON 数据
```

Go 项目通常不使用泛化的 `src` 目录，`cmd` 和 `internal` 就是标准源码目录。`internal` 中的包不会被项目外部直接导入。

## 当前范围

- 首页岗位与数据概览
- 关键词、城市、岗位方向筛选
- 岗位详情及官方来源跳转
- 公司详情及在招岗位
- 健康检查：`/healthz`、`/readyz`
- 基础安全响应头和结构化访问日志

演示岗位不代表公司当前真实招聘状态。当前版本已使用 SQLite 本地数据库 `data/db/jdshow.sqlite` 保存公司、类型、城市、行业和岗位数据；启动时会执行 `db/migrations/001_init.sql` 并幂等导入现有 JSON/演示数据。真实数据源、Worker 和管理员后台仍按 `plan.md` 后续逐模块实现。

SQLite 数据表包括：

```text
company_types       公司类型
industries          行业/职业类型
cities              城市
companies           公司基础信息、人数、背景、官网、Logo
company_cities      公司与城市关联
company_industries  公司与行业关联
job_types           岗位类型
jobs                职位详情及来源
```

公司详情优先展示 Logo；没有 Logo 时使用公司名称首字母。
