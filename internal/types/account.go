package types

type AccountBalanceDetail struct {
	Ccy                   string `json:"ccy"`
	AvailBal              string `json:"availBal"`
	AvailEq               string `json:"availEq"`
	BorrowFroz            string `json:"borrowFroz"`
	CashBal               string `json:"cashBal"`
	CrossLiab             string `json:"crossLiab"`
	CollateralEnabled     bool   `json:"collateralEnabled"`
	CollateralRestrict    bool   `json:"collateralRestrict"`
	ColBorrAutoConversion string `json:"colBorrAutoConversion"`
	DisEq                 string `json:"disEq"`
	Eq                    string `json:"eq"`
	EqUsd                 string `json:"eqUsd"`
	SmtSyncEq             string `json:"smtSyncEq"`
	SpotCopyTradingEq     string `json:"spotCopyTradingEq"`
	FixedBal              string `json:"fixedBal"`
	FrozenBal             string `json:"frozenBal"`
	Imr                   string `json:"imr"`
	Interest              string `json:"interest"`
	IsoEq                 string `json:"isoEq"`
	IsoLiab               string `json:"isoLiab"`
	IsoUpl                string `json:"isoUpl"`
	Liab                  string `json:"liab"`
	MaxLoan               string `json:"maxLoan"`
	MgnRatio              string `json:"mgnRatio"`
	Mmr                   string `json:"mmr"`
	NotionalLever         string `json:"notionalLever"`
	OrdFrozen             string `json:"ordFrozen"`
	RewardBal             string `json:"rewardBal"`
	SpotInUseAmt          string `json:"spotInUseAmt"`
	ClSpotInUseAmt        string `json:"clSpotInUseAmt"`
	MaxSpotInUse          string `json:"maxSpotInUse"`
	SpotIsoBal            string `json:"spotIsoBal"`
	StgyEq                string `json:"stgyEq"`
	Twap                  string `json:"twap"`
	UTime                 string `json:"uTime"`
	Upl                   string `json:"upl"`
	UplLiab               string `json:"uplLiab"`
	SpotBal               string `json:"spotBal"`
	OpenAvgPx             string `json:"openAvgPx"`
	AccAvgPx              string `json:"accAvgPx"`
	SpotUpl               string `json:"spotUpl"`
	SpotUplRatio          string `json:"spotUplRatio"`
	TotalPnl              string `json:"totalPnl"`
	TotalPnlRatio         string `json:"totalPnlRatio"`
}

type AccountBalance struct {
	AdjEq                 string                 `json:"adjEq"`
	AvailEq               string                 `json:"availEq"`
	BorrowFroz            string                 `json:"borrowFroz"`
	Details               []AccountBalanceDetail `json:"details"`
	Imr                   string                 `json:"imr"`
	IsoEq                 string                 `json:"isoEq"`
	MgnRatio              string                 `json:"mgnRatio"`
	Mmr                   string                 `json:"mmr"`
	NotionalUsd           string                 `json:"notionalUsd"`
	NotionalUsdForBorrow  string                 `json:"notionalUsdForBorrow"`
	NotionalUsdForFutures string                 `json:"notionalUsdForFutures"`
	NotionalUsdForOption  string                 `json:"notionalUsdForOption"`
	NotionalUsdForSwap    string                 `json:"notionalUsdForSwap"`
	OrdFroz               string                 `json:"ordFroz"`
	TotalEq               string                 `json:"totalEq"`
	UTime                 string                 `json:"uTime"`
	Upl                   string                 `json:"upl"`
}

type AccountPosition struct {
	Adl                    string        `json:"adl"`
	AvailPos               string        `json:"availPos"`
	AvgPx                  string        `json:"avgPx"`
	BaseBal                string        `json:"baseBal"`
	BaseBorrowed           string        `json:"baseBorrowed"`
	BaseInterest           string        `json:"baseInterest"`
	BePx                   string        `json:"bePx"`
	BizRefId               string        `json:"bizRefId"`
	BizRefType             string        `json:"bizRefType"`
	CTime                  string        `json:"cTime"`
	Ccy                    string        `json:"ccy"`
	ClSpotInUseAmt         string        `json:"clSpotInUseAmt"`
	CloseOrderAlgo         []interface{} `json:"closeOrderAlgo"`
	DeltaBS                string        `json:"deltaBS"`
	DeltaPA                string        `json:"deltaPA"`
	Fee                    string        `json:"fee"`
	FundingFee             string        `json:"fundingFee"`
	GammaBS                string        `json:"gammaBS"`
	GammaPA                string        `json:"gammaPA"`
	IdxPx                  string        `json:"idxPx"`
	Imr                    string        `json:"imr"`
	InstId                 string        `json:"instId"`
	InstType               string        `json:"instType"`
	Interest               string        `json:"interest"`
	Last                   string        `json:"last"`
	Lever                  string        `json:"lever"`
	Liab                   string        `json:"liab"`
	LiabCcy                string        `json:"liabCcy"`
	LiqPenalty             string        `json:"liqPenalty"`
	LiqPx                  string        `json:"liqPx"`
	Margin                 string        `json:"margin"`
	MarkPx                 string        `json:"markPx"`
	MaxSpotInUseAmt        string        `json:"maxSpotInUseAmt"`
	MgnMode                string        `json:"mgnMode"`
	MgnRatio               string        `json:"mgnRatio"`
	Mmr                    string        `json:"mmr"`
	NotionalUsd            string        `json:"notionalUsd"`
	OptVal                 string        `json:"optVal"`
	PendingCloseOrdLiabVal string        `json:"pendingCloseOrdLiabVal"`
	Pnl                    string        `json:"pnl"`
	Pos                    string        `json:"pos"`
	PosCcy                 string        `json:"posCcy"`
	PosId                  string        `json:"posId"`
	PosSide                string        `json:"posSide"`
	QuoteBal               string        `json:"quoteBal"`
	QuoteBorrowed          string        `json:"quoteBorrowed"`
	QuoteInterest          string        `json:"quoteInterest"`
	RealizedPnl            string        `json:"realizedPnl"`
	SpotInUseAmt           string        `json:"spotInUseAmt"`
	SpotInUseCcy           string        `json:"spotInUseCcy"`
	ThetaBS                string        `json:"thetaBS"`
	ThetaPA                string        `json:"thetaPA"`
	TradeId                string        `json:"tradeId"`
	UTime                  string        `json:"uTime"`
	Upl                    string        `json:"upl"`
	UplLastPx              string        `json:"uplLastPx"`
	UplRatio               string        `json:"uplRatio"`
	UplRatioLastPx         string        `json:"uplRatioLastPx"`
	UsdPx                  string        `json:"usdPx"`
	VegaBS                 string        `json:"vegaBS"`
	VegaPA                 string        `json:"vegaPA"`
	NonSettleAvgPx         string        `json:"nonSettleAvgPx"`
	SettledPnl             string        `json:"settledPnl"`
}

type AccountPositionRisk struct {
	AdjEq   string                `json:"adjEq"`
	BalData []AccountRiskBalance  `json:"balData"`
	PosData []AccountRiskPosition `json:"posData"`
	Ts      string                `json:"ts"`
}

type AccountRiskBalance struct {
	Ccy   string `json:"ccy"`
	Eq    string `json:"eq"`
	DisEq string `json:"disEq"`
}

type AccountRiskPosition struct {
	BaseBal     string `json:"baseBal"`
	Ccy         string `json:"ccy"`
	InstId      string `json:"instId"`
	InstType    string `json:"instType"`
	MgnMode     string `json:"mgnMode"`
	NotionalCcy string `json:"notionalCcy"`
	NotionalUsd string `json:"notionalUsd"`
	Pos         string `json:"pos"`
	PosCcy      string `json:"posCcy"`
	PosId       string `json:"posId"`
	PosSide     string `json:"posSide"`
	QuoteBal    string `json:"quoteBal"`
}

type AccountConfig struct {
	AcctLv              string   `json:"acctLv"`
	AcctStpMode         string   `json:"acctStpMode"`
	AutoLoan            bool     `json:"autoLoan"`
	CtIsoMode           string   `json:"ctIsoMode"`
	EnableSpotBorrow    bool     `json:"enableSpotBorrow"`
	GreeksType          string   `json:"greeksType"`
	IP                  string   `json:"ip"`
	Type                string   `json:"type"`
	KycLv               string   `json:"kycLv"`
	Label               string   `json:"label"`
	Level               string   `json:"level"`
	LevelTmp            string   `json:"levelTmp"`
	LiquidationGear     string   `json:"liquidationGear"`
	MainUid             string   `json:"mainUid"`
	MgnIsoMode          string   `json:"mgnIsoMode"`
	OpAuth              string   `json:"opAuth"`
	Perm                string   `json:"perm"`
	PosMode             string   `json:"posMode"`
	RoleType            string   `json:"roleType"`
	SpotBorrowAutoRepay bool     `json:"spotBorrowAutoRepay"`
	SpotOffsetType      string   `json:"spotOffsetType"`
	SpotRoleType        string   `json:"spotRoleType"`
	SpotTraderInsts     []string `json:"spotTraderInsts"`
	TraderInsts         []string `json:"traderInsts"`
	Uid                 string   `json:"uid"`
}
