package types

// AssetValuationDetails 账户类型资产估值明细
type AssetValuationDetails struct {
	Funding string `json:"funding"` // 资金账户
	Trading string `json:"trading"` // 交易账户
	Classic string `json:"classic"` // 经典账户（已废弃）
	Earn    string `json:"earn"`    // 金融账户
}

// AssetValuation 资产估值响应
type AssetValuation struct {
	TotalBal string                `json:"totalBal"` // 账户总资产估值
	Ts       string                `json:"ts"`       // 更新时间戳
	Details  AssetValuationDetails `json:"details"`  // 账户类型明细
}
