package gonion

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"encoding/json"
  "log"
)

type Client struct {
    UserAgent string
    HttpClient *http.Client
}

func (c *Client) Summary(args Params) SSummary {
  req, err := c.SendRequest("/summary", args)

  resp, err := c.HttpClient.Do(req)
  if err != nil {
    log.Fatal(err)
  }

  defer resp.Body.Close()

  if err != nil {
    log.Fatal(err)
  }
  var Sum SSummary
  body, err := ioutil.ReadAll(resp.Body)
  json.Unmarshal([]byte(body), &Sum)
  return Sum
}

func (c *Client) Details(args Params) SDetails {

  req, err := c.SendRequest("/details", args)

  resp, err := c.HttpClient.Do(req)
  if err != nil {
    log.Fatal(err)
  }

  defer resp.Body.Close()

  if err != nil {
    log.Fatal(err)
  }
  var Det SDetails
  body, err := ioutil.ReadAll(resp.Body)
  json.Unmarshal([]byte(body), &Det)
  return Det

}

func (c *Client) Bandwidth(args Params) SBandwidth {

  req, err := c.SendRequest("/bandwidth", args)

  resp, err := c.HttpClient.Do(req)
  if err != nil {
    log.Fatal(err)
  }

  defer resp.Body.Close()

  if err != nil {
    log.Fatal(err)
  }
  var Ban SBandwidth
  body, err := ioutil.ReadAll(resp.Body)
  json.Unmarshal([]byte(body), &Ban)
  return Ban

}
func (c *Client) Weights(args Params) SWeights {

  req, err := c.SendRequest("/weights", args)

  resp, err := c.HttpClient.Do(req)
  if err != nil {
    log.Fatal(err)
  }

  defer resp.Body.Close()

  if err != nil {
    log.Fatal(err)
  }
  var Wei SWeights
  body, err := ioutil.ReadAll(resp.Body)
  json.Unmarshal([]byte(body), &Wei)
  return Wei

}
func (c *Client) Clients(args Params) SClients {

  req, err := c.SendRequest("/clients", args)

  resp, err := c.HttpClient.Do(req)
  if err != nil {
    log.Fatal(err)
  }

  defer resp.Body.Close()

  if err != nil {
    log.Fatal(err)
  }
  var Cli SClients
  body, err := ioutil.ReadAll(resp.Body)
  json.Unmarshal([]byte(body), &Cli)
  return Cli

}
func (c *Client) Uptime(args Params) SUptime {

  req, err := c.SendRequest("/bandwidth", args)

  resp, err := c.HttpClient.Do(req)
  if err != nil {
    log.Fatal(err)
  }

  defer resp.Body.Close()

  if err != nil {
    log.Fatal(err)
  }
  var Upt SUptime
  body, err := ioutil.ReadAll(resp.Body)
  json.Unmarshal([]byte(body), &Upt)
  return Upt

}

func (c *Client)SendRequest(path string, args Params) (*http.Request, error){
	  BaseURL, err := url.Parse("https://onionoo.torproject.org")
    if err != nil {
      log.Fatal(err)
    }
    rel := &url.URL{Path: path}
    u := BaseURL.ResolveReference(rel)
    req, err := http.NewRequest("GET", u.String(), nil)
    if err != nil {
      log.Fatal(err)
    }
    req.Header.Set("Accept", "application/json")
    req.Header.Set("User-Agent", c.UserAgent)
    q, err := args.QueryParams()
    if err != nil{
      log.Fatal(err)
    }
    req.URL.RawQuery = q.Encode()

    return req, nil
}
