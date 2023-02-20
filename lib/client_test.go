package lib_test

import (
	"net/http"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"github.com/petewall/device-service/lib"
)

var _ = Describe("Client", func() {
	var (
		client     *lib.Client
		server     *ghttp.Server
		statusCode int
	)

	BeforeEach(func() {
		server = ghttp.NewServer()
		client = &lib.Client{
			URL: server.URL(),
		}
	})

	AfterEach(func() {
		server.Close()
	})

	Describe("GetAllDevices", func() {
		BeforeEach(func() {
			deviceList := []*lib.Device{
				&lib.Device{
					MAC:              "aa:bb:cc:dd:ee:ff",
					Name:             "device 1",
					Firmware:         "bootstrap",
					Version:          "1.2.3",
					AssignedFirmware: "clock",
					AssignedVersion:  "~1.0",
					LastUpdate:       time.Now().Unix(),
				},
				&lib.Device{
					MAC:        "aa:aa:aa:aa:aa:aa",
					Name:       "device 2",
					Firmware:   "bootstrap",
					Version:    "2.0.0",
					LastUpdate: time.Now().Unix(),
				},
			}

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/"),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, &deviceList),
				),
			)
		})

		When("the request succeeds", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
			})

			It("sends the right request", func() {
				deviceList, err := client.GetAllDevices()
				Expect(err).ToNot(HaveOccurred())
				Expect(server.ReceivedRequests()).To(HaveLen(1))
				Expect(deviceList).To(HaveLen(2))
			})
		})

		When("the request fails", func() {
			BeforeEach(func() {
				statusCode = http.StatusTeapot
			})

			It("returns an error", func() {
				_, err := client.GetAllDevices()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("device list request failed: 418 I'm a teapot"))
			})
		})
	})

	Describe("GetDevice", func() {
		BeforeEach(func() {
			device := &lib.Device{
				MAC:              "aa:bb:cc:dd:ee:ff",
				Name:             "device 1",
				Firmware:         "bootstrap",
				Version:          "1.2.3",
				AssignedFirmware: "clock",
				AssignedVersion:  "~1.0",
				LastUpdate:       time.Now().Unix(),
			}

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/aa:bb:cc:dd:ee:ff"),
					ghttp.RespondWithJSONEncodedPtr(&statusCode, &device),
				),
			)
		})

		When("the request succeeds", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
			})

			It("sends the right request", func() {
				device, err := client.GetDevice("aa:bb:cc:dd:ee:ff")
				Expect(err).ToNot(HaveOccurred())
				Expect(server.ReceivedRequests()).To(HaveLen(1))
				Expect(device.Name).To(Equal("device 1"))
			})
		})

		When("the request fails", func() {
			BeforeEach(func() {
				statusCode = http.StatusTeapot
			})

			It("returns an error", func() {
				_, err := client.GetDevice("aa:bb:cc:dd:ee:ff")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("device aa:bb:cc:dd:ee:ff request failed: 418 I'm a teapot"))
			})
		})
	})

	Describe("UpdateDevice", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/aa:bb:cc:dd:ee:ff"),
					ghttp.VerifyHeader(http.Header{
						"content-type": []string{"application/json"},
					}),
					ghttp.VerifyJSON(`{"firmware":"clock","version":"1.0"}`),
					ghttp.RespondWithPtr(&statusCode, nil),
				),
			)
		})

		When("the request succeeds", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
			})

			It("sends the right request", func() {
				err := client.UpdateDevice("aa:bb:cc:dd:ee:ff", "clock", "1.0")
				Expect(err).ToNot(HaveOccurred())
				Expect(server.ReceivedRequests()).To(HaveLen(1))
			})
		})

		When("the request fails", func() {
			BeforeEach(func() {
				statusCode = http.StatusTeapot
			})

			It("returns an error", func() {
				err := client.UpdateDevice("aa:bb:cc:dd:ee:ff", "clock", "1.0")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("update device aa:bb:cc:dd:ee:ff request failed: 418 I'm a teapot"))
			})
		})
	})

	Describe("SetDeviceName", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/aa:bb:cc:dd:ee:ff/name", "val=device%202"),
					ghttp.RespondWithPtr(&statusCode, nil),
				),
			)
		})

		When("the request succeeds", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
			})

			It("sends the right request", func() {
				err := client.SetDeviceName("aa:bb:cc:dd:ee:ff", "device 2")
				Expect(err).ToNot(HaveOccurred())
				Expect(server.ReceivedRequests()).To(HaveLen(1))
			})
		})

		When("the request fails", func() {
			BeforeEach(func() {
				statusCode = http.StatusTeapot
			})

			It("returns an error", func() {
				err := client.SetDeviceName("aa:bb:cc:dd:ee:ff", "device 2")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("set device aa:bb:cc:dd:ee:ff name to device 2 request failed: 418 I'm a teapot"))
			})
		})
	})

	Describe("SetDeviceFirmwareType", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/aa:bb:cc:dd:ee:ff/firmware", "val=laser-x-target"),
					ghttp.RespondWithPtr(&statusCode, nil),
				),
			)
		})

		When("the request succeeds", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
			})

			It("sends the right request", func() {
				err := client.SetDeviceFirmwareType("aa:bb:cc:dd:ee:ff", "laser-x-target")
				Expect(err).ToNot(HaveOccurred())
				Expect(server.ReceivedRequests()).To(HaveLen(1))
			})
		})

		When("the request fails", func() {
			BeforeEach(func() {
				statusCode = http.StatusTeapot
			})

			It("returns an error", func() {
				err := client.SetDeviceFirmwareType("aa:bb:cc:dd:ee:ff", "laser-x-target")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("set device aa:bb:cc:dd:ee:ff firmware type to laser-x-target request failed: 418 I'm a teapot"))
			})
		})
	})

	Describe("SetDeviceFirmwareVersion", func() {
		BeforeEach(func() {
			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("POST", "/aa:bb:cc:dd:ee:ff/version", "val=0.1.3"),
					ghttp.RespondWithPtr(&statusCode, nil),
				),
			)
		})

		When("the request succeeds", func() {
			BeforeEach(func() {
				statusCode = http.StatusOK
			})

			It("sends the right request", func() {
				err := client.SetDeviceFirmwareVersion("aa:bb:cc:dd:ee:ff", "0.1.3")
				Expect(err).ToNot(HaveOccurred())
				Expect(server.ReceivedRequests()).To(HaveLen(1))
			})
		})

		When("the request fails", func() {
			BeforeEach(func() {
				statusCode = http.StatusTeapot
			})

			It("returns an error", func() {
				err := client.SetDeviceFirmwareVersion("aa:bb:cc:dd:ee:ff", "0.1.3")
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("set device aa:bb:cc:dd:ee:ff firmware version to 0.1.3 request failed: 418 I'm a teapot"))
			})
		})
	})
})
