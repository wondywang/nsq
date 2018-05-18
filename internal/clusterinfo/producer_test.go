package clusterinfo

import "testing"

func TestHostNameAddresses(t *testing.T) {
	p := &Producer{
		BroadcastAddress: "host.domain.com",
		TCPPort:          19850,
		HTTPPort:         19851,
	}

	if p.HTTPAddress() != "host.domain.com:19851" {
		t.Errorf("Incorrect HTTPAddress: %s", p.HTTPAddress())
	}
	if p.TCPAddress() != "host.domain.com:19850" {
		t.Errorf("Incorrect TCPAddress: %s", p.TCPAddress())
	}
}

func TestIPv4Addresses(t *testing.T) {
	p := &Producer{
		BroadcastAddress: "192.168.1.17",
		TCPPort:          19850,
		HTTPPort:         19851,
	}

	if p.HTTPAddress() != "192.168.1.17:19851" {
		t.Errorf("Incorrect IPv4 HTTPAddress: %s", p.HTTPAddress())
	}
	if p.TCPAddress() != "192.168.1.17:19850" {
		t.Errorf("Incorrect IPv4 TCPAddress: %s", p.TCPAddress())
	}
}

func TestIPv6Addresses(t *testing.T) {
	p := &Producer{
		BroadcastAddress: "fd4a:622f:d2f2::1",
		TCPPort:          19850,
		HTTPPort:         19851,
	}
	if p.HTTPAddress() != "[fd4a:622f:d2f2::1]:19851" {
		t.Errorf("Incorrect IPv6 HTTPAddress: %s", p.HTTPAddress())
	}
	if p.TCPAddress() != "[fd4a:622f:d2f2::1]:19850" {
		t.Errorf("Incorrect IPv6 TCPAddress: %s", p.TCPAddress())
	}
}
