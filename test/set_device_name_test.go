package test_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("SetDeviceName", func() {
	Context("The device already exists", func() {
		It("the device is updated", func() {
			err := client.SetDeviceName("aa:bb:cc:dd:ee:ff", "newly updated device")
			Expect(err).ToNot(HaveOccurred())

			device, err := client.GetDevice("aa:bb:cc:dd:ee:ff")
			Expect(err).ToNot(HaveOccurred())
			Expect(device.Name).To(Equal("newly updated device"))
		})
	})

	Context("The device does not exist", func() {
		It("the device is added", func() {
			err := client.SetDeviceName("bb:bb:bb:bb:bb:bb", "newly updated device")
			Expect(err).ToNot(HaveOccurred())

			device, err := client.GetDevice("bb:bb:bb:bb:bb:bb")
			Expect(err).ToNot(HaveOccurred())
			Expect(device.Name).To(Equal("newly updated device"))
		})
	})
})
