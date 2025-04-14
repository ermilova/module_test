package Hello

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEnglish(t *testing.T) {
	name := "Ben"
	language := "english"
	expected := "Hello Ben"
	actual, err := Hello(name, language)

	assert.Nil(t, err)
	assert.Equal(t, expected, actual)
}
