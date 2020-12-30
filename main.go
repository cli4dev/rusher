package main

import (
	"time"

	"github.com/micro-plat/hydra"
	"github.com/micro-plat/hydra/hydra/servers/cron"
	"github.com/micro-plat/hydra/hydra/servers/http"
	"github.com/micro-plat/lib4go/sysinfo/pipes"
)

var app = hydra.NewApp(
	hydra.WithPlatName("rusher"),
	hydra.WithSystemName("rusher"),
	hydra.WithServerTypes(cron.CRON, http.API),
	hydra.WithClusterName("prod"),
)

func main() {
	app.Micro("/request", maotai)
	app.CRON("/request", maotai, "43/1 17/1 * * ?")
	app.Start()
}

func maotai(ctx hydra.IContext) interface{} {

	ctx.Log().Info("1. 解锁屏幕")
	pipes.RunString("adb shell input keyevent 82")

	ctx.Log().Info("2. 上滑解锁")
	pipes.RunString("adb shell input swipe 600 1200 600 600")
	time.Sleep(time.Millisecond * 2000)

	ctx.Log().Info("3. home进入第一页")
	pipes.RunString("adb shell input keyevent 3")

	ctx.Log().Info("4. 进入第二屏")
	pipes.RunString("adb shell input swipe 1000 1200 100 1200")

	ctx.Log().Info("5. 打开京东")
	pipes.RunString("adb shell input tap 500 800")
	time.Sleep(time.Millisecond * 1000)

	ctx.Log().Info("6. 进入我的")
	pipes.RunString("adb shell input tap 1100 2350")
	time.Sleep(time.Millisecond * 100)

	ctx.Log().Info("7. 进入收藏")
	pipes.RunString("adb shell input tap 300 420")
	time.Sleep(time.Millisecond * 400)

	ctx.Log().Info("8. 点击第一个收藏")
	pipes.RunString("adb shell input tap 500 600")
	time.Sleep(time.Millisecond * 400)

	ctx.Log().Info("9. 点击预约")
	pipes.RunString("adb shell input tap 1100 2350")
	return nil
}
