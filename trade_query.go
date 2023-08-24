package goalipay

import (
	"fmt"
	"github.com/jellycheng/gcurl"
	"github.com/jellycheng/gosupport"
)

// alipay.trade.query(统一收单交易查询):https://opendocs.alipay.com/open/02ivbt?scene=common&pathHash=8abc6ffe
// https://opendocs.alipay.com/open/82ea786a_alipay.trade.query?scene=23&pathHash=0745ecea
// https://opendocs.alipay.com/open/bff76748_alipay.trade.query?scene=23&pathHash=e3ddce1d
func TradeQuery(config *AplipayConfig, bizMap gosupport.BodyMap, commonMap gosupport.BodyMap) (*TradeQueryResponseDto, string, error) {
	bodyStr := ""
	aliRespDto := new(TradeQueryResponseDto)
	if bizMap.GetString("out_trade_no") == "" && bizMap.GetString("trade_no") == "" {
		return aliRespDto, bodyStr, fmt.Errorf("%s", "缺少必需的参数,out_trade_no、trade_no不能同时为空")
	}
	commonMap["method"] = "alipay.trade.query"
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
