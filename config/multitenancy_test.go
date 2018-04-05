package config_test

import (
	"bytes"

	. "github.com/delving/rapid/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
)

// any approach to require this configuration into your program.
var tomlExample = []byte(`
[all]
orgID = "all"

[default]
orgID  = "default"

[default.lod]
# enable the lod routing point
enabled = true
## resource routing point
resource = "resource"
# html routing point
html = "page"
# rdf data routing point
rdf = "data"
`)

var _ = Describe("Multitenancy", func() {

	Describe("When creating a Tenant", func() {

		Context("from configuration", func() {
			tc := viper.New()
			tc.SetConfigType("toml")
			err := tc.ReadConfig(bytes.NewBuffer(tomlExample))

			It("should not be empty", func() {
				Expect(err).ToNot(HaveOccurred())
				Expect(tc).ToNot(BeNil())
				subTenant := tc.Sub("all")
				t, err := NewTenant(subTenant)
				Expect(err).ToNot(HaveOccurred())
				Expect(t).ToNot(BeNil())
				Expect(t.OrgID).To(Equal("all"))
			})

			It("should read different sub-tenants", func() {
				subTenant := tc.Sub("default")
				t, err := NewTenant(subTenant)
				Expect(err).ToNot(HaveOccurred())
				Expect(t).ToNot(BeNil())
				Expect(t.OrgID).To(Equal("default"))
			})

			It("should set LOD defaults", func() {
				subTenant := tc.Sub("all")
				t, err := NewTenant(subTenant)
				Expect(err).ToNot(HaveOccurred())
				Expect(t.LOD.Enabled).To(BeTrue())
				Expect(t.LOD.HTML).To(Equal("page"))
				Expect(t.LOD.Resource).To(Equal("resource"))
				Expect(t.LOD.RDF).To(Equal("data"))
				Expect(t.LOD.HTMLRedirectRegex).To(BeEmpty())
			})

			It("should throw an error on missing Org config", func() {
				t, err := NewTenant(viper.New())
				Expect(t).To(BeNil())
				Expect(err).To(HaveOccurred())
			})
		})

	})

})
