package delivery

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/server/web"
)

// Cookie is cookie model
type Cookie struct {
	MaxAge   int64
	Path     string
	Secure   bool
	HttpOnly bool
	Domain   string
	SameSite string
	Value    string
	Name     string
}

var cookieNameSanitizer = strings.NewReplacer("\n", "-", "\r", "-")

// SetCookie is function to setting up cookie
func SetCookie(controller *web.Controller, cookie Cookie) {
	var b bytes.Buffer
	fmt.Fprintf(&b, "%s=%s", cookieNameSanitizer.Replace(cookie.Name), cookieNameSanitizer.Replace(cookie.Value))

	maxAge := cookie.MaxAge
	switch {
	case maxAge > 0:
		fmt.Fprintf(&b, "; Expires=%s; Max-Age=%d", time.Now().Add(time.Duration(maxAge)*time.Second).UTC().Format(time.RFC1123), maxAge)
	case maxAge < 0:
		fmt.Fprintf(&b, "; Max-Age=0")
	}

	if cookie.Path != "" {
		fmt.Fprintf(&b, "; Path=%s", cookieNameSanitizer.Replace(cookie.Path))
	} else {
		fmt.Fprintf(&b, "; Path=%s", "/")
	}

	if cookie.Domain != "" {
		fmt.Fprintf(&b, "; Domain=%s", cookieNameSanitizer.Replace(cookie.Domain))
	}

	if cookie.Secure {
		fmt.Fprintf(&b, "; Secure")
	}

	if cookie.HttpOnly {
		fmt.Fprintf(&b, "; HttpOnly")
	}

	if cookie.SameSite != "" {
		fmt.Fprintf(&b, "; SameSite="+cookie.SameSite)
	}

	controller.Ctx.Output.Context.ResponseWriter.Header().Add("Set-Cookie", b.String())
}

// SetSecureCookie is function to setting up secure cookie
func SetSecureCookie(controller *web.Controller, Secret string, cookie Cookie) {
	vs := base64.URLEncoding.EncodeToString([]byte(cookie.Value))
	timestamp := strconv.FormatInt(time.Now().UnixNano(), 10)
	h := hmac.New(sha256.New, []byte(Secret))
	fmt.Fprintf(h, "%s%s", vs, timestamp)
	sig := fmt.Sprintf("%02x", h.Sum(nil))
	cookie.Value = strings.Join([]string{vs, timestamp, sig}, "|")
	SetCookie(controller, cookie)
}
