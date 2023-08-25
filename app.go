package goalipay

import "github.com/jellycheng/gosupport"

// 获取app下单参数：https://opendocs.alipay.com/open/cd12c885_alipay.trade.app.pay?scene=20&pathHash=c0e35284
// alipay.trade.app.pay(app支付接口2.0)
func AppPay(config *AplipayConfig, bizMap gosupport.BodyMap, commonMap gosupport.BodyMap) (string, error) {
	err := bizMap.CheckEmpty("out_trade_no", "total_amount", "subject")
	if err != nil {
		return "", err
	}
	commonMap["method"] = "alipay.trade.app.pay"
	formParam, _ := CommonParamsHandle(config, bizMap, commonMap)
	return formParam.UrlEncode(), nil
}
