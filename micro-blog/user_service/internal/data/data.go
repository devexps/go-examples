package data

import (
	"context"
	"github.com/devexps/go-examples/micro-blog/api/gen/go/common/conf"
	"github.com/devexps/go-examples/micro-blog/user_service/internal/data/ent"
	"github.com/devexps/go-examples/micro-blog/user_service/internal/data/ent/migrate"
	"github.com/devexps/go-micro/v2/log"

	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Data .
type Data struct {
	log *log.Helper

	db *ent.Client
}

// NewData .
func NewData(logger log.Logger, entClient *ent.Client) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	d := &Data{
		log: l,
		db:  entClient,
	}
	return d, func() {
		l.Info("closing the data resources")
		if err := d.db.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

// NewEntClient creates new ent client
func NewEntClient(logger log.Logger, cfg *conf.Bootstrap) *ent.Client {
	l := log.NewHelper(log.With(logger, "module", "user_service/data/ent"))

	client, err := ent.Open(cfg.Data.Database.Driver, cfg.Data.Database.Source)
	if err != nil {
		l.Fatalf("failed opening connection to db: %v", err)
	}
	if cfg.Data.Database.Migrate {
		if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
			l.Fatalf("failed creating schema resources: %v", err)
		}
	}
	return client
}
