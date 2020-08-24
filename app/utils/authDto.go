package utils

import "net/url"

//一旦最低限で！！

//認証リクエスト Getの場合はqueryStringで　Postの場合は application/x-www-form-urlencodedでbodyに含める
type AuthenticationRequestDto struct {
	//　openid を必ず含める その他公開したいパラメータを含める
	Scope string `json:"scope"`
	// だいたい code でいい 認証(IDtoken) => 認可(access_token) => userInfo
	ResponseType string `json:"response_type"`
	// クライアントを識別する 事前に登録するイメージ??
	ClientId string `json:"client_id"`
	//　認証後に返すURL これも認証サーバーに事前に登録する必要がある
	RedirectURI string `json:"redirect_url"`
	State       string `json:"state"`
}

func (dto AuthenticationRequestDto) GetQuery() string {
	loginPagePath := "/"
	query := url.Values{}
	query.Add("scope", dto.Scope)
	query.Add("redirect_uri", dto.RedirectURI)
	query.Add("response_type", dto.ResponseType)
	query.Add("client_id", dto.ClientId)
	query.Add("state", dto.State)
	return loginPagePath + "?" + query.Encode()
}

//認証リクエストをバインドさせる
type LoginRequestDto struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	RedirectURI string `json:"redirect_uri"`
	State       string `json:"state"`
}

//認可リクエスト
type AuthorizationRequestDto struct{}

//認証認可フローでキャッシュが必要になった場合にやり取りする入れ物
type AuthFlowInfo struct {
	Code        string `redis:"code"`
	State       string `redis:"state"`
	RedirectURI string `redis:"redirect_uri"`
}

func (info AuthFlowInfo) GetQuery() string {
	url := url.Values{}
	url.Set("code", info.Code)
	url.Set("state", info.State)
	url.Set("redirect_uri", info.RedirectURI)
	return info.RedirectURI + "?" + url.Encode()
}
