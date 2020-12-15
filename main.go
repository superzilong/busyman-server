package main

import (
	"fmt"
	"gg/dao"
	"gg/pkg/logger"
	"gg/pkg/settings"
	"gg/pkg/snowflake"
	"gg/routes"

	"go.uber.org/zap"
)

func main() {
	// 1. Load config
	if err := settings.Init(); err != nil {
		fmt.Printf("init settings failed, err: %v\n", err)
		return
	}
	// 2. Init logger
	if err := logger.Init(); err != nil {
		fmt.Printf("init logger failed, err: %v\n", err)
		return
	}
	defer zap.L().Sync()
	zap.L().Info("logger init success...")

	// 3. Init gorm.
	if err := dao.Init(); err != nil {
		fmt.Printf("init dao failed, err: %v\n", err)
		return
	}
	defer dao.Close()
	// 4. Init Redis

	// 5. Init snowflake
	if err := snowflake.Init(settings.Conf.StartTime, settings.Conf.MachineID); err != nil {
		fmt.Printf("init snowflake failed, err:%v\n", err)
		return
	}

	// 6. Register router
	r := routes.Setup()
	// 7. Startup server
	r.Run(fmt.Sprintf(":%d", settings.Conf.Port))
	// 8. Graceful shutdown and reboot
}
