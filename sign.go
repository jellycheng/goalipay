package goalipay

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/jellycheng/gosupport"
	"hash"
	"reflect"
	"sort"
	"strings"
)

func GetAlipaySign(bm gosupport.BodyMap, signType string, privateKey *rsa.PrivateKey) (sign string, err error) {
	var (
		h              hash.Hash
		hashs          crypto.Hash
		encryptedBytes []byte
	)

	switch signType {
	case RSA:
		h = sha1.New()
		hashs = crypto.SHA1
	case RSA2:
		h = sha256.New()
		hashs = crypto.SHA256
	default:
		h = sha256.New()
		hashs = crypto.SHA256
	}
	signParams := EncodeAliPaySignParams(bm)
	if _, err = h.Write([]byte(signParams)); err != nil {
		return
	}
	if encryptedBytes, err = rsa.SignPKCS1v15(rand.Reader, privateKey, hashs, h.Sum(nil)); err != nil {
		return "", fmt.Errorf("[%w]: %+v", SignErr, err)
	}
	sign = base64.StdEncoding.EncodeToString(encryptedBytes)
	return
}

func EncodeAliPaySignParams(bm gosupport.BodyMap) string {
	if bm == nil {
		return ""
	}
	var (
		buf     strings.Builder
		keyList []string
	)
	for k := range bm {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		if v := bm.GetString(k); v != "" {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
			buf.WriteByte('&')
		}
	}
	if buf.Len() <= 0 {
		return ""
	}
	return buf.String()[:buf.Len()-1]
}

// 验证api数据：https://opendocs.alipay.com/common/02mse7#自行实现验签
func VerifySignByCert(sign, signData string, signType string, aliPayPublicKey *rsa.PublicKey) error {
	var (
		h     hash.Hash
		hashs crypto.Hash
		err   error
	)
	//aliPayPublicKey, err := xcrypto.DecodePublicKey(pemContent)
	//if err != nil {
	//	return err
	//}

	if aliPayPublicKey != nil {
		signBytes, _ := base64.StdEncoding.DecodeString(sign)
		switch signType {
		case RSA:
			hashs = crypto.SHA1
		case RSA2:
			hashs = crypto.SHA256
		default:
			hashs = crypto.SHA256
		}
		h = hashs.New()
		h.Write([]byte(signData))
		if err = rsa.VerifyPKCS1v15(aliPayPublicKey, hashs, h.Sum(nil), signBytes); err != nil {
			return fmt.Errorf("[验证签名错误]: %v", err)
		}
		return nil
	} else {
		return errors.New("没有设置支付宝公钥")
	}

}

func VerifyNotifySign(notifyContent interface{}, aliPayPublicKey *rsa.PublicKey) (bool, error) {
	if aliPayPublicKey == nil || notifyContent == nil {
		return false, errors.New("notifyContent or aliPayPublicKeyis nil")
	}
	var (
		bodySign     string
		bodySignType string
		signData     string
		bm           = gosupport.NewBodyMap()
	)
	if reflect.ValueOf(notifyContent).Kind() == reflect.Map {
		if bmObj, ok := notifyContent.(gosupport.BodyMap); ok {
			bm = bmObj
		} else {
			return false, fmt.Errorf("map类型错误：%s", notifyContent)
		}
	} else if err := gosupport.JsonUnmarshal(notifyContent.(string), &bm); err != nil {
		return false, fmt.Errorf("json.Unmarshal(%s)：%w", notifyContent, err)
	}
	bodySign = bm.GetString("sign")
	bodySignType = bm.GetString("sign_type")
	bm.Remove("sign")
	bm.Remove("sign_type")
	signData = EncodeAliPaySignParams(bm)
	if err := VerifySignByCert(bodySign, signData, bodySignType, aliPayPublicKey); err != nil {
		return false, err
	}
	return true, nil
}
