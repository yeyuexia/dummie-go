package generator

import (
	"math/rand"
	"reflect"

	"github.com/yeyuexia/dummie-go/constant"
)

type BoolGenerator struct {
	T reflect.Type
}

func (g *BoolGenerator) GetType() reflect.Type {
	return g.T
}

func (g *BoolGenerator) Generate(strategy constant.GenerateStrategy, _ string) any {
	switch strategy {
	case constant.GenerateStrategy_Random:
		return rand.Int()&1 == 0
	case constant.GenerateStrategy_Default:
		return true
	default:
		return true
	}
}
