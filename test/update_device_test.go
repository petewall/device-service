package test_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("UpdateDevice", func() {
	Context("The device already exists", func() {
		It("the device is updated", func() {
			device, err := client.GetDevice("aa:bb:cc:dd:ee:ff")
			Expect(err).ToNot(HaveOccurred())
			firstUpdateTime := device.LastUpdate

			err = client.UpdateDevice("aa:bb:cc:dd:ee:ff", "clock", "1.2.3")
			Expect(err).ToNot(HaveOccurred())

			device, err = client.GetDevice("aa:bb:cc:dd:ee:ff")
			Expect(err).ToNot(HaveOccurred())
			Expect(device.Firmware).To(Equal("clock"))
			Expect(device.Version).To(Equal("1.2.3"))
			Expect(device.LastUpdate).To(BeNumerically(">=", firstUpdateTime))
		})
	})

	Context("The device does not exist", func() {
		It("the device is added", func() {
			err := client.UpdateDevice("bb:bb:bb:bb:bb:bb", "clock", "1.2.3")
			Expect(err).ToNot(HaveOccurred())

			device, err := client.GetDevice("bb:bb:bb:bb:bb:bb")
			Expect(err).ToNot(HaveOccurred())
			Expect(device.Firmware).To(Equal("clock"))
			Expect(device.Version).To(Equal("1.2.3"))
		})
	})
})
