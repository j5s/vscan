package pkg

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"sort"
	"strings"
	"time"
)

type Response struct {
	Status        string
	StatusCode    int
	Body          string
	Header        http.Header
	ContentLength int
	RequestUrl    string
	Location      string
}

var HttpProxy string

func HttpRequsetBasic(username string, password string, urlstring string, method string, postdata string, isredirect bool, headers map[string]string) (*Response, error) {
	var tr *http.Transport
	if HttpProxy != "" {
		uri, _ := url.Parse(HttpProxy)
		tr = &http.Transport{
			MaxIdleConnsPerHost: -1,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives:   true,
			Proxy:               http.ProxyURL(uri),
		}
	} else {
		tr = &http.Transport{
			MaxIdleConnsPerHost: -1,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives:   true,
		}
	}

	client := &http.Client{
		Timeout:   time.Duration(10) * time.Second,
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	if isredirect {
		jar, _ := cookiejar.New(nil)
		client = &http.Client{
			Timeout:   time.Duration(10) * time.Second,
			Transport: tr,
			Jar:       jar,
		}
	}
	req, err := http.NewRequest(strings.ToUpper(method), urlstring, strings.NewReader(postdata))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	for v, k := range headers {
		req.Header[v] = []string{k}
	}
	resp, err := client.Do(req)
	if err != nil {
		//防止空指针
		return &Response{"999", 999, "", nil, 0, "", ""}, err
	}
	var location string
	var reqbody string
	defer resp.Body.Close()
	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		reqbody = string(body)
	}
	if resplocation, err := resp.Location(); err == nil {
		location = resplocation.String()
	}
	return &Response{resp.Status, resp.StatusCode, reqbody, resp.Header, len(reqbody), resp.Request.URL.String(), location}, nil
}

func HttpRequset(urlstring string, method string, postdata string, isredirect bool, headers map[string]string) (*Response, error) {
	var tr *http.Transport
	if HttpProxy != "" {
		uri, _ := url.Parse(HttpProxy)
		tr = &http.Transport{
			MaxIdleConnsPerHost: -1,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives:   true,
			Proxy:               http.ProxyURL(uri),
		}
	} else {
		tr = &http.Transport{
			MaxIdleConnsPerHost: -1,
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives:   true,
		}
	}

	client := &http.Client{
		Timeout:   time.Duration(10) * time.Second,
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		}}
	if isredirect {
		jar, _ := cookiejar.New(nil)
		client = &http.Client{
			Timeout:   time.Duration(10) * time.Second,
			Transport: tr,
			Jar:       jar,
		}
	}
	req, err := http.NewRequest(strings.ToUpper(method), urlstring, strings.NewReader(postdata))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")
	for v, k := range headers {
		req.Header[v] = []string{k}
	}
	resp, err := client.Do(req)
	if err != nil {
		//防止空指针
		return &Response{"999", 999, "", nil, 0, "", ""}, err
	}
	var location string
	var reqbody string
	defer resp.Body.Close()
	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		reqbody = string(body)
	}
	if resplocation, err := resp.Location(); err == nil {
		location = resplocation.String()
	}
	return &Response{resp.Status, resp.StatusCode, reqbody, resp.Header, len(reqbody), resp.Request.URL.String(), location}, nil
}

func IntInSlice(i int, slice []int) bool {
	if slice == nil {
		return false
	}
	sort.Ints(slice)
	index := sort.SearchInts(slice, i)
	if index < len(slice) && slice[index] == i {
		return true
	}
	return false
}

func StringInSlice(str string, slice []string) bool {
	if slice == nil {
		return false
	}
	sort.Strings(slice)
	index := sort.SearchStrings(slice, str)
	if index < len(slice) && slice[index] == str {
		return true
	}
	return false
}

func SliceInString(str string, slice []string) bool {
	if slice == nil {
		return false
	}
	for _, v := range slice {
		if strings.Contains(str, v) {
			return true
		}
	}
	return false
}
