package goalipay

import (
	"github.com/jellycheng/gosupport"
)

// 获取h5下单链接地址
func WapPay(config *AplipayConfig, bizMap gosupport.BodyMap, commonMap gosupport.BodyMap) (string, error) {
	err := bizMap.CheckEmpty("out_trade_no", "total_amount", "subject")
	if err != nil {
		return "", err
	}
	commonMap["method"] = "alipay.trade.wap.pay"
	formParam, _ := CommonParamsHandle(config, bizMap, commonMap)
	return AlipayGatewayUrl + "?" + formParam.UrlEncode(), nil
}
