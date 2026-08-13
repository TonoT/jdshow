# JDShow 公司与招聘入口解析资产清单

> 文档状态：公司来源资产初始登记稿  
> 更新日期：2026-08-10  
> 适用项目：`personwork/jdshow`  
> 重要说明：本文件用于后续建立 `CompanySource` / Crawler Adapter 注册表，不等于已获准抓取清单。官网和招聘入口会变化，所有 URL 必须在启用自动采集前重新核验 robots、服务条款、访问权限、字段授权和请求频率。

## 1. 目标与完成口径

目标候选池：

- 中国互联网 + AI + 科技公司：500 家。
- 中国国企/央企及其科技、数字化、信息化子企业：200 家。
- 外企中国区或亚洲科技相关雇主：100 家。

“完成一家公司”必须同时具备以下字段：

```text
company_name
company_type
official_site_url
official_company_intro_url
official_career_url
career_url_status
source_acquisition_method
parser_strategy
last_verified_at
verification_evidence
notes
```

当前文档只将能够从官方页面或官方域名确认的信息标为 `confirmed`。没有逐站核验的公司采用 `candidate`，URL 采用 `待核验`，不得直接注册为生产 Adapter。以下首批样例是开发前的种子数据，800 家全量目录应按第 6 节的批次流程补齐并逐条复核。

## 2. 状态定义

| 状态 | 含义 | 是否允许生产采集 |
|---|---|---|
| `confirmed` | 官网和招聘入口已通过官方域名核验，且获取方式已记录 | 仍需完成条款/robots 审核后才允许 |
| `candidate` | 公司属于目标范围，但官网招聘入口尚未逐站确认 | 否 |
| `pending_terms` | 入口已找到，服务条款、robots 或展示/存储范围待确认 | 否 |
| `blocked` | 需要登录、验证码、付费权限或明确禁止自动化访问 | 否 |
| `manual_only` | 暂无合规自动方式，只允许人工录入或用户提交链接 | 否 |
| `retired` | 链接失效、公司合并或来源不再提供职位 | 否 |

严禁将搜索引擎结果、招聘聚合站、培训机构页面、个人文章或社交媒体帖子作为 `official_career_url`。它们可以作为发现线索，但不能作为最终来源。

## 3. 统一解析方法

### 3.1 方法编号

| 方法 | 适用页面 | 实现方式 | 关键字段 | 风险与限制 |
|---|---|---|---|---|
| `OFFICIAL_API` | 官方公开职位 API | 仅使用公开或授权 API；保存请求版本、分页游标和响应时间 | 外部职位 ID、标题、地点、发布时间、状态、详情 URL | 必须确认授权、频率和字段保存范围 |
| `ATS_API` | 官方站点托管的 ATS 接口 | 只调用页面自身公开使用的接口；域名白名单和低频请求 | ID、职位详情、分类、地点、更新时间 | 接口可能变更，不能猜接口或绕过签名 |
| `JSON_LD` | 职位详情页含 `JobPosting` | 解析 JSON-LD，保留页面 URL 和抓取时间 | `title`、`datePosted`、`validThrough`、`hiringOrganization`、`jobLocation` | JSON-LD 可能不完整，必须回退 HTML |
| `EMBEDDED_JSON` | SSR/前端页面内嵌状态 | 使用 JSON 解析器读取明确的状态对象，不用正则拼 JSON | 外部 ID、列表分页、详情字段 | 不执行页面脚本，不解密或绕过访问控制 |
| `STATIC_HTML` | 服务端渲染职位列表/详情 | HTML 解析器 + 稳定语义选择器；保存解析器版本 | 标题、链接、正文、地点、日期 | 结构漂移；需 fixture 和字段完整率告警 |
| `SITEMAP_RSS` | 官方 Sitemap/RSS | 解析标准 XML，详情页再按低频策略处理 | URL、更新时间、分类 | 只发现 URL，不假设 URL 一定是职位 |
| `ATS_LINK_DISCOVERY` | 官网跳转到 Greenhouse/Workday/Lever 等 ATS | 只记录官方页面明确跳转的 ATS URL；Adapter 按 ATS 版本实现 | 公司映射、外部 ID、岗位字段 | ATS 域名和授权关系必须保留证据 |
| `MANUAL_SUBMISSION` | 没有公开自动化入口 | 管理员或用户提交官方岗位链接，定期人工复核 | URL、提交人、核验状态、发现时间 | 不自动抓取；保留人工证据 |

### 3.2 通用流水线

```text
Source Registry
  -> 合规状态检查
  -> Fetcher（域名白名单、超时、响应大小、限速）
  -> Parse（API / XML / JSON-LD / 内嵌 JSON / HTML）
  -> Normalize（公司、城市、职类、时间、薪资）
  -> Deduplicate（来源内确定性去重）
  -> Duplicate Candidate（跨来源候选）
  -> Job Storage（版本、状态、来源证据）
  -> Quality Check（字段完整率和解析漂移）
```

### 3.3 禁止行为

- 不绕过验证码、登录、付费墙、签名校验、访问控制或频率限制。
- 不使用账号池、代理池或伪造来源规避封禁。
- 不抓取求职者简历、手机号、邮箱、私信或其他用户隐私信息。
- 不使用未经授权的第三方职位聚合页面替代官方来源。
- 不把大模型生成的公司属性直接写为事实；必须有证据和人工确认。

## 4. 公司信息标准

### 4.1 官网公司介绍

优先级：官方“关于我们/公司简介/投资者关系/年报” > 官方招聘页中的企业简介 > 权威公开登记信息 > 待核验。

保存：

```text
intro_url
intro_title
intro_text_original
intro_text_normalized
intro_source_type
intro_published_or_updated_at
last_verified_at
verification_evidence
```

简介只保存与产品、行业、技术方向和招聘判断有关的最小必要内容；版权和展示范围待来源条款确认。

### 4.2 公司标签

以下标签必须带证据，不允许仅凭公司名称或模型猜测：

```text
internet_company
ai_company
cloud_infrastructure
technology_company
state_owned
central_state_owned
foreign_company
outsourcing_company
listed_company
```

每个标签保存 `source_url`、`source_type`、`confidence`、`last_verified_at`、`verified_by` 和 `is_manual_override`。

## 5. 首批官方入口种子清单

以下是已经通过公开检索发现的开发种子。由于页面可能动态变化，`last_verified_at` 只代表本轮检索时间；生产启用前仍需二次打开核验。

### 5.1 中国互联网、AI 与科技公司

| 编号 | 公司 | 官网 | 公司介绍入口 | 招聘入口 | 入口状态 | 解析策略 |
|---|---|---|---|---|---|---|
| CN-I001 | 百度 | https://www.baidu.com | https://home.baidu.com | 待核验 | candidate | `ATS_LINK_DISCOVERY` + `EMBEDDED_JSON` |
| CN-I002 | 阿里巴巴 | https://www.alibabagroup.com | https://www.alibabagroup.com/about-alibaba | https://talent.alibaba.com | pending_terms | `ATS_API` / `EMBEDDED_JSON` |
| CN-I003 | 腾讯 | https://www.tencent.com | https://www.tencent.com/zh-cn/about.html | https://join.qq.com | pending_terms | `STATIC_HTML` + `EMBEDDED_JSON` |
| CN-I004 | 字节跳动 | https://www.bytedance.com | https://www.bytedance.com/zh/about | https://jobs.bytedance.com | pending_terms | `ATS_API` / `EMBEDDED_JSON` |
| CN-I005 | 美团 | https://about.meituan.com | https://about.meituan.com/about | https://zhaopin.meituan.com | pending_terms | `EMBEDDED_JSON` |
| CN-I006 | 京东 | https://corporate.jd.com | https://corporate.jd.com/about | https://zhaopin.jd.com | candidate | `STATIC_HTML` / `EMBEDDED_JSON` |
| CN-I007 | 小米 | https://www.mi.com | https://www.mi.com/about | https://hr.xiaomi.com | candidate | `ATS_LINK_DISCOVERY` + `STATIC_HTML` |
| CN-I008 | 快手 | https://www.kuaishou.com | https://www.kuaishou.com/about | https://zhaopin.kuaishou.cn | candidate | `EMBEDDED_JSON` |
| CN-I009 | 网易 | https://www.163.com | https://ir.netease.com | https://hr.163.com | candidate | `STATIC_HTML` |
| CN-I010 | 拼多多 | https://www.pinduoduo.com | https://www.pinduoduo.com/about.html | 待核验 | candidate | `MANUAL_SUBMISSION`，入口确认后再实现 |
| CN-I011 | 小红书 | https://www.xiaohongshu.com | 待核验 | https://job.xiaohongshu.com | candidate | `EMBEDDED_JSON` |
| CN-I012 | 滴滴 | https://www.didiglobal.com | https://www.didiglobal.com/about | https://talent.didiglobal.com | candidate | `ATS_API` / `EMBEDDED_JSON` |
| CN-I013 | 携程 | https://www.trip.com | https://group.trip.com/about | https://campus.ctrip.com | candidate | `STATIC_HTML` / `EMBEDDED_JSON` |
| CN-I014 | 华为 | https://www.huawei.com | https://www.huawei.com/cn/about-huawei | https://career.huawei.com | pending_terms | `ATS_LINK_DISCOVERY` / `STATIC_HTML` |
| CN-I015 | 联想 | https://www.lenovo.com | https://www.lenovo.com/cn/zh/about | https://jobs.lenovo.com | candidate | `ATS_API` |
| CN-I016 | 科大讯飞 | https://www.iflytek.com | https://www.iflytek.com/about | https://campus.iflytek.com | candidate | `STATIC_HTML` |
| CN-I017 | 商汤科技 | https://www.sensetime.com | https://www.sensetime.com/cn/about | https://careers.sensetime.com | candidate | `ATS_LINK_DISCOVERY` |
| CN-I018 | 旷视科技 | https://www.megvii.com | https://www.megvii.com/about | 待核验 | candidate | `MANUAL_SUBMISSION` |
| CN-I019 | 云知声 | https://www.unisound.com | https://www.unisound.com/about | 待核验 | candidate | `STATIC_HTML` |
| CN-I020 | 智谱 AI | https://www.zhipuai.cn | https://www.zhipuai.cn/about | 待核验 | candidate | `MANUAL_SUBMISSION` |
| CN-I021 | 月之暗面 | https://www.moonshot.cn | 待核验 | 待核验 | candidate | `MANUAL_SUBMISSION` |
| CN-I022 | 零一万物 | https://01.ai | 待核验 | 待核验 | candidate | `MANUAL_SUBMISSION` |
| CN-I023 | MiniMax | https://www.minimaxi.com | 待核验 | 待核验 | candidate | `MANUAL_SUBMISSION` |
| CN-I024 | 百川智能 | https://www.baichuan-ai.com | 待核验 | 待核验 | candidate | `MANUAL_SUBMISSION` |
| CN-I025 | 深度求索 DeepSeek | https://www.deepseek.com | 待核验 | 待核验 | candidate | `MANUAL_SUBMISSION` |

### 5.2 国企与央企科技相关公司

| 编号 | 公司 | 官网 | 公司介绍入口 | 招聘入口 | 入口状态 | 解析策略 |
|---|---|---|---|---|---|---|
| CN-S001 | 中国电子信息产业集团有限公司 | https://www.cec.com.cn | 待核验 | 待核验 | candidate | `STATIC_HTML` / `SITEMAP_RSS` |
| CN-S002 | 中国电子科技集团有限公司 | https://www.cetc.com.cn | 待核验 | 待核验 | candidate | `STATIC_HTML` / `MANUAL_SUBMISSION` |
| CN-S003 | 中国电信 | https://www.chinatelecom.com.cn | 待核验 | 待核验 | candidate | `STATIC_HTML` / `EMBEDDED_JSON` |
| CN-S004 | 中国移动 | https://www.chinamobileltd.com | 待核验 | 待核验 | candidate | `STATIC_HTML` |
| CN-S005 | 中国联通 | https://www.chinaunicom.com | 待核验 | 待核验 | candidate | `STATIC_HTML` |
| CN-S006 | 国家电网 | https://www.sgcc.com.cn | 待核验 | 待核验 | candidate | `MANUAL_SUBMISSION`，各省入口需单独登记 |
| CN-S007 | 中国石油 | https://www.cnpc.com.cn | 待核验 | 待核验 | candidate | `STATIC_HTML` |
| CN-S008 | 中国石化 | https://www.sinopec.com | 待核验 | 待核验 | candidate | `STATIC_HTML` |
| CN-S009 | 中国海油 | https://www.cnooc.com.cn | 待核验 | 待核验 | candidate | `STATIC_HTML` |
| CN-S010 | 中国航天科技集团 | https://www.spacechina.com | 待核验 | 待核验 | candidate | `STATIC_HTML` / `MANUAL_SUBMISSION` |
| CN-S011 | 中国航天科工集团 | https://www.casic.com.cn | 待核验 | 待核验 | candidate | `STATIC_HTML` |
| CN-S012 | 中国航空工业集团 | https://www.avic.com | 待核验 | 待核验 | candidate | `STATIC_HTML` |
| CN-S013 | 中国船舶集团 | https://www.cssc.net.cn | 待核验 | 待核验 | candidate | `STATIC_HTML` |
| CN-S014 | 中国兵器工业集团 | https://www.norincogroup.com.cn | 待核验 | 待核验 | candidate | `MANUAL_SUBMISSION` |
| CN-S015 | 中国国家铁路集团 | https://www.china-railway.com.cn | 待核验 | 待核验 | candidate | `STATIC_HTML` |
| CN-S016 | 中国建筑 | https://www.cscec.com | 待核验 | 待核验 | candidate | `STATIC_HTML` |
| CN-S017 | 中国交通建设 | https://www.ccccltd.cn | 待核验 | 待核验 | candidate | `STATIC_HTML` |
| CN-S018 | 中国核工业集团 | https://www.cnnc.com.cn | 待核验 | 待核验 | candidate | `STATIC_HTML` |
| CN-S019 | 中国广核 | https://www.cgnpc.com.cn | 待核验 | 待核验 | candidate | `STATIC_HTML` |
| CN-S020 | 中国华能 | https://www.chng.com.cn | 待核验 | 待核验 | candidate | `STATIC_HTML` |
| CN-S021 | 中国大唐 | https://www.china-datang.com | 待核验 | 待核验 | candidate | `STATIC_HTML` |
| CN-S022 | 国家能源集团 | https://www.chnenergy.com.cn | 待核验 | 待核验 | candidate | `STATIC_HTML` |
| CN-S023 | 中国投资有限责任公司 | https://www.china-inv.cn | 待核验 | 待核验 | candidate | `MANUAL_SUBMISSION` |
| CN-S024 | 中国邮政 | https://www.chinapost.com.cn | 待核验 | 待核验 | candidate | `STATIC_HTML` |
| CN-S025 | 浪潮集团 | https://www.inspur.com | https://www.inspur.com/about | 待核验 | candidate | `STATIC_HTML` / `EMBEDDED_JSON` |

### 5.3 外企科技相关公司

| 编号 | 公司 | 官网 | 公司介绍入口 | 招聘入口 | 入口状态 | 解析策略 |
|---|---|---|---|---|---|---|
| INTL-001 | Microsoft | https://www.microsoft.com | https://www.microsoft.com/about | https://jobs.careers.microsoft.com | candidate | `ATS_API` / `EMBEDDED_JSON` |
| INTL-002 | Google | https://about.google | https://about.google/intl/zh-CN/ | https://careers.google.com | candidate | `ATS_API` |
| INTL-003 | Amazon | https://www.amazon.com | https://www.aboutamazon.com | https://www.amazon.jobs | candidate | `EMBEDDED_JSON` |
| INTL-004 | Apple | https://www.apple.com | https://www.apple.com/leadership | https://jobs.apple.com | candidate | `EMBEDDED_JSON` |
| INTL-005 | IBM | https://www.ibm.com | https://www.ibm.com/about | https://www.ibm.com/careers | candidate | `ATS_LINK_DISCOVERY` |
| INTL-006 | Oracle | https://www.oracle.com | https://www.oracle.com/corporate | https://careers.oracle.com | candidate | `ATS_API` |
| INTL-007 | SAP | https://www.sap.com | https://www.sap.com/about/company.html | https://jobs.sap.com | candidate | `ATS_API` |
| INTL-008 | Siemens | https://www.siemens.com | https://www.siemens.com/global/en/company/about.html | https://jobs.siemens.com | candidate | `ATS_API` |
| INTL-009 | Intel | https://www.intel.com | https://www.intel.com/content/www/us/en/company-overview/company-overview.html | https://jobs.intel.com | candidate | `ATS_API` |
| INTL-010 | NVIDIA | https://www.nvidia.com | https://www.nvidia.com/en-us/about-nvidia | https://www.nvidia.com/en-us/about-nvidia/careers | candidate | `STATIC_HTML` / `EMBEDDED_JSON` |
| INTL-011 | AMD | https://www.amd.com | https://www.amd.com/en/corporate | https://careers.amd.com | candidate | `ATS_API` |
| INTL-012 | Qualcomm | https://www.qualcomm.com | https://www.qualcomm.com/company | https://www.qualcomm.com/company/careers | candidate | `ATS_LINK_DISCOVERY` |
| INTL-013 | Cisco | https://www.cisco.com | https://www.cisco.com/c/en/us/about.html | https://jobs.cisco.com | candidate | `ATS_API` |
| INTL-014 | Dell Technologies | https://www.dell.com | https://www.dell.com/en-us/dt/corporate/about-us.htm | https://jobs.dell.com | candidate | `ATS_API` |
| INTL-015 | Accenture | https://www.accenture.com | https://www.accenture.com/us-en/about | https://www.accenture.com/us-en/careers | candidate | `STATIC_HTML` / `EMBEDDED_JSON` |
| INTL-016 | SAP Concur | https://www.concur.com | https://www.concur.com/about | https://jobs.sap.com | candidate | `ATS_API`，母公司 ATS |
| INTL-017 | Bosch | https://www.bosch.com | https://www.bosch.com/company | https://www.bosch.com/careers | candidate | `ATS_LINK_DISCOVERY` |
| INTL-018 | BMW Group | https://www.bmwgroup.com | https://www.bmwgroup.com/en/company.html | https://www.bmwgroup.jobs | candidate | `ATS_API` |
| INTL-019 | Mercedes-Benz | https://group.mercedes-benz.com | https://group.mercedes-benz.com/company | https://group.mercedes-benz.com/careers | candidate | `STATIC_HTML` |
| INTL-020 | Toyota | https://global.toyota | https://global.toyota/en/company | https://global.toyota/en/careers | candidate | `STATIC_HTML` |

## 6. 800 家候选池的补齐方案

本轮已登记 70 家种子公司，不把未逐站核验的名称、简介和 URL 伪装成已完成数据。要达到 500 + 200 + 100，按以下批次补齐：

### 6.1 中国互联网 + AI 500 家

按优先级建立 500 家候选：

- 一级平台与互联网公司：综合平台、电商、内容、社交、出行、物流、游戏、在线教育、生活服务。
- AI 与大模型：基础模型、计算机视觉、语音、机器人、具身智能、自动驾驶、AI 应用。
- 云计算与基础设施：公有云、数据库、中间件、数据中心、服务器、芯片、网络安全。
- 科技制造：智能终端、半导体、工业软件、通信、智能汽车和新能源数字化。
- 区域科技公司：按北京、上海、深圳、杭州、广州、成都、苏州、南京、武汉、西安等城市建立候选池。

每批 50 家，先查官方公司站，再查官方招聘站；不以“行业知名”替代官网核验。建议顺序：一级平台 50、AI 基础模型 50、云和基础设施 50、芯片/硬件 100、智能汽车/机器人 50、游戏/内容/电商 100、区域科技公司 100。

### 6.2 国企/央企 200 家

分为：

- 央企集团总部及数字科技子企业 60 家。
- 电信、能源、电力、交通、航空航天、军工、金融科技等央企科技单位 80 家。
- 省属国企、城市投资和地方数字科技平台 60 家。

央企身份优先以国务院国资委或公司正式披露为证据；地方国企保存国资监管机构或公司年报证据。集团官网不一定是招聘入口，子企业招聘页必须单独建记录，不能把集团 URL 复制给所有子公司。

### 6.3 外企 100 家

分为：

- 软件、云服务、互联网和企业 IT 30 家。
- 芯片、通信、硬件和自动化 25 家。
- 汽车、工业、能源和智能制造 25 家。
- 医疗科技、咨询、消费科技和研发机构 20 家。

优先选择中国区或亚洲区存在官方技术岗位的雇主。全球 careers 页面若无法筛选中国地区，只能登记为候选入口，并补充中国区官方站或地区 ATS。

### 6.4 每批交付物

每个 50 家批次必须提交：

- 公司名称、别名、类别、总部/中国区范围。
- 官方官网、公司介绍页、官方招聘页和招聘地区。
- 入口状态、最后核验时间、证据 URL。
- 解析方法编号和预计字段覆盖率。
- robots/条款核验结果、限速建议和人工回退方案。
- 至少 3 个脱敏职位详情样本，用于 Parser Fixture；若无法合法保存原文，保存字段结构和哈希，不保存正文。

## 7. 数据文件建议

后续不建议把 800 家信息只维护在 Markdown 表格中。`parse.md` 用于规则、审核和人工阅读，机器注册表建议使用：

```text
internal/source/registry.yaml
internal/source/companies.yaml
internal/source/fixtures/<source>/<sample>.json
```

建议字段：

```yaml
id: cn-i-001
name: 百度
category: internet_ai
official_site_url: https://www.baidu.com
company_intro_url: https://home.baidu.com
career_url: null
career_url_status: candidate
acquisition_method: manual_submission
parser_strategy: [static_html, embedded_json]
terms_review_status: pending
robots_reviewed_at: null
last_verified_at: 2026-08-10
verification_evidence: []
rate_limit_policy: null
fallback: official_career_link_or_manual_submission
notes: 生产启用前必须人工打开并复核
```

公司介绍、招聘入口和解析器应解耦：同一公司可有多个招聘来源；同一 ATS 解析器可服务多个公司，但每家公司仍保存自己的官方跳转证据。

## 8. 解析器实现规范

### 8.1 API/ATS

- 使用请求 DTO 和响应 DTO，不把第三方字段直接暴露给领域层。
- 保存外部 ID、分页游标、字段版本和原始响应哈希。
- 使用 schema 校验；字段缺失进入质量告警，不静默填充。
- 分页必须有最大页数和最大职位数，避免异常接口导致无限抓取。
- 认证信息只能从环境变量/Secret 读取；未授权 API 不接入。

### 8.2 HTML/JSON-LD

- 优先提取 JSON-LD `JobPosting`，其次是页面明确的内嵌 JSON，再回退到语义化 HTML。
- 选择器按版本管理，至少覆盖标题、详情链接、地点、发布时间和正文。
- 日期按页面时区解析，写入 UTC，并保留原始日期文本。
- 富文本先做 HTML 白名单清洗；不得执行页面脚本。
- 对列表页和详情页分别做 fixture 测试，记录解析器版本。

### 8.3 失败处理

- HTTP 4xx/5xx、超时、空响应、结构漂移、字段校验失败分别分类。
- 解析失败不删除已有岗位，不触发批量 `EXPIRED`。
- 同一来源连续失败时自动暂停生产采集，生成告警并转人工复核。
- Adapter 只访问配置的官方域名和路径；禁止根据页面内容任意跳转抓取。

## 9. 核验与维护流程

1. 发现公司：通过官方站、监管名录或公司公开披露建立候选记录。
2. 官方确认：打开官网“关于/公司简介”和招聘入口，记录 URL、页面标题、时间和证据。
3. 合规检查：核验 robots、条款、接口授权、字段使用和频率限制。
4. 技术探查：只观察公开页面自身请求，选择 API、JSON-LD、内嵌 JSON 或 HTML 策略。
5. 样本测试：人工选取列表和详情样本，确认字段、时区、状态和分页。
6. 入库评审：状态变为 `pending_terms` 或 `confirmed`，由管理员批准是否启用。
7. 定期复核：默认每 30 天复核入口和条款；来源失败或结构变化立即复核。
8. 变更追踪：保存 URL、解析器版本、字段覆盖率、最后成功运行和停用原因。

## 10. 当前待确认项

- 800 家全量候选池的具体公司名单和优先级。
- 首批真正允许自动采集的 2～3 个官方来源。
- 各公司招聘入口是否支持社会招聘，还是只有校园招聘。
- 招聘内容是否允许缓存、结构化展示和保存历史版本。
- 公司简介与 Logo 的版权、引用和缓存范围。
- 央国企子公司归属口径，以及是否把集团与子企业分别计数。
- 外企“中国区/亚洲区”范围和职位地域筛选口径。
- 需要采集的城市、岗位类别、更新频率和数据保留周期。

在上述问题确认前，所有未逐站核验记录保持 `candidate` 或 `pending_terms`，不创建生产采集任务。

## 11. 参考发现来源

以下链接仅用于发现和交叉核验，不作为招聘数据源：

- 国务院国资委招聘公告示例：https://www.sasac.gov.cn/
- 智联招聘（聚合平台，当前项目不直接自动采集）：https://www.zhaopin.com/
- 中国政府网及国资监管机构公开信息：以具体官方公告 URL 为准。
- 外企官网必须回到各公司官方域名的 About/Careers 页面核验。

## 12. 下一步

建议先完成 10 家官方来源的逐站核验和 Parser Fixture，再批量补齐 800 家目录。执行顺序：

1. 确认首批公司和数据获取授权。
2. 完善公司注册表和证据字段。
3. 为同一 ATS 或同一页面类型实现通用 Parser。
4. 通过 3 个来源验证标准化、去重和状态流转。
5. 每批 50 家扩展公司清单，逐条核验后再启用。

本文件当前是“来源规划和解析登记文档”，不是绕过任何网站访问控制的抓取方案。
