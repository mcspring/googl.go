// Copyright 2011 Mc.Spring. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"
	"bytes"
	"strings"
	"io/ioutil"
	"json"
	"http"
	"fmt"
)

const (
	GOOGL_V1 = "https://www.googleapis.com/urlshortener/v1/url"
)

type Googl struct {
	Key string
}

// NewGoogl returns a new Googl, an API key is optionally
func NewGoogl(params ...string) *Googl {
	if 1 == len(params) {
		return &Googl{params[0]}
	}

	return &Googl{}
}

// Shorten shortens long url to short url.
// It returns the shortened url and nil if success, or empty
// string and an os.Error when failed.
func (req *Googl) Shorten(url string) (string, os.Error) {
	res, err := req.post(url)
	if err != nil {
		return "", err
	}

	var temp map[string]string
	if err = json.Unmarshal([]byte(res), &temp); err != nil {
		return "", os.Error(err)
	}

	if _, ok := temp["code"]; ok {
		return "", os.NewError(temp["message"])
	}

	id, ok := temp["id"]
	if !ok {
		return "", os.NewError("Invalid response!")
	}

	return id, nil
}

// Expand expands a shortened url to long url.
// It returns the orgin url and nil if success, or empty string 
// and an os.Error when failed.
func (req *Googl) Expand(url string) (string, os.Error) {
	res, err := req.get(url)
	if err != nil {
		return "", err
	}

	var temp map[string]string
	if err = json.Unmarshal([]byte(res), &temp); err != nil {
		return "", os.Error(err)
	}

	if _, ok := temp["code"]; ok {
		return "", os.NewError(temp["message"])
	}

	if status, _ := temp["status"]; "OK" != status {
		return "", os.NewError(status)
	}

	url, ok := temp["longUrl"]
	if !ok {
		return "", os.NewError("Invalid response")
	}

	return url, nil
}

// Project like Expand above.
// But this will return with short URL's analytics as a 
// map[string]interface{} type.
func (req *Googl) Project(url, proj string) (map[string]interface{}, os.Error) {
	proj = strings.ToUpper(proj)
	if "FULL"!=proj && "ANALYTICS_CLICKS"!=proj && "ANALYTICS_TOP_STRINGS"!=proj {
		return nil, os.EINVAL
	}

	res, err := req.get(url, map[string]string{"projection": proj})
	if err != nil {
		return nil, err
	}

	var temp map[string]interface{}
	if err = json.Unmarshal([]byte(res), &temp); err != nil {
		return nil, os.Error(err)
	}

	if _, ok := temp["code"]; ok {
		msg, _ := temp["message"].(string)

		return nil, os.NewError(msg)
	}

	if status, _ := temp["status"]; "OK" != status {
		msg, _ := status.(string)

		return nil, os.NewError(msg)
	}

	return temp, nil
}

func (req *Googl) History() {
	// TODO
}

func (req *Googl) get(url string, params ...map[string]string) (string, os.Error) {
	if !strings.Contains(url, "://goo.gl/") {
		return "", os.EINVAL
	}

	req_url := GOOGL_V1 + "?shortUrl=" + url
	if 0 < len(params) {
		for i:=0; i<len(params); i++ {
			req_url += "&" + toQuery(params[i])
		}
	}

	if "" != req.Key {
		req_url += "&key=" + req.Key
	}

	res, _, err := http.Get(req_url)
	defer res.Body.Close()
	if err != nil {
		return "", os.Error(err)
	}

	body, _ := ioutil.ReadAll(res.Body)

	return string(body), nil
}

func (req *Googl) post(url string) (string, os.Error) {
	if strings.Contains(url, "://goo.gl/") {
		return fmt.Sprintf("{\n \"kind\": \"urlshortener#url\",\n \"id\": \"%s\",\n \"longUrl\": \"%s\"\n}", url, url), nil
	}

	req_url := GOOGL_V1
	if "" != req.Key {
		req_url += "?key=" + req.Key
	}

	buf := bytes.NewBuffer(nil)
	buf.WriteString(fmt.Sprintf(`{"longUrl": "%s"}`, url))

	res, err := http.Post(req_url, "application/json", buf)
	if err != nil {
		return "", os.Error(err)
	}

	body, _ := ioutil.ReadAll(res.Body)

	return string(body), nil
}

func toQuery(params map[string]string) string {
	buf := bytes.NewBuffer(nil)

	flag := false
	for key, val := range params {
		if flag {
			buf.WriteString("&" + key + "=" + http.URLEscape(val))
		} else {
			flag = true
			buf.WriteString(key + "=" + http.URLEscape(val))
		}
	}

	return buf.String()
}
