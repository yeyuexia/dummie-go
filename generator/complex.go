package generator

import (
	"math/rand"
	"reflect"

	"github.com/yeyuexia/dummie-go/constant"
)

type Complex64Generator struct {
	T reflect.Type
}

func (g *Complex64Generator) GetType() reflect.Type {
	return g.T
}

func (g *Complex64Generator) Generate(strategy constant.GenerateStrategy, _ string) any {
	switch strategy {
	case constant.GenerateStrategy_Random:
		return complex(rand.Float32(), rand.Float32())
	case constant.GenerateStrategy_Default:
		return complex(1, 1)
	default:
		return complex(1, 1)
	}
}

type Complex128Generator struct {
	T reflect.Type
}

func (g *Complex128Generator) GetType() reflect.Type {
	return g.T
}

func (g *Complex128Generator) Generate(strategy constant.GenerateStrategy, _ string) any {
	switch strategy {
	case constant.GenerateStrategy_Random:
		return complex(rand.Float64(), rand.Float64())
	case constant.GenerateStrategy_Default:
		return complex(float64(1), float64(1))
	default:
		return complex(float64(1), float64(1))
	}
}
