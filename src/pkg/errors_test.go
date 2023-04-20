package pkg

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"net/http"
	"testing"
)

func TestError(t *testing.T) {
	cookie := make([]byte, 10)
	_, _ = rand.Read(cookie)

	responseErr := NewResponseErr(
		marshalParseErr,
		"Method not allowed",
		http.StatusText(http.StatusMethodNotAllowed),
		http.StatusMethodNotAllowed,
		map[string]string{
			"cookie": base64.URLEncoding.EncodeToString(cookie),
		})

	errs := responseErr.Error()
	if errs != "" {
		t.Logf("marshal ret is:[%+v]", errs)
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(errs)
	if err != nil {
		panic(err)
	}

	t.Logf("buf bytes:%v", buf.Bytes())

	t.Logf("marshal szie:[%d], buf size:[%d]", len(errs), buf.Len())
}

func NewResponseErr(code int, p string, err string, statusCode int, headers map[string]string) *ResponseError {
	return &ResponseError{
		Code:       code,
		Err:        err,
		Type:       p,
		StatusCode: statusCode,
		Headers:    headers,
	}
}

type ResponseError struct {
	Code       int               `json:"code"`
	Err        string            `json:"err"`
	Type       string            `json:"type"`
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers"`
}

const marshalParseErr = 1001

func (re *ResponseError) Error() string {
	ret, err := json.Marshal(*re)
	if err != nil {
		return statusText(re.Code)
	}

	return string(ret)
}

func statusText(code int) string {
	switch code {
	case marshalParseErr:
		return "marshal parse is error"
	default:
		return "invalid status code"
	}
}
