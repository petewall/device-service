package test_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SetDeviceFirmwareType", func() {
	Context("The device already exists", func() {
		It("the device is updated", func() {
			err := client.SetDeviceFirmwareType("aa:bb:cc:dd:ee:ff", "solar-panel")
			Expect(err).ToNot(HaveOccurred())

			device, err := client.GetDevice("aa:bb:cc:dd:ee:ff")
			Expect(err).ToNot(HaveOccurred())
			Expect(device.AssignedFirmware).To(Equal("solar-panel"))
		})
	})

	Context("The device does not exist", func() {
		It("the device is added", func() {
			err := client.SetDeviceFirmwareType("bb:bb:bb:bb:bb:bb", "solar-panel")
			Expect(err).ToNot(HaveOccurred())

			device, err := client.GetDevice("bb:bb:bb:bb:bb:bb")
			Expect(err).ToNot(HaveOccurred())
			Expect(device.AssignedFirmware).To(Equal("solar-panel"))
		})
	})
})
