package main

import (
	"github.com/busy-cloud/boat/apps"
	"github.com/busy-cloud/boat/boot"
	"github.com/busy-cloud/boat/log"
	"github.com/busy-cloud/boat/web"
	_ "github.com/busy-cloud/user/internal"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	//测试
	apps.Pages().Dir("pages", "")
}

func main() {
	viper.SetConfigName("user")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs

		//关闭web，出发
		_ = web.Shutdown()
	}()

	//安全退出
	defer boot.Shutdown()

	err := boot.Startup()
	if err != nil {
		log.Fatal(err)
		return
	}

	err = web.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
