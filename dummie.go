package dummie

import (
	"reflect"

	"github.com/yeyuexia/dummie-go/constant"
)

type Dummie struct {
	Configuration    *Configuration
	Cache            *KeyValueCache
	GeneratorManager *GeneratorManager
}

func Inflate(ptr any) error {
	return InflateWithConfiguration(ptr, &Configuration{})
}

func InflateWithConfiguration(ptr any, configuration *Configuration) error {
	instance := &Dummie{
		Configuration:    configuration,
		Cache:            NewCache(configuration),
		GeneratorManager: NewGenerators(),
	}

	return instance.inflateValue(reflect.ValueOf(ptr), []string{})
}

func (d *Dummie) inflateValue(target reflect.Value, path []string) (err error) {
	if !target.CanSet() {
		if target.Kind() == reflect.Pointer {
			d.inflateValue(target.Elem(), path)
		}
		return
	}
	if !d.tryInflateDirectly(target, path) {
		switch target.Kind() {
		case reflect.Pointer, reflect.UnsafePointer:
			if target.IsNil() && target.CanAddr() {
				target.Set(reflect.New(target.Type().Elem()))
			}
			err = d.inflateValue(target.Elem(), path)
		case reflect.Array, reflect.Slice:
			err = d.generateArrayData(target, path)
		case reflect.Map:
			//TODO: implementation
		case reflect.Struct:
			err = d.generateStructData(target, path)
		}
	}
	return
}

func (d *Dummie) generateStructData(target reflect.Value, path []string) error {
	structType := target.Type()
	for i := 0; i < target.NumField(); i++ {
		field := target.Field(i)

		if err := d.inflateValue(field, append([]string{structType.Field(i).Name}, path...)); err != nil {
			return err
		}
	}
	return nil
}

func (d *Dummie) generateArrayData(target reflect.Value, path []string) (err error) {
	if target.IsNil() && target.CanAddr() {
		target.Set(reflect.MakeSlice(target.Type(), 0, 1))
	}
	path = append([]string{"[]"}, path...)
	elemType := target.Type().Elem()
	var elem reflect.Value
	if elemType.Kind() == reflect.Pointer || elemType.Kind() == reflect.UnsafePointer {
		elem = reflect.New(elemType.Elem())
		err = d.inflateValue(elem.Elem(), path)
		target.Set(reflect.Append(target, elem))
	} else {
		elem = reflect.New(elemType)
		err = d.inflateValue(elem.Elem(), path)
		target.Set(reflect.Append(target, elem.Elem()))
	}
	return
}

func (d *Dummie) tryInflateDirectly(target reflect.Value, path []string) bool {
	v := d.getCachedValue(target.Type(), path)
	if v == nil {
		return false
	}
	switch v.(type) {
	case bool:
		target.SetBool(v.(bool))
	case int:
		target.SetInt(int64(v.(int)))
	case int8:
		target.SetInt(int64(v.(int8)))
	case int16:
		target.SetInt(int64(v.(int16)))
	case int32:
		target.SetInt(int64(v.(int32)))
	case int64:
		target.SetInt(v.(int64))
	case uint:
		target.SetUint(uint64(v.(uint)))
	case uint8:
		target.SetUint(uint64(v.(uint8)))
	case uint16:
		target.SetUint(uint64(v.(uint16)))
	case uint32:
		target.SetUint(uint64(v.(uint32)))
	case uint64:
		target.SetUint(v.(uint64))
	case float32:
		target.SetFloat(float64(v.(float32)))
	case float64:
		target.SetFloat(v.(float64))
	case string:
		target.SetString(v.(string))
	case complex64:
		target.SetComplex(complex128(v.(complex64)))
	case complex128:
		target.SetComplex(v.(complex128))
	default:
		target.Set(reflect.ValueOf(v))
	}
	return true
}

func (d *Dummie) getCachedValue(t reflect.Type, path []string) any {
	fieldName := getFieldName(path)
	value := d.Cache.GetValue(t, fieldName)
	if value != nil {
		return value
	}
	value = d.GeneratorManager.GenerateValue(t, d.getStrategy(t, fieldName), path)
	if value != nil {
		d.Cache.SetValue(t.Name(), fieldName, value)
	}
	return value
}

func (d *Dummie) getStrategy(t reflect.Type, fieldName string) constant.GenerateStrategy {
	if strategy, ok := d.Configuration.Strategy.FieldStrategies[fieldName]; ok {
		return strategy
	}
	if strategy, ok := d.Configuration.Strategy.TypeStrategies[t.Name()]; ok {
		return strategy
	}
	return d.Configuration.Strategy.GlobalStrategy
}

func getFieldName(elems []string) string {
	if len(elems) == 0 {
		return ""
	}
	for _, elem := range elems {
		if elem != "[]" {
			return elem
		}
	}
	return ""
}
