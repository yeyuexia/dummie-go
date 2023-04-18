package dummie

import (
	"reflect"
	"strings"

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
	pathString := getPathString(path)
	value := d.Cache.GetValue(t, pathString)
	if value != nil {
		return value
	}
	value = d.GeneratorManager.GenerateValue(t, d.getStrategy(t, pathString), path)
	if value != nil {
		d.Cache.SetValue(t.Name(), pathString, value)
	}
	return value
}

func (d *Dummie) getStrategy(t reflect.Type, path string) constant.GenerateStrategy {
	if strategy, ok := d.Configuration.Strategy.FieldStrategies[path]; ok {
		return strategy
	}
	if strategy, ok := d.Configuration.Strategy.TypeStrategies[t.Name()]; ok {
		return strategy
	}
	return d.Configuration.Strategy.GlobalStrategy
}

func getPathString(elems []string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return elems[0]
	}

	n := len(elems) - 1
	for _, elem := range elems {
		n += len(elem)
	}

	var b strings.Builder
	b.Grow(n)
	for i := len(elems) - 1; i > 0; i-- {
		b.WriteString(elems[i])
		b.WriteString(".")
	}
	b.WriteString(elems[0])
	return b.String()
}
