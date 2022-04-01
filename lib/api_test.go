package lib_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/petewall/device-service/v2/lib"
	. "github.com/petewall/device-service/v2/lib/libfakes"
)

var _ = Describe("API", Label("api"), func() {
	var (
		api *API
		db  *FakeDBInterface
		res *httptest.ResponseRecorder
	)

	BeforeEach(func() {
		db = &FakeDBInterface{}
		api = &API{
			DB: db,
		}
		res = httptest.NewRecorder()
	})

	Describe("GetDevices", func() {
		When("there are no devices", func() {
			BeforeEach(func() {
				db.GetDevicesReturns([]*Device{}, nil)
			})
			It("returns an empty list", func() {
				req, err := http.NewRequest("GET", "/", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetMux().ServeHTTP(res, req)
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
						CurrentFirmware:  "bootstrap",
						CurrentVersion:   "1.2.3",
						AssignedFirmware: "bootstrap",
					},
					{
						MAC:               "a1:b2:c3:d4:e5:f6",
						CurrentFirmware:   "lightswtich",
						CurrentVersion:    "0.0.1-rc1",
						AssignedFirmware:  "lightswtich",
						AcceptsPrerelease: true,
					},
				}, nil)
			})

			It("returns the list of devices", func() {
				req, err := http.NewRequest("GET", "/", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetMux().ServeHTTP(res, req)
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

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusInternalServerError))
				Expect(res.Body.String()).To(Equal("failed to request devices from the database"))
			})
		})
	})
})
