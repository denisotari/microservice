package proto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloIn_Validate(t *testing.T) {
	var in User
	assert.Error(t, in.Validate())

	in.Name = "Wesley"
	assert.NoError(t, in.Validate())
}
