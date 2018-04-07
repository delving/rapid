// Copyright © 2017 Delving B.V. <info@delving.eu>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package config_test

import (
	. "github.com/delving/rapid/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {

	Describe("after initialisation", func() {

		Context("when calling initConfig", func() {

			It("should be initialised with defaults", func() {
				InitConfig()
				Expect(Config.HTTP.Port).To(Equal(3001))
				//Expect(Config.OrgID).ToNot(BeEmpty())
			})

		})
	})

	Describe("build information", func() {

		Context("when initializing", func() {

			It("should set devBuild  on empty", func() {
				info := NewBuildVersionInfo(
					"", "", "ginkgo", "",
				)
				Expect(info).ToNot(BeNil())
				Expect(info.Version).To(Equal("devBuild"))
				Expect(info.BuildAgent).To(Equal("ginkgo"))
			})

			It("should pretty print to JSON", func() {
				info := NewBuildVersionInfo(
					"", "", "ginkgo", "",
				)
				Expect(info).ToNot(BeNil())
				json, err := info.JSON(true)
				Expect(err).ToNot(HaveOccurred())
				Expect(json).To(ContainSubstring("ginkgo"))
			})
		})
	})

})
