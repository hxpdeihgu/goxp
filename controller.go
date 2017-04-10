package goxp

import "encoding/base64"

type Model struct {
	
}

func (this *Model) Md5(s string) string {
	return _md5(s)
}

func (this *Model) Sha1(s string) string {
	return _sha1(s)
}

func (this *Model) Sha256(s string) string {
	return _sha256(s)
}

func (this *Model) Sha512(s string) string {
	return _sha512(s)
}

func (this *Model) Base64Decode(s string) (string,error) {
	v,err:=base64.StdEncoding.DecodeString(s)
	if err !=nil {
		return "",err
	}
	return string(v),nil
}

func (this *Model) Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}