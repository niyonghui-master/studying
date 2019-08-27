package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

func init() {
	viper.SetConfigName("config") // 指定配置文件的文件名称（不需要指定配置文件的扩展名）
	viper.AddConfigPath(".")      // 设置配置文件和可执行二进制文件在同一个目录
	err := viper.ReadInConfig()   // 根据以上配置读取加载配置文件
	if err != nil {
		log.Fatal(err) // 读取配置文件出现致命错误
	}
}

func main() {
	fmt.Println("获取配置文件的string", viper.GetString(`app.name`))
	fmt.Println("获取配置文件的int", viper.GetInt(`app.foo`))
	fmt.Println("获取配置文件的bool", viper.GetBool(`app.bar`))
	fmt.Println("获取配置文件的map[string]string", viper.GetStringMapString(`app`))

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		// viper配置发生变化了，执行响应的操作
		fmt.Println("config file changed:", in.Name)
	})

	time.Sleep(10 * time.Second)

	fmt.Println("获取配置文件的string", viper.GetString(`app.name`))
	fmt.Println("获取配置文件的int", viper.GetInt(`app.foo`))
	fmt.Println("获取配置文件的bool", viper.GetBool(`app.bar`))
	fmt.Println("获取配置文件的map[string]string", viper.GetStringMapString(`app`))
}
