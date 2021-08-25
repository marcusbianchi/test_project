package algoritms_test

import (
	"github.com/go-playground/assert/v2"
	"github.com/marcusbianchi/test_project/internal/algoritms"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	actual := algoritms.Hello("Marcus")
	assert.Equal(t, "Hello Marcus", actual)
}
