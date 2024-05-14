package media_model

import "xuetang/kitex_gen/xuetang"

func ValidFail(msg string) *xuetang.RestResponse {
	return &xuetang.RestResponse{Code: -1, Msg: msg}
}

func ValidFailWithData(result string, msg string) *xuetang.RestResponse {
	return &xuetang.RestResponse{Code: -1, Result_: &result, Msg: msg}
}

func Success(result string) *xuetang.RestResponse {
	return &xuetang.RestResponse{Result_: &result}
}

func SuccessWithMsg(result string, msg string) *xuetang.RestResponse {
	return &xuetang.RestResponse{Result_: &result, Msg: msg}
}

func SuccessEmpty() *xuetang.RestResponse {
	return &xuetang.RestResponse{}
}

func IsSuccessful(r *xuetang.RestResponse) bool {
	return r.Code == 0
}
