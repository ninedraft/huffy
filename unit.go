package huffy

import (
	"fmt"
	"reflect"
)

func unit(u interface{}) func(args ...interface{}) bool {
	var unit = reflect.ValueOf(u)
	var unitT = reflect.TypeOf(u)
	if unitT.Kind() != reflect.Func {
		var pm = fmt.Sprintf("[huffy.unit] expected to get func(a A, [b A]*) bool, got %q", unitT)
		panic(pm)
	}
	if unitT.NumOut() != 1 || unitT.Out(0).Kind() != reflect.Bool {
		var pm = fmt.Sprintf("[huffy.unit] expected to get func(a A, [b A]*) bool, got %q", unitT)
		panic(pm)
	}
	var argsN = unitT.NumIn()
	var argsTypes = make([]reflect.Type, 0, argsN)
	for i := range r(argsN) {
		argsTypes = append(argsTypes, unitT.In(i))
	}
	var validSpec = argsSpec(argsTypes)
	return func(args ...interface{}) bool {
		var reflectedArgs = reflectSlice(args)
		if len(args) != argsN {
			var invalidInput = typesOfValues(reflectedArgs)
			panic(fmt.Sprintf("expected to get %d args %v, got %v", argsN, validSpec, argsSpec(invalidInput)))
		}
		var ret = unit.Call(reflectedArgs)
		return ret[0].Bool()
	}
}

func reflectSlice(args []interface{}) []reflect.Value {
	var reflected = make([]reflect.Value, 0, len(args))
	for _, arg := range args {
		reflected = append(reflected, reflect.ValueOf(arg))
	}
	return reflected
}

func argsSpec(args []reflect.Type) []string {
	var spec = make([]string, 0, len(args))
	for _, arg := range args {
		spec = append(spec, arg.String())
	}
	return spec
}

func typesOfValues(values []reflect.Value) []reflect.Type {
	var types = make([]reflect.Type, 0, len(values))
	for _, value := range values {
		types = append(types, value.Type())
	}
	return types
}
