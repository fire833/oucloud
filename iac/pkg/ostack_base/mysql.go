package ostack_base

import (
	"github.com/pulumi/pulumi-mysql/sdk/v3/go/mysql"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MySQLConfiguration struct {
	pulumi.ResourceState
}

func NewMySQLConfiguration(ctx *pulumi.Context, name string, opts ...pulumi.ResourceOption) (*MySQLConfiguration, error) {
	c := &MySQLConfiguration{}
	if e := ctx.RegisterComponentResource("oucloud:ostack_base:MySQLConfiguration", name, c, opts...); e != nil {
		return nil, e
	}

	for _, component := range []string{"keystone", "glance", "nova", "nova_api", "zun", "neutron", "cinder", "swift", "magnum"} {
		_, e := mysql.NewDatabase(ctx, component, &mysql.DatabaseArgs{
			Name:                pulumi.String(component),
			DefaultCharacterSet: pulumi.String("utf8"),
			DefaultCollation:    pulumi.String("utf8GeneralCi"),
		}, pulumi.Parent(c))
		if e != nil {
			return nil, e
		}
	}

	return c, nil
}
