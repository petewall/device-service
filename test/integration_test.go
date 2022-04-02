package test_test

import (
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

		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", dbConfig.Host, dbConfig.Port),
			Password: "",
			DB:       0,
		})
		client.FlushAll(context.Background())
	})

	BeforeEach(func() {
		log = gbytes.NewBuffer()
		api = &API{
			DB:        Connect(dbConfig),
			LogOutput: log,
		}
		res = httptest.NewRecorder()
	})

	Describe("GET /", func() {
		When("there are no devices", func() {
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
	})
})
