package main

import (
	"NYH/Day20190829/Example_01/logging"
	"log"
	"os"
)

func main() {
	NLog := logging.NewNLog(os.Stdout, logging.LevelInfo, "Test", log.LstdFlags)

	NLog.Debug("Debug-->测试....")
	NLog.Info("Info-->测试....")
	NLog.Warning("Warning-->测试....")
	NLog.Error("Error-->测试....")
	NLog.Fatal("Fatal-->测试....")
}
