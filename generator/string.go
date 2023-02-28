package generator

import (
	"math/rand"
	"reflect"

	"github.com/yeyuexia/dummie-go/constant"
)

type StringGenerator struct {
	T reflect.Type
}

func (g *StringGenerator) GetType() reflect.Type {
	return g.T
}

func (g *StringGenerator) Generate(strategy constant.GenerateStrategy, path string) any {
	switch strategy {
	case constant.Random:
		return RandStringRunes(18)
	case constant.Default:
		return path
	default:
		return path
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-=`~!@#$%^&*()_+")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
