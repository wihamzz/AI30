package restserv

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rest_tp5/comsoc"
)

type RestClientAgent struct {
	id       string
	url      string
	alt 	 comsoc.Alternative
}

func NewRestClientAgent(id string, url string, alt comsoc.Alternative) *RestClientAgent {
	return &RestClientAgent{id, url, alt}
}

func (rca *RestClientAgent) treatResponse(r *http.Response) int {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)

	var resp Response
	json.Unmarshal(buf.Bytes(), &resp)

	return resp.
}


