package constants

const (
	NoteTableName           = "note"
	UserTableName           = "user"
	SecretKey               = "secret key"
	IdentityKey             = "id"
	Total                   = "total"
	Notes                   = "notes"
	NoteID                  = "note_id"
	ApiServiceName          = "demoapi"
	NoteServiceName         = "demonote"
	UserServiceName         = "demouser"
	MySQLDefaultDSN         = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	EtcdAddress             = "127.0.0.1:2379"
	CPURateLimit    float64 = 80.0
	DefaultLimit            = 10

	SrvConfigFileName =  "cmd/user/douyin-debug.yaml"
	PubSrvConfigFileName = "cmd/publish/douyin-debug.yaml"
	ApiConfigFileName =  "cmd/user_api/douyin-debug.yaml"
	PubApiConfigFileName = "cmd/publish_api/douyin-debug.yaml"
	LogDir = "tmp/naocs/log"
	CacheDir = "tmp/naocs/cache"
	LogLevel = "debug"
)