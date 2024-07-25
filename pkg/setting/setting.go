package setting

import (
	"time"
)

type App struct {
	JwtSecret string
	PageSize  int
	PrefixUrl string

	RuntimeRootPath string

	ImageSavePath  string
	ImageMaxSize   int
	ImageAllowExts []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseSetting = &Database{}

// Setup initialize the configuration instance
func Setup() {
	AppSetting.JwtSecret = "NozfkxSLpMQACMgKwm#rfcFEfSMytD2B@&FJT!XQ6qdf5#uGN5"
	AppSetting.PageSize = 10
	AppSetting.PrefixUrl = "http://127.0.0.1:1604"

	// AppSetting.RuntimeRootPath = "runtime/"

	AppSetting.LogSavePath = "logs/"
	AppSetting.LogSaveName = "log"
	AppSetting.LogFileExt = "log"
	AppSetting.TimeFormat = "20060102"

	ServerSetting.RunMode = "debug"
	ServerSetting.HttpPort = 8000
	ServerSetting.ReadTimeout = 60
	ServerSetting.WriteTimeout = 60

	DatabaseSetting.Type = "sqlserver"
	DatabaseSetting.User = "sa"
	DatabaseSetting.Password = "123"
	DatabaseSetting.Host = "HUNG:1433"
	DatabaseSetting.Name = "RETAIL_DEMO_V0"
	//DatabaseSetting.Host = "113.161.84.61:1433"
	//DatabaseSetting.Name = "RETAIL_DEMO"
	DatabaseSetting.TablePrefix = ""

	AppSetting.ImageMaxSize = 10 * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
	AppSetting.ImageSavePath = "upload/images/"

}
