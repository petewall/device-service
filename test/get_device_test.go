package test_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("GetDevice", func() {
	It("returns the device", func() {
		device, err := client.GetDevice("aa:bb:cc:dd:ee:ff")
		Expect(err).ToNot(HaveOccurred())
		Expect(device.MAC).To(Equal("aa:bb:cc:dd:ee:ff"))
		Expect(device.Name).To(Equal("test device"))
		Expect(device.Firmware).To(Equal("bootstrap"))
		Expect(device.Version).To(Equal("1.0"))
		Expect(device.AssignedFirmware).To(Equal("bootstrap"))
		Expect(device.LastUpdate).To(BeNumerically(">", 0))
	})

	Context("Device does not exist", func() {
		BeforeEach(func() {
			RemoveAllDevices()
		})

		It("returns not found", func() {
			_, err := client.GetDevice("ff:ff:ff:ff:ff:ff")
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("device ff:ff:ff:ff:ff:ff request failed: 404 Not Found"))
		})
	})

})
