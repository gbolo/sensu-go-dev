package mockstore

import (
	"context"

	types "github.com/sensu/sensu-go/api/core/v2"
)

// CreateClusterRoleBinding ...
func (s *MockStore) CreateClusterRoleBinding(ctx context.Context, ClusterRoleBinding *types.ClusterRoleBinding) error {
	args := s.Called(ctx, ClusterRoleBinding)
	return args.Error(0)
}

// CreateOrUpdateClusterRoleBinding ...
func (s *MockStore) CreateOrUpdateClusterRoleBinding(ctx context.Context, ClusterRoleBinding *types.ClusterRoleBinding) error {
	args := s.Called(ctx, ClusterRoleBinding)
	return args.Error(0)
}

// DeleteClusterRoleBinding ...
func (s *MockStore) DeleteClusterRoleBinding(ctx context.Context, name string) error {
	args := s.Called(ctx, name)
	return args.Error(0)
}

// GetClusterRoleBinding ...
func (s *MockStore) GetClusterRoleBinding(ctx context.Context, name string) (*types.ClusterRoleBinding, error) {
	args := s.Called(ctx, name)
	err := args.Error(1)

	if clusterRoleBinding, ok := args.Get(0).(*types.ClusterRoleBinding); ok {
		return clusterRoleBinding, err
	}
	return nil, err
}

// ListClusterRoleBindings ...
func (s *MockStore) ListClusterRoleBindings(ctx context.Context) ([]*types.ClusterRoleBinding, error) {
	args := s.Called(ctx)
	return args.Get(0).([]*types.ClusterRoleBinding), args.Error(1)
}

// UpdateClusterRoleBinding ...
func (s *MockStore) UpdateClusterRoleBinding(ctx context.Context, clusterRoleBinding *types.ClusterRoleBinding) error {
	args := s.Called(ctx, clusterRoleBinding)
	return args.Error(0)
}
