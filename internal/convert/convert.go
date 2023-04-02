package convert

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

func Marshal(v any) ([][]string, error) {
	sliceVal := reflect.ValueOf(v)
	if sliceVal.Kind() != reflect.Slice {
		return nil, errors.New("must be a slice of structs")
	}
	structType := sliceVal.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return nil, errors.New("must be a slice of structs")
	}
	var out [][]string
	header := marshalHeader(structType)
	out = append(out, header)
	for i := 0; i < sliceVal.Len(); i++ {
		row, err := marshalOne(sliceVal.Index(i))
		if err != nil {
			return nil, err
		}
		out = append(out, row)
	}
	return out, nil
}

func marshalHeader(vt reflect.Type) []string {
	var row []string
	for i := 0; i < vt.NumField(); i++ {
		field := vt.Field(i)
		if curTag, ok := field.Tag.Lookup("csv"); ok {
			row = append(row, curTag)
		}
	}
	return row
}

func marshalOne(vv reflect.Value) ([]string, error) {
	var row []string
	vt := vv.Type()
	for i := 0; i < vv.NumField(); i++ {
		fieldVal := vv.Field(i)
		if _, ok := vt.Field(i).Tag.Lookup("csv"); !ok {
			continue
		}
		switch fieldVal.Kind() {
		case reflect.String:
			row = append(row, fieldVal.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			row = append(row, strconv.FormatInt(fieldVal.Int(), 10))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			row = append(row, strconv.FormatUint(fieldVal.Uint(), 10))
		case reflect.Float32, reflect.Float64:
			row = append(row, strconv.FormatFloat(fieldVal.Float(), 'f', -1, 64))
		case reflect.Bool:
			row = append(row, strconv.FormatBool(fieldVal.Bool()))
		default:
			return nil, fmt.Errorf("unsupported type: %v", fieldVal.Kind())
		}
	}
	return row, nil
}

func Unmarshal(data [][]string, v any) error {
	sliceValPtr := reflect.ValueOf(v)
	if sliceValPtr.Kind() != reflect.Ptr {
		return errors.New("must be a pointer to a slice of structs")
	}
	sliceVal := sliceValPtr.Elem()
	if sliceVal.Kind() != reflect.Slice {
		return errors.New("must be a pointer to a slice of structs")
	}
	structType := sliceVal.Type().Elem()
	if structType.Kind() != reflect.Struct {
		return errors.New("must be a pointer to a slice of structs")
	}

	header := data[0]
	namePos := make(map[string]int, len(header))
	for i, name := range header {
		namePos[name] = i
	}
	for _, row := range data[1:] {
		newVal := reflect.New(structType).Elem()
		err := unmarshalOne(row, namePos, newVal)
		if err != nil {
			return err
		}
		sliceVal.Set(reflect.Append(sliceVal, newVal))
	}
	return nil
}

func unmarshalOne(row []string, namePos map[string]int, vv reflect.Value) error {
	vt := vv.Type()
	for i := 0; i < vv.NumField(); i++ {
		typeField := vt.Field(i)
		pos, ok := namePos[typeField.Tag.Get("csv")]
		if !ok {
			continue
		}
		val := row[pos]
		field := vv.Field(i)
		switch field.Kind() {
		case reflect.String:
			field.SetString(val)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			n, err := strconv.ParseInt(val, 10, 64)
			if err != nil {
				return err
			}
			field.SetInt(n)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			n, err := strconv.ParseUint(val, 10, 64)
			if err != nil {
				return err
			}
			field.SetUint(n)
		case reflect.Float32, reflect.Float64:
			n, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return err
			}
			field.SetFloat(n)
		case reflect.Bool:
			n, err := strconv.ParseBool(val)
			if err != nil {
				return err
			}
			field.SetBool(n)
		default:
			return fmt.Errorf("unsupported type: %v", field.Kind())
		}
	}
	return nil
}
