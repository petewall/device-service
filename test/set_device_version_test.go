package test_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SetDeviceFirmwareVersion", func() {
	Context("The device already exists", func() {
		It("the device is updated", func() {
			err := client.SetDeviceFirmwareVersion("aa:bb:cc:dd:ee:ff", "1.0-rc.1")
			Expect(err).ToNot(HaveOccurred())

			device, err := client.GetDevice("aa:bb:cc:dd:ee:ff")
			Expect(err).ToNot(HaveOccurred())
			Expect(device.AssignedVersion).To(Equal("1.0-rc.1"))
		})
	})

	Context("The device does not exist", func() {
		It("the device is added", func() {
			err := client.SetDeviceFirmwareVersion("bb:bb:bb:bb:bb:bb", "~1.0")
			Expect(err).ToNot(HaveOccurred())

			device, err := client.GetDevice("bb:bb:bb:bb:bb:bb")
			Expect(err).ToNot(HaveOccurred())
			Expect(device.AssignedVersion).To(Equal("~1.0"))
		})
	})
})
