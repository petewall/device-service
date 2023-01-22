package test_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetAll", func() {
	It("returns the list of devices", func() {
		deviceList, err := client.GetAllDevices()
		Expect(err).ToNot(HaveOccurred())
		Expect(deviceList).To(HaveLen(2))
	})

	Context("No existing firmware", func() {
		BeforeEach(func() {
			RemoveAllDevices()
		})

		It("returns an empty list", func() {
			deviceList, err := client.GetAllDevices()
			Expect(err).ToNot(HaveOccurred())
			Expect(deviceList).To(BeEmpty())
		})
	})
})
