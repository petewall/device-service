package test_test

import (
	"context"
	"fmt"
	"os/exec"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
	"github.com/petewall/device-service/v2/lib"
	"github.com/phayes/freeport"
)

func TestIntegration(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Feature test suite")
}

var (
	client               lib.Client
	dbHost               string = "localhost"
	dbPort               int    = 6379
	deviceService        string
	deviceServiceSession *gexec.Session
	deviceServiceURL     string
)

var _ = BeforeSuite(func() {
	var err error
	deviceService, err = gexec.Build("github.com/petewall/device-service/v2")
	Expect(err).ToNot(HaveOccurred())
})

var _ = AfterSuite(func() {
	gexec.CleanupBuildArtifacts()
})

var _ = BeforeEach(func() {
	port, err := freeport.GetFreePort()
	Expect(err).ToNot(HaveOccurred())
	deviceServiceURL = fmt.Sprintf("http://localhost:%d", port)
	client = lib.Client{
		URL: deviceServiceURL,
	}
	args := []string{
		"--db-host", dbHost,
		"--db-port", fmt.Sprintf("%d", dbPort),
		"--port", fmt.Sprintf("%d", port),
	}
	command := exec.Command(deviceService, args...)
	deviceServiceSession, err = gexec.Start(command, GinkgoWriter, GinkgoWriter)
	Expect(err).ToNot(HaveOccurred())
	Eventually(deviceServiceSession.Out, 10*time.Second).Should(Say(fmt.Sprintf("Listening on port %d", port)))

	Seed()
})

var _ = AfterEach(func() {
	deviceServiceSession.Terminate().Wait()
	Eventually(deviceServiceSession).Should(gexec.Exit())
})

func Seed() {
	RemoveAllDevices()
	Expect(client.UpdateDevice("aa:aa:aa:aa:aa:aa", "bootstrap", "1.0")).To(Succeed())
	Expect(client.UpdateDevice("aa:bb:cc:dd:ee:ff", "bootstrap", "1.0")).To(Succeed())
	Expect(client.SetDeviceName("aa:bb:cc:dd:ee:ff", "test device")).To(Succeed())
	Expect(client.SetDeviceFirmwareType("aa:bb:cc:dd:ee:ff", "bootstrap")).To(Succeed())
}

func RemoveAllDevices() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", dbHost, dbPort),
		Password: "",
		DB:       0,
	})
	res := redisClient.FlushAll(context.Background())
	Expect(res.Err()).ToNot(HaveOccurred())
}
