package inject

import (
	"reflect"
)

type module struct {
	bindings map[bindingKey]binding
}

func createModule() *module {
	return &module{make(map[bindingKey]binding)}
}

func (this *module) Bind(from interface{}) Builder {
	fromReflectType := reflect.TypeOf(from)
	if fromReflectType == nil {
		return newPropogatedErrorBuilder(newErrorBuilder(InjectErrorTypeNil).build())
	}
	return newBuilder(this, fromReflectType)
}

func (this *module) BindTagged(from interface{}, tag string) Builder {
	fromReflectType := reflect.TypeOf(from)
	if fromReflectType == nil {
		return newPropogatedErrorBuilder(newErrorBuilder(InjectErrorTypeNil).build())
	}
	if tag == "" {
		return newPropogatedErrorBuilder(newErrorBuilder(InjectErrorTypeTagEmpty).build())
	}
	return newTaggedBuilder(this, fromReflectType, tag)
}

func (this *module) assignBinding(bindingKey bindingKey, binding binding) error {
	this.bindings[bindingKey] = binding
}