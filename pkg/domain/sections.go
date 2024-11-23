package domain

const (
	SectionFullAllAccess            = "full_all_access"
	SectionFullCompanyAccess        = "full_company_access"
	SectionFullAccess               = "full_access"
	SectionOrderCardAccess          = "order_card_access"
	SectionProductionDataAccess     = "production_data_access"
	SectionStatusAndCalculateAccess = "status_and_calculate_access"
	SectionPurchasePlanningAccess   = "purchase_planning_access"
)

type Section struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
