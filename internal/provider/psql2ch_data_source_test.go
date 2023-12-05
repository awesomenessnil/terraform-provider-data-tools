// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccPsql2ChDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Read testing
			{
				Config: testAccPsql2ChDataSourceConfig,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("data.datatools_psql2ch.test", "clickhouse_primarykey", "key_id"),
					resource.TestCheckResourceAttr("data.datatools_psql2ch.test", "clickhouse_columns.0.name", "key_id"),
					resource.TestCheckResourceAttr("data.datatools_psql2ch.test", "clickhouse_columns.0.type", "Nullable(Int)"),
				),
			},
		},
	})
}

const testAccPsql2ChDataSourceConfig = `
data "datatools_psql2ch" "test" {
  postgres_columns = [{
	name                     = "key_id"
	type                     = "int4"  
	character_maximum_length = 0
	is_primary_key           = true
	numeric_precision        = 32
	numeric_scale            = 0
	datetime_precision       = 0
	is_nullable 			 = true
  }]
}
`
