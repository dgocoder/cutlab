package fstransform

import (
	"reflect"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
)

const (
	tagName      = "json"
	tagOmitEmpty = "omitempty"
	delimiter    = ","
)

type FirestoreMap map[string]interface{}

func ToFirestoreMap(value interface{}) FirestoreMap {
	var result = parseData(value)
	return result.(FirestoreMap)
}

func isZeroOfUnderlyingType(x interface{}) bool {
	return reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func parseData(value interface{}) interface{} {
	if value == nil {
		return nil
	}
	var firestoreMap = FirestoreMap{}
	var tag string

	var fieldCount int
	var val = reflect.ValueOf(value)

	switch v := value.(type) {
	case time.Time, *time.Time:
		return v
	case uuid.UUID:
		return v.String()
	}
	switch val.Kind() {
	case reflect.Map:
		for _, key := range val.MapKeys() {
			val := val.MapIndex(key)
			firestoreMap[key.String()] = parseData(val.Interface())
		}
		return firestoreMap
	case reflect.Ptr:
		if val.IsNil() {
			return nil
		}
		fieldCount = val.Elem().NumField()
		for i := 0; i < fieldCount; i++ {
			tag = val.Elem().Type().Field(i).Tag.Get(tagName)
			value = val.Elem().Field(i).Interface()
			setValue(firestoreMap, tag, value)
		}
		return firestoreMap
	case reflect.Struct, reflect.Interface:
		fieldCount = val.NumField()
		for i := 0; i < fieldCount; i++ {
			tag = val.Type().Field(i).Tag.Get(tagName)
			value = val.Field(i).Interface()
			setValue(firestoreMap, tag, value)
		}
		return firestoreMap
	}
	return value
}

func setValue(firestoreMap FirestoreMap, tag string, value interface{}) {
	if tag == "" || tag == "-" || value == nil {
		return
	}

	tagValues := strings.Split(tag, delimiter)
	if strings.Contains(tag, tagOmitEmpty) {
		if isZeroOfUnderlyingType(value) {
			return
		}
	}
	firestoreMap[tagValues[0]] = parseData(value)
}

func DataTo(input, output interface{}) error {
	config := &mapstructure.DecoderConfig{
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			stringToUUIDHookFunc(),
		),
		Result: &output,
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func stringToUUIDHookFunc() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(uuid.UUID{}) {
			return data, nil
		}

		return uuid.Parse(data.(string))
	}
}
