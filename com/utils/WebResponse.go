package utils

import "fmt"

type Resp struct {
	Code    string `json:"code"`
	Msg     string `json:"msg"`
	Data     string `json:"Data"`
}

func CodeMsgFactory(code int,msg string) (int,string){

	return code,msg
}

func CodeMsg() {

	//var CodeMsgMap map[string]string
	CodeMsgMap :=make(map[string]map[string]map[int]string)
	CodeMsgTypeMap :=make(map[string]map[int]string)
	CodeMsgInfoMap :=make(map[int]string)

	CodeMsgInfoMap[0] = "操作成功"
	CodeMsgTypeMap["SUCCESS"] = CodeMsgInfoMap





	CodeMsgMap[0] = "操作成功"
	CodeMsgMap[1] = "操作成功1"

	 sunccess := CodeMsgFactory(0,"操作成功")

	return sunccess
}

