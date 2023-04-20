package generator

import (
	"math"
	"math/rand"
	"reflect"

	"github.com/yeyuexia/dummie-go/constant"
)

type IntGenerator struct {
	T reflect.Type
}

func (g *IntGenerator) GetType() reflect.Type {
	return g.T
}

func (g *IntGenerator) Generate(strategy constant.GenerateStrategy, _ string) any {
	switch strategy {
	case constant.GenerateStrategy_Random:
		return rand.Int()
	case constant.GenerateStrategy_Default:
		return 1
	default:
		return 1
	}
}

type Int8Generator struct {
	T reflect.Type
}

func (g *Int8Generator) GetType() reflect.Type {
	return g.T
}

func (g *Int8Generator) Generate(strategy constant.GenerateStrategy, _ string) any {
	switch strategy {
	case constant.GenerateStrategy_Random:
		return int8(rand.Intn(math.MaxInt8-math.MinInt8) + math.MinInt8)
	case constant.GenerateStrategy_Default:
		return 1
	default:
		return 1
	}
}

type Int16Generator struct {
	T reflect.Type
}

func (g *Int16Generator) GetType() reflect.Type {
	return g.T
}

func (g *Int16Generator) Generate(strategy constant.GenerateStrategy, _ string) any {
	switch strategy {
	case constant.GenerateStrategy_Random:
		return int16(rand.Intn(math.MaxInt16-math.MinInt16) + math.MinInt16)
	case constant.GenerateStrategy_Default:
		return 1
	default:
		return 1
	}
}

type Int32Generator struct {
	T reflect.Type
}

func (g *Int32Generator) GetType() reflect.Type {
	return g.T
}

func (g *Int32Generator) Generate(strategy constant.GenerateStrategy, _ string) any {
	switch strategy {
	case constant.GenerateStrategy_Random:
		return rand.Int31()
	case constant.GenerateStrategy_Default:
		return 1
	default:
		return 1
	}
}

type Int64Generator struct {
	T reflect.Type
}

func (g *Int64Generator) GetType() reflect.Type {
	return g.T
}

func (g *Int64Generator) Generate(strategy constant.GenerateStrategy, _ string) any {
	switch strategy {
	case constant.GenerateStrategy_Random:
		return rand.Int63()
	case constant.GenerateStrategy_Default:
		return 1
	default:
		return 1
	}
}
