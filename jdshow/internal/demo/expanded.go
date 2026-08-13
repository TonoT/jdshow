package demo

import (
	"fmt"
	"jdshow/internal/company"
)

func ExpandedCompanies() []company.Company {
	out := make([]company.Company, 0, 140)
	privateNames := []string{"上海悦商信息科技", "上海智链云科", "上海云帆互联", "上海数桥科技", "上海星图网络", "上海启明软件", "上海凌云数据", "上海新锐智能", "上海卓越云创", "上海云栈科技", "上海数启信息", "上海智汇网络", "上海极客互联", "上海云脉科技", "上海睿联科技", "上海数舟软件", "上海智谷信息", "上海云鹤科技", "上海卓数科技", "上海启航互联", "上海视界智能", "上海链动科技", "上海麦穗网络", "上海灵犀数据", "上海云图软件", "上海橙果科技", "上海青禾信息", "上海蓝鲸网络", "上海知行科技", "上海远见数据", "上海新程软件", "上海睿思智能", "上海繁星互联", "上海创客云", "上海方舟数据", "上海飞鱼科技", "上海云杉软件", "上海明日网络", "上海数联智能", "上海星河科技", "上海万象互联", "上海智策数据", "上海启辰软件", "上海云翼网络", "上海深蓝智能", "上海原力科技", "上海极光数据", "上海云途信息", "上海智源互联", "上海新境科技", "上海数智未来", "上海云帆数据", "上海互联新创", "上海智能工场", "上海码上科技", "上海云创空间", "上海数科在线", "上海悦动网络", "上海极客云图", "上海智联新媒", "上海链云科技", "上海橙子互联", "上海新火科技", "上海云端智造", "上海数融科技", "上海慧谷网络", "上海星辰软件", "上海风行数据", "上海云象科技", "上海智见信息", "上海蓝海互联", "上海新途软件", "上海数梦智能", "上海云上未来", "上海睿云科技", "上海启点网络", "上海联数科技", "上海知微智能", "上海云科在线", "上海新翼互联", "上海极数科技", "上海云桥信息", "上海智绘网络", "上海数聚科技", "上海启智云", "上海航迹数据", "上海云辰科技", "上海未来互联", "上海数航软件", "上海智核科技", "上海新锐数据", "上海云雀互联", "上海数见科技", "上海启元智能", "上海极致网络", "上海云麓科技", "上海智潮信息", "上海新联软件", "上海数云空间", "上海云海互联"}
	for i, n := range privateNames {
		out = append(out, candidate(fmt.Sprintf("private-new-%03d", i+1), n, "互联网 / AI", "上海互联网与科技", "上海"))
	}
	stateNames := []string{"上海数据集团", "上海联和投资", "上海仪电", "上海国有资本数科", "上海临港科技", "上海城投数科", "上海申通地铁科技", "上海机场信息网络", "上海国际港务数字科技", "浦东发展集团数科", "张江集团科技", "上海国盛数字科技", "上海久事智慧交通", "上海华谊信息科技", "上海建工数科", "上海隧道股份数字科技", "上海联交所信息科技", "上海环境数字科技", "上海文广科技", "上海东方网股份"}
	for i, n := range stateNames {
		out = append(out, candidate(fmt.Sprintf("state-new-%03d", i+1), n, "国企 / 央企科技", "数字基础设施与城市科技", "上海"))
	}
	foreignNames := []string{"Adobe", "Salesforce", "VMware", "Red Hat", "ServiceNow", "Snowflake", "Databricks", "Cloudflare", "Palantir", "Atlassian", "GitLab", "MongoDB", "Elastic", "Twilio", "Zoom", "Autodesk", "Dassault Systemes", "Honeywell", "ABB", "Rockwell Automation"}
	for i, n := range foreignNames {
		out = append(out, candidate(fmt.Sprintf("foreign-new-%03d", i+1), n, "外企科技", "软件、云计算与工业科技", "上海"))
	}
	return out
}
func candidate(id, name, typ, industry, city string) company.Company {
	return company.Company{ID: id, Name: name, EnglishName: name, Type: typ, Industry: industry, City: city, Description: name + " 已纳入公司候选库，主营业务和产品信息待官网核验。", EmployeeSize: "待核验", Background: typ + "候选公司，工商、股权和上市背景待核验", DataStatus: "candidate_pending_verification", Tags: []string{typ, industry, "来源待核验"}}
}
