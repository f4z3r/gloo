package test

import (
	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"
	. "github.com/solo-io/go-utils/manifesttestutils"
)

var _ = Describe("SVC Accnt Test", func() {
	var (
		testManifest    TestManifest
		resourceBuilder ResourceBuilder
	)

	prepareMakefile := func(name string, helmFlags []string) {
		resourceBuilder.Name = name
		resourceBuilder.Labels["gloo"] = name

		tm, err := renderManifest(namespace, helmValues{
			valuesArgs: helmFlags,
		})
		Expect(err).NotTo(HaveOccurred(), "Should be able to render the manifest in the service account unit test")
		testManifest = tm
	}

	BeforeEach(func() {
		resourceBuilder = ResourceBuilder{
			Namespace: namespace,
			Labels: map[string]string{
				"app": "gloo",
			},
		}
	})

	It("gloo", func() {
		prepareMakefile("gloo", []string{"rbac.namespaced=false"})
		testManifest.ExpectServiceAccount(resourceBuilder.GetServiceAccount())
	})

	It("discovery", func() {
		prepareMakefile("discovery", []string{"rbac.namespaced=false"})
		testManifest.ExpectServiceAccount(resourceBuilder.GetServiceAccount())
	})

	It("gateway", func() {
		prepareMakefile("gateway", []string{"rbac.namespaced=false"})
		testManifest.ExpectServiceAccount(resourceBuilder.GetServiceAccount())
	})

	It("gateway-proxy", func() {
		prepareMakefile("gateway-proxy", []string{"rbac.namespaced=false"})
		svcAccount := resourceBuilder.GetServiceAccount()
		testManifest.ExpectServiceAccount(svcAccount)
	})

	It("gateway-proxy disables svc account", func() {
		prepareMakefile("gateway-proxy", []string{"rbac.namespaced=false", "gateway.proxyServiceAccount.disableAutomount=true"})
		svcAccount := resourceBuilder.GetServiceAccount()
		falze := false
		svcAccount.AutomountServiceAccountToken = &falze
		testManifest.ExpectServiceAccount(svcAccount)
	})

})
