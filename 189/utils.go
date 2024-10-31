package _189

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	math_rand "math/rand"
	"net/http"
	"strings"
	"time"
)
import "regexp"

var (
	_shareCodePart1 = regexp.MustCompile(`https://cloud.189.cn/t/(\w+)`)
	_shareCodePart2 = regexp.MustCompile(`https://cloud.189.cn/web/share\?code=(\w+)`)
)
var b64map = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
var bI_RM = "0123456789abcdefghijklmnopqrstuvwxyz"

// 获取随机数
func random() string {
	return fmt.Sprintf("0.%17v", math_rand.New(math_rand.NewSource(time.Now().UnixNano())).Int63n(100000000000000000))
}
func int2char(a int) string {
	return strings.Split(bI_RM, "")[a]
}
func RsaEncode(origData []byte, j_rsakey string) string {
	publicKey := []byte("-----BEGIN PUBLIC KEY-----\n" + j_rsakey + "\n-----END PUBLIC KEY-----")
	block, _ := pem.Decode(publicKey)
	pubInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	pub := pubInterface.(*rsa.PublicKey)
	b, err := rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
	if err != nil {
		fmt.Printf("err: %s", err.Error())
	}
	return b64tohex(base64.StdEncoding.EncodeToString(b))
}
func b64tohex(a string) string {
	d := ""
	e := 0
	c := 0
	for i := 0; i < len(a); i++ {
		m := strings.Split(a, "")[i]
		if m != "=" {
			v := strings.Index(b64map, m)
			if 0 == e {
				e = 1
				d += int2char(v >> 2)
				c = 3 & v
			} else if 1 == e {
				e = 2
				d += int2char(c<<2 | v>>4)
				c = 15 & v
			} else if 2 == e {
				e = 3
				d += int2char(c)
				d += int2char(v >> 2)
				c = 3 & v
			} else {
				e = 0
				d += int2char(c<<2 | v>>4)
				d += int2char(15 & v)
			}
		}
	}
	if e == 1 {
		d += int2char(c << 2)
	}
	return d
}

func ParseShareCode(url string) string {
	var matched []string
	if matched = _shareCodePart1.FindStringSubmatch(url); len(matched) > 1 {
		return matched[1]
	}
	if matched = _shareCodePart2.FindStringSubmatch(url); len(matched) > 1 {
		return matched[1]
	}
	return ""
}
func cookiesToString(cookies []*http.Cookie) string {
	var result string
	for _, cookie := range cookies {
		result += cookie.Name + "=" + cookie.Value + ";"
	}
	return result
}
