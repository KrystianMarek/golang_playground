package configurator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfiguration(t *testing.T) {
	Configuration := Configuration{}
	err := Configuration.GetConfiguration("chat.conf")
	if err != nil {
		panic(err)
	}

	assert.Equal(t, "127.0.0.1:2100", Configuration.RemoteAddr, "")
	assert.Equal(t, true, Configuration.TCP, "")
	assert.Equal(t, "Jack", Configuration.Name, "")
}
