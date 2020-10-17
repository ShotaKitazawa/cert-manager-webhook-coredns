package main

import (
	"os"
	"testing"

	"github.com/ShotaKitazawa/cert-manager-webhook-coredns/coredns"
	"github.com/jetstack/cert-manager/test/acme/dns"
)

var (
	zone       = os.Getenv("TEST_ZONE_NAME")
	retryCount = 6
)

func TestRunsSuite(t *testing.T) {
	// The manifest path should contain a file named config.json that is a
	// snippet of valid configuration that should be included on the
	// ChallengeRequest passed as part of the test cases.

	fixture := dns.NewFixture(
		&coredns.CustomDNSProviderSolver{},
		dns.SetBinariesPath("__main__/hack/bin"),
		dns.SetResolvedZone(zone),
		dns.SetDNSServer("192.168.0.50:53"),
		dns.SetAllowAmbientCredentials(false),
		dns.SetManifestPath("./testdata/my-custom-solver"),
	)

	fixture.RunConformance(t)
}
