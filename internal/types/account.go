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
