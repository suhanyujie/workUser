// Code generated by goctl. DO NOT EDIT.
package types

type CreateUserReq struct {
	UserName string `json:"userName"`
	Pwd      string `json:"pwd"`
}

type CreateUserResp struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    OneUser `json:"data"`
}

type FindUserResp struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    OneUser `json:"data"`
}

type OneUser struct {
	Id       int64  `json:"id"`
	UserName string `json:"userName"`
}

type UserListReq struct {
	Ids      []int64 `json:"ids,optional"`
	NickName string  `json:"nickName,optional"`
}

type UserListRespVo struct {
	Code    int                `json:"code"`
	Message string             `json:"message"`
	Data    UserListRespVoData `json:"data"`
}

type UserListRespVoData struct {
	List  []OneUser `json:"list"`
	Total int64     `json:"total"`
	Page  int       `json:"page"`
	Size  int       `json:"size"`
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

type ResponseVo struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type IsOkResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
