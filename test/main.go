package main

import (
	"encoding/json"
	"github.com/busy-cloud/boat/apps"
	"github.com/busy-cloud/boat/boot"
	"github.com/busy-cloud/boat/log"
	"github.com/busy-cloud/boat/store"
	"github.com/busy-cloud/boat/web"
	_ "github.com/busy-cloud/user/internal"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	manifest, err := os.ReadFile("manifest.json")
	if err != nil {
		log.Fatal(err)
	}

	//注册为内部插件
	var a apps.App
	err = json.Unmarshal(manifest, &a)
	if err != nil {
		log.Fatal(err)
	}

	a.AssetsFS = store.Dir("assets")
	a.PagesFS = store.Dir("pages")
	a.TablesFS = store.Dir("tables")

	apps.Register(&a)
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
