# 公司岗位独立数据目录

每家公司一个 JSON 文件，岗位信息放在该公司的 `jobs` 数组中。

## 目录

```text
company_jobs/
├── private/ai/                         # 已有 AI 公司独立岗位文件
├── private/shanghai_internet/          # 上海互联网/科技民营公司候选池
├── state_owned/telecom_cloud/          # 国企/央企科技公司
└── foreign/software_cloud/             # 外企软件和云计算公司
```

## 状态说明

- `verified_source_pending_terms`：官方来源已找到并核验，自动采集条款待确认。
- `verified_announcement_pending_terms`：官方招聘公告已核验，招聘栏目本身仍需人工复核。
- `candidate_pending_verification`：公司属于候选池，官网、招聘入口和岗位尚未逐站核验。

候选公司文件中的 `jobs` 为空是有意设计：没有官方岗位详情证据时不创建虚假 JD。完成官网核验后，才将真实岗位写入对应公司的 JSON 文件。

## 文件字段

```text
company_id
company_name
company_type
primary_industry
city_focus
data_status
official_site_url
company_intro_url
career_url
parser_strategy
source_notes
jobs
```

每条岗位至少包含：

```text
id
title
city
category
experience
education
salary
status
source_name
source_url
description
requirements
skills
```

上海候选池的 `index.json` 记录公司数量、公司名称、ID 和对应文件名。当前目标数量为 200 家，均需后续逐站核验后才能进入自动采集。
