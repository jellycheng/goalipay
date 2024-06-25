# goalipay
```
golang封装支付宝支付
术语：
    应用私钥
    应用公钥
    支付宝公钥
    
```

## 查询订单
```
package main

import (
	"fmt"
	"github.com/jellycheng/goalipay"
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/xcrypto"
)

func main() {
	appid := "2018091561417031" //应用ID
	pk, _ := gosupport.FileGetContents("应用私钥文件地址")
	cfg, err := goalipay.NewClient(appid, pk)
	if err != nil {
		fmt.Println(err)
		return
	}
	cfg.IsVerifySign = true
	pubK, _ := gosupport.FileGetContents("支付宝公钥文件地址")
	pubKStr := xcrypto.RsaContent2PublicKey(pubK)
	var e error
	cfg.AliPayPublicKey, e = xcrypto.DecodePublicKey([]byte(pubKStr))
	if e != nil {
		fmt.Println(e)
		return
	}
	// 请求参数
	bizMap := gosupport.NewBodyMap()
	bizMap.Set("out_trade_no", "GZ20190908174341CjDk") // 商户订单号
	if resData, bodyStr, err := goalipay.TradeQuery(cfg, bizMap, make(gosupport.BodyMap)); err == nil {
		fmt.Println(bodyStr)
		fmt.Println(fmt.Sprintf("%+v", resData.Response))
		if goalipay.IsError(resData.Response.Code) {
			fmt.Println("查询错误：", resData.Response.Msg)
		} else { // 订单查询结果
			fmt.Println(fmt.Sprintf("订单号： %s， 订单金额：%s，", resData.Response.OutTradeNo, resData.Response.TotalAmount))
		}
	} else {
		fmt.Println(err)
	}
}


```

## 发起退款
```
package main

import (
	"fmt"
	"github.com/jellycheng/goalipay"
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/xcrypto"
)

func main() {
	appid := "2018091561417031" //应用ID
	pk, _ := gosupport.FileGetContents("应用私钥文件地址")
	cfg, err := goalipay.NewClient(appid, pk)
	if err != nil {
		fmt.Println(err)
		return
	}
	cfg.IsVerifySign = true
	pubK, _ := gosupport.FileGetContents("支付宝公钥文件地址")
	pubKStr := xcrypto.RsaContent2PublicKey(pubK)
	var e error
	cfg.AliPayPublicKey, e = xcrypto.DecodePublicKey([]byte(pubKStr))
	if e != nil {
		fmt.Println(e)
		return
	}
	// 请求参数
	bizMap := gosupport.NewBodyMap()
	// 要退款的正向订单号和退款金额
	bizMap.Set("out_trade_no", "GZ20190908174341CjDk").Set("refund_amount", 0.01)
	bizMap.Set("out_request_no", "orn2023" + gosupport.GetRandString(6)) // 退款请求号，唯一
	if resData, bodyStr, err := goalipay.TradeRefund(cfg, bizMap, make(gosupport.BodyMap)); err == nil {
		fmt.Println(bodyStr)
		fmt.Println(fmt.Sprintf("%+v", resData.Response))
		if goalipay.IsError(resData.Response.Code) {
			fmt.Println("发起退款错误：", resData.Response.Msg)
		} else { // 退款发起成功，同步知道退款结果
			fmt.Println(fmt.Sprintf("订单号： %s， 退款金额：%s，", resData.Response.OutTradeNo, resData.Response.RefundFee))
		}
	} else {
		fmt.Println(err)
	}
}

```
[https://opendocs.alipay.com/open/6c0cdd7d_alipay.trade.refund?scene=common&pathHash=4081e89c](https://opendocs.alipay.com/open/6c0cdd7d_alipay.trade.refund?scene=common&pathHash=4081e89c)

## 退款结果查询
```
package main

import (
	"fmt"
	"github.com/jellycheng/goalipay"
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/xcrypto"
)

func main() {
	appid := "2018091561417031" //应用ID
	pk, _ := gosupport.FileGetContents("应用私钥文件地址")
	cfg, err := goalipay.NewClient(appid, pk)
	if err != nil {
		fmt.Println(err)
		return
	}
	cfg.IsVerifySign = true
	pubK, _ := gosupport.FileGetContents("支付宝公钥文件地址")
	pubKStr := xcrypto.RsaContent2PublicKey(pubK)
	var e error
	cfg.AliPayPublicKey, e = xcrypto.DecodePublicKey([]byte(pubKStr))
	if e != nil {
		fmt.Println(e)
		return
	}
	// 请求参数
	bizMap := gosupport.NewBodyMap()
	// 支付订单号和退款单号
	bizMap.Set("out_trade_no", "GZ201909081743JVZkjB").Set("out_request_no", "orn2023abc")
	if resData, bodyStr, err := goalipay.TradeFastpayRefundQuery(cfg, bizMap, make(gosupport.BodyMap)); err == nil {
		fmt.Println(bodyStr)
		fmt.Println(fmt.Sprintf("%+v", resData.Response))
		if goalipay.IsError(resData.Response.Code) {
			fmt.Println("退款查询错误：", resData.Response.Msg)
		} else { // 退款查询结果
			fmt.Println(fmt.Sprintf("支付订单号： %s，退款单号： %s， 退款金额：%s，", resData.Response.OutTradeNo, resData.Response.OutRequestNo, resData.Response.RefundAmount))
		}
	} else {
		fmt.Println(err)
	}
}

```

## pc支付文档
```
拼接pc web支付地址
package main

import (
	"fmt"
	"github.com/jellycheng/goalipay"
	"github.com/jellycheng/gosupport"
)

func main() {
	appid := "2018091561417031" //应用ID
	pk, _ := gosupport.FileGetContents("应用私钥文件地址")
	cfg, err := goalipay.NewClient(appid, pk)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 请求业务参数
	bizMap := gosupport.NewBodyMap()
	bizMap.Set("subject", "购买商品").
		Set("out_trade_no", "GZ201909081743"+gosupport.GetRandString(6)).
		Set("total_amount", "0.01")

	// 公共参数，设置通知地址和跳转地址
	commonMap := gosupport.NewBodyMap()
	commonMap.Set("notify_url", "https://支付成功通知地址").
		Set("return_url", "https://支付成功pc页面跳转地址")

	if pcUrl, err := goalipay.PagePay(cfg, bizMap, commonMap); err == nil {
		fmt.Println(pcUrl)
	} else {
		fmt.Println(err)
	}
}

支付成功后在跳转地址中追加参数如下：
    domain?charset=utf-8&out_trade_no=商户订单号&method=alipay.trade.page.pay.return&total_amount=支付金额&sign=签名&trade_no=支付宝订单号&auth_app_id=应用ID&version=1.0&app_id=应用ID&sign_type=RSA2&seller_id=购买者&timestamp=时间

```
[https://opendocs.alipay.com/open/repo-0038oa](https://opendocs.alipay.com/open/repo-0038oa) <br>

## 手机支付文档
```
拼接h5地址
package main

import (
	"fmt"
	"github.com/jellycheng/goalipay"
	"github.com/jellycheng/gosupport"
)

func main() {
	appid := "2018091561417031" //应用ID
	pk, _ := gosupport.FileGetContents("应用私钥文件地址")
	cfg, err := goalipay.NewClient(appid, pk)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 请求参数
	bizMap := gosupport.NewBodyMap()
	bizMap.Set("subject", "购买商品").
		Set("out_trade_no", "ono20230824"+gosupport.GetRandString(6)).
		Set("total_amount", "0.02")

	commonMap := gosupport.NewBodyMap()
	commonMap.Set("notify_url", "https://支付成功通知地址").
		Set("return_url", "https://支付成功h5页面跳转地址")

	if h5Url, err := goalipay.WapPay(cfg, bizMap, commonMap); err == nil {
		fmt.Println(h5Url)
	} else {
		fmt.Println(err)
	}
}

支付成功h5页面跳转地址追加参数如下：
    domain?charset=utf-8&out_trade_no=商户订单号&method=alipay.trade.wap.pay.return&total_amount=金额&sign=签名&trade_no=支付宝单号&auth_app_id=应用ID&version=1.0&app_id=应用ID&sign_type=RSA2&seller_id=购买者ID&timestamp=时间

```
[https://opendocs.alipay.com/open/repo-0038v7](https://opendocs.alipay.com/open/repo-0038v7) <br>

## app支付文档
```
package main

import (
	"fmt"
	"github.com/jellycheng/goalipay"
	"github.com/jellycheng/gosupport"
)

func main() {
	appid := "2018091561417031" //应用ID
	pk, _ := gosupport.FileGetContents("应用私钥文件地址")
	cfg, err := goalipay.NewClient(appid, pk)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 请求参数
	bizMap := gosupport.NewBodyMap()
	bizMap.Set("subject", "购买商品").
		Set("out_trade_no", "GZ201909081743JVZkjB").
		Set("total_amount", 0.01)

	commonMap := gosupport.NewBodyMap()
	commonMap.Set("notify_url", "https://支付成功通知地址")
	if resData, err := goalipay.AppPay(cfg, bizMap, commonMap); err == nil {
		fmt.Println(resData) //传给app的内容
	} else {
		fmt.Println(err)
	}
}

```
[https://opendocs.alipay.com/open/repo-0038v9](https://opendocs.alipay.com/open/repo-0038v9) <br>

## 查询支付宝商户号余额
```
package main

import (
	"fmt"
	"github.com/jellycheng/goalipay"
	"github.com/jellycheng/gosupport"
	"github.com/jellycheng/gosupport/xcrypto"
)

func main() {
	appid := "" // 开放平台创建的应用ID
	pk, _ := gosupport.FileGetContents("应用私钥文件地址")
	cfg, err := goalipay.NewClient(appid, pk)
	if err != nil {
		fmt.Println(err)
		return
	}
	cfg.IsVerifySign = true
	pubK, _ := gosupport.FileGetContents("支付宝公钥文件地址")
	pubKStr := xcrypto.RsaContent2PublicKey(pubK)
	var e error
	cfg.AliPayPublicKey, e = xcrypto.DecodePublicKey([]byte(pubKStr))
	if e != nil {
		fmt.Println(e)
		return
	}
	// 请求参数
	bizMap := make(gosupport.BodyMap)
	if resData, bodyStr, err := goalipay.BalanceQuery(cfg, bizMap, make(gosupport.BodyMap)); err == nil {
		fmt.Println("支付宝返回原始内容：", bodyStr)
		if goalipay.IsError(resData.Response.Code) {
			fmt.Println("查询错误：", resData.Response.Msg)
		} else { // 查询正确
			fmt.Println(fmt.Sprintf("查询结果： %+v", resData.Response))
		}
	} else {
		fmt.Println(err)
	}

}


```

## 工具
```

```
[支付宝提供密钥工具：https://opendocs.alipay.com/common/02kipk?pathHash=0d20b438](https://opendocs.alipay.com/common/02kipk?pathHash=0d20b438) <br>
[api文档：https://open.alipay.com/api](https://open.alipay.com/api) <br>

