package types

type SignatureData struct {
	PenandaTangan   string `json:"penanda-tangan"`
	DocName         string `json:"doc-name"`
	PemilikDocument string `json:"pemilik_document"`
}

type SignatureDataNew struct {
	PenandaTangan   string `json:"penanda-tangan"`
	DocName         string `json:"doc-name"`
	PemilikDocument string `json:"pemilik_document"`
	Universitas     string `json:"universitas"`
}

type RequestData struct {
	Id   string        `json:"id"`
	Data SignatureData `json:"data"`
}

type RequestDataNew struct {
	Id   string           `json:"id"`
	Data SignatureDataNew `json:"data"`
}
type TokenResp struct {
	Token string `json:"token"`
}
