package repository

import (
	"context"
	"fmt"

	"github.com/ddr4869/k8s-kafka/config"
	"github.com/ddr4869/k8s-kafka/ent"
)

type Repository struct {
	entClient *ent.Client
}

func (r *Repository) NewEntClient(dbcfg config.DBConf) error {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbcfg.DBHost, dbcfg.DBPort, dbcfg.DBUser, dbcfg.DBName, dbcfg.DBPassword)
	fmt.Println(dataSource)
	client, err := ent.Open("postgres", dataSource)
	if err != nil {
		return err
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		return err
	}
	r.entClient = client
	return nil
}
