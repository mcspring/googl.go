// Copyright 2011 Mc.Spring. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goo

import (
	"os"
	"bytes"
	"io/ioutil"
	"http"
	"net"
	"strings"
)

const (
	GOO_URL_V1 = "https://www.googleapis.com/urlshortener/v1/url"
)

type GURLShort struct {
	APIKey string
	Params map[string]string
}

func (req *GURLShort) Get(url string, params ...map[string]string) (json string, os.Error) {
	if !strings.Contains(url, "goo.gl") {
		return "", os.EINVAL
	}

	if len(params) {
		for i:=0; i<len(params); i++ {
			key, val := range params[i]

			req.Params[key] = val
		}
	}

	req_url = GOO_URL_V1 + "?" + toQuery(req.Params)

	res, _, err := http.Get(req_url)
	defer res.Body.Close()
	if err != nil {
		return "", os.Errno(err)
	}

	body, _ := ioutil.ReadAll(res.Body)

	return string(body), nil
}

func (req *GURLShort) Post(url string, params map[string]string) (json string, os.Error) {
	// TODO
}

func toQuery(params map[string]string) string {
	flag := false
	buf := bytes.NewBuffer(nil)
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

type HTTPError struct {
	Errno int // HTTP response code
	Error string // HTTP response error message, empty if none error occurred
	Trace interface
}

func (err HTTPError) String() string {
	return err.Message
}

