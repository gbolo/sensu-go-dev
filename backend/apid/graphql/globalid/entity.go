package globalid

import (
	types "github.com/sensu/sensu-go/api/core/v2"
)

//
// Entity
//

var entityName = "entities"

// EntityTranslator global ID resource
var EntityTranslator = commonTranslator{
	name:       entityName,
	encodeFunc: standardEncoder(entityName, "Name"),
	decodeFunc: standardDecoder,
	isResponsibleFunc: func(record interface{}) bool {
		_, ok := record.(*types.Entity)
		return ok
	},
}

// Register entity encoder/decoder
func init() { registerTranslator(EntityTranslator) }
