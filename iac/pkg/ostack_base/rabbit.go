package ostack_base

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

type RabbitConfiguration struct {
	pulumi.ResourceState
}

func NewRabbitConfiguration(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*RabbitConfiguration, error) {
	c := &RabbitConfiguration{}
	if e := ctx.RegisterComponentResource("oucloud:ostack_base:RabbitConfiguration", name, c, opts...); e != nil {
		return nil, e
	}

	// TODO: Add child resources here

	return c, nil
}
