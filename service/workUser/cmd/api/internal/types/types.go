// Code generated by goctl. DO NOT EDIT.
package types

type CreateUserReq struct {
	Name string `json:"name"`
	Pwd  string `json:"name"`
}

type CreateUserResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type FindUserResp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type GetUserListReq struct {
	Keyword string `json:"keyword"`
	Page    int    `json:"page"`
	Size    int    `json:"size"`
}

type GetUserListResp struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
}

type Request struct {
	Name string `path:"name,options=you|me"`
}

type Response struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}
