package main

const serverErrMsg = "Server is down, we are working towards fixing it"
const invalidParams = "Invalid params"
const invalidQuery = "Invalid query"
const invalidRequestBody = "Invalud request body. Please check"

type ErrResp struct {
	Err string `json:"error"`
}

func newErrResp(errMsg string) ErrResp {
	return ErrResp{
		errMsg,
	}
}
