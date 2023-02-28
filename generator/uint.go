package generator

import (
	"math"
	"math/rand"
	"reflect"

	"github.com/yeyuexia/dummie-go/constant"
)

type UintGenerator struct {
	T reflect.Type
}

func (g *UintGenerator) GetType() reflect.Type {
	return g.T
}

func (g *UintGenerator) Generate(strategy constant.GenerateStrategy, _ string) any {
	switch strategy {
	case constant.Random:
		return uint(rand.Int31())
	case constant.Default:
		return uint(1)
	default:
		return uint(1)
	}
}

type Uint8Generator struct {
	T reflect.Type
}

func (g *Uint8Generator) GetType() reflect.Type {
	return g.T
}

func (g *Uint8Generator) Generate(strategy constant.GenerateStrategy, _ string) any {
	switch strategy {
	case constant.Random:
		return uint8(rand.Uint32() % math.MaxUint8)
	case constant.Default:
		return uint8(1)
	default:
		return uint8(1)
	}
}

type Uint16Generator struct {
	T reflect.Type
}

func (g *Uint16Generator) GetType() reflect.Type {
	return g.T
}

func (g *Uint16Generator) Generate(strategy constant.GenerateStrategy, _ string) any {
	switch strategy {
	case constant.Random:
		return uint16(rand.Uint32() % math.MaxUint16)
	case constant.Default:
		return uint16(1)
	default:
		return uint16(1)
	}
}

type Uint32Generator struct {
	T reflect.Type
}

func (g *Uint32Generator) GetType() reflect.Type {
	return g.T
}

func (g *Uint32Generator) Generate(strategy constant.GenerateStrategy, _ string) any {
	switch strategy {
	case constant.Random:
		return rand.Uint32()
	case constant.Default:
		return uint32(1)
	default:
		return uint32(1)
	}
}

type Uint64Generator struct {
	T reflect.Type
}

func (g *Uint64Generator) GetType() reflect.Type {
	return g.T
}

func (g *Uint64Generator) Generate(strategy constant.GenerateStrategy, _ string) any {
	switch strategy {
	case constant.Random:
		return rand.Uint64()
	case constant.Default:
		return uint64(1)
	default:
		return uint64(1)
	}
}
