package lib_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/petewall/device-service/lib"
)

var _ = Describe("Device", func() {
	Describe("IsDifferent", func() {
		var device *Device

		BeforeEach(func() {
			device = &Device{
				Firmware: "bootstrap",
				Version:  "1.2.3",
			}
		})

		When("the values are the same", func() {
			It("returns false", func() {
				Expect(device.IsDifferent("bootstrap", "1.2.3")).To(BeFalse())
			})
		})

		When("the type is different", func() {
			It("returns true", func() {
				Expect(device.IsDifferent("lightswitch", "1.2.3")).To(BeTrue())
			})
		})

		When("the version is different", func() {
			It("returns true", func() {
				Expect(device.IsDifferent("bootstrap", "2.3.4")).To(BeTrue())
			})
		})
	})

	XDescribe("IsOlderThan", func() {

	})
})
