package goalipay

type TradeSettleDetail struct {
	OperationType     string `json:"operation_type,omitempty"`
	OperationSerialNo string `json:"operation_serial_no,omitempty"`
	OperationDt       string `json:"operation_dt,omitempty"`
	TransOut          string `json:"trans_out,omitempty"`
	TransIn           string `json:"trans_in,omitempty"`
	Amount            string `json:"amount,omitempty"`
	OriTransOut       string `json:"ori_trans_out,omitempty"`
	OriTransIn        string `json:"ori_trans_in,omitempty"`
}

type TradeSettleInfo struct {
	TradeSettleDetailList []*TradeSettleDetail `json:"trade_settle_detail_list,omitempty"`
}

type HbFqPayInfo struct {
	UserInstallNum   string `json:"user_install_num,omitempty"`
	CreditPayMode    string `json:"credit_pay_mode,omitempty"`
	CreditBizOrderId string `json:"credit_biz_order_id,omitempty"`
	HybAmount        string `json:"hyb_amount,omitempty"`
}

type TradeQueryDto struct {
	ErrorResponseDto
	TradeNo             string           `json:"trade_no,omitempty"`
	OutTradeNo          string           `json:"out_trade_no,omitempty"`
	BuyerLogonId        string           `json:"buyer_logon_id,omitempty"`
	TradeStatus         string           `json:"trade_status,omitempty"`
	TotalAmount         string           `json:"total_amount,omitempty"`
	TransCurrency       string           `json:"trans_currency,omitempty"`
	SettleCurrency      string           `json:"settle_currency,omitempty"`
	SettleAmount        string           `json:"settle_amount,omitempty"`
	PayCurrency         string           `json:"pay_currency,omitempty"`
	PayAmount           string           `json:"pay_amount,omitempty"`
	SettleTransRate     string           `json:"settle_trans_rate,omitempty"`
	TransPayRate        string           `json:"trans_pay_rate,omitempty"`
	BuyerPayAmount      string           `json:"buyer_pay_amount,omitempty"`
	PointAmount         string           `json:"point_amount,omitempty"`
	InvoiceAmount       string           `json:"invoice_amount,omitempty"`
	SendPayDate         string           `json:"send_pay_date,omitempty"`
	ReceiptAmount       string           `json:"receipt_amount,omitempty"`
	StoreId             string           `json:"store_id,omitempty"`
	TerminalId          string           `json:"terminal_id,omitempty"`
	FundBillList        []*TradeFundBill `json:"fund_bill_list"`
	StoreName           string           `json:"store_name,omitempty"`
	BuyerUserId         string           `json:"buyer_user_id,omitempty"`
	ChargeAmount        string           `json:"charge_amount,omitempty"`
	ChargeFlags         string           `json:"charge_flags,omitempty"`
	SettlementId        string           `json:"settlement_id,omitempty"`
	TradeSettleInfo     *TradeSettleInfo `json:"trade_settle_info,omitempty"`
	AuthTradePayMode    string           `json:"auth_trade_pay_mode,omitempty"`
	BuyerUserType       string           `json:"buyer_user_type,omitempty"`
	MdiscountAmount     string           `json:"mdiscount_amount,omitempty"`
	DiscountAmount      string           `json:"discount_amount,omitempty"`
	Subject             string           `json:"subject,omitempty"`
	Body                string           `json:"body,omitempty"`
	AlipaySubMerchantId string           `json:"alipay_sub_merchant_id,omitempty"`
	ExtInfos            string           `json:"ext_infos,omitempty"`
	HbFqPayInfo         *HbFqPayInfo     `json:"hb_fq_pay_info,omitempty"`
	CreditPayMode       string           `json:"credit_pay_mode,omitempty"`
	CreditBizOrderId    string           `json:"credit_biz_order_id,omitempty"`
}

type TradeQueryResponseDto struct {
	Response     *TradeQueryDto `json:"alipay_trade_query_response"`
	AlipayCertSn string         `json:"alipay_cert_sn,omitempty"`
	SignData     string         `json:"-"`
	Sign         string         `json:"sign"`
}
