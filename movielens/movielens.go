package movielens

import (
	"errors"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/mozillazg/request"
)

const (
	defaultBaseURL = "https://movielens.org/api"
	requestTimeout = 30000

	headerAccept         = "application/json, text/plain, */*"
	headerAcceptEncoding = "gzip, deflate"
	headerAcceptLanguage = "en-US,en;q=0.5"
	headerCacheControl   = "no-cache"
	headerContentType    = "application/json;charset=utf-8"
	headerDNT            = "1"
	headerHost           = "movielens.org"
	headerPragma         = "no-cache"
	headerUserAgent      = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.10; rv:42.0) Gecko/20100101 Firefox/42.0"
)

// A Client manages communication with the Movielens API.
type Client struct {
	// HTTP client used to communicate with the API.
	client *http.Client
}

// NewClient returns a new API client. If a nil httpClient is
// provided, a default client will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		cookieJar, _ := cookiejar.New(nil)
		httpClient = &http.Client{
			Jar: cookieJar,
		}
	}
	c := &Client{client: httpClient}
	return c
}

// Login logs the user in and set cookies if successful, returns
// error otherwise
func (c *Client) Login(username string, password string) (*Client, error) {
	url := strings.Join([]string{defaultBaseURL, "sessions"}, "/")
	req := request.NewRequest(c.client)
	req.Headers = map[string]string{
		"Accept":          headerAccept,
		"Accept-Encoding": headerAcceptEncoding,
		"Accept-Language": headerAcceptLanguage,
		"Cache-Control":   headerCacheControl,
		"Content-Type":    headerContentType,
		"DNT":             headerDNT,
		"Host":            headerHost,
		"Pragma":          headerPragma,
		"User-Agent":      headerUserAgent,
	}
	req.Json = map[string]string{
		"userName": username,
		"password": password,
	}
	resp, err := req.Post(url)
	if err != nil {
		return nil, err
	}
	if !resp.OK() {
		return nil, errors.New(resp.Response.Status)
	}
	return c, nil
}
