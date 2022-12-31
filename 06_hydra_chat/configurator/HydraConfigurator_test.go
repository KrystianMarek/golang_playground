package configurator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfiguration(t *testing.T) {
	configurator := NewConfiguration()
	err := configurator.GetConfiguration("chat.conf")
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "127.0.0.1:2100", configurator.RemoteAddr, "")
	assert.Equal(t, true, configurator.TCP, "")
	assert.Equal(t, "Jack", configurator.Name, "")
}
