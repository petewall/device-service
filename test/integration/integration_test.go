package integration_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/go-redis/redis/v8"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	. "github.com/petewall/device-service/v2/lib"
)

func flushAll(config *DBConfig) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: "",
		DB:       0,
	})
	client.FlushAll(context.Background())
}

var _ = Describe("Integration test", Ordered, func() {
	var (
		api      *API
		dbConfig *DBConfig
		res      *httptest.ResponseRecorder
		log      *gbytes.Buffer
	)

	BeforeAll(func() {
		dbConfig = &DBConfig{
			Host: "localhost",
			Port: 6379,
		}
		flushAll(dbConfig)
	})

	AfterAll(func() {
		flushAll(dbConfig)
	})

	BeforeEach(func() {
		log = gbytes.NewBuffer()
		api = &API{
			DB:        Connect(dbConfig),
			LogOutput: log,
		}
		res = httptest.NewRecorder()
	})

	When("getting the list of devices", func() {
		It("returns an empty list", func() {
			req, err := http.NewRequest("GET", "/", nil)
			Expect(err).ToNot(HaveOccurred())

			api.GetMux().ServeHTTP(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))

			devices := []*Device{}
			err = json.Unmarshal(res.Body.Bytes(), &devices)
			Expect(err).ToNot(HaveOccurred())
			Expect(devices).To(BeEmpty())
		})
	})

	When("getting details for a single device", func() {
		It("returns 404", func() {
			req, err := http.NewRequest("GET", "/aa:bb:cc:dd:ee:ff", nil)
			Expect(err).ToNot(HaveOccurred())

			api.GetMux().ServeHTTP(res, req)
			Expect(res.Code).To(Equal(http.StatusNotFound))
			Expect(res.Body.String()).To(Equal("no device found with MAC aa:bb:cc:dd:ee:ff"))
		})
	})

	When("updating a device", func() {
		var validDeviceBody []byte
		BeforeEach(func() {
			device := &Device{
				MAC:               "aa:bb:cc:dd:ee:ff",
				CurrentFirmware:   "bootstrap",
				CurrentVersion:    "1.2.3",
				AssignedFirmware:  "bootstrap",
				AcceptsPrerelease: true,
			}

			var err error
			validDeviceBody, err = json.Marshal(device)
			Expect(err).ToNot(HaveOccurred())
		})

		It("adds the device", func() {
			req, err := http.NewRequest("POST", "/aa:bb:cc:dd:ee:ff", bytes.NewReader(validDeviceBody))
			Expect(err).ToNot(HaveOccurred())

			api.GetMux().ServeHTTP(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))
		})
	})

	When("getting the list of devices", func() {
		It("returns a list with only the new device", func() {
			req, err := http.NewRequest("GET", "/", nil)
			Expect(err).ToNot(HaveOccurred())

			api.GetMux().ServeHTTP(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))

			devices := []*Device{}
			err = json.Unmarshal(res.Body.Bytes(), &devices)
			Expect(err).ToNot(HaveOccurred())

			Expect(devices).To(HaveLen(1))
			Expect(devices[0].MAC).To(Equal("aa:bb:cc:dd:ee:ff"))
			Expect(devices[0].CurrentFirmware).To(Equal("bootstrap"))
			Expect(devices[0].CurrentVersion).To(Equal("1.2.3"))
			Expect(devices[0].AssignedFirmware).To(Equal("bootstrap"))
			Expect(devices[0].AssignedVersion).To(BeEmpty())
			Expect(devices[0].AcceptsPrerelease).To(BeTrue())
		})
	})

	When("getting details for a single device", func() {
		It("returns the new device", func() {
			req, err := http.NewRequest("GET", "/aa:bb:cc:dd:ee:ff", nil)
			Expect(err).ToNot(HaveOccurred())

			api.GetMux().ServeHTTP(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))

			var device *Device
			err = json.Unmarshal(res.Body.Bytes(), &device)
			Expect(err).ToNot(HaveOccurred())

			Expect(device.MAC).To(Equal("aa:bb:cc:dd:ee:ff"))
			Expect(device.CurrentFirmware).To(Equal("bootstrap"))
			Expect(device.CurrentVersion).To(Equal("1.2.3"))
			Expect(device.AssignedFirmware).To(Equal("bootstrap"))
			Expect(device.AssignedVersion).To(BeEmpty())
			Expect(device.AcceptsPrerelease).To(BeTrue())
		})
	})

	When("updating a device", func() {
		var validDeviceBody []byte
		BeforeEach(func() {
			device := &Device{
				MAC:               "aa:bb:cc:dd:ee:ff",
				CurrentFirmware:   "bootstrap",
				CurrentVersion:    "9.9.9",
				AssignedFirmware:  "lightswitch",
				AcceptsPrerelease: false,
			}

			var err error
			validDeviceBody, err = json.Marshal(device)
			Expect(err).ToNot(HaveOccurred())
		})

		It("updates the device", func() {
			req, err := http.NewRequest("POST", "/aa:bb:cc:dd:ee:ff", bytes.NewReader(validDeviceBody))
			Expect(err).ToNot(HaveOccurred())

			api.GetMux().ServeHTTP(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))
		})
	})

	When("getting the list of devices", func() {
		It("returns a list with only the updated device", func() {
			req, err := http.NewRequest("GET", "/", nil)
			Expect(err).ToNot(HaveOccurred())

			api.GetMux().ServeHTTP(res, req)
			Expect(res.Code).To(Equal(http.StatusOK))

			devices := []*Device{}
			err = json.Unmarshal(res.Body.Bytes(), &devices)
			Expect(err).ToNot(HaveOccurred())
			Expect(devices).To(HaveLen(1))
			Expect(devices[0].MAC).To(Equal("aa:bb:cc:dd:ee:ff"))
			Expect(devices[0].CurrentFirmware).To(Equal("bootstrap"))
			Expect(devices[0].CurrentVersion).To(Equal("9.9.9"))
			Expect(devices[0].AssignedFirmware).To(Equal("lightswitch"))
			Expect(devices[0].AssignedVersion).To(BeEmpty())
			Expect(devices[0].AcceptsPrerelease).To(BeFalse())
		})
	})
})
