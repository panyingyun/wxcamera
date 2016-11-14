package httpserver

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type Request struct {
	//request id
	Id string `json:"id"`
	//system info
	Ver   string `json:"ver"`
	Sign  string `json:"sign"`
	Appid string `json:"appid"`
	Time  string `json:"time"`
	//https method
	Method string `json:"method"`
	//request parameters
	DeviceId    string `json:"deviceId"`
	AccessToken string `json:"accessToken"`
	UserId      string `json:"userId"`
	Phone       string `json:"phone"`
}

func SignReq(ptrReq interface{}, appKey string) string {
	sortmap := make(map[string]interface{})
	s := reflect.ValueOf(ptrReq).Elem()
	typeOfReq := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		sortmap[string(typeOfReq.Field(i).Tag.Get("json"))] = f.Interface()
	}

	var keys []string
	for k := range sortmap {
		if k != "sign" {
			keys = append(keys, k)
		}
	}
	sort.Strings(keys)

	//slice
	var slice []string = make([]string, 0)
	for _, v := range keys {
		value := sortmap[v]
		//append parameters to slice
		slice = append(slice, fmt.Sprintf("%v:%v", v, value))
	}
	slice = append(slice, "appKey"+":"+appKey)
	sign := strings.Join(slice, ",")
	//md5
	hasher := md5.New()
	hasher.Write([]byte(sign))
	return hex.EncodeToString(hasher.Sum(nil))
}
