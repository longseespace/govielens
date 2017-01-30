package movielens

import (
	"encoding/json"
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

func (c *Client) newRequest() *request.Request {
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
	return req
}

// Login logs the user in and set cookies if successful, returns
// error otherwise
func (c *Client) Login(username string, password string) error {
	url := strings.Join([]string{defaultBaseURL, "sessions"}, "/")
	req := c.newRequest()
	req.Json = map[string]string{
		"userName": username,
		"password": password,
	}
	resp, err := req.Post(url)
	if err != nil {
		return err
	}
	if !resp.OK() {
		return errors.New(resp.Response.Status)
	}
	return nil
}

// GetMe fetchs user details
// Gets your user information such as: Number of Ratings, Email, User Name, Preferences, and Recommender Type.
func (c *Client) GetMe() (*User, error) {
	url := strings.Join([]string{defaultBaseURL, "users/me"}, "/")
	req := c.newRequest()
	resp, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	if !resp.OK() {
		return nil, errors.New(resp.Response.Status)
	}
	j, err := resp.Json()
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	user := MakeUserFromJSON(j)
	return user, nil
}

// Explore search for movies.
func (c *Client) Explore(params map[string]string) (*ExploreResponse, error) {
	url := strings.Join([]string{defaultBaseURL, "movies/explore"}, "/")
	req := c.newRequest()
	req.Params = params
	resp, err := req.Get(url)
	if err != nil {
		return nil, err
	}
	if !resp.OK() {
		return nil, errors.New(resp.Response.Status)
	}
	b, err := resp.Content()
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	var exploreResponse ExploreResponseWrapper
	err = json.Unmarshal(b, &exploreResponse)
	return exploreResponse.Data, nil
}
