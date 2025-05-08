package internal

import (
	"github.com/harluo/di"
)

func init() {
	di.New().Instance().Put(newDefault).Build().Apply()
}
