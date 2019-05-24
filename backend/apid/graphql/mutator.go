package graphql

import (
	v2 "github.com/sensu/sensu-go/api/core/v2"
	"github.com/sensu/sensu-go/backend/apid/graphql/schema"
	"github.com/sensu/sensu-go/graphql"
)

var _ schema.MutatorFieldResolvers = (*mutatorImpl)(nil)

//
// Implement MutatorFieldResolvers
//

type mutatorImpl struct {
	schema.MutatorAliases
}

// IsTypeOf is used to determine if a given value is associated with the type
func (*mutatorImpl) IsTypeOf(s interface{}, p graphql.IsTypeOfParams) bool {
	_, ok := s.(*v2.Mutator)
	return ok
}

// ToJSON implements response to request for 'toJSON' field.
func (*mutatorImpl) ToJSON(p graphql.ResolveParams) (interface{}, error) {
	return v2.WrapResource(p.Source.(v2.Resource)), nil
}
