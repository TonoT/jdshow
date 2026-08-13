# 官网数据收集与核验规范

> 更新日期：2026-08-10  
> 用途：指导 JDShow 后续扩展公司官网、公司介绍页和官方招聘来源。  
> 数据目录：`data/companies/<company_type>/<industry>/*.json`

## 1. 收集原则

1. 优先从公司官方网站发现公司介绍和招聘入口。
2. 官方首页明确跳转的独立招聘域名或 ATS，可以作为官方招聘来源，但必须保存跳转证据。
3. 搜索引擎只用于发现候选链接，最终必须回到官方域名页面核验。
4. 招聘平台、培训机构、百科、媒体和个人文章不能作为官网或官方招聘入口证据。
5. 不绕过登录、验证码、付费墙、签名、频率限制和其他访问控制。
6. 不采集求职者姓名、联系方式、简历或其他隐私数据。
7. 每次核验都保存日期、证据 URL、HTTP 结果、最终 URL 和核验说明。

## 2. 目录分类

第一层按公司类型区分：

```text
private       民营、互联网和 AI 公司
state_owned   国企、央企及其科技子企业
foreign       外企及其中国区/亚洲招聘来源
```

第二层按主要行业区分，例如：

```text
ai
internet_platform
software_cloud
telecom_cloud
semiconductor
security
robotics
smart_vehicle
industrial_technology
```

一家公司有多个行业时，只放入主行业文件，其他行业写入 `industry_tags`，避免重复维护。集团和独立招聘的子公司分别建立记录，并使用 `parent_company_id` 关联。

## 3. 标准数据结构

每个 JSON 文件包含同一公司类型和主行业下的记录：

```json
{
  "schema_version": "1.0",
  "company_type": "private",
  "primary_industry": "ai",
  "companies": []
}
```

公司记录必填字段：

```text
id
name
english_name
company_type
primary_industry
industry_tags
official_site_url
company_intro_url
career_url
career_scope
parser_strategy
verification_status
verified_at
verification_evidence
```

`verification_evidence` 至少包含：

```text
url
kind               official_home / official_intro / official_career / official_announcement
http_status
final_url
checked_at
notes
```

## 4. 收集流程

### 4.1 官网确认

1. 搜索公司全称和“官网”，找到候选域名。
2. 打开候选官网，确认页面品牌、公司名称、备案主体或版权主体。
3. 保存规范化 HTTPS 首页 URL；去除追踪参数和无意义锚点。
4. 若集团官网与产品官网不同，以法律主体或官方公司站作为 `official_site_url`。

### 4.2 公司介绍确认

按以下顺序选择：

1. 官网“关于我们/公司简介”。
2. 官网投资者关系或正式年报中的公司介绍。
3. 官网招聘页面中的企业介绍。
4. 官网没有独立介绍页时，使用首页，并在 `notes` 说明。

简介字段只保存事实性短摘要，必须能从证据页面直接支持，不能由大模型补充推断。

### 4.3 招聘入口确认

1. 从官网导航、页脚或官方招聘公告寻找“招聘/加入我们/Careers/Jobs”。
2. 如果跳转到独立域名，记录官网跳转页和最终招聘 URL。
3. 确认招聘页面展示公司名称或品牌，并支持目标地区/招聘类型。
4. 只支持校园招聘时，`career_scope` 写 `campus`；社会招聘写 `social`；两者都有写 `social_and_campus`。
5. 入口需要登录才能浏览、出现验证码或条款不明确时，不继续探查，状态设为 `manual_only` 或 `pending_terms`。

### 4.4 解析方式判定

按优先级记录，不在收集阶段实现绕过：

- `OFFICIAL_API`：官方文档公开或明确授权的 API。
- `JSON_LD`：职位详情包含标准 `JobPosting` JSON-LD。
- `EMBEDDED_JSON`：公开 HTML 中的服务端内嵌 JSON。
- `STATIC_HTML`：服务端 HTML 列表和详情。
- `ATS_LINK_DISCOVERY`：官网明确跳转到 ATS，具体 Adapter 后续复核。
- `MANUAL_SUBMISSION`：没有稳定、合规的自动入口。

不得仅凭页面由 JavaScript 渲染就推断存在可调用 API。网络接口的使用方式、条款和频率需要独立审核。

## 5. 准确性校验

每条记录执行以下检查：

1. **格式校验**：JSON 可解析，ID 唯一，枚举和必填字段完整。
2. **URL 校验**：只允许 `https`，无用户名密码、无本地/私网地址。
3. **可达性校验**：记录 HTTP 状态和最终跳转 URL；`2xx/3xx` 只表示可达，不等于官方归属正确。
4. **归属校验**：官网页面品牌/备案主体与公司匹配；招聘独立域名必须由官网链接或官方公告背书。
5. **内容校验**：介绍摘要能由官方证据支持；招聘页面确实包含职位或招聘说明。
6. **时间校验**：`verified_at` 与证据 `checked_at` 使用 ISO 8601 UTC 时间。
7. **交叉校验**：网站当前使用的公司与岗位来源必须在数据文件中存在。

校验状态：

- `verified`：官方归属、入口和内容均确认。
- `verified_manual_only`：官网和招聘入口确认，但当前只允许人工使用链接。
- `pending_terms`：入口确认，自动采集条款或 robots 待审核。
- `candidate`：只确认公司，不确认招聘入口。
- `invalid`：链接错误、归属不符或已失效。

## 6. 更新和失效处理

- 默认每 30 天复核一次已启用招聘入口。
- HTTP 失败不立即删除记录，先记录失败时间并连续复核。
- 官网更换招聘域名时保留旧 URL 和失效原因，新增证据后再切换。
- 解析结构变化只暂停该来源，不触发岗位批量失效。
- 公司更名、合并或拆分时保留别名和主体关系，不复用旧 ID 表示新主体。

## 7. 当前网站数据约束

当前网站使用百度、月之暗面、中国电信和 Microsoft 四家公司。对应数据文件是：

```text
data/companies/private/ai/companies.json
data/companies/state_owned/telecom_cloud/companies.json
data/companies/foreign/software_cloud/companies.json
```

这些文件是后续数据库种子和来源注册表的候选输入。当前 Go 网站尚未改为从 JSON 加载，避免在本次来源整理中扩大业务改动范围。
