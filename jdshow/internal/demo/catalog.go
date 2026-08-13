package demo

import (
	"jdshow/internal/company"
	"jdshow/internal/job"
	"strconv"
	"time"
)

func Catalog(companies map[string]company.Company, jobs []job.Job, now time.Time) (map[string]company.Company, []job.Job) {
	groups := []struct {
		prefix, typ, industry string
		names                 []string
	}{
		{"tech", "互联网 / AI", "互联网与 AI", []string{"阿里巴巴", "腾讯", "字节跳动", "美团", "京东", "小米", "快手", "网易", "拼多多", "小红书", "滴滴", "携程", "科大讯飞", "商汤科技", "旷视科技", "智谱 AI", "零一万物", "MiniMax", "百川智能", "DeepSeek"}},
		{"state", "国企 / 央企科技", "云计算与数字基础设施", []string{"中国移动", "中国联通", "中国电子", "中国电科", "国家电网", "中国石油", "中国石化", "中国海油", "航天科技", "航天科工", "航空工业", "中国船舶", "中国兵器", "中国国家铁路", "中国建筑", "中国交建", "中国核工业", "中国广核", "国家能源集团", "浪潮集团"}},
		{"intl", "外企科技", "软件、云计算与工业科技", []string{"Google", "Amazon", "Apple", "IBM", "Oracle", "SAP", "Siemens", "Intel", "NVIDIA", "AMD", "Qualcomm", "Cisco", "Dell Technologies", "Accenture", "Bosch", "BMW Group", "Mercedes-Benz", "Toyota", "Ericsson", "Schneider Electric"}},
		{"private-extra", "互联网 / AI", "上海互联网与科技", []string{"上海悦商信息科技"}},
	}
	cities := []string{"北京", "上海", "深圳", "杭州", "广州", "成都"}
	cats := []string{"后端开发", "算法 / AI", "云计算 / 基础设施", "数据工程", "产品技术"}
	for _, g := range groups {
		for i, name := range g.names {
			id := g.prefix + "-" + strconv.Itoa(i+1)
			if _, ok := companies[id]; ok {
				continue
			}
			city := cities[i%len(cities)]
			if g.prefix == "private-extra" {
				city = "上海"
			}
			companies[id] = company.Company{ID: id, Name: name, EnglishName: name, Type: g.typ, Industry: g.industry, City: city, Description: name + " 已纳入来源候选池，主要业务和产品信息待官网核验。", EmployeeSize: "待核验", Background: g.typ + "候选公司，工商与股权背景待核验", Tags: []string{g.typ, g.industry, "来源待核验"}, DataStatus: "candidate_pending_verification"}
			jobs = append(jobs, job.Job{ID: "catalog-" + id, Title: cats[i%len(cats)] + " - " + name, CompanyID: id, City: cities[i%len(cities)], Category: cats[i%len(cats)], Experience: "经验要求待核验", Education: "学历要求待核验", Salary: "薪资待核验", Status: "NEW", PublishedAt: now.Add(-time.Duration(i+1) * 12 * time.Hour), UpdatedAt: now.Add(-time.Duration(i+1) * 12 * time.Hour), SourceName: "官方招聘入口待核验", Description: "该岗位为来源目录展示记录，真实职责和任职要求需从官方招聘页面核验后替换。", Requirements: []string{"官方岗位详情待核验"}, Skills: []string{"来源待核验"}})
		}
	}
	return companies, jobs
}
