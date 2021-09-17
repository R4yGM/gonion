package gonion

import(
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Client struct {

    UserAgent string
    HttpClient *http.Client
}

func (c *Client) Summary() SSummary {
	BaseURL, err := url.Parse("https://onionoo.torproject.org")
    rel := &url.URL{Path: "/summary"}
    u := BaseURL.ResolveReference(rel)
    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
		panic(err)
    }
    req.Header.Set("Accept", "application/json")
    req.Header.Set("User-Agent", c.UserAgent)
    resp, err := c.HttpClient.Do(req)
    if err != nil {
		panic(err)
    }
    defer resp.Body.Close()
	var Sum SSummary
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(body), &Sum)
    return Sum
}

func (c *Client) Details() SDetails {
	BaseURL, err := url.Parse("https://onionoo.torproject.org")
    rel := &url.URL{Path: "/details"}
    u := BaseURL.ResolveReference(rel)
    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
		panic(err)
    }
    req.Header.Set("Accept", "application/json")
    req.Header.Set("User-Agent", c.UserAgent)
 
    resp, err := c.HttpClient.Do(req)
    if err != nil {
		panic(err)
    }
    defer resp.Body.Close()
	var Sum SDetails
	body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	json.Unmarshal([]byte(body), &Sum)
	//fmt.Println(Sum.Relays)
    return Sum
}

func Test() {
    fmt.Println("testing aaaaaaaa")
}

