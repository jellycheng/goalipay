package goalipay

type BalanceQueryDto struct {
	ErrorResponseDto
	FreezeAmount    string `json:"freeze_amount"`           //冻结金额,单位（元）
	TotalAmount     string `json:"total_amount"`            //支付宝账户余额,单位（元）
	AvailableAmount string `json:"available_amount"`        //账户可用余额,单位（元）
	SettleAmount    string `json:"settle_amount,omitempty"` //待结算金额，单位（元）,可选
}

// 账户余额查询响应参数
type BalanceQueryResponseDto struct {
	Response     *BalanceQueryDto `json:"alipay_data_bill_balance_query_response"`
	AlipayCertSn string           `json:"alipay_cert_sn,omitempty"`
	SignData     string           `json:"-"`
	Sign         string           `json:"sign"`
}
