package signature

import (
	"encoding/json"
	"fmt"
	"github.com/aiteung/document/client"
	"github.com/aiteung/document/types"
	"github.com/skip2/go-qrcode"
)

func CreateToken(docid, url string, data types.SignatureData) (token string) {
	resp := new(types.TokenResp)
	body := new(types.RequestData)

	body.Id = docid
	body.Data = data

	res, err := client.CreateRequestHTTP().
		SetBody(body).
		Post(url)

	if err != nil {
		return "error ni kakak sistem akademiknya silahkan hubungi admin yaaaaa........"
	}

	defer res.Body.Close()
	_ = json.Unmarshal(res.Bytes(), &resp)

	token = resp.Token

	return token
}

func CreateQRCode(link string, filename string) error {
	// Generate QR code
	err := qrcode.WriteFile(link, qrcode.Highest, 256, filename)
	if err != nil {
		return err
	}

	fmt.Printf("QR code generated and saved to %s\n", filename)
	return nil
}

func CreateQRCodeMedium(link string, filename string) error {
	// Generate QR code
	err := qrcode.WriteFile(link, qrcode.Medium, 256, filename)
	if err != nil {
		return err
	}

	fmt.Printf("QR code generated and saved to %s\n", filename)
	return nil
}
