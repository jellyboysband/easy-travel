package request

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
)

type Response struct {
	*http.Response
}

func (r *Response) Body() ([]byte, error) {
	return readBody(r.Response.Body)
}

func (r *Response) JSON(v interface{}) error {
	buf, err := readBody(r.Response.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(buf, v)
}

func (r *Response) XML(v interface{}) error {
	buf, err := readBody(r.Response.Body)
	if err != nil {
		return err
	}

	return xml.Unmarshal(buf, v)
}

func readBody(body io.ReadCloser) ([]byte, error) {
	buf, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	return buf, body.Close()
}
