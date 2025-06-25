package types

// Instrument defines the structure for a tradable instrument.
type Instrument struct {
	Alias            string `json:"alias"`
	AuctionEndTime   string `json:"auctionEndTime"`
	BaseCcy          string `json:"baseCcy"`
	Category         string `json:"category"`
	CtMult           string `json:"ctMult"`
	CtType           string `json:"ctType"`
	CtVal            string `json:"ctVal"`
	CtValCcy         string `json:"ctValCcy"`
	ContTdSwTime     string `json:"contTdSwTime"`
	ExpTime          string `json:"expTime"`
	FutureSettlement bool   `json:"futureSettlement"`
	InstFamily       string `json:"instFamily"`
	InstID           string `json:"instId"`
	InstType         string `json:"instType"`
	Lever            string `json:"lever"`
	ListTime         string `json:"listTime"`
	LotSz            string `json:"lotSz"`
	MaxIcebergSz     string `json:"maxIcebergSz"`
	MaxLmtAmt        string `json:"maxLmtAmt"`
	MaxLmtSz         string `json:"maxLmtSz"`
	MaxMktAmt        string `json:"maxMktAmt"`
	MaxMktSz         string `json:"maxMktSz"`
	MaxStopSz        string `json:"maxStopSz"`
	MaxTriggerSz     string `json:"maxTriggerSz"`
	MaxTwapSz        string `json:"maxTwapSz"`
	MinSz            string `json:"minSz"`
	OptType          string `json:"optType"`
	OpenType         string `json:"openType"`
	QuoteCcy         string `json:"quoteCcy"`
	RuleType         string `json:"ruleType"`
	SettleCcy        string `json:"settleCcy"`
	State            string `json:"state"`
	Stk              string `json:"stk"`
	TickSz           string `json:"tickSz"`
	Uly              string `json:"uly"`
}
