package client

import (
	"net/url"
	"strings"
)

func HasPort(s string) bool { return strings.LastIndex(s, ":") > strings.LastIndex(s, "]") }

func validateWsConnectionURL(str1 string) bool {
	Uri, _ := url.Parse(str1)
	return Uri.IsAbs() && HasPort(Uri.Host) && (Uri.Scheme == "ws" || Uri.Scheme == "wss")
}
