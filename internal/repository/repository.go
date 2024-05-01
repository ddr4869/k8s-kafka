package repository

import (
	"context"

	"github.com/ddr4869/k8s-kafka/ent"
)

type Repository struct {
	entClient *ent.Client
}

func (r *Repository) NewEntClient() error {
	client, err := ent.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=postgres password=1821 sslmode=disable")
	if err != nil {
		return err
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		return err
	}
	r.entClient = client
	return nil
}
