package request

import (
	"aviasales/internal/request/header"
	"golang.org/x/net/context"
	"golang.org/x/net/context/ctxhttp"
	"io"
	"net/http"
	"net/url"
)

type Request struct {
	method string
	url    string
	client *Client
	header *header.Header
	params url.Values
	body   io.Reader
}

func New() *Request {
	return &Request{
		client: &Client{},
	}
}

func Get(url string) *Request {
	return New().Method(http.MethodGet).URL(url)
}

func (r *Request) Method(method string) *Request {
	r.method = method
	return r
}

func (r *Request) URL(url string) *Request {
	r.url = url
	return r
}

func (r *Request) Header(h *header.Header) *Request {
	r.header = h
	return r
}

func (r *Request) SetToken(key, val string) *Request {
	r.header.Set(key, val)
	return r
}

func (r *Request) Do(ctx context.Context) (*Response, error) {
	path, err := r.prepareURL()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(r.method, path, r.body)
	if err != nil {
		return nil, err
	}
	req.Header = r.header.Header

	response, err := ctxhttp.Do(ctx, r.client.Client, req)
	if err != nil {
		return nil, err
	}

	return &Response{response}, nil
}

func (r *Request) prepareURL() (path string, err error) {
	URL, err := url.Parse(r.url)
	if err != nil {
		return "", err
	}
	if r.params != nil {
		URL.RawQuery = r.params.Encode()
	}

	return URL.String(), nil
}
