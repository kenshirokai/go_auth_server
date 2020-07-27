package utils

//一旦最低限で！！

//認証リクエスト Getの場合はqueryStringで　Postの場合は application/x-www-form-urlencodedでbodyに含める
type AuthenticationRequestDto struct {
	//　openid を必ず含める その他公開したいパラメータを含める
	Scope string `json: "scope"`
	// だいたい code でいい 認証(IDtoken) => 認可(access_token) => userInfo
	ResponseType string `json: "response_type"`
	// クライアントを識別する 事前に登録するイメージ??
	ClientId string `json: "client_id"`
	//　認証後に返すURL これも認証サーバーに事前に登録する必要がある
	RedirectURI string `json: "redirect_url"`
	State string `json: "state"`
}

//認可リクエスト
type AuthorizationRequestDto struct {}