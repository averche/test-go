package main

import (
	"github.com/hashicorp/vault"
)

func main() {
	cluster := vault.NewTestCluster(nil, nil, nil)
	cluster.Start()
	defer cluster.Cleanup()
}
