package auth

import (
	"apiserver/internal/usecase/componet/auth/credential"
	"bytes"
	"common/response"
	"common/util"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// CallbackValidator send a http request to validte access, will cost 40ms+ delay
type CallbackValidator struct {
	cfg *CallbackConfig
}

func NewCallbackValidator(cfg *CallbackConfig) *CallbackValidator {
	return &CallbackValidator{cfg: cfg}
}

func (cv *CallbackValidator) Verify(token Credential) error {
	body := bytes.NewBuffer([]byte(token.GetUsername()))
	uri := fmt.Sprint(cv.cfg.Url, "?", url.Values(token.GetExtra()).Encode())
	resp, err := http.Post(uri, "application/json", body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return response.NewError(resp.StatusCode, response.MessageFromJSONBody(resp.Body))
	}
	return nil
}

func (cv *CallbackValidator) Middleware(c *gin.Context) (bool, error) {
	if !cv.cfg.Enable {
		return false, nil
	}
	var bucket string
	idx := strings.IndexByte(c.Request.Host, '.')
	if idx > 0 {
		bucket = c.Request.Host[:idx]
	} else if bucket = c.GetHeader("Bucket"); bucket == "" {
		return false, response.NewError(http.StatusBadRequest, "Host or Header does not contains bucket")
	}
	token := &credential.CallbackToken{
		Bucket:   bucket,
		Region:   c.GetHeader("Region"),
		FileName: c.Param("name"),
		Version:  util.ToInt(c.Query("version")),
		Method:   c.Request.Method,
		Extra:    make(map[string][]string),
	}
	for _, key := range cv.cfg.Params {
		arr, _ := c.GetQueryArray(key)
		if val := c.GetHeader(key); val != "" {
			arr = append(arr, val)
		}
		token.Extra[key] = arr
	}
	err := cv.Verify(token)
	return err == nil, err
}
