package globalid

import types "github.com/sensu/sensu-go/api/core/v2"

//
// Roles
//
var roleName = "roles"

// RoleTranslator global ID resource
var RoleTranslator = commonTranslator{
	name:       roleName,
	encodeFunc: standardEncoder(roleName, "Name"),
	decodeFunc: standardDecoder,
	isResponsibleFunc: func(record interface{}) bool {
		_, ok := record.(*types.Role)
		return ok
	},
}

// Register entity encoder/decoder
func init() { registerTranslator(RoleTranslator) }
