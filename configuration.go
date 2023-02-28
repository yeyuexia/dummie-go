package dummie

import (
	"reflect"

	"github.com/yeyuexia/dummie-go/constant"
)

type Strategy struct {
	GlobalStrategy  constant.GenerateStrategy
	FieldStrategies map[string]constant.GenerateStrategy
	TypeStrategies  map[string]constant.GenerateStrategy
}

type Configuration struct {
	Strategy      Strategy
	DefaultValues map[string]map[string]any
}

func NewConfiguration() *Configuration {
	return &Configuration{}
}

func (c *Configuration) Override(target, value any) *Configuration {
	targetType := reflect.ValueOf(target).Type().Name()
	if _, ok := c.DefaultValues[targetType]; !ok {
		c.DefaultValues[targetType] = map[string]any{}
	}
	c.DefaultValues[targetType][""] = value
	return c
}

func (c *Configuration) OverrideWithPath(path string, value any) *Configuration {
	targetType := reflect.ValueOf(value).Type().Name()
	if _, ok := c.DefaultValues[targetType]; !ok {
		c.DefaultValues[targetType] = map[string]any{}
	}
	c.DefaultValues[targetType][path] = value
	return c
}

func (c *Configuration) GlobalGenerateStrategy(strategy constant.GenerateStrategy) *Configuration {
	c.Strategy.GlobalStrategy = strategy
	return c
}

func (c *Configuration) GenerateStrategy(path string, strategy constant.GenerateStrategy) *Configuration {
	c.Strategy.FieldStrategies[path] = strategy
	return c
}
