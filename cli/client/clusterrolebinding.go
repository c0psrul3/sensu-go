package client

import (
	corev2 "github.com/sensu/sensu-go/api/core/v2"
	"github.com/sensu/sensu-go/types"
)

var clusterRoleBindingsPath = CreateBasePath(coreAPIGroup, coreAPIVersion, "clusterrolebindings")

// CreateClusterRoleBinding with the given cluster role binding
func (client *RestClient) CreateClusterRoleBinding(clusterRoleBinding *types.ClusterRoleBinding) error {
	return client.Post(clusterRoleBindingsPath(), clusterRoleBinding)
}

// DeleteClusterRoleBinding with the given name
func (client *RestClient) DeleteClusterRoleBinding(name string) error {
	return client.Delete(clusterRoleBindingsPath(name))
}

// FetchClusterRoleBinding with the given name
func (client *RestClient) FetchClusterRoleBinding(name string) (*types.ClusterRoleBinding, error) {
	clusterRoleBinding := &types.ClusterRoleBinding{}
	if err := client.Get(clusterRoleBindingsPath(name), clusterRoleBinding); err != nil {
		return nil, err
	}
	return clusterRoleBinding, nil
}

// ListClusterRoleBindings in the cluster
func (client *RestClient) ListClusterRoleBindings(options ListOptions) ([]corev2.ClusterRoleBinding, string, error) {
	var header string
	clusterRoleBindings := []corev2.ClusterRoleBinding{}

	header, err := client.List(clusterRoleBindingsPath(), &clusterRoleBindings, options)
	if err != nil {
		return clusterRoleBindings, header, err
	}

	return clusterRoleBindings, header, nil
}
