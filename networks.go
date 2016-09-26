package realip

const (
	// LoopbackNetworkIPv4 is the local loopback network for IPv4
	LoopbackNetworkIPv4 = "127.0.0.0/8"

	// LoopbackNetworkIPv6 is the local loopback network for IPv6
	LoopbackNetworkIPv6 = "::1/128"

	// Private24BitNetworkIPv4 is the private 24-bit network for IPv4
	// RFC 1918 (https://tools.ietf.org/html/rfc1918)
	Private24BitNetworkIPv4 = "10.0.0.0/8"

	// Private20BitNetworkIPv4 is the private 20-bit network for IPv4
	// RFC 1918 (https://tools.ietf.org/html/rfc1918)
	Private20BitNetworkIPv4 = "172.16.0.0/12"

	// Private16BitNetworkIPv4 is the private 16-bit network for IPv4
	// RFC 1918 (https://tools.ietf.org/html/rfc1918)
	Private16BitNetworkIPv4 = "192.168.0.0/16"

	// ULANetworkIPv6 is the unique local addresses for IPv6
	// RFC 4193 (https://tools.ietf.org/html/rfc4193)
	ULANetworkIPv6 = "fc00::/7"

	// LinkLocalNetworkIPv4 is the link-local network for IPv4
	LinkLocalNetworkIPv4 = "169.254.0.0/16"

	// LinkLocalNetworkIPv6 is the link-local network for IPv6
	LinkLocalNetworkIPv6 = "fe80::/10"
)

var (
	// AllReservedNetworks is a slice of all of the reserved addresses for
	// local and private networks for IPv4 and IPv6
	AllReservedNetworks = []string{
		LoopbackNetworkIPv4,
		Private24BitNetworkIPv4,
		Private20BitNetworkIPv4,
		Private16BitNetworkIPv4,
		LinkLocalNetworkIPv4,
		LoopbackNetworkIPv6,
		ULANetworkIPv6,
		LinkLocalNetworkIPv6,
	}
)
