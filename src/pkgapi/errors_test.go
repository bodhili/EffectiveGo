package pkgapi

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"os"
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

const marshalParseErr = 1001

type ResponseError struct {
	Code       int               `json:"code"`
	Err        string            `json:"err"`
	Type       string            `json:"type"`
	StatusCode int               `json:"status_code"`
	Headers    map[string]string `json:"headers"`
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

func TestErrorsJoin(t *testing.T) {
	err1 := errors.New("err1")
	err2 := errors.New("err2")
	err := errors.Join(err1, err2)

	if errors.Is(err, err1) {
		t.Log("err is err1")
	}

	if errors.Is(err, err2) {
		t.Log("err is err2")
	}
}

func TestErrorsUnWarp(t *testing.T) {
	err1 := errors.New("error1")
	err2 := fmt.Errorf("error2: [%w]", err1)
	fmt.Println(errors.Unwrap(err2))
}

func TestErrorAsAndIs(t *testing.T) {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *fs.PathError
		if errors.As(err, &pathError) {
			fmt.Println("Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}

		if errors.Is(err, fs.ErrNotExist) {
			fmt.Println("file does not exist")
		} else {
			fmt.Println(err)
		}
	}
}
