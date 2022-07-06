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

	UserSrvConfigFileName =  "cmd/user/douyin-debug.yaml"
	PubSrvConfigFileName = "cmd/publish/douyin-debug.yaml"
	UserApiConfigFileName =  "cmd/user_api/douyin-debug.yaml"
	PubApiConfigFileName = "cmd/publish_api/douyin-debug.yaml"
	OssApiConfigFileName = "cmd/oss_api/douyin-debug.yaml"
	FeedSrvConfigFileName ="cmd/feed/douyin-debug.yaml"
	FeedApiConfigFileName ="cmd/feed_api/douyin-debug.yaml"
	FavoriteSrvConfigFileName = "cmd/favorite/douyin-debug.yaml"
	FavoriteApiConfigFileName = "cmd/favorite_api/douyin-debug.yaml"
	CommentSrvConfigFileName ="cmd/comment/douyin-debug.yaml"
	CommentApiConfigFileName = "cmd/comment_api/douyin-debug.yaml"
	RelationSrvConfigFileName = "cmd/relation/douyin-debug.yaml"
	RelationApiConfigFileName = "cmd/relation_api/douyin-debug.yaml"




	LogDir = "tmp/naocs/log"
	CacheDir = "tmp/naocs/cache"
	LogLevel = "debug"
	Limit = 30
)