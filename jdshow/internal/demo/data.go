package demo

import (
	"jdshow/internal/company"
	"jdshow/internal/job"
	"time"
)

func Seed(now time.Time) (map[string]company.Company, []job.Job) {
	companies := map[string]company.Company{
		"baidu":     {ID: "baidu", Name: "百度", EnglishName: "Baidu", Type: "互联网 / AI", Industry: "搜索、AI、云计算", City: "北京", Description: "以搜索、智能云和人工智能为核心能力的科技公司。", Tags: []string{"AI", "智能云", "互联网"}, Website: "https://www.baidu.com"},
		"moonshot":  {ID: "moonshot", Name: "月之暗面", EnglishName: "Moonshot AI", Type: "AI 公司", Industry: "大模型、智能应用", City: "北京", Description: "专注于大模型与智能应用探索的 AI 公司。", Tags: []string{"大模型", "AI", "智能应用"}, Website: "https://www.moonshot.cn"},
		"telecom":   {ID: "telecom", Name: "中国电信", EnglishName: "China Telecom", Type: "央企科技", Industry: "云计算、通信、算力", City: "全国", Description: "面向云改数转智惠持续建设云网融合和数字基础设施。", Tags: []string{"央企", "云计算", "基础设施"}, Website: "https://www.chinatelecom.com.cn"},
		"microsoft": {ID: "microsoft", Name: "Microsoft", EnglishName: "Microsoft", Type: "外企科技", Industry: "软件、云计算、AI", City: "上海", Description: "面向全球用户提供软件、云服务和人工智能产品。", Tags: []string{"外企", "云计算", "AI"}, Website: "https://www.microsoft.com"},
	}
	jobs := []job.Job{
		{ID: "job-001", Title: "高级后端工程师 - 智能云", CompanyID: "baidu", City: "北京", Category: "后端开发", Experience: "3-8 年", Education: "本科", Salary: "30-55K·15薪", Status: "ACTIVE", PublishedAt: now.Add(-6 * time.Hour), UpdatedAt: now.Add(-2 * time.Hour), SourceName: "百度官方招聘", SourceURL: "https://talent.baidu.com/jobs/social-list", Description: "负责智能云核心服务的架构设计、研发与稳定性建设，参与高并发分布式系统的持续演进。", Requirements: []string{"熟悉 Go 或 C++，具备扎实的数据结构与系统设计能力", "有分布式系统、云原生或基础设施研发经验", "重视工程质量，能够推动跨团队协作"}, Skills: []string{"Go", "分布式系统", "Kubernetes", "高并发"}},
		{ID: "job-002", Title: "大模型算法工程师", CompanyID: "moonshot", City: "北京", Category: "算法 / AI", Experience: "1-5 年", Education: "硕士", Salary: "35-70K", Status: "NEW", PublishedAt: now.Add(-18 * time.Hour), UpdatedAt: now.Add(-18 * time.Hour), SourceName: "月之暗面官方招聘", SourceURL: "https://careers.kimi.com/", Description: "参与大语言模型训练、评测和推理优化，将前沿研究转化为可靠的模型能力。", Requirements: []string{"熟悉 Transformer、预训练或微调方法", "掌握 Python 和至少一种深度学习框架", "有大模型、NLP 或多模态项目经验"}, Skills: []string{"LLM", "PyTorch", "NLP", "模型训练"}},
		{ID: "job-003", Title: "云平台研发工程师", CompanyID: "telecom", City: "广州", Category: "云计算 / 基础设施", Experience: "应届 / 1-3 年", Education: "本科", Salary: "15-30K", Status: "ACTIVE", PublishedAt: now.Add(-2 * 24 * time.Hour), UpdatedAt: now.Add(-24 * time.Hour), SourceName: "中国电信官方招聘", SourceURL: "https://www.chinatelecom.com.cn/ct/zp/165144.html", Description: "参与云平台、智能云网和算力基础设施的研发与交付，服务企业数字化转型。", Requirements: []string{"熟悉 Go、Java 或 Python 之一", "了解 Linux、容器和网络基础", "具备良好的问题分析与沟通能力"}, Skills: []string{"Go", "云计算", "Linux", "容器"}},
		{ID: "job-004", Title: "Software Engineer, Azure AI", CompanyID: "microsoft", City: "上海", Category: "算法 / AI", Experience: "3-8 年", Education: "本科", Salary: "面议", Status: "UPDATED", PublishedAt: now.Add(-3 * 24 * time.Hour), UpdatedAt: now.Add(-5 * time.Hour), SourceName: "Microsoft Careers", SourceURL: "https://careers.microsoft.com/v2/global/en/locations/gcr.html", Description: "Build reliable AI platform capabilities for enterprise customers across the Asia Pacific region.", Requirements: []string{"Strong software engineering fundamentals and experience with cloud services", "Experience with machine learning systems or distributed infrastructure", "Professional English communication skills"}, Skills: []string{"Azure", "Python", "Machine Learning", "Distributed Systems"}},
		{ID: "job-005", Title: "数据工程师 - 推荐平台", CompanyID: "baidu", City: "深圳", Category: "数据工程", Experience: "3-8 年", Education: "本科", Salary: "25-45K·14薪", Status: "ACTIVE", PublishedAt: now.Add(-5 * 24 * time.Hour), UpdatedAt: now.Add(-4 * 24 * time.Hour), SourceName: "百度官方招聘", SourceURL: "https://talent.baidu.com/jobs/social-list", Description: "建设面向推荐和搜索业务的数据平台，提升数据质量、实时性和可观测性。", Requirements: []string{"熟悉 SQL、数据仓库和流式计算", "有大规模数据处理经验", "能够用工程化方法解决数据质量问题"}, Skills: []string{"SQL", "数据仓库", "Flink", "实时计算"}},
		{ID: "job-006", Title: "技术产品经理 - AI 应用", CompanyID: "moonshot", City: "杭州", Category: "产品技术", Experience: "3-8 年", Education: "本科", Salary: "30-50K", Status: "ACTIVE", PublishedAt: now.Add(-8 * 24 * time.Hour), UpdatedAt: now.Add(-7 * 24 * time.Hour), SourceName: "月之暗面官方招聘", SourceURL: "https://careers.kimi.com/", Description: "围绕 AI 原生产品设计用户体验和增长路径，推动模型能力落地为可用的产品。", Requirements: []string{"理解 AI 产品的技术边界和用户需求", "能够与算法、工程和设计团队协作", "具备清晰的产品文档和项目推进能力"}, Skills: []string{"AI 产品", "用户研究", "项目管理", "Prompt"}},
	}
	return companies, jobs
}
