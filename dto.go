package goalipay

type ErrorResponseDto struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	SubCode string `json:"sub_code,omitempty"`
	SubMsg  string `json:"sub_msg,omitempty"`
}

// 退款查询、交易查询 返回
type TradeFundBill struct {
	FundChannel string `json:"fund_channel,omitempty"` // 同步通知里是 fund_channel
	Amount      string `json:"amount,omitempty"`
	RealAmount  string `json:"real_amount,omitempty"`
	FundType    string `json:"fund_type,omitempty"`
}
