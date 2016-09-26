package realip_test

import (
	"net/http"
	"testing"

	"github.com/lostghost/realip"
)

// optionsSets is a slice of realip.Options. Each set of options will
// be used to parse the set of testRequests below.
var optionsSets = []realip.Options{
	{
		TrustedNetworks: realip.AllReservedNetworks,
	},
	{
		TrustedNetworks: []string{"0.0.0.0/0"},
	},
	{
		TrustedNetworks: append(realip.AllReservedNetworks, "56.123.43.24/24"),
	},
	{
		TrustedNetworks: []string{"192.168.0.0/16"},
	},
}

// testRequests is a slice of tests. Each test in will be passed to the
// parse function with each of the option sets defined above. The out
// attribute for each test is a slice of values corresponding to the
// expected output for the option set with the corresponding index.
var testRequests = []struct {
	in  *http.Request
	out []string
}{
	{
		// Simple request with no headers
		in: &http.Request{
			RemoteAddr: "53.23.53.123:13245",
		},
		out: []string{
			"53.23.53.123:13245",
			"53.23.53.123:13245",
			"53.23.53.123:13245",
			"53.23.53.123:13245",
		},
	},
	{
		in: &http.Request{
			RemoteAddr: "192.168.21.34:12345",
			Header: http.Header{
				"X-Forwarded-For": []string{"53.23.53.123"},
			},
		},
		out: []string{
			"53.23.53.123",
			"53.23.53.123",
			"53.23.53.123",
			"53.23.53.123",
		},
	},
	{
		in: &http.Request{
			RemoteAddr: "192.168.21.34:12345",
			Header: http.Header{
				"X-Forwarded-For": []string{"53.23.53.123, 10.10.1.2, 127.0.0.1"},
			},
		},
		out: []string{
			"53.23.53.123",
			"53.23.53.123",
			"53.23.53.123",
			"127.0.0.1",
		},
	},
	{
		in: &http.Request{
			RemoteAddr: "56.123.43.24:567",
			Header: http.Header{
				"X-Forwarded-For": []string{"10.7.12.97, 53.23.53.123:12345, 10.10.1.2"},
			},
		},
		out: []string{
			"56.123.43.24:567",
			"10.7.12.97",
			"53.23.53.123:12345",
			"56.123.43.24:567",
		},
	},
	{
		in: &http.Request{
			RemoteAddr: "10.1.23.153:345",
			Header: http.Header{
				"X-Forwarded-For": []string{"10.34.21.243, 192.168.168.97, 172.16.27.102"},
			},
		},
		out: []string{
			"10.34.21.243",
			"10.34.21.243",
			"10.34.21.243",
			"10.1.23.153:345",
		},
	},
}

// TestParse will test the realip.Parse function with a matrix of requests
// and options.
func TestParse(t *testing.T) {

	for optionsIdx, optionsSet := range optionsSets {

		for _, rt := range testRequests {
			ip := realip.Parse(rt.in, optionsSet)
			if ip != rt.out[optionsIdx] {
				t.Errorf("realip.Parse(%+v, %+v) => %q, want %q", rt.in, optionsSet, ip, rt.out[optionsIdx])
			}
		}

	}

}

func BenchmarkParse(b *testing.B) {
	req := &http.Request{
		RemoteAddr: "56.123.43.24:567",
		Header: http.Header{
			"X-Forwarded-For": []string{"10.7.12.97, 53.23.53.123:12345, 10.10.1.2"},
		},
	}
	options := realip.Options{
		TrustedNetworks: realip.AllReservedNetworks,
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = realip.Parse(req, options)
	}
}
