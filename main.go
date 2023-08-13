package main

import (
	"encoding/json"
	"fmt"
	"os"

	messaging "github.com/calculi-corp/common/pkg/messaging"
	"github.com/calculi-corp/common/pkg/messaging/subjects"
	"github.com/calculi-corp/config"
	testutil "github.com/calculi-corp/grpc-testutil"
	"github.com/calculi-corp/log"
	nats "github.com/calculi-corp/nats"
	"github.com/calculi-corp/nats/testdata"
)

func main() {
	fmt.Println("This is experimental repo, have fun!")
	MyNats1()
	for {
	}
}

func forTest(a, b int) int {
	r := 10
	r = a + b*r
	if r == 10 {
		return r
	}
	fmt.Println("Testing GX-4884")
	fmt.Println("Testing GX-4884")
	fmt.Println("Testing GX-4884")
	return r
}

type awscreds struct {
	Access   string `json:"access"`
	Secret   string `json:"secret"`
	IAMRole  string `json:"iam_role"`
	Password string `json:"password"`
}

func VStrings() {
	creds := &awscreds{
		Access:   "somekey",
		Secret:   "secret",
		Password: "myawesomepassword",
	}

	log.Infof("creds: %v", creds)
}

func MyNats1() {
	testutil.SetUnitTestConfig()
	config.Config.SetCliFlags()
	config.Config.Set("nats.server", "tls://localhost:4222")
	log.Info("Starting NATS test")
	natsClient, err := nats.NewMessagingClient()
	if log.CheckErrorf(err, "Failed to create NATS client") {
		os.Exit(1)
	}
	natsClient.Subscribe(subjects.HttpEvent_InstrumentationUI, myTestHander)
}

func myTestHander(msg messaging.Message) error {
	log.Infof("raw message %s", msg)
	tstMsg := &testdata.TestMessage{}
	err := json.Unmarshal(msg.Data(), tstMsg)
	if log.CheckErrorf(err, "Failed to unmarshal message") {
		return err
	}
	log.Infof("Received message %s", tstMsg)
	return nil
}
