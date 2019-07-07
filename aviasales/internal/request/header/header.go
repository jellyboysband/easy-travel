package header

import "net/http"

type Header struct {
	http.Header
}

func New(key, val string) *Header {
	header := http.Header{}
	header.Set(key, val)
	return &Header{header}
}

func (h *Header) Set(key, val string) *Header {
	h.Header.Set(key, val)
	return h
}

func (h *Header) Add(key, val string) *Header {
	h.Header.Add(key, val)
	return h
}

func (h *Header) Get(key string) string {
	return h.Header.Get(key)
}

func (h *Header) Del(key string) *Header {
	h.Header.Del(key)
	return h
}

func (h *Header) Clean() *Header {
	h.Header = make(map[string][]string)
	return h
}