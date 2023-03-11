package dummie

import (
	"fmt"
	"reflect"
)

type KeyValueCache struct {
	Cache map[string]map[string]any
}

func NewCache(configuration *Configuration) *KeyValueCache {
	cache := KeyValueCache{
		Cache: map[string]map[string]any{},
	}

	for targetType, values := range configuration.DefaultValues {
		for name, value := range values {
			cache.SetValue(targetType, name, value)
		}
	}
	fmt.Println(cache.Cache)
	return &cache
}

func (c *KeyValueCache) GetValue(t reflect.Type, fieldName string) any {
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return c.tryGetCachedValue(t.Name(), fieldName, "int")
	case reflect.Float32, reflect.Float64:
		return c.tryGetCachedValue(t.Name(), fieldName, "float")
	case reflect.Bool:
		return c.tryGetCachedValue(t.Name(), fieldName, "bool")
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return c.tryGetCachedValue(t.Name(), fieldName, "uint")
	case reflect.Struct, reflect.Map, reflect.Array, reflect.Slice, reflect.Chan, reflect.Pointer:
		return c.tryGetCachedValue(t.Name(), fieldName, t.Name())
	case reflect.String:
		return c.tryGetCachedValue(t.Name(), fieldName, "string")
	case reflect.Complex64, reflect.Complex128:
		return c.tryGetCachedValue(t.Name(), fieldName, "complex128")
	}
	return c.Cache[t.Name()][fieldName]
}

func (c *KeyValueCache) SetValue(typeName string, name string, value any) {
	if _, ok := c.Cache[typeName]; !ok {
		c.Cache[typeName] = map[string]any{}
		c.Cache[typeName][""] = value
	}
	c.Cache[typeName][name] = value
}

func (c *KeyValueCache) tryGetCachedValue(primaryKey string, secondKey string, defaultKey string) any {
	fmt.Printf("tryGetCachedValue, %s %s %s\n", primaryKey, secondKey, defaultKey)
	if valMap, ok := c.Cache[primaryKey]; ok {
		if val, ok := valMap[secondKey]; ok {
			return val
		} else if secondKey != "" {
			if val, ok := valMap[""]; ok {
				return val
			}
		}
	}
	if val, ok := c.Cache[defaultKey][secondKey]; ok {
		return val
	}
	if val, ok := c.Cache[defaultKey][""]; ok {
		return val
	}
	return nil
}
