package dummie

import (
	"reflect"

	"github.com/yeyuexia/dummie-go/constant"
	"github.com/yeyuexia/dummie-go/generator"
)

type Generator interface {
	GetType() reflect.Type
	Generate(Strategy constant.GenerateStrategy, fieldName string) any
}

type GeneratorManager struct {
	Generators map[reflect.Kind]Generator
}

func NewGenerators() *GeneratorManager {
	return &GeneratorManager{
		Generators: map[reflect.Kind]Generator{
			reflect.Bool:       &generator.BoolGenerator{T: reflect.TypeOf(false)},
			reflect.Int:        &generator.IntGenerator{T: reflect.TypeOf(1)},
			reflect.Int8:       &generator.Int8Generator{T: reflect.TypeOf(int8(1))},
			reflect.Int16:      &generator.Int16Generator{T: reflect.TypeOf(int16(1))},
			reflect.Int32:      &generator.Int32Generator{T: reflect.TypeOf(int32(1))},
			reflect.Int64:      &generator.Int64Generator{T: reflect.TypeOf(int64(1))},
			reflect.Uint:       &generator.UintGenerator{T: reflect.TypeOf(uint(1))},
			reflect.Uint8:      &generator.Uint8Generator{T: reflect.TypeOf(uint8(1))},
			reflect.Uint16:     &generator.Uint16Generator{T: reflect.TypeOf(uint16(1))},
			reflect.Uint32:     &generator.Uint32Generator{T: reflect.TypeOf(uint32(1))},
			reflect.Uint64:     &generator.Uint64Generator{T: reflect.TypeOf(uint64(1))},
			reflect.Float32:    &generator.Float32Generator{T: reflect.TypeOf(float32(1))},
			reflect.Float64:    &generator.Float64Generator{T: reflect.TypeOf(float64(1))},
			reflect.Complex64:  &generator.Complex64Generator{T: reflect.TypeOf(complex(1, 1))},
			reflect.Complex128: &generator.Complex128Generator{T: reflect.TypeOf(complex(float64(1), float64(1)))},
			reflect.String:     &generator.StringGenerator{T: reflect.TypeOf("")},
		},
	}
}

func (g *GeneratorManager) GenerateValue(t reflect.Type, strategy constant.GenerateStrategy, path []string) any {
	if generator, ok := g.Generators[t.Kind()]; ok {
		if len(path) == 0 {
			return generator.Generate(strategy, ".")
		}
		return generator.Generate(strategy, path[0])
	}
	return nil
}
