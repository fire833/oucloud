package main

import (
	"github.com/fire833/oucloud/iac/pkg/ostack_infra"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, e := ostack_infra.NewMySQLConfiguration(ctx, "mysql")
		if e != nil {
			return e
		}

		_, e = ostack_infra.NewRabbitConfiguration(ctx, "rabbit")
		if e != nil {
			return e
		}

		return nil
	})
}
