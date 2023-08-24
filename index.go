package goalipay

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/xcrypto"
	"strings"
	"time"
)

type AplipayConfig struct {
	AppId              string
	SignType           string
	Charset            string
	PrivateKey         *rsa.PrivateKey // 应用私钥
	AliPayPublicKey    *rsa.PublicKey  // 支付宝证书公钥内容
	AliPayPublicCertSN string
	ReturnUrl          string
	NotifyUrl          string
	Format             string
	IsVerifySign       bool // 是否验证支付宝响应结果签名
}

// 设置支付后的ReturnUrl
func (m *AplipayConfig) SetReturnUrl(url string) *AplipayConfig {
	m.ReturnUrl = url
	return m
}

// 设置支付宝通知地址
func (m *AplipayConfig) SetNotifyUrl(url string) *AplipayConfig {
	m.NotifyUrl = url
	return m
}

// 设置编码格式
func (m *AplipayConfig) SetCharset(charset string) *AplipayConfig {
	if charset != "" {
		m.Charset = charset
	}
	return m
}

func (m *AplipayConfig) SetSignType(signType string) *AplipayConfig {
	if signType != "" {
		m.SignType = signType
	}
	return m
}

func (m *AplipayConfig) GetSignData(bodyStr string, alipayCertSN string) (signData string, err error) {
	var (
		indexStart = strings.Index(bodyStr, `_response":`)
		indexEnd   int
	)
	indexStart = indexStart + 11
	bsLen := len(bodyStr)
	if alipayCertSN != "" {
		// 公钥证书模式
		if alipayCertSN != m.AliPayPublicCertSN {
			return "", fmt.Errorf("[证书不匹配], 当前使用的支付宝公钥证书SN[%s]与网关响应报文中的SN[%s]不匹配", m.AliPayPublicCertSN, alipayCertSN)
		}
		indexEnd = strings.Index(bodyStr, `,"alipay_cert_sn":`)
		if indexEnd > indexStart && bsLen > indexStart {
			signData = bodyStr[indexStart:indexEnd]
			return
		}
		return "", fmt.Errorf("[原签名数据错误], value: %s", bodyStr)
	}
	// 普通公钥模式
	indexEnd = strings.Index(bodyStr, `,"sign":`)
	if indexEnd > indexStart && bsLen > indexStart {
		signData = bodyStr[indexStart:indexEnd]
		return
	}
	return "", fmt.Errorf("[原签名数据错误], value: %s", bodyStr)
}

func (m *AplipayConfig) AutoVerifySignByCert(sign, signData string, signDataErr error) (err error) {
	if m.IsVerifySign && m.AliPayPublicKey != nil {
		if signDataErr != nil {
			return signDataErr
		}
		return VerifySignByCert(sign, signData, m.SignType, m.AliPayPublicKey)
	}
	return nil
}

func NewClient(appid, privateKey string) (*AplipayConfig, error) {
	if appid == "" || privateKey == "" {
		return nil, errors.New("appid，privateKey参数不能为空")
	}
	key := xcrypto.RsaContent2PrivateKey(privateKey)
	priKey, err := xcrypto.DecodePrivateKey([]byte(key))
	if err != nil {
		return nil, err
	}
	cfg := &AplipayConfig{
		AppId:      appid,
		SignType:   RSA2,
		Charset:    UTF8,
		PrivateKey: priKey,
		Format:     JsonFormat,
	}
	return cfg, nil
}

// 公共参数
func CommonParamsHandle(alipayCfg *AplipayConfig, bizMap gosupport.BodyMap, commonMap gosupport.BodyMap) (gosupport.BodyMap, error) {
	apiMethod := commonMap.GetString("method")
	format := commonMap.GetString("format")
	if format == "" {
		format = alipayCfg.Format
		if format == "" {
			format = "JSON"
		}
	}
	charset := commonMap.GetString("charset")
	if charset == "" {
		charset = alipayCfg.Charset
		if charset == "" {
			charset = UTF8
		}
	}

	signType := commonMap.GetString("sign_type")
	if signType == "" {
		signType = alipayCfg.SignType
		if signType == "" {
			signType = RSA2
		}
	}

	paramBody := make(gosupport.BodyMap)
	paramBody.Set("app_id", alipayCfg.AppId).
		Set("method", apiMethod). //接口名称
		Set("format", format).
		Set("charset", charset).
		Set("sign_type", signType).
		Set("version", "1.0").
		Set("timestamp", time.Now().Format(gosupport.TimeFormat))

	if version := commonMap.GetString("version"); version != "" {
		paramBody.Set("version", version)
	}

	if alipayCfg.ReturnUrl != "" {
		paramBody.Set("return_url", alipayCfg.ReturnUrl)
	}
	if returnUrl := commonMap.GetString("return_url"); returnUrl != "" {
		paramBody.Set("return_url", returnUrl)
	}

	if timestamp := commonMap.GetString("timestamp"); timestamp != "" {
		paramBody.Set("timestamp", timestamp)
	}

	if alipayCfg.NotifyUrl != "" {
		paramBody.Set("notify_url", alipayCfg.NotifyUrl)
	}
	if notifyUrl := commonMap.GetString("notify_url"); notifyUrl != "" {
		paramBody.Set("notify_url", notifyUrl)
	}
	if aat := commonMap.GetString("app_auth_token"); aat != "" {
		paramBody.Set("app_auth_token", aat)
	}
	bizContent := PinBizContent(bizMap)
	if bizContent != "" {
		paramBody.Set("biz_content", bizContent)
	}
	// 签名
	sign, err := GetAlipaySign(paramBody, paramBody.GetString("sign_type"), alipayCfg.PrivateKey)
	if err != nil {
		return paramBody, fmt.Errorf("签名错误: %w", err)
	}
	paramBody.Set("sign", sign)
	return paramBody, nil
}

func PinBizContent(bizmap gosupport.BodyMap) string {
	bizContent := gosupport.ToJson(bizmap)

	return bizContent
}
