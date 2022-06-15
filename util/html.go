package util

import (
	"github.com/microcosm-cc/bluemonday"
)

var p *bluemonday.Policy

func init() {
	p = bluemonday.UGCPolicy()
}

// AvoidXSS 避免XSS
func AvoidXSS(theHTML string) string {
	return p.Sanitize(theHTML)
}
