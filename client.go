package gic

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"errors"
)

var Version = "1.0.0"
var DefaultUserAgent = fmt.Sprintf("gic-cloud-sdk-go (%s; %s) Golang/%s Core/%s", runtime.GOOS, runtime.GOARCH, strings.Trim(runtime.Version(), "go"), Version)

type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	Token      string
	httpClient *http.Client
}

type LoginResult struct {
	Status      string `json:"status"`
	AccessToken string `json:"Access-Token"`
}

// the constructor for Client struct
// http://api2.capitalonline.net
func (c *Client) Init(baseUrl string, userAgent string) error {

	burl, err := url.Parse(baseUrl)
	if err != nil {
		log.Fatal(err)
		return err
	}
	c.BaseURL = burl
	if userAgent != "" {
		c.UserAgent = userAgent
	} else {
		c.UserAgent = DefaultUserAgent
	}
	c.httpClient = &http.Client{}

	return nil
}

// login with token 
func (c *Client) LoginWithToken(token string) error {
	if token == "" {
		return errors.New("token is empty")
	}
	c.Token =  token
	c.UserAgent = DefaultUserAgent
	return nil
}

// login with username and password
func (c *Client) Login(username string, pwd string) error {
	rel := &url.URL{Path: "/gic/v1/get_token/"}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Printf("unable to new http request with error: %+v\n", err)
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("username", username)
	req.Header.Set("password", pwd)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("failed to login with error: %+v\n", err)
		return err
	}
	defer resp.Body.Close()
	var loginResult LoginResult
	err = json.NewDecoder(resp.Body).Decode(&loginResult)
	if err != nil {
		log.Printf("failed to parse login result with error: %+v\n", err)
		return err
	}
	if loginResult.AccessToken != "" {
		// fmt.Printf("token is %s\n", loginResult.AccessToken)
		c.Token = loginResult.AccessToken
	} else {
		log.Println("no token found")
	}
	return nil
}

// list the gic datacenter
func (c *Client) ListDataCenter() ([]DataCenter, error) {
	var dataCenters []DataCenter
	var dataCenterData DataCenterData
	rel := &url.URL{Path: "gic/v1/app/list/"}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Printf("unable to new http request with error: %+v\n", err)
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("token", c.Token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("failed to login with error: %+v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&dataCenterData)
	// err = json.NewDecoder(resp.Body).Decode(&dataCenterData)
	if err != nil {
		log.Printf("failed to parse list datacenter result with error: %+v\n", err)
		return nil, err
	}
	dataCenters = dataCenterData.Data

	return dataCenters, nil
}
