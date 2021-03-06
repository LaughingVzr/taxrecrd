package model

import (
	"github.com/astaxie/beego/orm"
)

/*工商机构表model*/
type Organize struct {
	OrgId        int    `orm:"auto"` // 自增id
	OrgUuid      string // 逻辑id
	OrgCreateT   string `orm:"column(org_create_time)"` // 记录创建时间
	OrgModifyT   string `orm:"column(org_modify_time)"` // 记录修改时间
	OrgStatus    int    `orm:"default(1)"`              // 记录状态(1.正常,0.删除)
	OrgSerialNum string // 纳税机构识别码
	OrgName      string // 纳税机构名称
	OrgLegal     string // 纳税机构法人
	OrgRegT      string `orm:"column(org_reg_time)"` // 纳税机构注册时间
	OrgAddr      string `orm:"column(org_address)"`  // 机构地址
	OrgRegCap    int    // 机构注册资金
	OrgTaxOffice string // 机构主管税务机关
	OrgIndus     string `orm:"column(org_industry)"` // 机构所属行业
}

/*年度税务记录*/
type TaxRecord struct {
	TaxId         int     `orm:"auto"` // 自增id
	TaxUuid       string  // 逻辑id
	TaxCreateT    string  `orm:"column(tax_create_time)"` // 记录创建时间
	TaxModifyT    string  `orm:"column(tax_modify_time)"` // 记录修改时间
	TaxStatus     int     `orm:"default(1)"`              // 记录状态(1.正常,0.删除)
	TaxOrgSeriNum string  // 纳税机构识别码
	TaxOrgUuid    string  // 纳税机构逻辑id
	TaxTime       string  // 纳税年份
	TaxIncome     float32 // 应纳税收入
	TaxVat        float32 // 增值税
	TaxIncomeVal  float32 // 应缴税所得额
	TaxIncomeTax  float32 // 应缴纳所得税
}

/*N年纳税统计*/
type TaxStat struct {
	StatId         int     `orm:"auto"` // 自增id
	StatUuid       string  // 逻辑id
	StatCreateT    string  `orm:"column(Stat_create_time)"` // 记录创建时间
	StatModifyT    string  `orm:"column(Stat_modify_time)"` // 记录修改时间
	StatStatus     int     `orm:"default(1)"`               // 记录状态(1.正常,0.删除)
	StatOrgSeriNum string  // 纳税机构识别码
	StatTaxSum     float32 // 应纳税总额
	StatTaxAvg     float32 // 平均应纳税额
	StatIsCheck    int     // 是否检查
	StatCheckT     string  `orm:"column(stat_check_time)"` // 检查时间
}

/*一张表搞定源文件字段(增值纳税人)*/
type TaxRecordRef struct {
	Id            int     `orm:"auto"` // 自增ID
	OrgName       string  // 纳税机构名称
	OrgSerialNum  string  // 纳税机构识别码
	OrgIndus      string  `orm:"column(org_industry)"` // 机构所属行业
	OrgBusScope   string  // 经营范围
	OrgLegal      string  // 纳税机构法人
	OrgRegT       int64   `orm:"column(org_reg_time)"` // 纳税机构注册时间
	OrgAddr       string  `orm:"column(org_address)"`  // 生产经营地址
	OrgRegCap     float64 // 机构注册资本
	OrgTaxOffice  string  // 机构主管税务机关
	OrgIsExport   string  // 是否是出口企业
	TaxIncome1    float64 // 应纳税收入
	TaxExIncome1  float64 // 出口收入
	TaxVat1       float64 // 增值税
	TaxIncomeTax1 float64 // 所得税
	TaxSum1       float64 // 本年度税额总和
	TaxIncome2    float64 // 应纳税收入
	TaxExIncome2  float64 // 出口收入
	TaxVat2       float64 // 增值税
	TaxIncomeTax2 float64 // 所得税
	TaxSum2       float64 // 本年度税额总和
	TaxIncome3    float64 // 应纳税收入
	TaxExIncome3  float64 // 出口收入
	TaxVat3       float64 // 增值税
	TaxIncomeTax3 float64 // 应缴纳所得税
	TaxSum3       float64 // 本年度税额总和
	StatTaxSum    float64 // 三年纳税统计
	StatCheckT    int64   `orm:"column(stat_check_time)"` // 稽查时间
	StatYear      string  `orm:"column(stat_year)"`       // 稽查年度
	IsImportant   string  `orm:"column(is_important)"`    // 是否是重点税源
	Status        int     // 记录状态

}

func init() {
	// 注册模型
	orm.RegisterModel(new(TaxRecordRef))
}
