package internal_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	. "github.com/petewall/device-service/v2/internal"
	. "github.com/petewall/device-service/v2/internal/internalfakes"
	. "github.com/petewall/device-service/v2/lib"
)

var _ = Describe("API", Label("api"), func() {
	var (
		api *API
		db  *FakeDBInterface
		res *httptest.ResponseRecorder
		log *gbytes.Buffer
	)

	BeforeEach(func() {
		db = &FakeDBInterface{}
		log = gbytes.NewBuffer()
		api = &API{
			DB:        db,
			LogOutput: log,
		}
		res = httptest.NewRecorder()
	})

	Describe("GET /", func() {
		When("there are no devices", func() {
			BeforeEach(func() {
				db.GetDevicesReturns([]*Device{}, nil)
			})
			It("returns an empty list", func() {
				req, err := http.NewRequest("GET", "/", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetHttpHandler().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusOK))

				devices := []*Device{}
				err = json.Unmarshal(res.Body.Bytes(), &devices)
				Expect(err).ToNot(HaveOccurred())

				Expect(devices).To(BeEmpty())

				Expect(db.GetDevicesCallCount()).To(Equal(1))
			})
		})

		When("there are some devices", func() {
			BeforeEach(func() {
				db.GetDevicesReturns([]*Device{
					{
						MAC:              "aa:bb:cc:dd:ee:ff",
						Name:             "test device 1",
						Firmware:         "bootstrap",
						Version:          "1.2.3",
						AssignedFirmware: "bootstrap",
					},
					{
						MAC:              "a1:b2:c3:d4:e5:f6",
						Name:             "test device 2",
						Firmware:         "lightswtich",
						Version:          "0.0.1-rc1",
						AssignedFirmware: "lightswtich",
					},
				}, nil)
			})

			It("returns the list of devices", func() {
				req, err := http.NewRequest("GET", "/", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetHttpHandler().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusOK))

				devices := []*Device{}
				err = json.Unmarshal(res.Body.Bytes(), &devices)
				Expect(err).ToNot(HaveOccurred())

				Expect(devices).To(HaveLen(2))

				Expect(db.GetDevicesCallCount()).To(Equal(1))
			})
		})

		When("the database has an error", func() {
			BeforeEach(func() {
				db.GetDevicesReturns(nil, errors.New("db error"))
			})
			It("returns a 500 error", func() {
				req, err := http.NewRequest("GET", "/", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetHttpHandler().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusInternalServerError))
				Expect(res.Body.String()).To(Equal("failed to request devices from the database"))

				Expect(db.GetDevicesCallCount()).To(Equal(1))
			})
		})
	})

	Describe("GET /<mac>", func() {
		When("the device does not exist", func() {
			BeforeEach(func() {
				db.GetDeviceReturns(nil, nil)
			})
			It("returns not found", func() {
				req, err := http.NewRequest("GET", "/aa:bb:cc:dd:ee:ff", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetHttpHandler().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusNotFound))
				Expect(res.Body.String()).To(Equal("no device found with MAC aa:bb:cc:dd:ee:ff"))

				Expect(db.GetDeviceCallCount()).To(Equal(1))
			})
		})

		When("the device exists", func() {
			BeforeEach(func() {
				db.GetDeviceReturns(&Device{
					MAC:              "aa:bb:cc:dd:ee:ff",
					Name:             "test device",
					Firmware:         "bootstrap",
					Version:          "1.2.3",
					AssignedFirmware: "bootstrap",
				}, nil)
			})

			It("returns the device", func() {
				req, err := http.NewRequest("GET", "/aa:bb:cc:dd:ee:ff", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetHttpHandler().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusOK))

				var device *Device
				err = json.Unmarshal(res.Body.Bytes(), &device)
				Expect(err).ToNot(HaveOccurred())

				Expect(device.MAC).To(Equal("aa:bb:cc:dd:ee:ff"))
				Expect(device.Name).To(Equal("test device"))
				Expect(device.Firmware).To(Equal("bootstrap"))
				Expect(device.Version).To(Equal("1.2.3"))
				Expect(device.AssignedFirmware).To(Equal("bootstrap"))
				Expect(device.AssignedVersion).To(BeEmpty())

				Expect(db.GetDeviceCallCount()).To(Equal(1))
			})
		})

		When("the database has an error", func() {
			BeforeEach(func() {
				db.GetDeviceReturns(nil, errors.New("db error"))
			})
			It("returns a 500 error", func() {
				req, err := http.NewRequest("GET", "/aa:bb:cc:dd:ee:ff", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetHttpHandler().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusInternalServerError))
				Expect(res.Body.String()).To(Equal("failed to request device from the database"))

				Expect(db.GetDeviceCallCount()).To(Equal(1))
			})
		})
	})

	Describe("POST /<mac>", func() {
		var validDeviceBody []byte
		BeforeEach(func() {
			device := &Device{
				Firmware: "bootstrap",
				Version:  "1.2.3",
			}

			var err error
			validDeviceBody, err = json.Marshal(device)
			Expect(err).ToNot(HaveOccurred())
		})

		It("updates the device in the DB", func() {
			req, err := http.NewRequest("POST", "/aa:bb:cc:dd:ee:ff", bytes.NewReader(validDeviceBody))
			Expect(err).ToNot(HaveOccurred())

			api.GetHttpHandler().ServeHTTP(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))

			Expect(db.UpdateDeviceCallCount()).To(Equal(1))
			mac, firmwareType, firmwareVersion := db.UpdateDeviceArgsForCall(0)
			Expect(mac).To(Equal("aa:bb:cc:dd:ee:ff"))
			Expect(firmwareType).To(Equal("bootstrap"))
			Expect(firmwareVersion).To(Equal("1.2.3"))
		})

		When("no payload is sent", func() {
			It("returns a 400 error", func() {
				req, err := http.NewRequest("POST", "/aa:bb:cc:dd:ee:ff", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetHttpHandler().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusBadRequest))
				Expect(res.Body.String()).To(Equal("device payload required"))

				Expect(db.UpdateDeviceCallCount()).To(Equal(0))
			})
		})

		When("an invalid payload is sent", func() {
			It("returns a 400 error", func() {
				req, err := http.NewRequest("POST", "/aa:bb:cc:dd:ee:ff", strings.NewReader("this is not valid json"))
				Expect(err).ToNot(HaveOccurred())

				api.GetHttpHandler().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusBadRequest))
				Expect(res.Body.String()).To(Equal("invalid device payload"))

				Expect(db.UpdateDeviceCallCount()).To(Equal(0))
			})
		})

		When("the database has an error", func() {
			BeforeEach(func() {
				db.UpdateDeviceReturns(errors.New("db error"))
			})
			It("returns a 500 error", func() {
				req, err := http.NewRequest("POST", "/aa:bb:cc:dd:ee:ff", bytes.NewReader(validDeviceBody))
				Expect(err).ToNot(HaveOccurred())

				api.GetHttpHandler().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusInternalServerError))
				Expect(res.Body.String()).To(Equal("failed to update device in the database"))

				Expect(db.UpdateDeviceCallCount()).To(Equal(1))
				mac, firmwareType, firmwareVersion := db.UpdateDeviceArgsForCall(0)
				Expect(mac).To(Equal("aa:bb:cc:dd:ee:ff"))
				Expect(firmwareType).To(Equal("bootstrap"))
				Expect(firmwareVersion).To(Equal("1.2.3"))
			})
		})
	})

	Describe("POST /<mac>/name", func() {
		It("sets the name of the device", func() {
			req, err := http.NewRequest("POST", "/aa:bb:cc:dd:ee:ff/name?val=garage", nil)
			Expect(err).ToNot(HaveOccurred())

			api.GetHttpHandler().ServeHTTP(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))

			Expect(db.SetDeviceFieldCallCount()).To(Equal(1))
			mac, key, value := db.SetDeviceFieldArgsForCall(0)
			Expect(mac).To(Equal("aa:bb:cc:dd:ee:ff"))
			Expect(key).To(Equal("name"))
			Expect(value).To(Equal("garage"))
		})

		When("no value is set", func() {
			It("returns a 400 error", func() {
				req, err := http.NewRequest("POST", "/aa:bb:cc:dd:ee:ff/name", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetHttpHandler().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusBadRequest))
				Expect(res.Body.String()).To(Equal("missing device name"))

				Expect(db.SetDeviceFieldCallCount()).To(Equal(0))
			})
		})

		When("the database has an error", func() {
			BeforeEach(func() {
				db.SetDeviceFieldReturns(errors.New("db error"))
			})
			It("returns a 500 error", func() {
				req, err := http.NewRequest("POST", "/aa:bb:cc:dd:ee:ff/name?val=garage", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetHttpHandler().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusInternalServerError))
				Expect(res.Body.String()).To(Equal("failed to set device name in the database"))

				Expect(db.SetDeviceFieldCallCount()).To(Equal(1))
				mac, key, value := db.SetDeviceFieldArgsForCall(0)
				Expect(mac).To(Equal("aa:bb:cc:dd:ee:ff"))
				Expect(key).To(Equal("name"))
				Expect(value).To(Equal("garage"))
			})
		})
	})

	Describe("POST /<mac>/firmware", func() {
		It("sets the assigned firmware type of the device", func() {
			req, err := http.NewRequest("POST", "/aa:bb:cc:dd:ee:ff/firmware?val=clock", nil)
			Expect(err).ToNot(HaveOccurred())

			api.GetHttpHandler().ServeHTTP(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))

			Expect(db.SetDeviceFieldCallCount()).To(Equal(1))
			mac, key, value := db.SetDeviceFieldArgsForCall(0)
			Expect(mac).To(Equal("aa:bb:cc:dd:ee:ff"))
			Expect(key).To(Equal("assignedFirmware"))
			Expect(value).To(Equal("clock"))
		})

		When("no value is sent", func() {
			It("returns a 400 error", func() {
				req, err := http.NewRequest("POST", "/aa:bb:cc:dd:ee:ff/firmware", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetHttpHandler().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusBadRequest))
				Expect(res.Body.String()).To(Equal("missing device firmware type value"))

				Expect(db.SetDeviceFieldCallCount()).To(Equal(0))
			})
		})

		When("the database has an error", func() {
			BeforeEach(func() {
				db.SetDeviceFieldReturns(errors.New("db error"))
			})
			It("returns a 500 error", func() {
				req, err := http.NewRequest("POST", "/aa:bb:cc:dd:ee:ff/firmware?val=clock", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetHttpHandler().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusInternalServerError))
				Expect(res.Body.String()).To(Equal("failed to set device assigned firmware type in the database"))

				Expect(db.SetDeviceFieldCallCount()).To(Equal(1))
				mac, key, value := db.SetDeviceFieldArgsForCall(0)
				Expect(mac).To(Equal("aa:bb:cc:dd:ee:ff"))
				Expect(key).To(Equal("assignedFirmware"))
				Expect(value).To(Equal("clock"))
			})
		})
	})

	Describe("POST /<mac>/version", func() {
		It("sets the assigned firmware version of the device", func() {
			req, err := http.NewRequest("POST", "/aa:bb:cc:dd:ee:ff/version?val=1.2.3", nil)
			Expect(err).ToNot(HaveOccurred())

			api.GetHttpHandler().ServeHTTP(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))

			Expect(db.SetDeviceFieldCallCount()).To(Equal(1))
			mac, key, value := db.SetDeviceFieldArgsForCall(0)
			Expect(mac).To(Equal("aa:bb:cc:dd:ee:ff"))
			Expect(key).To(Equal("assignedVersion"))
			Expect(value).To(Equal("1.2.3"))
		})

		When("no value is sent", func() {
			It("returns a 400 error", func() {
				req, err := http.NewRequest("POST", "/aa:bb:cc:dd:ee:ff/version", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetHttpHandler().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusBadRequest))
				Expect(res.Body.String()).To(Equal("missing device firmware version value"))

				Expect(db.SetDeviceFieldCallCount()).To(Equal(0))
			})
		})

		When("the database has an error", func() {
			BeforeEach(func() {
				db.SetDeviceFieldReturns(errors.New("db error"))
			})
			It("returns a 500 error", func() {
				req, err := http.NewRequest("POST", "/aa:bb:cc:dd:ee:ff/version?val=1.2.3", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetHttpHandler().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusInternalServerError))
				Expect(res.Body.String()).To(Equal("failed to set device assigned firmware version in the database"))

				Expect(db.SetDeviceFieldCallCount()).To(Equal(1))
				mac, key, value := db.SetDeviceFieldArgsForCall(0)
				Expect(mac).To(Equal("aa:bb:cc:dd:ee:ff"))
				Expect(key).To(Equal("assignedVersion"))
				Expect(value).To(Equal("1.2.3"))
			})
		})
	})
})
