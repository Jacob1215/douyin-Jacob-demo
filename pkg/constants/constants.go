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

	UserSrvConfigFileName =  "cmd/srv/user/douyin-debug.yaml"
	PubSrvConfigFileName = "cmd/srv/publish/douyin-debug.yaml"
	UserApiConfigFileName =  "cmd/api/user_api/douyin-debug.yaml"
	PubApiConfigFileName = "cmd/api/publish_api/douyin-debug.yaml"
	OssApiConfigFileName = "cmd/api/oss_api/douyin-debug.yaml"
	FeedSrvConfigFileName ="cmd/srv/feed/douyin-debug.yaml"
	FeedApiConfigFileName ="cmd/api/feed_api/douyin-debug.yaml"
	FavoriteSrvConfigFileName = "cmd/srv/favorite/douyin-debug.yaml"
	FavoriteApiConfigFileName = "cmd/api/favorite_api/douyin-debug.yaml"
	CommentSrvConfigFileName ="cmd/srv/comment/douyin-debug.yaml"
	CommentApiConfigFileName = "cmd/api/comment_api/douyin-debug.yaml"
	RelationSrvConfigFileName = "cmd/srv/relation/douyin-debug.yaml"
	RelationApiConfigFileName = "cmd/api/relation_api/douyin-debug.yaml"




	LogDir = "tmp/naocs/log"
	CacheDir = "tmp/naocs/cache"
	LogLevel = "debug"
	Limit = 30
)