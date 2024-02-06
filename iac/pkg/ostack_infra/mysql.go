package ostack_infra

import (
	"errors"

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
		db, e := mysql.NewDatabase(ctx, component+"-db", &mysql.DatabaseArgs{
			Name:                pulumi.String(component),
			DefaultCharacterSet: pulumi.String("utf8"),
			DefaultCollation:    pulumi.String("utf8GeneralCi"),
		}, pulumi.Parent(c))
		if e != nil {
			return nil, e
		}

		pgpKey, found := ctx.GetConfig("pgp:pubkey")
		if !found {
			return nil, errors.New("pgp:pubkey value must be set")
		}

		user, e := mysql.NewUserPassword(ctx, component+"-user", &mysql.UserPasswordArgs{
			User:   pulumi.String(component),
			PgpKey: pulumi.String(pgpKey),
			Host:   pulumi.String("localhost"),
		}, pulumi.Parent(c))
		if e != nil {
			return nil, e
		}

		_, e = mysql.NewGrant(ctx, component+"-grant", &mysql.GrantArgs{
			Database:   db.Name,
			User:       user.User,
			Host:       user.Host,
			Privileges: pulumi.ToStringArray([]string{"ALL"}),
		}, pulumi.Parent(c), pulumi.DependsOn([]pulumi.Resource{user}))

	}

	return c, nil
}
