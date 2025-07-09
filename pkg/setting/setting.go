package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File
)

type Server struct {
	RunMode      string
	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type App struct {
	PageSize        int
	JwtSecret       string
	EnableCron      bool
	EnableLog       bool
	RunTimePath     string
	LogSavePath     string
	LogFileExt      string
	LogTimeFormat   string
	UploadAllowExts []string
	UploadMaxSize   int
	ImageSavePath   string
}

var AppSetting = &App{}

type Redis struct {
	Host        string
	Port        int
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisSetting = &Redis{}

func SetUp() {
	loadFile()
	loadServer()
	loadApp()
	loadRedis()
}

func loadFile() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
}

func loadServer() {
	err := Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
	}
}

func loadApp() {
	err := Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}
}

func loadRedis() {
	err := Cfg.Section("redis").MapTo(RedisSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
