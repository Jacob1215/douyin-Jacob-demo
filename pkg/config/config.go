package config

type ServerConfig struct {
	Name        string        `mapstructure:"name" json:"name"`
	Host string `mapstructure:"host" json:"host"`
	Tags        []string       `mapstructure:"tags" json:"tags"`
	Port        int           `mapstructure:"port" json:"port"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv" json:"user_srv"`
	PublishSrvInfo PublishSrvConfig `mapstructure:"publish_srv" json:"publish_srv"`
	FeedSrvInfo FeedSrvConfig `mapstructure:"feed_srv" json:"feed_srv"`
	FavoriteInfo FavoriteSrvConfig `mapstructure:"favorite_srv" json:"favorite_srv"`
	CommentInfo CommentSrvConfig `mapstructure:"comment_srv" json:"comment_srv"`
	RelationInfo RelationConfig `mapstructure:"relation_srv" json:"relation_srv"`

	JWTInfo     JWTConfig     `mapstructure:"jwt" json:"jwt"`
	ConsulInfo  ConsulConfig  `mapstructure:"consul" json:"consul"`
	JaegerInfo  JaegerConfig   `mapstructure:"jaeger" json:"jaeger"`
	MysqlInfo  MysqlConfig  `mapstructure:"mysql" json:"mysql"`
	OssInfo OssConfig `mapstructure:"oss" json:"oss"`
}

type RelationConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type UserSrvConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type PublishSrvConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type FeedSrvConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type FavoriteSrvConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}

type CommentSrvConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}


type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}


type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}



//nacos配置。
type NacosConfig struct {
	Host      string `mapstructure:"host"`
	Port      uint64    `mapstructure:"port"`
	Namespace string `mapstructure:"namespace"`
	User      string `mapstructure:"user"`
	Password  string `mapstructure:"password"`
	DataId    string `mapstructure:"dataid"`
	Group     string `mapstructure:"group"`
}

type JaegerConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}


type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}


type OssConfig struct {
	EndPoint string `mapstructure:"endpoint" json:"endpoint"`
	ApiKey string `mapstructure:"accessKeyId" json:"accessKeyId"`
	ApiSecrect string `mapstructure:"secrect" json:"secrect"`
	Host string `mapstructure:"host" json:"host"`
	CallBackUrl string `mapstructure:"callback_url" json:"callback_url"`
	UploadDir string `mapstructure:"upload_dir" json:"upload_dir"`
}









