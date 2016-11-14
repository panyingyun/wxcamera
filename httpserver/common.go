package httpserver

import (
	"encoding/json"
	"reflect"
)

const (
	APIURL    = "http://wx.michaelapp.com"
	APPID     = "wx5f2139191e42b3f0"
	APPSECRET = "104ed693b599b24df164fdb2b9cb5cd4"
)

//common func define here!
type CommonResult struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CommonResp struct {
	ID     string       `json:"id"`
	Result CommonResult `json:"result"`
}

func GenerateCommonResp(code int, msg string, id string) []byte {
	var resp CommonResp
	resp.ID = id
	resp.Result.Code = code
	resp.Result.Msg = msg
	bytes, _ := json.Marshal(resp)
	return bytes
}

func Convert(src, dst interface{}) {
	sval := reflect.ValueOf(src).Elem()
	dval := reflect.ValueOf(dst).Elem()

	for i := 0; i < sval.NumField(); i++ {
		value := sval.Field(i)
		name := sval.Type().Field(i).Name

		dvalue := dval.FieldByName(name)
		if dvalue.IsValid() == false {
			continue
		}
		dvalue.Set(value)
	}
}
