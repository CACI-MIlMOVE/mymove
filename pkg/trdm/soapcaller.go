package trdm

import "github.com/tiaguinho/gosoap"

type SoapCaller interface {
	Call(m string, p gosoap.SoapParams) (res *gosoap.Response, err error)
}
