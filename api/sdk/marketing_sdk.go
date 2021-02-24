package sdk

type MarketingSDK interface {
	Auth // 授权接口
	//ADDelivery
	//Account
	Report
}

type ADDelivery interface {
	BudgetOperation
	CampaignOperation
	ADGroupOperation
	CreativeOperation
}

type BudgetOperation interface {
	GetBudget(input interface{}) (interface{}, error)
	UpdateBudget(input interface{}) (interface{}, error)
}

type CampaignOperation interface {
}

type ADGroupOperation interface {
}

type CreativeOperation interface {
}

type Account interface {
}

type Report interface {
	GetReport(reportInput *GetReportInput) (*GetReportOutput, error)
}