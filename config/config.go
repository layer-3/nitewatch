package config

import (
	"fmt"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// Load populates the given struct with values from environment variables
// based on struct tags. It supports env, env-default, and env-required tags.
func Load(cfg interface{}) error {
	v := reflect.ValueOf(cfg)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("cfg must be a pointer to a struct")
	}

	return loadStruct(v.Elem())
}

func loadStruct(v reflect.Value) error {
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldValue := v.Field(i)

		if !fieldValue.CanSet() {
			continue
		}

		// Handle nested structs
		if fieldValue.Kind() == reflect.Struct {
			if err := loadStruct(fieldValue); err != nil {
				return err
			}
			continue
		}

		envTag := field.Tag.Get("env")
		defaultTag := field.Tag.Get("env-default")
		requiredTag := field.Tag.Get("env-required")

		var value string
		if envTag != "" {
			value = os.Getenv(envTag)
		}

		if value == "" && defaultTag != "" {
			value = defaultTag
		}

		if value == "" && requiredTag == "true" {
			return fmt.Errorf("required environment variable %s is not set", envTag)
		}

		if value == "" {
			continue
		}

		if err := setFieldValue(fieldValue, value); err != nil {
			return fmt.Errorf("error setting field %s: %w", field.Name, err)
		}
	}

	return nil
}

func setFieldValue(field reflect.Value, value string) error {
	fieldType := field.Type()

	// Handle special types first
	switch fieldType {
	case reflect.TypeOf(time.Duration(0)):
		duration, err := time.ParseDuration(value)
		if err != nil {
			return fmt.Errorf("invalid duration: %w", err)
		}
		field.Set(reflect.ValueOf(duration))
		return nil

	case reflect.TypeOf(&time.Location{}):
		location, err := time.LoadLocation(value)
		if err != nil {
			return fmt.Errorf("invalid location: %w", err)
		}
		field.Set(reflect.ValueOf(location))
		return nil

	case reflect.TypeOf(url.URL{}):
		parsedURL, err := url.Parse(value)
		if err != nil {
			return fmt.Errorf("invalid URL: %w", err)
		}
		field.Set(reflect.ValueOf(*parsedURL))
		return nil
	}

	// Handle basic types
	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Bool:
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		field.SetBool(boolValue)
	case reflect.Int:
		intValue, err := strconv.ParseInt(value, 10, 0)
		if err != nil {
			return err
		}
		field.SetInt(intValue)
	case reflect.Int8:
		intValue, err := strconv.ParseInt(value, 10, 8)
		if err != nil {
			return err
		}
		field.SetInt(intValue)
	case reflect.Int16:
		intValue, err := strconv.ParseInt(value, 10, 16)
		if err != nil {
			return err
		}
		field.SetInt(intValue)
	case reflect.Int32:
		intValue, err := strconv.ParseInt(value, 10, 32)
		if err != nil {
			return err
		}
		field.SetInt(intValue)
	case reflect.Int64:
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(intValue)
	case reflect.Uint:
		uintValue, err := strconv.ParseUint(value, 10, 0)
		if err != nil {
			return err
		}
		field.SetUint(uintValue)
	case reflect.Uint8:
		uintValue, err := strconv.ParseUint(value, 10, 8)
		if err != nil {
			return err
		}
		field.SetUint(uintValue)
	case reflect.Uint16:
		uintValue, err := strconv.ParseUint(value, 10, 16)
		if err != nil {
			return err
		}
		field.SetUint(uintValue)
	case reflect.Uint32:
		uintValue, err := strconv.ParseUint(value, 10, 32)
		if err != nil {
			return err
		}
		field.SetUint(uintValue)
	case reflect.Uint64:
		uintValue, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return err
		}
		field.SetUint(uintValue)
	case reflect.Float32:
		floatValue, err := strconv.ParseFloat(value, 32)
		if err != nil {
			return err
		}
		field.SetFloat(floatValue)
	case reflect.Float64:
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return err
		}
		field.SetFloat(floatValue)
	case reflect.Slice:
		if field.Type().Elem().Kind() == reflect.String {
			field.Set(reflect.ValueOf(strings.Split(value, ",")))
		} else {
			return fmt.Errorf("unsupported slice type: %s", field.Type())
		}
	default:
		return fmt.Errorf("unsupported field type: %s", field.Kind())
	}
	return nil
}
