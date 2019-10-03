package db

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserPassport struct {
	UID       int64  `bson:"_id"`
	DeviceID  int64  `bson:"device_id"`
	Name      string `bson:"name"`
	Password  string `bson:"password"`
	Salt      string `bson:"salt"`
	Mobile    string `bson:"mobile"`
	Email     string `bson:"email"`
	WeChatID  string `bson:"wechat_id"`
	CreatedAt int64  `bson:"created_at"`
	UpdatedAt int64  `bson:"updated_at"`
}

type userPassportOp struct{
	client *mongo.Client
	table string
}

var UserPassportOp = &userPassportOp{
	table: "user_auth",
}

func (op *userPassportOp) Init() {
	// 读取配置
	conf := &MongoConfig{}
	err := config.Get("mongodb").Scan(conf)
	if err != nil {
		logrus.Fatal(err)
	}

	uri := fmt.Sprintf("mongodb://%s:%s@%s/test?w=majority", conf.User, conf.Password, conf.Addr)
	op.client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		logrus.Fatal(err)
	}
}

func (op *userPassportOp) Insert(ctx context.Context, up *UserPassport) error {
	_, err := op.client.Database(DBName).Collection(op.table).InsertOne(ctx, up)
	if err != nil {
		return err
	}

	return nil
}

func (op *userPassportOp) Get(ctx context.Context, query bson.D) (up *UserPassport, err error) {
	up = &UserPassport{}
	err = op.client.Database(DBName).Collection(op.table).FindOne(ctx, query).Decode(up)
	if err != nil {
		return nil, err
	}

	return up, nil
}
