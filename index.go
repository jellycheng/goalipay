package goalipay

import "github.com/jellycheng/gosupport"

type BodyMap map[string]interface{}

func (m BodyMap) Set(key string, value interface{}) BodyMap {
	m[key] = value
	return m
}

func (m BodyMap) GetInterface(key string) interface{} {
	if m == nil {
		return nil
	}
	return m[key]
}

func (m BodyMap) Remove(key string) {
	delete(m, key)
}

func (m BodyMap) ToJson() string {
	return gosupport.ToJson(m)
}
