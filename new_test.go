package boot_test

import (
	"testing"

	"github.com/hetue/boot"
)

func TestNew(t *testing.T) {
	boot.New().Build()
}
