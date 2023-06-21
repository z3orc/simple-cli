package cli

import (
	"fmt"
	"reflect"
)

func validateCommand(command any) error {
	structType := reflect.TypeOf(command)
	if structType.Kind() != reflect.Struct {
		return fmt.Errorf("input command is not a struct")
	}

	commandStruct := reflect.ValueOf(command)
	fieldCount := commandStruct.NumField()

	nameField := commandStruct.Field(0)
	if nameField.IsZero() {
		return fmt.Errorf("command defined without a name")
	}

	for i := 0; i < fieldCount; i++ {
		field := commandStruct.Field(i)
		fieldName := structType.Field(i).Name

		if fieldName == "Help" ||
			fieldName == "Options" ||
			fieldName == "Debug" ||
			fieldName == "Usage" ||
			fieldName == "Arguments" {
			continue
		}

		if !field.IsValid() || field.IsZero() {
			return fmt.Errorf("required field: '%s' of command: '%s' was not properly defined", fieldName, nameField.String())
		}
	}

	return nil
}

func validateOption(option Option) error {
	structType := reflect.TypeOf(option)
	if structType.Kind() != reflect.Struct {
		return fmt.Errorf("input command is not a struct")
	}

	commandStruct := reflect.ValueOf(option)
	fieldCount := commandStruct.NumField()

	nameField := commandStruct.Field(0)
	if nameField.IsZero() {
		return fmt.Errorf("option defined without name")
	}

	for i := 0; i < fieldCount; i++ {
		field := commandStruct.Field(i)
		fieldName := structType.Field(i).Name

		if fieldName == "Usage" ||
			fieldName == "Arguments" {
			continue
		}

		if !field.IsValid() || field.IsZero() {
			return fmt.Errorf("required field: '%s' of option: '%s' was not properly defined", fieldName, nameField.String())
		}
	}

	return nil
}
