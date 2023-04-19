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
	return &Configuration{
		Strategy: Strategy{
			GlobalStrategy:  constant.Default,
			FieldStrategies: make(map[string]constant.GenerateStrategy),
			TypeStrategies:  make(map[string]constant.GenerateStrategy),
		},
		DefaultValues: make(map[string]map[string]any),
	}
}

func (c *Configuration) Override(target, value any) *Configuration {
	targetType := reflect.ValueOf(target).Type().Name()
	if _, ok := c.DefaultValues[targetType]; !ok {
		c.DefaultValues[targetType] = map[string]any{}
	}
	c.DefaultValues[targetType][""] = value
	return c
}

func (c *Configuration) OverrideWithFieldName(fieldName string, value any) *Configuration {
	targetType := reflect.ValueOf(value).Type().Name()
	if _, ok := c.DefaultValues[targetType]; !ok {
		c.DefaultValues[targetType] = map[string]any{}
	}
	c.DefaultValues[targetType][fieldName] = value
	return c
}

func (c *Configuration) GlobalGenerateStrategy(strategy constant.GenerateStrategy) *Configuration {
	c.Strategy.GlobalStrategy = strategy
	return c
}

func (c *Configuration) GenerateStrategy(fieldName string, strategy constant.GenerateStrategy) *Configuration {
	c.Strategy.FieldStrategies[fieldName] = strategy
	return c
}
