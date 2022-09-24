package ftx

import (
	"net/http"
	"net/url"
)

type FtxClient struct {
	Client     *http.Client
	Api        string
	Secret     []byte
	Subaccount string
}

func New(api string, secret string, subaccount string) *FtxClient {
	return &FtxClient{Client: &http.Client{}, Api: api, Secret: []byte(secret), Subaccount: url.PathEscape(subaccount)}
}

func (client *FtxClient) ProxyDo(req *http.Request) (*http.Response, error) {
	proxyURL, _ := url.Parse("http://127.0.0.1:7890")
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	httpClient := &http.Client{
		Transport: transport,
	}
	resp, erro := httpClient.Do(req)
	return resp, erro
}
