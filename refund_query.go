package goalipay

import (
	"fmt"
	"github.com/jellycheng/gcurl"
	"github.com/jellycheng/gosupport"
)

// alipay.trade.fastpay.refund.query(统一收单交易退款查询)
// https://opendocs.alipay.com/open/357441a2_alipay.trade.fastpay.refund.query?scene=common&pathHash=01981dca
// https://opendocs.alipay.com/open/8c776df6_alipay.trade.fastpay.refund.query?scene=common&pathHash=fb6e1894
func TradeFastpayRefundQuery(config *AplipayConfig, bizMap gosupport.BodyMap, commonMap gosupport.BodyMap) (*TradeFastpayRefundQueryResponseDto, string, error) {
	bodyStr := ""
	aliRespDto := new(TradeFastpayRefundQueryResponseDto)
	if bizMap.GetString("out_trade_no") == "" && bizMap.GetString("trade_no") == "" {
		return aliRespDto, bodyStr, fmt.Errorf("%s", "缺少必需的参数,out_trade_no、trade_no不能同时为空")
	}
	if err := bizMap.CheckEmpty("out_request_no"); err != nil {
		return aliRespDto, bodyStr, err
	}
	commonMap["method"] = "alipay.trade.fastpay.refund.query"
	formParam, _ := CommonParamsHandle(config, bizMap, commonMap)
	resp, err := gcurl.Post(AlipayGatewayUrl, gcurl.Options{
		Headers: map[string]interface{}{
			"Content-Type": gcurl.ContentTypeForm,
			"User-Agent":   "gcurl/1.0",
		},
		FormParams: formParam,
	})
	if err != nil {
		return aliRespDto, bodyStr, err
	} else {
		body, _ := resp.GetBody()
		bodyStr = body.String()
		if err = gosupport.JsonUnmarshal(bodyStr, aliRespDto); err != nil || aliRespDto.Response == nil {
			return aliRespDto, bodyStr, fmt.Errorf("[json unmarshal error], string: %s", bodyStr)
		}
		signData, signDataErr := config.GetSignData(bodyStr, aliRespDto.AlipayCertSn)
		aliRespDto.SignData = signData
		return aliRespDto, bodyStr, config.AutoVerifySignByCert(aliRespDto.Sign, signData, signDataErr)
	}
}
