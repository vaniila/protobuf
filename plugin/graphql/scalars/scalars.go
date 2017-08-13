package scalars

import (
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/vaniila/protobuf/types"
)

func coerceString(value interface{}) interface{} {
	switch value.(type) {
	case string:
		return value
	case []byte:
		return string(value.([]byte))
	}
	return fmt.Sprintf("%v", value)
}

// String is the GraphQL string type definition
var ByteString *graphql.Scalar = graphql.NewScalar(graphql.ScalarConfig{
	Name:       "ByteString",
	Serialize:  coerceString,
	ParseValue: coerceString,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return valueAST.Value
		}
		return nil
	},
})

func coerceInt(value interface{}) interface{} {
	switch value := value.(type) {
	case bool:
		if value == true {
			return 1
		}
		return 0
	case int:
		return value
	case int8:
		return int(value)
	case int16:
		return int(value)
	case int32:
		return int(value)
	case int64:
		if value < int64(math.MinInt32) || value > int64(math.MaxInt32) {
			return nil
		}
		return int(value)
	case uint:
		return int(value)
	case uint8:
		return int(value)
	case uint16:
		return int(value)
	case uint32:
		if value > uint32(math.MaxInt32) {
			return nil
		}
		return int(value)
	case uint64:
		if value > uint64(math.MaxInt32) {
			return nil
		}
		return int(value)
	case float32:
		if value < float32(math.MinInt32) || value > float32(math.MaxInt32) {
			return nil
		}
		return int(value)
	case float64:
		if value < float64(math.MinInt32) || value > float64(math.MaxInt32) {
			return nil
		}
		return int(value)
	case string:
		val, err := strconv.ParseFloat(value, 0)
		if err != nil {
			return nil
		}
		return coerceInt(val)
	case time.Time:
		return int(value.UnixNano())
	}

	// If the value cannot be transformed into an int, return nil instead of '0'
	// to denote 'no integer found'
	return nil
}

// Timestamp is the GraphQL Timestamp type definition.
var Timestamp *graphql.Scalar = graphql.NewScalar(graphql.ScalarConfig{
	Name:       "Timestamp",
	Serialize:  coerceInt,
	ParseValue: coerceInt,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.IntValue:
			if intValue, err := strconv.Atoi(valueAST.Value); err == nil {
				return intValue
			}
		}
		return nil
	},
})

var DateTime *graphql.Scalar = graphql.NewScalar(graphql.ScalarConfig{
	Name: "DateTime",
	Serialize: func(value interface{}) interface{} {
		switch value := value.(type) {
		case time.Time:
			return value.UTC().Format(time.RFC3339)
		case int:
			return time.Unix(int64(value), 0).UTC().Format(time.RFC3339)
		case int64:
			return time.Unix(value, 0).UTC().Format(time.RFC3339)
		}
		return "0001-01-01T00:00:00Z"
	},
	ParseValue: func(value interface{}) interface{} {
		switch tvalue := value.(type) {
		case string:
			tval, err := time.Parse(time.RFC3339, tvalue)
			if err == nil {
				return tval.UTC()
			}
		case int:
			return time.Unix(int64(tvalue), 0).UTC()
		case int64:
			return time.Unix(tvalue, 0).UTC()
		}
		return time.Time{}
	},
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return valueAST.Value
		}
		return nil
	},
})

func coerceMap(value interface{}) interface{} {
	return value
}

// Map is the GraphQL Map type definition.
var Map *graphql.Scalar = graphql.NewScalar(graphql.ScalarConfig{
	Name:       "Map",
	Serialize:  coerceMap,
	ParseValue: coerceMap,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return valueAST.Value
		}
		return nil
	},
})

func coerceAny(value interface{}) interface{} {
	a, ok := value.(*types.Any)
	if ok {
		return string(a.Value)
	}

	return nil
}

// Any is the GraphQL Any type definition.
var Any *graphql.Scalar = graphql.NewScalar(graphql.ScalarConfig{
	Name:       "Any",
	Serialize:  coerceAny,
	ParseValue: coerceAny,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return valueAST.Value
		}
		return nil
	},
})

var Error *graphql.Scalar = graphql.NewScalar(graphql.ScalarConfig{
	Name:       "Error",
	Serialize:  coerceError,
	ParseValue: coerceError,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return valueAST.Value
		}
		return nil
	},
})

func coerceError(value interface{}) interface{} {
	a, ok := value.(*types.Error)
	if ok {
		return string(a.Error())
	}

	return nil
}

var Permission *graphql.Scalar = graphql.NewScalar(graphql.ScalarConfig{
	Name:       "Permission",
	Serialize:  coercePermission,
	ParseValue: coercePermission,
	ParseLiteral: func(valueAST ast.Value) interface{} {
		switch valueAST := valueAST.(type) {
		case *ast.StringValue:
			return valueAST.Value
		}
		return nil
	},
})

func coercePermission(value interface{}) interface{} {
	a, ok := value.(*types.Permission)
	if ok {
		return a.Permissions()
	}

	return nil
}
