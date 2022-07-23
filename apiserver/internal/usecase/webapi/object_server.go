package webapi

import (
	"apiserver/internal/usecase/pool"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

func DeleteTmpObject(locate, id string) {
	req, _ := http.NewRequest(http.MethodDelete, tempRest(locate, id), nil)
	resp, e := pool.Http.Do(req)
	if resp.StatusCode == http.StatusBadRequest {
		if content, e := io.ReadAll(resp.Body); e == nil {
			log.Errorf("patch temp object id=%v, return code=%v\n", id, string(content))
		}
	}
	if e != nil || resp.StatusCode != http.StatusOK {
		log.Println(e, resp.StatusCode)
	}
}

func PostTmpObject(ip, name string, size int64) (string, error) {
	req, _ := http.NewRequest(http.MethodPost, tempRest(ip, name), nil)
	req.Header.Add("Size", fmt.Sprint(size))
	resp, e := pool.Http.Do(req)
	if e != nil {
		return "", e
	}
	res, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return "", fmt.Errorf("post temp object name=%v, return error response body, status=%v", name, resp.StatusCode)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("post temp object name=%v, return code=%v, content=%s", name, resp.Status, string(res))
	}
	return string(res), nil
}

func PatchTmpObject(ip, id string, body io.Reader) error {
	req, _ := http.NewRequest(http.MethodPatch, tempRest(ip, id), body)
	resp, e := pool.Http.Do(req)
	if e != nil {
		return e
	}
	if resp.StatusCode == http.StatusBadRequest {
		if content, e := io.ReadAll(resp.Body); e == nil {
			return fmt.Errorf("patch temp object id=%v, return content=%v", id, string(content))
		}
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("patch temp object id=%v, return code=%v", id, resp.Status)
	}
	return nil
}

func PutTmpObject(ip, id, name string) error {
	form := make(url.Values)
	form.Set("name", name)
	req, _ := http.NewRequest(http.MethodPut, tempRest(ip, id), strings.NewReader(form.Encode()))
	resp, e := pool.Http.Do(req)
	if e != nil {
		return e
	}
	if resp.StatusCode == http.StatusBadRequest {
		if content, e := io.ReadAll(resp.Body); e == nil {
			return fmt.Errorf("patch temp object id=%v, return content=%v", id, string(content))
		}
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("put temp object id=%v, return code=%v", id, resp.Status)
	}
	return nil
}

func HeadTmpObject(ip, id string) (int64, error) {
	resp, e := http.Head(tempRest(ip, id))
	if e != nil {
		return 0, e
	}
	if resp.StatusCode == http.StatusBadRequest {
		if content, e := io.ReadAll(resp.Body); e == nil {
			return 0, fmt.Errorf("patch temp object id=%v, return content=%v", id, string(content))
		}
	}
	if resp.StatusCode == http.StatusNotFound {
		return 0, nil
	}
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("head temp object id=%v, return code=%v", id, resp.Status)
	}
	if str := resp.Header.Get("Size"); len(str) > 0 {
		size, e := strconv.ParseInt(str, 10, 0)
		if e != nil {
			return 0, fmt.Errorf("parse size string %s error: %v", str, e)
		}
		return size, nil
	}
	return 0, fmt.Errorf("response doesn't contains size")
}

func GetObject(ip, name string) (*http.Response, error) {
	return pool.Http.Get(objectRest(ip, name))
}

func objectRest(ip, id string) string {
	return fmt.Sprintf("http://%s/objects/%s", ip, id)
}

func tempRest(ip, id string) string {
	return fmt.Sprintf("http://%s/temp/%s", ip, id)
}