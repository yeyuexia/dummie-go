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
			GlobalStrategy:  constant.GenerateStrategy_Default,
			FieldStrategies: make(map[string]constant.GenerateStrategy),
			TypeStrategies:  make(map[string]constant.GenerateStrategy),
		},
		DefaultValues: make(map[string]map[string]any),
	}
}

func (c *Configuration) Override(fieldName string, value any) *Configuration {
	targetType := reflect.TypeOf(value).Name()
	if _, ok := c.DefaultValues[targetType]; !ok {
		c.DefaultValues[targetType] = map[string]any{}
	}
	c.DefaultValues[targetType][fieldName] = value
	return c
}

func (c *Configuration) OverrideType(target, value any) *Configuration {
	targetType := reflect.TypeOf(target).Name()
	if _, ok := c.DefaultValues[targetType]; !ok {
		c.DefaultValues[targetType] = map[string]any{}
	}
	c.DefaultValues[targetType][""] = value
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

func (c *Configuration) Random(fieldNames ...string) *Configuration {
	for _, fieldName := range fieldNames {
		c.GenerateStrategy(fieldName, constant.GenerateStrategy_Default)
	}
	return c
}

func (c *Configuration) RandomType(fieldType any) *Configuration {
	c.Strategy.TypeStrategies[reflect.TypeOf(fieldType).Name()] = constant.GenerateStrategy_Default
	return c
}
