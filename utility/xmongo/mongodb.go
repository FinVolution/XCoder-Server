package xmongo

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"sync"
)

type Mongodb struct {
	DBName  string
	configs map[string]*Config
	db      sync.Map
	lock    sync.Mutex
}

func NewFromFile() (*Mongodb, error) {
	var (
		configs = make(map[string]*Config)
		ctx     = gctx.New()
	)
	mongodbYAML, err := g.Cfg().Get(ctx, "mongodb.default")
	if err != nil {
		return nil, err
	}

	dbName, ok := mongodbYAML.Map()["dbname"].(string)
	if !ok {
		return nil, fmt.Errorf("dbname is empty")
	}

	config := dbConfig{}
	err = gconv.Struct(mongodbYAML, &config)
	if err != nil {
		return nil, err
	}
	g.Log().Infof(ctx, "mongodb config: %+v", config)

	configs[dbName] = &Config{
		c: &config,
	}

	return &Mongodb{
		DBName:  dbName,
		configs: configs,
	}, nil
}

func (m *Mongodb) GetClient(ctx context.Context, name string) (*mongo.Client, error) {
	// check cache
	if client, ok := m.db.Load(name); ok {
		return client.(*mongo.Client), nil
	}
	c, ok := m.configs[name]
	if !ok {
		return nil, fmt.Errorf("%s has no mongodb config", name)
	}

	m.lock.Lock()
	defer m.lock.Unlock()
	// may has cache
	if client, ok := m.db.Load(name); ok {
		return client.(*mongo.Client), nil
	}

	opts := options.Client().
		ApplyURI("mongodb://" + c.c.Hosts).
		SetAuth(options.Credential{
			Username:   c.c.Username,
			Password:   c.c.Password,
			AuthSource: c.c.AuthSource,
		}).
		SetReadPreference(readpref.PrimaryPreferred())

	if c.c.ReplicaSet != "" {
		opts.SetReplicaSet(c.c.ReplicaSet)
	}

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	m.db.Store(name, client)
	return client, nil
}

func (m *Mongodb) GetDBName(ctx context.Context) string {
	return m.DBName
}

func (m *Mongodb) Close() {
	m.db.Range(func(_, v interface{}) bool {
		v.(*mongo.Client).Disconnect(context.Background())
		return true
	})
}
