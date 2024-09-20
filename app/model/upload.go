package model

//将word文档类容录入数据库
type UploadDoc struct {
	//表格0
	CompanyName string `orm:"CompanyName"   json:"CompanyName"` // 公司名称
	FBD         string `orm:"FBD"   json:"FBD"`
	Region      string `orm:"Region"   json:"Region"`
	VentureLab  string `orm:"VentureLab"   json:"VentureLab"`
	TeamSize    string `orm:"TeamSize"   json:"TeamSize"`
	LastUpdate  string `orm:"LastUpdate"   json:"LastUpdate"`

	//表格1
	InvestmentThesis string `orm:"InvestmentThesis"   json:"InvestmentThesis"`

	//表格2 --由于是多条记录，存json
	CapTable string `orm:"CapTable"   json:"CapTable"`

	//表格3 --由于是多条记录，存json
	OwnSharesOrSeatOnBoard string `orm:"OwnSharesOrSeatOnBoard"   json:"OwnSharesOrSeatOnBoard"`

	//表格4 与tencent合作历史
	ProductTitle1          string `orm:"ProductTitle1"   json:"ProductTitle1"`
	CooperationMode1       string `orm:"CooperationMode1"   json:"CooperationMode1"`
	ParticipantTencent1    string `orm:"ParticipantTencent1"   json:"ParticipantTencent1"`
	CooperationPeriodFrom1 string `orm:"CooperationPeriodFrom1"   json:"CooperationPeriodFrom1"`
	CooperationPeriodTo1   string `orm:"CooperationPeriodTo1"   json:"CooperationPeriodTo1"`
	AchievementsProblems1  string `orm:"AchievementsProblems1"   json:"AchievementsProblems1"`

	ProductTitle2          string `orm:"ProductTitle2"   json:"ProductTitle2"`
	CooperationMode2       string `orm:"CooperationMode2"   json:"CooperationMode2"`
	ParticipantTencent2    string `orm:"ParticipantTencent2"   json:"ParticipantTencent2"`
	CooperationPeriodFrom2 string `orm:"CooperationPeriodFrom2"   json:"CooperationPeriodFrom2"`
	CooperationPeriodTo2   string `orm:"CooperationPeriodTo2"   json:"CooperationPeriodTo2"`
	AchievementsProblems2  string `orm:"AchievementsProblems2"   json:"AchievementsProblems2"`

	ProductTitle3          string `orm:"ProductTitle3"   json:"ProductTitle3"`
	CooperationMode3       string `orm:"CooperationMode3"   json:"CooperationMode3"`
	ParticipantTencent3    string `orm:"ParticipantTencent3"   json:"ParticipantTencent3"`
	CooperationPeriodFrom3 string `orm:"CooperationPeriodFrom3"   json:"CooperationPeriodFrom3"`
	CooperationPeriodTo3   string `orm:"CooperationPeriodTo3"   json:"CooperationPeriodTo3"`
	AchievementsProblems3  string `orm:"AchievementsProblems3"   json:"AchievementsProblems3"`

	//表格5
	WeRespond string `orm:"WeRespond"   json:"WeRespond"`

	//表格6 	Company Overall
	CurentFocus      string `orm:"CurentFocus"   json:"CurentFocus"`
	YearBusinessPlan string `orm:"YearBusinessPlan"   json:"YearBusinessPlan"`
	PlanCollaborate  string `orm:"PlanCollaborate"   json:"PlanCollaborate"`

	//表格7 Significant New Hires in the past 12 months  存josn
	SignificantNewHires string `orm:"SignificantNewHires"   json:"SignificantNewHires"`

	//表格8  Team Dynamics
	RiskAndSolution string `orm:"RiskAndSolution"   json:"RiskAndSolution"`
	ReductionPlan   string `orm:"ReductionPlan"   json:"ReductionPlan"`

	//表格9  Current Product Performance
	CurrentProducTitle1          string `orm:"CurrentProducTitle1"   json:"CurrentProducTitle1"`
	CurrentPerformanceHighlight1 string `orm:"CurrentPerformanceHighlight1"   json:"CurrentPerformanceHighlight1"`
	CurrentProducTitle2          string `orm:"CurrentProducTitle2"   json:"CurrentProducTitle2"`
	CurrentPerformanceHighlight2 string `orm:"CurrentPerformanceHighlight2"   json:"CurrentPerformanceHighlight2"`
	CurrentProducTitle3          string `orm:"CurrentProducTitle3"   json:"CurrentProducTitle3"`
	CurrentPerformanceHighlight3 string `orm:"CurrentPerformanceHighlight3"   json:"CurrentPerformanceHighlight3"`

	//表格10  Pipeline Product in Development
	PipelineProductTitle1            string `orm:"PipelineProductTitle1"   json:"PipelineProductTitle1"`
	PipelineProductLaunchDate1       string `orm:"PipelineProductLaunchDate1"   json:"PipelineProductLaunchDate1"`
	PipelineProductGenre1            string `orm:"PipelineProductGenre1"   json:"PipelineProductGenre1"`
	PipelineProductPlatform1         string `orm:"PipelineProductPlatform1"   json:"PipelineProductPlatform1"`
	PipelineProductIp1               string `orm:"PipelineProductIp1"   json:"PipelineProductIp1"`
	PipelineProductCurrentMilestone1 string `orm:"PipelineProductCurrentMilestone1"   json:"PipelineProductCurrentMilestone1"`
	PipelineProductCurrentProcess1   string `orm:"PipelineProductCurrentProcess1"   json:"PipelineProductCurrentProcess1"`
	PipelineProductPublisher1        string `orm:"PipelineProductPublisher1"   json:"PipelineProductPublisher1"`
	PipelineProductFte1              string `orm:"PipelineProductFte1"   json:"PipelineProductFte1"`
	PipelineProductDevCost1          string `orm:"PipelineProductDevCost1"   json:"PipelineProductDevCost1"`
	PipelineProductPerformance1      string `orm:"PipelineProductPerformance1"   json:"PipelineProductPerformance1"`
	PipelineProductRemark1           string `orm:"PipelineProductRemark1"   json:"PipelineProductRemark1"`

	PipelineProductTitle2            string `orm:"PipelineProductTitle2"   json:"PipelineProductTitle2"`
	PipelineProductLaunchDate2       string `orm:"PipelineProductLaunchDate2"   json:"PipelineProductLaunchDate2"`
	PipelineProductGenre2            string `orm:"PipelineProductGenre2"   json:"PipelineProductGenre2"`
	PipelineProductPlatform2         string `orm:"PipelineProductPlatform2"   json:"PipelineProductPlatform2"`
	PipelineProductIp2               string `orm:"PipelineProductIp2"   json:"PipelineProductIp2"`
	PipelineProductCurrentMilestone2 string `orm:"PipelineProductCurrentMilestone2"   json:"PipelineProductCurrentMilestone2"`
	PipelineProductCurrentProcess2   string `orm:"PipelineProductCurrentProcess2"   json:"PipelineProductCurrentProcess2"`
	PipelineProductPublisher2        string `orm:"PipelineProductPublisher2"   json:"PipelineProductPublisher2"`
	PipelineProductFte2              string `orm:"PipelineProductFte2"   json:"PipelineProductFte2"`
	PipelineProductDevCost2          string `orm:"PipelineProductDevCost2"   json:"PipelineProductDevCost2"`
	PipelineProductPerformance2      string `orm:"PipelineProductPerformance2"   json:"PipelineProductPerformance2"`
	PipelineProductRemark2           string `orm:"PipelineProductRemark2"   json:"PipelineProductRemark2"`

	PipelineProductTitle3            string `orm:"PipelineProductTitle3"   json:"PipelineProductTitle3"`
	PipelineProductLaunchDate3       string `orm:"PipelineProductLaunchDate3"   json:"PipelineProductLaunchDate3"`
	PipelineProductGenre3            string `orm:"PipelineProductGenre3"   json:"PipelineProductGenre3"`
	PipelineProductPlatform3         string `orm:"PipelineProductPlatform3"   json:"PipelineProductPlatform3"`
	PipelineProductIp3               string `orm:"PipelineProductIp3"   json:"PipelineProductIp3"`
	PipelineProductCurrentMilestone3 string `orm:"PipelineProductCurrentMilestone3"   json:"PipelineProductCurrentMilestone3"`
	PipelineProductCurrentProcess3   string `orm:"PipelineProductCurrentProcess3"   json:"PipelineProductCurrentProcess3"`
	PipelineProductPublisher3        string `orm:"PipelineProductPublisher3"   json:"PipelineProductPublisher3"`
	PipelineProductFte3              string `orm:"PipelineProductFte3"   json:"PipelineProductFte3"`
	PipelineProductDevCost3          string `orm:"PipelineProductDevCost3"   json:"PipelineProductDevCost3"`
	PipelineProductPerformance3      string `orm:"PipelineProductPerformance3"   json:"PipelineProductPerformance3"`
	PipelineProductRemark3           string `orm:"PipelineProductRemark3"   json:"PipelineProductRemark3"`

	//表格11 Cash Flow (report in $mn)
	CashBalance   string `orm:"CashBalance"   json:"CashBalance"`
	CashBurnRate  string `orm:"CashBurnRate"   json:"CashBurnRate"`
	LasterUpdate  string `orm:"LasterUpdate"   json:"LasterUpdate"`
	FundingNeeded string `orm:"FundingNeeded"   json:"FundingNeeded"`
	AmountNeeded  string `orm:"AmountNeeded"   json:"AmountNeeded"`
	RequiredTime  string `orm:"RequiredTime"   json:"RequiredTime"`
	UseOfFunds    string `orm:"UseOfFunds"   json:"UseOfFunds"`

	//表格12 P&L (report in $mn)
	Revenue2020         string `orm:"Revenue2020"   json:"Revenue2020"`
	Revenue2021H1       string `orm:"Revenue2021H1"   json:"Revenue2021H1"`
	Revenue2021Forecast string `orm:"Revenue2021Forecast"   json:"Revenue2021Forecast"`

	Ebitda2020         string `orm:"Ebitda2020"   json:"Ebitda2020"`
	Ebitda2021H1       string `orm:"Ebitda2021H1"   json:"Ebitda2021H1"`
	Ebitda2020Forecast string `orm:"Ebitda2020Forecast"   json:"Ebitda2020Forecast"`

	//表格13 Support  存json
	//CompanyChallenges string `orm:"CompanyChallenges"   json:"CompanyChallenges"` // 废弃
	CompanyChallenge1 string `orm:"CompanyChallenge1"   json:"CompanyChallenge1"`
	CompanyPriority1  string `orm:"CompanyPriority1"   json:"CompanyPriority1"`
	CompanyChallenge2 string `orm:"CompanyChallenge2"   json:"CompanyChallenge2"`
	CompanyPriority2  string `orm:"CompanyPriority2"   json:"CompanyPriority2"`
	CompanyChallenge3 string `orm:"CompanyChallenge3"   json:"CompanyChallenge3"`
	CompanyPriority3  string `orm:"CompanyPriority3"   json:"CompanyPriority3"`

	//所有类容总和
	AllText string `orm:"AllText"   json:"AllText"`
}
