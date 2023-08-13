package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/calculi-corp/common/pkg/messaging/subjects"
	"github.com/calculi-corp/config"
	testutil "github.com/calculi-corp/grpc-testutil"
	nats "github.com/calculi-corp/nats"
	"github.com/calculi-corp/nats/testdata"
)

func TestForTest(t *testing.T) {
	r := forTest(1, 1)
	if r != 11 {
		t.Error("expected return value of 100, got ", r)
	}
}

func TestForTest2(t *testing.T) {
	r := forTest(0, 1)
	if r != 10 {
		t.Error("expected return value of 100, got ", r)
	}
}

func TestPublish(t *testing.T) {
	testutil.SetUnitTestConfig()
	config.Config.SetCliFlags()
	config.Config.Set("nats.server", "tls://localhost:4222")

	natsClient, err := nats.NewMessagingClient()
	require.NoError(t, err)
	require.NotNil(t, natsClient)
	tstMsg := &testdata.TestMessage{Name: "myTestMessage", Value: "myAwesomeValue"}
	tb, err := json.Marshal(tstMsg)
	require.NoError(t, err)
	natsClient.Publish(subjects.HttpEvent_InstrumentationUI, tb)
}
