// Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.

package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

const (
	LoadBalancerShapeResourceConfig = LoadBalancerShapeResourceDependencies + `

`
	LoadBalancerShapePropertyVariables = `

`
	LoadBalancerShapeResourceDependencies = ""
)

func TestLoadBalancerLoadBalancerShapeResource_basic(t *testing.T) {
	provider := testAccProvider
	config := testProviderConfig()

	compartmentId2 := getRequiredEnvSetting("compartment_id_for_update")
	compartmentIdVariableStr2 := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId2)

	datasourceName := "data.oci_load_balancer_shapes.test_load_balancer_shapes"

	resource.Test(t, resource.TestCase{
		Providers: map[string]terraform.ResourceProvider{
			"oci": provider,
		},
		Steps: []resource.TestStep{
			// verify datasource
			{
				Config: config + `

data "oci_load_balancer_shapes" "test_load_balancer_shapes" {
	#Required
	compartment_id = "${var.compartment_id}"
}
                ` + compartmentIdVariableStr2 + LoadBalancerShapeResourceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId2),

					resource.TestCheckResourceAttrSet(datasourceName, "shapes.#"),
					resource.TestCheckResourceAttrSet(datasourceName, "shapes.0.name"),
				),
			},
		},
	})
}
