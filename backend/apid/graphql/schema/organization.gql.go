// Code generated by scripts/gengraphql.go. DO NOT EDIT.

package schema

import (
	errors "errors"
	graphql1 "github.com/graphql-go/graphql"
	graphql "github.com/sensu/sensu-go/graphql"
)

// OrganizationIDFieldResolver implement to resolve requests for the Organization's id field.
type OrganizationIDFieldResolver interface {
	// ID implements response to request for id field.
	ID(p graphql.ResolveParams) (string, error)
}

// OrganizationDescriptionFieldResolver implement to resolve requests for the Organization's description field.
type OrganizationDescriptionFieldResolver interface {
	// Description implements response to request for description field.
	Description(p graphql.ResolveParams) (string, error)
}

// OrganizationNameFieldResolver implement to resolve requests for the Organization's name field.
type OrganizationNameFieldResolver interface {
	// Name implements response to request for name field.
	Name(p graphql.ResolveParams) (string, error)
}

// OrganizationEnvironmentsFieldResolver implement to resolve requests for the Organization's environments field.
type OrganizationEnvironmentsFieldResolver interface {
	// Environments implements response to request for environments field.
	Environments(p graphql.ResolveParams) (interface{}, error)
}

// OrganizationIconIDFieldResolver implement to resolve requests for the Organization's iconId field.
type OrganizationIconIDFieldResolver interface {
	// IconID implements response to request for iconId field.
	IconID(p graphql.ResolveParams) (Icon, error)
}

//
// OrganizationFieldResolvers represents a collection of methods whose products represent the
// response values of the 'Organization' type.
//
// == Example SDL
//
//   """
//   Dog's are not hooman.
//   """
//   type Dog implements Pet {
//     "name of this fine beast."
//     name:  String!
//
//     "breed of this silly animal; probably shibe."
//     breed: [Breed]
//   }
//
// == Example generated interface
//
//   // DogResolver ...
//   type DogFieldResolvers interface {
//     DogNameFieldResolver
//     DogBreedFieldResolver
//
//     // IsTypeOf is used to determine if a given value is associated with the Dog type
//     IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
//   }
//
// == Example implementation ...
//
//   // DogResolver implements DogFieldResolvers interface
//   type DogResolver struct {
//     logger logrus.LogEntry
//     store interface{
//       store.BreedStore
//       store.DogStore
//     }
//   }
//
//   // Name implements response to request for name field.
//   func (r *DogResolver) Name(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     return dog.GetName()
//   }
//
//   // Breed implements response to request for breed field.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     breed := r.store.GetBreed(dog.GetBreedName())
//     return breed
//   }
//
//   // IsTypeOf is used to determine if a given value is associated with the Dog type
//   func (r *DogResolver) IsTypeOf(p graphql.IsTypeOfParams) bool {
//     // ... implementation details ...
//     _, ok := p.Value.(DogGetter)
//     return ok
//   }
//
type OrganizationFieldResolvers interface {
	OrganizationIDFieldResolver
	OrganizationDescriptionFieldResolver
	OrganizationNameFieldResolver
	OrganizationEnvironmentsFieldResolver
	OrganizationIconIDFieldResolver
}

// OrganizationAliases implements all methods on OrganizationFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
//
// == Example SDL
//
//    type Dog {
//      name:   String!
//      weight: Float!
//      dob:    DateTime
//      breed:  [Breed]
//    }
//
// == Example generated aliases
//
//   type DogAliases struct {}
//   func (_ DogAliases) Name(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Weight(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Dob(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//
// == Example Implementation
//
//   type DogResolver struct { // Implements DogResolver
//     DogAliases
//     store store.BreedStore
//   }
//
//   // NOTE:
//   // All other fields are satisified by DogAliases but since this one
//   // requires hitting the store we implement it in our resolver.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) interface{} {
//     dog := v.(*Dog)
//     return r.BreedsById(dog.BreedIDs)
//   }
//
type OrganizationAliases struct{}

// ID implements response to request for 'id' field.
func (_ OrganizationAliases) ID(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'id'")
	}
	return ret, err
}

// Description implements response to request for 'description' field.
func (_ OrganizationAliases) Description(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'description'")
	}
	return ret, err
}

// Name implements response to request for 'name' field.
func (_ OrganizationAliases) Name(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'name'")
	}
	return ret, err
}

// Environments implements response to request for 'environments' field.
func (_ OrganizationAliases) Environments(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// IconID implements response to request for 'iconId' field.
func (_ OrganizationAliases) IconID(p graphql.ResolveParams) (Icon, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := Icon(val.(string)), true
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'iconId'")
	}
	return ret, err
}

// OrganizationType Organization represents a Sensu organization in RBAC
var OrganizationType = graphql.NewType("Organization", graphql.ObjectKind)

// RegisterOrganization registers Organization object type with given service.
func RegisterOrganization(svc *graphql.Service, impl OrganizationFieldResolvers) {
	svc.RegisterObject(_ObjectTypeOrganizationDesc, impl)
}
func _ObjTypeOrganizationIDHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(OrganizationIDFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.ID(frp)
	}
}

func _ObjTypeOrganizationDescriptionHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(OrganizationDescriptionFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Description(frp)
	}
}

func _ObjTypeOrganizationNameHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(OrganizationNameFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Name(frp)
	}
}

func _ObjTypeOrganizationEnvironmentsHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(OrganizationEnvironmentsFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Environments(frp)
	}
}

func _ObjTypeOrganizationIconIDHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(OrganizationIconIDFieldResolver)
	return func(frp graphql1.ResolveParams) (interface{}, error) {

		val, err := resolver.IconID(frp)
		return string(val), err
	}
}

func _ObjectTypeOrganizationConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "Organization represents a Sensu organization in RBAC",
		Fields: graphql1.Fields{
			"description": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Description is more information for an organization.",
				Name:              "description",
				Type:              graphql1.String,
			},
			"environments": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Environments that belong to the organization",
				Name:              "environments",
				Type:              graphql1.NewNonNull(graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("Environment")))),
			},
			"iconId": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "IconId. Experimental. Use graphical interfaces as symbolic reference to organization",
				Name:              "iconId",
				Type:              graphql1.NewNonNull(graphql.OutputType("Icon")),
			},
			"id": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "The globally unique identifier of the check.",
				Name:              "id",
				Type:              graphql1.NewNonNull(graphql1.ID),
			},
			"name": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Name is the unique identifier for an organization.",
				Name:              "name",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see OrganizationFieldResolvers.")
		},
		Name: "Organization",
	}
}

// describe Organization's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeOrganizationDesc = graphql.ObjectDesc{
	Config: _ObjectTypeOrganizationConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"description":  _ObjTypeOrganizationDescriptionHandler,
		"environments": _ObjTypeOrganizationEnvironmentsHandler,
		"iconId":       _ObjTypeOrganizationIconIDHandler,
		"id":           _ObjTypeOrganizationIDHandler,
		"name":         _ObjTypeOrganizationNameHandler,
	},
}

// Icon self descriptive
type Icon string

// Icons holds enum values
var Icons = _EnumTypeIconValues{
	BRIEFCASE:  "BRIEFCASE",
	DONUT:      "DONUT",
	EMOTICON:   "EMOTICON",
	ESPRESSO:   "ESPRESSO",
	EXPLORE:    "EXPLORE",
	FIRE:       "FIRE",
	HALFHEART:  "HALFHEART",
	HEART:      "HEART",
	MUG:        "MUG",
	POLYGON:    "POLYGON",
	VISIBILITY: "VISIBILITY",
}

// IconType self descriptive
var IconType = graphql.NewType("Icon", graphql.EnumKind)

// RegisterIcon registers Icon object type with given service.
func RegisterIcon(svc *graphql.Service) {
	svc.RegisterEnum(_EnumTypeIconDesc)
}
func _EnumTypeIconConfigFn() graphql1.EnumConfig {
	return graphql1.EnumConfig{
		Description: "self descriptive",
		Name:        "Icon",
		Values: graphql1.EnumValueConfigMap{
			"BRIEFCASE": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "BRIEFCASE",
			},
			"DONUT": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "DONUT",
			},
			"EMOTICON": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "EMOTICON",
			},
			"ESPRESSO": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "ESPRESSO",
			},
			"EXPLORE": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "EXPLORE",
			},
			"FIRE": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "FIRE",
			},
			"HALFHEART": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "HALFHEART",
			},
			"HEART": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "HEART",
			},
			"MUG": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "MUG",
			},
			"POLYGON": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "POLYGON",
			},
			"VISIBILITY": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "VISIBILITY",
			},
		},
	}
}

// describe Icon's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _EnumTypeIconDesc = graphql.EnumDesc{Config: _EnumTypeIconConfigFn}

type _EnumTypeIconValues struct {
	// BRIEFCASE - self descriptive
	BRIEFCASE Icon
	// DONUT - self descriptive
	DONUT Icon
	// EMOTICON - self descriptive
	EMOTICON Icon
	// ESPRESSO - self descriptive
	ESPRESSO Icon
	// EXPLORE - self descriptive
	EXPLORE Icon
	// FIRE - self descriptive
	FIRE Icon
	// HALFHEART - self descriptive
	HALFHEART Icon
	// HEART - self descriptive
	HEART Icon
	// MUG - self descriptive
	MUG Icon
	// POLYGON - self descriptive
	POLYGON Icon
	// VISIBILITY - self descriptive
	VISIBILITY Icon
}