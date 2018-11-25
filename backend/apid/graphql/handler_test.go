package graphql

import (
	"testing"

	client "github.com/sensu/sensu-go/backend/apid/graphql/mockclient"
	"github.com/sensu/sensu-go/graphql"
	types "github.com/sensu/sensu-go/api/core/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestHandlerTypeHandlersField(t *testing.T) {
	handler := types.FixtureHandler("my-handler")
	handler.Handlers = []string{"one", "two"}

	client, factory := client.NewClientFactory()
	impl := &handlerImpl{factory: factory}

	// Success
	client.On("ListHandlers", mock.Anything).Return([]types.Handler{
		*types.FixtureHandler("one"),
		*types.FixtureHandler("two"),
		*types.FixtureHandler("three"),
	}, nil).Once()
	res, err := impl.Handlers(graphql.ResolveParams{Source: handler})
	require.NoError(t, err)
	assert.NotEmpty(t, res)
}

func TestHandlerTypeMutatorField(t *testing.T) {
	mutator := types.FixtureMutator("my-mutator")
	handler := types.FixtureHandler("my-handler")
	handler.Mutator = mutator.Name

	client, factory := client.NewClientFactory()
	impl := &handlerImpl{factory: factory}

	// Success
	client.On("FetchMutator", mutator.Name).Return(mutator, nil).Once()
	res, err := impl.Mutator(graphql.ResolveParams{Source: handler})
	require.NoError(t, err)
	assert.NotEmpty(t, res)
}
