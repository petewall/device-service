package lib_test

import (
	"encoding/json"
	"net/http"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/ghttp"

	"github.com/petewall/device-service/v2/lib"
)

var _ = Describe("Client", func() {
	var (
		client *lib.Client
		server *ghttp.Server
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
			response := []*lib.Device{
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
			encoded, err := json.Marshal(response)
			Expect(err).ToNot(HaveOccurred())

			server.AppendHandlers(
				ghttp.CombineHandlers(
					ghttp.VerifyRequest("GET", "/"),
					ghttp.RespondWith(http.StatusOK, encoded),
				),
			)
		})

		It("sends the right request", func() {
			deviceList, err := client.GetAllDevices()
			Expect(err).ToNot(HaveOccurred())
			Expect(server.ReceivedRequests()).To(HaveLen(1))
			Expect(deviceList).To(HaveLen(2))
		})

		When("the request fails", func() {
			BeforeEach(func() {
				server.Reset()
				server.AppendHandlers(
					ghttp.CombineHandlers(
						ghttp.VerifyRequest("GET", "/"),
						ghttp.RespondWith(http.StatusTeapot, "I'm a little teapot"),
					),
				)
			})

			It("returns an error", func() {
				_, err := client.GetAllDevices()
				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("device list request failed: 418 I'm a teapot"))
			})
		})
	})

	XDescribe("GetDevice", func() {})
	XDescribe("UpdateDevice", func() {})
	XDescribe("SetDeviceName", func() {})
	XDescribe("SetDeviceFirmwareType", func() {})
	XDescribe("SetDeviceFirmwareVersion", func() {})
})
