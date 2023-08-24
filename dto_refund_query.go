package goalipay

type RefundRoyalty struct {
	RefundAmount  string `json:"refund_amount,omitempty"`
	RoyaltyType   string `json:"royalty_type,omitempty"`
	ResultCode    string `json:"result_code,omitempty"`
	TransOut      string `json:"trans_out,omitempty"`
	TransOutEmail string `json:"trans_out_email,omitempty"`
	TransIn       string `json:"trans_in,omitempty"`
	TransInEmail  string `json:"trans_in_email,omitempty"`
}

type DepositBackInfo struct {
	HasDepositBack     string `json:"has_deposit_back,omitempty"`
	DbackStatus        string `json:"dback_status,omitempty"`
	DbackAmount        string `json:"dback_amount,omitempty"`
	BankAckTime        string `json:"bank_ack_time,omitempty"`
	EstBankReceiptTime string `json:"est_bank_receipt_time,omitempty"`
}

type TradeRefundQueryDto struct {
	ErrorResponseDto
	TradeNo              string           `json:"trade_no,omitempty"`
	OutTradeNo           string           `json:"out_trade_no,omitempty"`
	OutRequestNo         string           `json:"out_request_no,omitempty"`
	RefundReason         string           `json:"refund_reason,omitempty"`
	TotalAmount          string           `json:"total_amount,omitempty"`
	RefundAmount         string           `json:"refund_amount,omitempty"`
	RefundStatus         string           `json:"refund_status,omitempty"`
	RefundRoyaltys       []*RefundRoyalty `json:"refund_royaltys,omitempty"`
	GmtRefundPay         string           `json:"gmt_refund_pay,omitempty"`
	RefundDetailItemList []*TradeFundBill `json:"refund_detail_item_list,omitempty"`
	SendBackFee          string           `json:"send_back_fee,omitempty"`
	DepositBackInfo      *DepositBackInfo `json:"deposit_back_info,omitempty"`
}

// 退款查询返回
type TradeFastpayRefundQueryResponseDto struct {
	Response     *TradeRefundQueryDto `json:"alipay_trade_fastpay_refund_query_response"`
	AlipayCertSn string               `json:"alipay_cert_sn,omitempty"`
	SignData     string               `json:"-"`
	Sign         string               `json:"sign"`
}
