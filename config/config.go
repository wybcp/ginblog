package config

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var (
	// RunMode 运行模式
	RunMode string
	// HTTPPort http 端口
	HTTPPort int
	// ReadTimeout 读超时
	ReadTimeout time.Duration
	// WriteTimeout 写超时
	WriteTimeout time.Duration
	// PageSize 页面大小
	PageSize int
	// JwtSecret Jwt验证加密字符串
	JwtSecret string
)

// LoadBase 加载运行参数
func LoadBase() {
	println(viper.ConfigFileUsed())
	if viper.IsSet("RUN_MODE") {
		RunMode = viper.GetString("RUN_MODE")
	} else {
		RunMode = "debug"
	}
}

// LoadServer 加载server参数
func LoadServer() {
	HTTPPort = viper.GetInt("server.HTTP_PORT")
	ReadTimeout = time.Duration(viper.GetInt("server.READ_TIMEOUT")) * time.Second
	WriteTimeout = time.Duration(viper.GetInt("server.WRITE_TIMEOUT")) * time.Second
}

// LoadApp 加载app参数
func LoadApp() {
	JwtSecret = viper.GetString("JWT_SECRET")
	PageSize = viper.GetInt("PAGE_SIZE")
}

// Init 初始化配置
func init() {

	// 初始化日志包
	// c.initLog()
	initConfig()
	LoadBase()
	LoadServer()
	LoadApp()
	watchConfig()
}

func initConfig() error {
	// 如果没有指定配置文件，则解析默认的配置文件
	viper.AddConfigPath("conf")
	viper.SetConfigName("config")
	// 设置配置文件格式为YAML
	viper.SetConfigType("yaml")
	// 读取匹配的环境变量
	viper.AutomaticEnv()
	// 读取环境变量的前缀为
	// viper.SetEnvPrefix("GINBLOG")
	// replacer := strings.NewReplacer(".", "_")
	// viper.SetEnvKeyReplacer(replacer)

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		// log.Infof("Config file changed:%s", e.Name)
		fmt.Printf("Config file changed:%s\n", e.Name)
	})
}
