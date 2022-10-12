package flags

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func Parse(args []string, t interface{}, toParse ...string) error {
	for _, variable := range toParse {
		field, ok := reflect.TypeOf(t).Elem().FieldByName(variable)
		if !ok {
			continue
		}
		tag := field.Tag.Get("argparse")
		val := reflect.ValueOf(t).Elem().FieldByName(variable)
		for index, arg := range args {
			if strings.Contains(tag, arg) {
				if val.CanSet() {
					kind := val.Kind()
					switch kind {
					default:
						err := fmt.Errorf("Unexpected type %v", kind)
						return err // %T prints whatever type t has
					case reflect.Bool:
						conv, err := strconv.ParseBool(args[index+1])
						if err != nil {
							return err
						}
						val.SetBool(conv)
					case reflect.Int:
						conv, err := strconv.Atoi(args[index+1])
						if err != nil {
							return err
						}
						val.SetInt(int64(conv))
					case reflect.String:
						val.SetString(args[index+1])
					case reflect.Struct:
						obj, err := parseJson(args[index+1])
						if err != nil {
							return err
						}
						val.Set(reflect.ValueOf(obj))
					}
				}
			}
		}

	}
	return nil
}

func parseJson(jsonString string) (interface{}, error) {
	info, err := os.Stat(jsonString)
	var data []byte
	if err != nil {
		data, err = os.ReadFile(info.Name())
		if err != nil {
			return nil, err
		}
	} else {
		data = []byte(jsonString)
	}

	if err != nil {
		return nil, err
	}
	val := &struct{}{}
	err = json.Unmarshal(data, val)
	return val, err

}

func ParseArgs(t interface{}, toParse ...string) error {
	return Parse(os.Args, t, toParse...)
}
