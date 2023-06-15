# dummie-go [![<Dummie GO>](https://circleci.com/gh/yeyuexia/dummie-go.svg?style=svg)](<LINK>)
### Overview
Image that you need to write test with some complex input parameters, and most of the fields in your test don't necessary and only need to fill some valid value. so you need to spend lot of time to initial the fake object but even you finish that, your test codes still hard to understand because of the given value has so many fields and it's hard to know which fields you really care about. In those situtation you may need **Dummie**.
**Dummie** is a simple data generator for preparing test data, it provides an easy way to construct complex data class as below:

```
instance := TargetClass{}
err := dummie.Inflate(&instance)
if err != nil {
    t.Fatalf("Dummie inflate data error: %+v.", err.Error())
}
```

### Usage

* Add Dependency
```groovy
go get github.com/yeyuexia/dummie-go&<latestVersion>
```

##### Basic usage

When you writing test code, you may want to create multi complicted data like:

```
type Employee struct {
	Id int64
	Name string
	Mobile string
	Email string
	Assets []Asset
    ...
}

type Asset struct {
	Id int64
	Name string
	Desc string
	...
}
```

The only thing you want do is create an Employee instance with fullfilled datas and deliver to other method. With **Dummie**, you can use simple one line code to get what your want:

```
employee := Employee{}
err := dummie.Inflate(&employee)
if err != nil {
    t.Fatalf("Dummie inflate data error: %+v.", err.Error())
}
```

Maybe you want some fields special in the dummy Employee.

##### Override type / field values
* Suppose in next test, we want all name field has same value, we can use `Override` method:

```
configuration := dummie.NewConfiguration().Override("Name", "value")
err := InflateWithConfiguration(&data, configuration)
if err != nil {
    t.Fatalf("Dummie inflate data error: %+v.", err.Error())
}
```

so all fields with name **Name** and **string type** should be filled to value **"value"**, dummie determine type based on `Override` second parameter type.

* If you want all fields with the type `int64` to be set 100, you can use `OverrideType`:

```
configuration := dummie.NewConfiguration().OverrideType(int64(1), 100)
err := InflateWithConfiguration(&data, configuration)
if err != nil {
    t.Fatalf("Dummie inflate data error: %+v.", err.Error())
}
```

##### Random fields
* Suppose we want set random values for all **Id** and **Desc** fields, we can use `Random` method:

```
configuration := dummie.NewConfiguration().Random("Id", "Desc")
err := InflateWithConfiguration(&data, configuration)
if err != nil {
    t.Fatalf("Dummie inflate data error: %+v.", err.Error())
}
```

so all fields named **Id** and **Desc** would be filled by random value.

* Like `Override`, random can also special a kind of type by `RandomType`:

```
configuration := dummie.NewConfiguration().RandomType("Id", "Desc")
err := InflateWithConfiguration(&data, configuration)
if err != nil {
    t.Fatalf("Dummie inflate data error: %+v.", err.Error())
}
```

##### Change generate strategy

**Dummie** uses static value for data generate by default. If you want all field values use random value generator, you should set GenerationStrategy:
```
configuration := dummie.NewConfiguration().GlobalGenerateStrategy(constant.GenerateStrategy_Random)
err := InflateWithConfiguration(&data, configuration)
if err != nil {
    t.Fatalf("Dummie inflate data error: %+v.", err.Error())
}
```

## Dummie for other language
java: https://github.com/dummie-java/dummie
