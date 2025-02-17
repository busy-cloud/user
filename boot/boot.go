package boot

import (
	"github.com/busy-cloud/boat/boot"
	_ "github.com/busy-cloud/user/apis"
)

func init() {
	boot.Register("user", &boot.Task{
		Startup:  Startup,
		Shutdown: nil,
		Depends:  []string{"database", "web"},
	})
}

func Startup() error {

	return nil
}
