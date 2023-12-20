package types

type SignatureData struct {
	PenandaTangan string `json:"penanda-tangan"`
	DocName       string `json:"doc-name"`
	NpmMahasiswa  string `json:"npm-mahasiswa"`
}

type RequestData struct {
	Id   string        `json:"id"`
	Data SignatureData `json:"data"`
}

type TokenResp struct {
	Token string `json:"token"`
}
