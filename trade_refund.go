package goalipay

import (
	"fmt"
	"github.com/jellycheng/gcurl"
	"github.com/jellycheng/gosupport"
)

// alipay.trade.refund(统一收单交易退款接口)
// https://opendocs.alipay.com/open/6c0cdd7d_alipay.trade.refund?scene=common&pathHash=4081e89c
func TradeRefund(config *AplipayConfig, bizMap gosupport.BodyMap, commonMap gosupport.BodyMap) (*TradeRefundResponseDto, string, error) {
	bodyStr := ""
	aliRespDto := new(TradeRefundResponseDto)
	if bizMap.GetString("out_trade_no") == "" && bizMap.GetString("trade_no") == "" {
		return aliRespDto, bodyStr, fmt.Errorf("%s", "缺少必需的参数,out_trade_no、trade_no不能同时为空")
	}

	if err := bizMap.CheckEmpty("refund_amount"); err != nil {
		return aliRespDto, bodyStr, err
	}

	commonMap["method"] = "alipay.trade.refund"
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
