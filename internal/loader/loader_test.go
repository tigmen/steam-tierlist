package loader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewLoader_Error(t *testing.T) {
	config := &Config{
		LogLevel: "z",
	}

	_, err := NewLoader(config)

	assert.Error(t, err)
}

// func Test_NewLoader_Success(t *testing.T) {
// 	config := &Config{
// 		LogLevel: "debug",
// 	}

// 	example_loader := &Loader{
// 		config: config,
// 		logger: logrus.New(),
// 	}
// 	example_loader.configureLogger()

// 	loader, err := NewLoader(config)

// 	assert.NoError(t, err)
// 	assert.Equal(t, loader, example_loader)
// }
