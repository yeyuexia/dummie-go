package generator

import (
	"math/rand"
	"reflect"

	"github.com/yeyuexia/dummie-go/constant"
)

type Float32Generator struct {
	T reflect.Type
}

func (g *Float32Generator) GetType() reflect.Type {
	return g.T
}

func (g *Float32Generator) Generate(strategy constant.GenerateStrategy, _ string) any {
	switch strategy {
	case constant.GenerateStrategy_Random:
		return rand.Float32()
	case constant.GenerateStrategy_Default:
		return 1.0
	default:
		return 1.0
	}
}

type Float64Generator struct {
	T reflect.Type
}

func (g *Float64Generator) GetType() reflect.Type {
	return g.T
}

func (g *Float64Generator) Generate(strategy constant.GenerateStrategy, _ string) any {
	switch strategy {
	case constant.GenerateStrategy_Random:
		return rand.Float64()
	case constant.GenerateStrategy_Default:
		return 1.0
	default:
		return 1.0
	}
}
