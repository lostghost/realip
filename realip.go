package realip

import (
	"log"
	"net"
	"net/http"
	"strings"
)

// Options type provides a type for specifying options used for parsing
// the real ip address from requests
type Options struct {
	TrustedNetworks []string
}

// Parse will parse a request with the provided options to return the
// real ip for the request
func Parse(r *http.Request, options Options) string {
	var po parseOptions
	for _, network := range options.TrustedNetworks {
		_, ipNet, err := net.ParseCIDR(network)
		if err != nil {
			log.Println(err)
		}
		if err == nil {
			po.trustedNetworks = append(po.trustedNetworks, ipNet)
		}
	}
	return parse(r, po)
}

type parseOptions struct {
	trustedNetworks []*net.IPNet
}

func parse(r *http.Request, options parseOptions) string {
	var ipChain []string

	xForwardedFor := r.Header.Get("x-forwarded-for")
	if len(xForwardedFor) > 0 {
		ipChain = strings.Split(xForwardedFor, ", ")
	}
	ipChain = append(ipChain, r.RemoteAddr)

	return lastUntrusted(ipChain, options.trustedNetworks)
}

func lastUntrusted(chain []string, networks []*net.IPNet) string {
	for i := len(chain) - 1; i > 0; i-- {
		if isTrusted(chain[i], networks) == false {
			return chain[i]
		}
	}
	return chain[0]
}

func isTrusted(address string, networks []*net.IPNet) bool {
	if host, _, err := net.SplitHostPort(address); err == nil {
		address = host
	}
	if ip := net.ParseIP(address); ip != nil {
		for _, network := range networks {
			if network.Contains(ip) {
				return true
			}
		}
	}
	return false
}
