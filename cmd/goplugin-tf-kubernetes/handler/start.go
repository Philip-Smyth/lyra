package handler

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/lyraproj/lyra/cmd/goplugin-tf-kubernetes/generated"
	"github.com/lyraproj/puppet-evaluator/eval"
	"github.com/lyraproj/servicesdk/grpc"
	"github.com/lyraproj/servicesdk/service"
	"github.com/terraform-providers/terraform-provider-kubernetes/kubernetes"
)

// Server configures the Terraform provider and creates an instance of the server
func Server(c eval.Context) *service.Server {
	sb := service.NewServerBuilder(c, "TerraformKubernetes")
	generated.Initialize(sb, kubernetes.Provider().(*schema.Provider))
	return sb.Server()
}

// Start this server running
func Start() {
	eval.Puppet.Do(func(c eval.Context) {
		grpc.Serve(c, Server(c))
	})
}
