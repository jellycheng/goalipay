package goalipay

type RefundPresetPaytool struct {
	Amount         []string `json:"amount,omitempty"`
	AssertTypeCode string   `json:"assert_type_code,omitempty"`
}

type TradeRefundDto struct {
	ErrorResponseDto
	TradeNo                      string                 `json:"trade_no,omitempty"`
	OutTradeNo                   string                 `json:"out_trade_no,omitempty"`
	BuyerLogonId                 string                 `json:"buyer_logon_id,omitempty"`
	FundChange                   string                 `json:"fund_change,omitempty"`
	RefundFee                    string                 `json:"refund_fee,omitempty"`
	RefundDetailItemList         []*TradeFundBill       `json:"refund_detail_item_list,omitempty"`
	StoreName                    string                 `json:"store_name,omitempty"`
	BuyerUserId                  string                 `json:"buyer_user_id,omitempty"`
	SendBackFee                  string                 `json:"send_back_fee,omitempty"`
	OpenId                       string                 `json:"open_id,omitempty"`
	RefundCurrency               string                 `json:"refund_currency,omitempty"`
	GmtRefundPay                 string                 `json:"gmt_refund_pay,omitempty"`
	RefundPresetPaytoolList      []*RefundPresetPaytool `json:"refund_preset_paytool_list,omitempty"`
	RefundChargeAmount           string                 `json:"refund_charge_amount,omitempty"`
	RefundSettlementId           string                 `json:"refund_settlement_id,omitempty"`
	PresentRefundBuyerAmount     string                 `json:"present_refund_buyer_amount,omitempty"`
	PresentRefundDiscountAmount  string                 `json:"present_refund_discount_amount,omitempty"`
	PresentRefundMdiscountAmount string                 `json:"present_refund_mdiscount_amount,omitempty"`
	HasDepositBack               string                 `json:"has_deposit_back,omitempty"`
	RefundHybAmount              string                 `json:"refund_hyb_amount,omitempty"`
}

// 发起退款 返回
type TradeRefundResponseDto struct {
	Response     *TradeRefundDto `json:"alipay_trade_refund_response"`
	AlipayCertSn string          `json:"alipay_cert_sn,omitempty"`
	SignData     string          `json:"-"`
	Sign         string          `json:"sign"`
}
