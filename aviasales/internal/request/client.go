package request

import "net/http"

type Client struct {
	*http.Client
}
