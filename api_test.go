package deviceservice_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/petewall/device-service/v2"
	. "github.com/petewall/device-service/v2/device-servicefakes"
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
				db.GetDevicesReturns(nil, []*Device{})
			})
			It("returns an empty list", func() {
				req, err := http.NewRequest("GET", "/", nil)
				Expect(err).ToNot(HaveOccurred())

				api.GetMux().ServeHTTP(res, req)
				Expect(res.Code).To(Equal(http.StatusOK))
				Expect(res.Body.String()).To(Equal("[]"))

				Expect(db.GetDevicesCallCount()).To(Equal(1))
			})
		})
	})
})
