package types

// PlaceOrderRequest 下单请求结构体
type PlaceOrderRequest struct {
	InstID     string `json:"instId"`
	TdMode     string `json:"tdMode"`
	Ccy        string `json:"ccy,omitempty"`
	ClOrdID    string `json:"clOrdId,omitempty"`
	Tag        string `json:"tag,omitempty"`
	Side       string `json:"side"`
	PosSide    string `json:"posSide,omitempty"`
	OrdType    string `json:"ordType"`
	Sz         string `json:"sz"`
	Px         string `json:"px,omitempty"`
	ReduceOnly bool   `json:"reduceOnly,omitempty"`
	TgtCcy     string `json:"tgtCcy,omitempty"`
	// 可根据需要添加更多止盈止损等高级参数
}

// PlaceOrderResponse 下单成功后的响应数据
type PlaceOrderResponse struct {
	OrdID   string `json:"ordId"`
	ClOrdID string `json:"clOrdId"`
	Tag     string `json:"tag"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
	Ts      string `json:"ts"`
}

// CancelOrderRequest 撤单请求结构体
type CancelOrderRequest struct {
	InstID  string `json:"instId"`
	OrdID   string `json:"ordId,omitempty"`
	ClOrdID string `json:"clOrdId,omitempty"`
}

// CancelOrderResponse 撤单响应数据
type CancelOrderResponse struct {
	ClOrdID string `json:"clOrdId"`
	OrdID   string `json:"ordId"`
	Ts      string `json:"ts"`
	SCode   string `json:"sCode"`
	SMsg    string `json:"sMsg"`
}

// PendingOrdersResponse 未完成订单响应数据
type PendingOrdersResponse struct {
	AccFillSz          string              `json:"accFillSz"`          // 累计成交数量
	AlgoClOrdId        string              `json:"algoClOrdId"`        // 客户自定义策略订单ID
	AlgoId             string              `json:"algoId"`             // 策略订单ID
	AttachAlgoClOrdId  string              `json:"attachAlgoClOrdId"`  // 下单附带止盈止损，客户自定义的策略订单ID
	AttachAlgoOrds     []map[string]string `json:"attachAlgoOrds"`     // 下单附带止盈止损信息
	AvgPx              string              `json:"avgPx"`              // 成交均价
	CTime              string              `json:"cTime"`              // 订单创建时间
	CancelSource       string              `json:"cancelSource"`       // 订单取消来源
	CancelSourceReason string              `json:"cancelSourceReason"` // 订单取消来源原因
	Category           string              `json:"category"`           // 订单类别
	Ccy                string              `json:"ccy"`                // 保证金币种
	ClOrdId            string              `json:"clOrdId"`            // 客户自定义订单ID
	Fee                string              `json:"fee"`                // 手续费
	FeeCcy             string              `json:"feeCcy"`             // 手续费币种
	FillPx             string              `json:"fillPx"`             // 最新成交价格
	FillSz             string              `json:"fillSz"`             // 最新成交数量
	FillTime           string              `json:"fillTime"`           // 最新成交时间
	InstId             string              `json:"instId"`             // 产品ID
	InstType           string              `json:"instType"`           // 产品类型
	IsTpLimit          string              `json:"isTpLimit"`          // 是否限价止盈
	Lever              string              `json:"lever"`              // 杠杆倍数
	LinkedAlgoOrd      map[string]string   `json:"linkedAlgoOrd"`      // 止损订单信息
	OrdId              string              `json:"ordId"`              // 订单ID
	OrdType            string              `json:"ordType"`            // 订单类型
	Pnl                string              `json:"pnl"`                // 收益
	PosSide            string              `json:"posSide"`            // 持仓方向
	Px                 string              `json:"px"`                 // 委托价格
	PxType             string              `json:"pxType"`             // 期权订单价格类型
	PxUsd              string              `json:"pxUsd"`              // 期权订单价格
	PxVol              string              `json:"pxVol"`              // 期权订单的隐含波动率
	QuickMgnType       string              `json:"quickMgnType"`       // 一键借币类型
	Rebate             string              `json:"rebate"`             // 返佣金额
	RebateCcy          string              `json:"rebateCcy"`          // 返佣币种
	ReduceOnly         string              `json:"reduceOnly"`         // 是否仅减仓
	Side               string              `json:"side"`               // 订单方向
	SlOrdPx            string              `json:"slOrdPx"`            // 止损委托价
	SlTriggerPx        string              `json:"slTriggerPx"`        // 止损触发价
	SlTriggerPxType    string              `json:"slTriggerPxType"`    // 止损触发价类型
	Source             string              `json:"source"`             // 订单来源
	State              string              `json:"state"`              // 订单状态
	StpId              string              `json:"stpId"`              // 自成交保护组ID
	StpMode            string              `json:"stpMode"`            // 自成交保护模式
	Sz                 string              `json:"sz"`                 // 委托数量
	Tag                string              `json:"tag"`                // 订单标签
	TdMode             string              `json:"tdMode"`             // 交易模式
	TgtCcy             string              `json:"tgtCcy"`             // 市价单委托数量的类型
	TpOrdPx            string              `json:"tpOrdPx"`            // 止盈委托价
	TpTriggerPx        string              `json:"tpTriggerPx"`        // 止盈触发价
	TpTriggerPxType    string              `json:"tpTriggerPxType"`    // 止盈触发价类型
	TradeId            string              `json:"tradeId"`            // 最新成交ID
	UTime              string              `json:"uTime"`              // 订单状态更新时间
}

// AmendOrderAttachAlgo 修改订单附带止盈止损信息
type AmendOrderAttachAlgo struct {
	AttachAlgoId         string `json:"attachAlgoId,omitempty"`         // 附带止盈止损的订单ID
	AttachAlgoClOrdId    string `json:"attachAlgoClOrdId,omitempty"`    // 客户自定义的策略订单ID
	NewTpTriggerPx       string `json:"newTpTriggerPx,omitempty"`       // 止盈触发价
	NewTpOrdPx           string `json:"newTpOrdPx,omitempty"`           // 止盈委托价
	NewTpOrdKind         string `json:"newTpOrdKind,omitempty"`         // 止盈订单类型
	NewSlTriggerPx       string `json:"newSlTriggerPx,omitempty"`       // 止损触发价
	NewSlOrdPx           string `json:"newSlOrdPx,omitempty"`           // 止损委托价
	NewTpTriggerPxType   string `json:"newTpTriggerPxType,omitempty"`   // 止盈触发价类型
	NewSlTriggerPxType   string `json:"newSlTriggerPxType,omitempty"`   // 止损触发价类型
	Sz                   string `json:"sz,omitempty"`                   // 新的张数
	AmendPxOnTriggerType string `json:"amendPxOnTriggerType,omitempty"` // 是否启用开仓价止损
}

// AmendOrderRequest 修改订单请求结构体
type AmendOrderRequest struct {
	InstId         string                 `json:"instId"`                   // 产品ID
	CxlOnFail      bool                   `json:"cxlOnFail,omitempty"`      // 修改失败时是否自动撤单
	OrdId          string                 `json:"ordId,omitempty"`          // 订单ID
	ClOrdId        string                 `json:"clOrdId,omitempty"`        // 用户自定义订单ID
	ReqId          string                 `json:"reqId,omitempty"`          // 用户自定义修改事件ID
	NewSz          string                 `json:"newSz,omitempty"`          // 修改的新数量
	NewPx          string                 `json:"newPx,omitempty"`          // 修改后的新价格
	NewPxUsd       string                 `json:"newPxUsd,omitempty"`       // 以USD价格进行期权改单
	NewPxVol       string                 `json:"newPxVol,omitempty"`       // 以隐含波动率进行期权改单
	AttachAlgoOrds []AmendOrderAttachAlgo `json:"attachAlgoOrds,omitempty"` // 修改附带止盈止损信息
}

// AmendOrderResponse 修改订单响应数据
type AmendOrderResponse struct {
	ClOrdId string `json:"clOrdId"` // 用户自定义订单ID
	OrdId   string `json:"ordId"`   // 订单ID
	ReqId   string `json:"reqId"`   // 用户自定义修改事件ID
	SCode   string `json:"sCode"`   // 事件执行结果的code，0代表成功
	SMsg    string `json:"sMsg"`    // 事件执行失败时的msg
	Ts      string `json:"ts"`      // 事件执行时间
}
