package demo

import "jdshow/internal/company"

func VerifiedCompanies() []company.Company {
	companies := []company.Company{
		verified("private-extra-1", "上海悦商信息科技", "Yueshang Technology", "互联网 / AI", "商业地产数字化与 AI", "上海", "专注存量资产运营数字化，提供商业地产和资产管理全链条系统、AI 风控及自动化运营产品。", "11-50 人（公开职业资料估算，待工商口径复核）", "民营科技公司；官网披露为腾讯投后生态及核心产品研发供应商、宝龙商业投后数字化服务商。", "https://yueworld.cn", []company.Source{
			{Type: "official", URL: "https://yueworld.cn", Title: "悦商科技官网", SupportsFields: "官网、主营业务、公司背景、上海地址", VerifiedAt: "2026-08-12"},
			{Type: "professional_network", URL: "https://cn.linkedin.com/company/%E4%B8%8A%E6%B5%B7%E6%82%A6%E5%95%86%E4%BF%A1%E6%81%AF%E7%A7%91%E6%8A%80%E6%9C%89%E9%99%90%E5%85%AC%E5%8F%B8", Title: "上海悦商信息科技有限公司 LinkedIn", SupportsFields: "公司主体、行业、上海地址、关联成员", VerifiedAt: "2026-08-12"},
		}),
		verified("021-company-021", "得物", "Dewu", "互联网 / AI", "潮流电商与社区", "上海", "提供正品潮流电商、商品鉴别和潮流生活社区服务，采用先鉴别后发货的平台模式。", "5001-10000 人（LinkedIn 公开口径）", "上海民营互联网平台，运营主体为上海得物信息集团有限公司。", "https://tech.dewu.com", []company.Source{
			{Type: "official", URL: "https://tech.dewu.com", Title: "得物技术官网", SupportsFields: "官网、业务介绍、技术方向、社会招聘入口", VerifiedAt: "2026-08-12"},
			{Type: "app_store", URL: "https://play.google.com/store/apps/details?id=com.shizhuang.duapp&hl=zh_CN", Title: "Google Play 得物应用页", SupportsFields: "运营主体、产品介绍、上海地址", VerifiedAt: "2026-08-12"},
			{Type: "professional_network", URL: "https://cn.linkedin.com/company/%E6%AF%92app", Title: "得物 LinkedIn", SupportsFields: "人员规模、成立年份、上海办公地点", VerifiedAt: "2026-08-12"},
		}),
		verified("195-company-195", "叠纸游戏", "Papergames", "互联网 / AI", "游戏与互动娱乐", "上海", "专注原创内容和互动娱乐，代表产品包括暖暖系列、恋与制作人、恋与深空和无限暖暖。", "2000 人以上（公开招聘/产业资料口径）", "民营游戏研发公司，总部位于上海并开展全球发行。", "https://www.papegames.com", []company.Source{
			{Type: "official", URL: "https://www.papegames.com", Title: "叠纸游戏官网", SupportsFields: "官网、产品线、公司愿景、招聘入口", VerifiedAt: "2026-08-12"},
			{Type: "university_career", URL: "https://career.cuhk.edu.cn/job/view/id/467190", Title: "香港中文大学（深圳）叠纸招聘信息", SupportsFields: "公司主体、上海总部、产品、招聘方向", VerifiedAt: "2026-08-12"},
		}),
		verified("196-company-196", "莉莉丝游戏", "Lilith Games", "互联网 / AI", "游戏研发与全球发行", "上海", "集游戏研发与全球发行为一体，代表产品包括小冰冰传奇、剑与远征、万国觉醒和剑与远征：启程。", "2300+ 人（官方招聘口径）", "上海民营游戏公司，在新加坡、韩国和美国等地设有分支机构。", "https://www.lilith.com", []company.Source{
			{Type: "official", URL: "https://www.lilith.com", Title: "莉莉丝游戏官网", SupportsFields: "官网、产品、公司主体、上海地址", VerifiedAt: "2026-08-12"},
			{Type: "official_career", URL: "https://jobs.lilith.com", Title: "莉莉丝招聘官网", SupportsFields: "公司介绍、2300+人员规模、办公城市、招聘入口", VerifiedAt: "2026-08-12"},
		}),
		verified("197-company-197", "米哈游", "miHoYo", "互联网 / AI", "游戏、AI 与数字内容", "上海", "研发和运营游戏及数字内容，代表产品包括崩坏系列、原神、崩坏：星穹铁道和绝区零，并投入动画、音乐和前沿技术研发。", "约5000 人（公开资料口径）", "上海民营科技和游戏公司，2012 年成立，开展全球化运营。", "https://www.mihoyo.com", []company.Source{
			{Type: "official", URL: "https://www.mihoyo.com", Title: "米哈游官网", SupportsFields: "官网、品牌和产品", VerifiedAt: "2026-08-12"},
			{Type: "public_reference", URL: "https://zh.wikipedia.org/zh-cn/%E7%B1%B3%E5%93%88%E6%B8%B8", Title: "米哈游公开资料", SupportsFields: "上海总部、成立时间、人员规模、产品", VerifiedAt: "2026-08-12"},
		}),
		verified("027-b", "B站", "Bilibili", "互联网 / AI", "视频社区与数字内容", "上海", "面向年轻用户的视频社区，业务覆盖视频内容、直播、游戏、广告、漫画和电商。", "8423 人（截至2025-12-31，官方投资者关系口径）", "纳斯达克和香港联交所双重主要上市公司，总部位于上海杨浦。", "https://www.bilibili.com", []company.Source{
			{Type: "official", URL: "https://www.bilibili.com/blackboard/aboutUs.html", Title: "哔哩哔哩关于我们", SupportsFields: "官网、业务介绍、上市背景", VerifiedAt: "2026-08-12"},
			{Type: "official_ir", URL: "https://ir.bilibili.com/cn/investor-resources", Title: "哔哩哔哩投资者资源", SupportsFields: "人员规模、上海总部、上市市场、公司定位", VerifiedAt: "2026-08-12"},
		}),
		verified("tech-9", "拼多多", "Pinduoduo", "互联网 / AI", "电子商务与农业科技", "上海", "新电商平台，通过社交拼购、农产品上行、供应链和农业科技投入连接消费者与生产者。", "8000+ 人（官网2021披露；当前规模待最新报告复核）", "上海创立的电商平台，所属 PDD Holdings 在纳斯达克上市。", "https://www.pinduoduo.com", []company.Source{
			{Type: "official", URL: "https://www.pinduoduo.com/home/about", Title: "拼多多关于我们", SupportsFields: "官网、业务模式、人员规模、发展历程", VerifiedAt: "2026-08-12"},
			{Type: "market_disclosure", URL: "https://emweb.eastmoney.com/pc_usf10/CompanyInfo/index?color=web&code=PDD", Title: "PDD 公司资料", SupportsFields: "上市主体、运营主体、行业和人员规模", VerifiedAt: "2026-08-12"},
		}),
		verified("state-new-001", "上海数据集团", "Shanghai Data Group", "国企 / 央企科技", "数据基础设施与数据要素", "上海", "承担上海城市大数据资源基础治理和公共数据授权运营，布局数据基础设施、数据流通、数据产品和数字化服务。", "1300+ 人（2024年上海市国资委公开口径）", "上海市委、市政府批准成立的市属一级国有企业，2022 年揭牌。", "https://www.shdatagroup.com", []company.Source{
			{Type: "official", URL: "https://www.shdatagroup.com", Title: "上海数据集团官网", SupportsFields: "官网、成立背景、主营业务、地址", VerifiedAt: "2026-08-12"},
			{Type: "government", URL: "https://www.gzw.sh.gov.cn/shgzw_xwzx_xwfb/20240723/7f70b22ead624ba9908dea3533efb78b.html", Title: "上海市国资委上海数据集团报道", SupportsFields: "国企背景、人员规模、公共数据职责", VerifiedAt: "2026-08-12"},
		}),
	}
	for i := range companies {
		companies[i].CareerName, companies[i].CareerURL = careerLink(companies[i].Name)
	}
	return companies
}

func ApplyCareerLinks(companies map[string]company.Company) {
	for id, c := range companies {
		c.CareerName, c.CareerURL = careerLink(c.Name)
		companies[id] = c
	}
}

func careerLink(name string) (string, string) {
	links := map[string][2]string{
		"360":               {"360招聘官网", "https://campus.360.cn"},
		"ABB":               {"ABB招聘官网", "https://careers.abb/china/zh/home"},
		"AMD":               {"AMD Careers", "https://careers.amd.com/careers-home/jobs"},
		"Accenture":         {"埃森哲中国招聘", "https://www.accenture.cn/cn-zh/careers/local/accenture-china-campus-page"},
		"Adobe":             {"Adobe Careers", "https://careers.adobe.com"},
		"Amazon":            {"Amazon Jobs", "https://www.amazon.jobs/content/zh/locations/china/beijing"},
		"Apple":             {"Apple中国招聘", "https://jobs.apple.com/zh-cn/search?location=china-CHNC"},
		"Atlassian":         {"Atlassian Careers", "https://www.atlassian.com/company/careers"},
		"Autodesk":          {"Autodesk Careers", "https://www.autodesk.com/careers"},
		"BMW Group":         {"BMW Group Careers", "https://jobs.bmwgroup.com"},
		"BOSS直聘":            {"BOSS直聘招聘", "https://www.zhipin.com/school/shixi"},
		"Bosch":             {"博世中国招聘", "https://www.bosch.com.cn/careers/job-offers"},
		"B站":                {"哔哩哔哩招聘", "https://jobs.bilibili.com/social/positions"},
		"Cisco":             {"Cisco Careers", "https://careers.cisco.com"},
		"Cloudflare":        {"Cloudflare Careers", "https://www.cloudflare.com/careers/jobs"},
		"Dassault Systemes": {"Dassault Systèmes Careers", "https://www.3ds.com/careers"},
		"Databricks":        {"Databricks Careers", "https://www.databricks.com/company/careers/open-positions"},
		"DeepSeek":          {"DeepSeek招聘", "https://talent.deepseek.com"},
		"Dell Technologies": {"Dell Careers", "https://enterpriseplatform.dell.com/hcmUI/CandidateExperience/en/sites/careers"},
	}
	link, ok := links[name]
	if !ok {
		return "", ""
	}
	return link[0], link[1]
}

func verified(id, name, english, typ, industry, city, summary, size, background, website string, sources []company.Source) company.Company {
	return company.Company{ID: id, Name: name, EnglishName: english, Type: typ, Industry: industry, City: city, Description: summary, EmployeeSize: size, Background: background, Website: website, DataStatus: "verified_multi_source", Tags: []string{typ, industry, "多方资料已核验"}, Sources: sources}
}
