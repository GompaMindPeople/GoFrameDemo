package utils

import "github.com/gogf/gf/encoding/gjson"



/**
	需要将gjson强制转换一下才是标准的json
	否则string 的时候会将 “” 一起转移成 \“
 */
func EntityToJson(entity interface{})(*gjson.Json, error){
	decode, e := gjson.Encode(entity)
	result, e := gjson.DecodeToJson(decode)
	return result,e
}