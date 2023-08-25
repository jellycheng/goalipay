package goalipay

import (
	"github.com/jellycheng/gosupport"
	"net/http"
	"net/url"
)

func ReturnSuccess() string {
	return SUCCESS
}

// 支付宝通知文档： https://opendocs.alipay.com/open/203/105286
// https://opendocs.alipay.com/open/204/105301
func NotifyContentToBodyMap(req *http.Request) (gosupport.BodyMap, error) {
	if err := req.ParseForm(); err != nil {
		return nil, err
	}
	// req.PostForm
	bm := NotifyContentToBodyMapByUrlValues(req.Form)
	return bm, nil
}

func NotifyContentToBodyMapByUrlValues(valMap url.Values) gosupport.BodyMap {
	bm := make(gosupport.BodyMap, len(valMap)+1)
	for k, v := range valMap {
		if len(v) == 1 {
			bm.Set(k, v[0])
		}
	}
	return bm
}
