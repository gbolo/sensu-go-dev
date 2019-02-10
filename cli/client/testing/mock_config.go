package testing

import (
	"github.com/sensu/sensu-go/types"
	"github.com/stretchr/testify/mock"
)

// MockConfig is a configuration used for CLI testing. When using the
// MockConfig in unit tests, stub out the behavior you wish to test against
// by assigning the appropriate function to the appropriate Func field. If you
// have forgotten to stub a particular function, the program will panic.
type MockConfig struct {
	mock.Mock
}

// APIUrl mocks the API URL config
func (m *MockConfig) APIUrl() string {
	args := m.Called()
	return args.String(0)
}

// CAFile mocks the Trusted-CA-File config
func (m *MockConfig) CAFile() string {
	args := m.Called()
	return args.String(0)
}

// Edition mocks the cluster edition
func (m *MockConfig) Edition() string {
	args := m.Called()
	return args.String(0)
}

// Format mocks the format config
func (m *MockConfig) Format() string {
	args := m.Called()
	return args.String(0)
}

// Namespace mocks the namespace config
func (m *MockConfig) Namespace() string {
	args := m.Called()
	return args.String(0)
}

// SaveAPIUrl mocks saving the API URL
func (m *MockConfig) SaveAPIUrl(url string) error {
	args := m.Called(url)
	return args.Error(0)
}

// SaveCAFile mocks saving the Trusted-CA-File
func (m *MockConfig) SaveCAFile(CAFile string) error {
	args := m.Called(CAFile)
	return args.Error(0)
}

// SaveEdition mocks saving the edition
func (m *MockConfig) SaveEdition(edition string) error {
	args := m.Called(edition)
	return args.Error(0)
}

// SaveFormat mocks saving the format
func (m *MockConfig) SaveFormat(format string) error {
	args := m.Called(format)
	return args.Error(0)
}

// SaveNamespace mocks saving the namespace
func (m *MockConfig) SaveNamespace(namespace string) error {
	args := m.Called(namespace)
	return args.Error(0)
}

// SaveTokens mocks saving the tokens
func (m *MockConfig) SaveTokens(tokens *types.Tokens) error {
	args := m.Called(tokens)
	return args.Error(0)
}

// Tokens mocks the tokens config
func (m *MockConfig) Tokens() *types.Tokens {
	args := m.Called()
	return args.Get(0).(*types.Tokens)
}
