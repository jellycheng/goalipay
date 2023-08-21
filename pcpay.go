package goalipay

import (
	"github.com/jellycheng/gosupport"
)

// pc端下单地址
func PagePay(config *AplipayConfig, bizMap gosupport.BodyMap, commonMap gosupport.BodyMap) (string, error) {
	err := bizMap.CheckEmpty("out_trade_no", "total_amount", "subject")
	if err != nil {
		return "", err
	}
	commonMap["method"] = "alipay.trade.page.pay"
	if _, ok := bizMap["product_code"]; !ok {
		bizMap["product_code"] = "FAST_INSTANT_TRADE_PAY"
	}
	formParam, _ := CommonParamsHandle(config, bizMap, commonMap)

	return AlipayGatewayUrl + "?" + formParam.UrlEncode(), nil
}
