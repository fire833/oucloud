package ostack_infra

import (
	"github.com/pulumi/pulumi-rabbitmq/sdk/v3/go/rabbitmq"
	"github.com/pulumi/pulumi-random/sdk/v4/go/random"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RabbitConfiguration struct {
	pulumi.ResourceState
}

func NewRabbitConfiguration(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*RabbitConfiguration, error) {
	c := &RabbitConfiguration{}
	if e := ctx.RegisterComponentResource("oucloud:ostack_base:RabbitConfiguration", name, c, opts...); e != nil {
		return nil, e
	}

	pass, e := random.NewRandomPassword(ctx, "ostack_pass", &random.RandomPasswordArgs{}, pulumi.Parent(c))
	if e != nil {
		return nil, e
	}

	vhost := "main-host"

	host, e := rabbitmq.NewVHost(ctx, vhost, &rabbitmq.VHostArgs{
		Name: pulumi.String(vhost),
	}, pulumi.Parent(c))
	if e != nil {
		return nil, e
	}

	acct := "openstack"

	user, e := rabbitmq.NewUser(ctx, acct, &rabbitmq.UserArgs{
		Name:     pulumi.String(acct),
		Password: pass.Result,
		Tags: pulumi.StringArray{
			pulumi.String("administrator"),
		},
	}, pulumi.Parent(c), pulumi.DependsOn([]pulumi.Resource{pass}))
	if e != nil {
		return nil, e
	}

	rabbitmq.NewPermissions(ctx, acct+"perms", &rabbitmq.PermissionsArgs{
		User:  pulumi.String(acct),
		Vhost: host.Name,
		Permissions: rabbitmq.PermissionsPermissionsArgs{
			Configure: pulumi.String(".*"),
			Read:      pulumi.String(".*"),
			Write:     pulumi.String(".*"),
		},
	}, pulumi.Parent(c), pulumi.DependsOn([]pulumi.Resource{user}))
	if e != nil {
		return nil, e
	}

	ctx.Export("ostack_user", pulumi.String(acct))
	ctx.Export("ostack_pass", pass.Result)

	return c, nil
}
