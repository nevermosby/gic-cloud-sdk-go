package gic

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"runtime"
	"strings"
)

const (
	CapitalonlineURL = "http://api2.capitalonline.net"
)

var Version = "1.0.0"
var DefaultUserAgent = fmt.Sprintf("gic-cloud-sdk-go (%s; %s) Golang/%s Core/%s", runtime.GOOS, runtime.GOARCH, strings.Trim(runtime.Version(), "go"), Version)

// Client for login
type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	Token      string
	httpClient *http.Client
}

// LoginResult for gic login result
type LoginResult struct {
	Status      string `json:"status"`
	AccessToken string `json:"Access-Token"`
}

// Init as the constructor for Client struct
// Accept baseURL and userAgent as parameters
func (c *Client) Init(baseURL string, userAgent string) error {
	if baseURL != "" {
		burl, err := url.Parse(baseURL)
		if err != nil {
			log.Fatal(err)
			return err
		}
		c.BaseURL = burl
	} else {
		burlReserved, _ := url.Parse(CapitalonlineURL)
		c.BaseURL = burlReserved
	}

	if userAgent != "" {
		c.UserAgent = userAgent
	} else {
		c.UserAgent = DefaultUserAgent
	}
	c.httpClient = &http.Client{}

	return nil
}

// LoginWithToken for login with existed token
func (c *Client) LoginWithToken(token string) error {
	if token == "" {
		err := errors.New("token is empty")
		log.Fatal(err)
		return err
	}
	c.Token = token
	c.UserAgent = DefaultUserAgent
	return nil
}

// Login with username and password
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

// ListDataCenter for list the gic virtual datacenters
func (c *Client) ListDataCenter() ([]DataCenterListData, error) {
	var dataCenterListData []DataCenterListData
	var dataCenterData DataCenterList
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
		log.Printf("failed to ListDataCenter with error: %+v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&dataCenterData)
	// err = json.NewDecoder(resp.Body).Decode(&dataCenterData)
	if err != nil {
		log.Printf("failed to parse ListDataCenter result with error: %+v\n", err)
		return nil, err
	}
	dataCenterListData = dataCenterData.Data

	return dataCenterListData, nil
}

// InfoDataCenter for show the single gic virtual datacenter info
func (c *Client) InfoDataCenter(datacenterID string) ([]DataCenterInfoData, error) {
	var dataCenterInfoDataList []DataCenterInfoData
	var dataCenterInfo DataCenterInfo
	rel := &url.URL{Path: "gic/v1/app/info"}
	u := c.BaseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Printf("unable to new http request with error: %+v\n", err)
		return dataCenterInfoDataList, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("token", c.Token)

	// add query params
	q := req.URL.Query()
	q.Add("app_id", datacenterID)
	req.URL.RawQuery = q.Encode()

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("failed to InfoDataCenter with error: %+v\n", err)
		return dataCenterInfoDataList, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&dataCenterInfo)
	if err != nil {
		log.Printf("failed to parse InfoDataCenter result with error: %+v\n", err)
		return dataCenterInfoDataList, err
	}
	dataCenterInfoDataList = dataCenterInfo.Data

	return dataCenterInfoDataList, nil
}

// ListOS for list the available os
func (c *Client) ListOS() ([]OSListData, error) {
	var osListData []OSListData
	var osList OSList
	rel := &url.URL{Path: "gic/v1/os/list/"}
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
		log.Printf("failed to list os with error: %+v\n", err)
		return nil, err
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&osList)
	// err = json.NewDecoder(resp.Body).Decode(&dataCenterData)
	if err != nil {
		log.Printf("failed to parse list os result with error: %+v\n", err)
		return nil, err
	}
	osListData = osList.Data

	return osListData, nil
}
