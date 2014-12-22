package fluent

import (
	"testing"
	"time"
)

func TestApplyDefaultValues(t *testing.T) {
	config := Config{}

	assertEqual(t, config.FluentHost, stringDefault)
	assertEqual(t, config.FluentPort, intDefault)
	assertEqual(t, config.ChannelLength, intDefault)
	assertEqual(t, config.BufferLength, intDefault)
	assertEqual(t, config.MaxTrialForConnection, intDefault)
	assertEqual(t, config.ConnectionTimeout, durationDefault)
	assertEqual(t, config.BufferingTimeout, durationDefault)
	assertEqual(t, config.TagPrefix, stringDefault)

	config.applyDefaultValues()
	assertEqual(t, config.FluentHost, DefaultFluentHost)
	assertEqual(t, config.FluentPort, DefaultFluentPort)
	assertEqual(t, config.ChannelLength, DefaultChannelLength)
	assertEqual(t, config.BufferLength, DefaultBufferLength)
	assertEqual(t, config.MaxTrialForConnection, DefaultMaxTrialForConnection)
	assertEqual(t, config.ConnectionTimeout, DefaultConnectionTimeout)
	assertEqual(t, config.BufferingTimeout, DefaultBufferingTimeout)
	assertEqual(t, config.TagPrefix, DefaultTagPrefix)

	config = Config{
		FluentHost:            "localhost",
		FluentPort:            80,
		ChannelLength:         1,
		BufferLength:          2,
		MaxTrialForConnection: 3,
		ConnectionTimeout:     2 * time.Second,
		BufferingTimeout:      2 * time.Second,
		TagPrefix:             "prefix",
	}

	config.applyDefaultValues()
	assertEqual(t, config.FluentHost, "localhost")
	assertEqual(t, config.FluentPort, 80)
	assertEqual(t, config.ChannelLength, 1)
	assertEqual(t, config.BufferLength, 2)
	assertEqual(t, config.MaxTrialForConnection, 3)
	assertEqual(t, config.ConnectionTimeout, 2*time.Second)
	assertEqual(t, config.BufferingTimeout, 2*time.Second)
	assertEqual(t, config.TagPrefix, "prefix")
}

func assertEqual(t *testing.T, actual interface{}, expect interface{}) {
	switch actual.(type) {
	case string:
		if actual.(string) != expect.(string) {
			t.Errorf("expected %s, but got %s\n", expect.(string), actual.(string))
		}
	case int:
		if actual.(int) != expect.(int) {
			t.Errorf("expected %d, but got %d\n", expect.(int), actual.(int))
		}
	case time.Duration:
		if actual.(time.Duration) != expect.(time.Duration) {
			t.Errorf("expected %d, but got %d\n", expect.(time.Duration), actual.(time.Duration))
		}
	}
}
