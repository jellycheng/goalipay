package goalipay

import (
	"fmt"
	"github.com/jellycheng/gcurl"
	"github.com/jellycheng/gosupport"
)

// 账户余额查询： https://opendocs.alipay.com/open/01inen#查询账户当前余额
func BalanceQuery(config *AplipayConfig, bizMap gosupport.BodyMap, commonMap gosupport.BodyMap) (*BalanceQueryResponseDto, string, error) {
	bodyStr := ""
	aliRespDto := new(BalanceQueryResponseDto)

	//if err := bizMap.CheckEmpty("bill_user_id"); err != nil {
	//	return aliRespDto, bodyStr, err
	//}
	commonMap["method"] = "alipay.data.bill.balance.query"
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
