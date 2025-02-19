package pkg

type JSONResponse struct {
	Code int `json:"code"`
	Data any `json:"data"`
}

type ResponseMessageData struct {
	Message   string `json: "message"`
	MessageZh string `json:"message_zh"`
}

func SuccessResponse(data any) JSONResponse {
	return JSONResponse{
		Code: 0,
		Data: data,
	}
}

func MessageResponse(code int, msg, msgZh string) JSONResponse {
	return JSONResponse{
		Code: code,
		Data: ResponseMessageData{
			Message:   msg,
			MessageZh: msgZh,
		},
	}
}
