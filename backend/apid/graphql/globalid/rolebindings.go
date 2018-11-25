package globalid

import types "github.com/sensu/sensu-go/api/core/v2"

//
// Role Bindings
//
var roleBindingName = "rolebindings"

// RoleBindingTranslator global ID resource
var RoleBindingTranslator = commonTranslator{
	name:       roleBindingName,
	encodeFunc: standardEncoder(roleBindingName, "Name"),
	decodeFunc: standardDecoder,
	isResponsibleFunc: func(record interface{}) bool {
		_, ok := record.(*types.RoleBinding)
		return ok
	},
}

// Register entity encoder/decoder
func init() { registerTranslator(RoleBindingTranslator) }
