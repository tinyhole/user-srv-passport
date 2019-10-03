package db

var (
	DBName = "platform_user_srv_passport"
)

type MongoConfig struct {
	Addr []string `json:"addr"`
	User string `json:"user"`
	Password string `json:"password"`
	PoolLimit int `json:"poolLimit"`
	ReplicaSet string `json:"replicaSet"`
}

func Init() {
	UserPassportOp.Init()
}