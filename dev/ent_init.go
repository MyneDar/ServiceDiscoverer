package dev

import (
	"context"
	"servicediscoverer/ent"
)

var LocalClient *ent.Client
var Ctx context.Context

// TODO fill database Data->Env file
func EntClientInit() (err error) {
	LocalClient, err = ent.Open("postgres", "host=localhost port=5432 user=serviceDisc dbname=serviceDisc password=serviceDisc sslmode=disable")
	if err != nil {
		return err
	}
	Ctx = context.Background()
	// Run the auto migration tool.
	if err = LocalClient.Schema.Create(Ctx); err != nil {
		return err
	}
	return err
}
