package restserv

import (
	"rest_tp5/comsoc"
)

type Request struct {
	Alt comsoc.Alternative `json:"alt"`
}

type Response struct {
	List comsoc.Count `json:"list"`
	Winner comsoc.Alternative `json:"winner"`
}
