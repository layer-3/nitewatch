package config

import (
	"net/url"
	"os"
	"reflect"
	"testing"
	"time"
)

type TestAppConfig struct {
	Number    int64  `env:"TEST_NUMBER" env-default:"1"`
	String    string `env:"TEST_STRING" env-default:"default"`
	NoDefault string `env:"TEST_NO_DEFAULT"`
	NoEnv     string `env-default:"default"`
	Required  string `env:"TEST_REQUIRED" env-required:"true"`
}

type NestedServerConfig struct {
	Host string `env:"SERVER_HOST" env-default:"localhost"`
	Port int    `env:"SERVER_PORT" env-default:"8080"`
}

type NestedDatabaseConfig struct {
	URL      string `env:"DATABASE_URL" env-required:"true"`
	MaxConns int    `env:"DATABASE_MAX_CONNS" env-default:"10"`
}

type TestNestedConfig struct {
	AppName  string `env:"APP_NAME" env-default:"TestApp"`
	Server   NestedServerConfig
	Database NestedDatabaseConfig
}

func LoadTestAppConfig() (*TestAppConfig, error) {
	cfg := &TestAppConfig{}
	if err := Load(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func TestLoad(t *testing.T) {
	tests := []struct {
		name    string
		setup   func()
		cleanup func()
		want    *TestAppConfig
		wantErr bool
	}{
		{
			name: "load with environment variables",
			setup: func() {
				os.Setenv("TEST_NUMBER", "42")
				os.Setenv("TEST_STRING", "hello")
				os.Setenv("TEST_NO_DEFAULT", "no default value")
				os.Setenv("TEST_REQUIRED", "required value")
			},
			cleanup: func() {
				os.Unsetenv("TEST_NUMBER")
				os.Unsetenv("TEST_STRING")
				os.Unsetenv("TEST_NO_DEFAULT")
				os.Unsetenv("TEST_REQUIRED")
			},
			want: &TestAppConfig{
				Number:    42,
				String:    "hello",
				NoDefault: "no default value",
				NoEnv:     "default",
				Required:  "required value",
			},
			wantErr: false,
		},
		{
			name: "load with default values",
			setup: func() {
				os.Setenv("TEST_REQUIRED", "required value")
			},
			cleanup: func() {
				os.Unsetenv("TEST_REQUIRED")
			},
			want: &TestAppConfig{
				Number:    1,
				String:    "default",
				NoDefault: "",
				NoEnv:     "default",
				Required:  "required value",
			},
			wantErr: false,
		},
		{
			name:    "missing required field",
			setup:   func() {},
			cleanup: func() {},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			defer tt.cleanup()

			cfg, err := LoadTestAppConfig()
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadTestAppConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && cfg != nil {
				if cfg.Number != tt.want.Number {
					t.Errorf("Number = %v, want %v", cfg.Number, tt.want.Number)
				}
				if cfg.String != tt.want.String {
					t.Errorf("String = %v, want %v", cfg.String, tt.want.String)
				}
				if cfg.NoDefault != tt.want.NoDefault {
					t.Errorf("NoDefault = %v, want %v", cfg.NoDefault, tt.want.NoDefault)
				}
				if cfg.NoEnv != tt.want.NoEnv {
					t.Errorf("NoEnv = %v, want %v", cfg.NoEnv, tt.want.NoEnv)
				}
				if cfg.Required != tt.want.Required {
					t.Errorf("Required = %v, want %v", cfg.Required, tt.want.Required)
				}
			}
		})
	}
}

func TestSetFieldValue(t *testing.T) {
	type TestStruct struct {
		StringField   string
		BoolField     bool
		IntField      int
		Int8Field     int8
		Int16Field    int16
		Int32Field    int32
		Int64Field    int64
		UintField     uint
		Uint8Field    uint8
		Uint16Field   uint16
		Uint32Field   uint32
		Uint64Field   uint64
		Float32Field  float32
		Float64Field  float64
		SliceField    []string
		DurationField time.Duration
		LocationField *time.Location
		URLField      url.URL
	}

	tests := []struct {
		name      string
		fieldName string
		value     string
		expected  interface{}
	}{
		{"string field", "StringField", "test", "test"},
		{"bool field true", "BoolField", "true", true},
		{"bool field false", "BoolField", "false", false},
		{"int field", "IntField", "123", 123},
		{"int8 field", "Int8Field", "127", int8(127)},
		{"int16 field", "Int16Field", "32767", int16(32767)},
		{"int32 field", "Int32Field", "2147483647", int32(2147483647)},
		{"int64 field", "Int64Field", "456", int64(456)},
		{"uint field", "UintField", "789", uint(789)},
		{"uint8 field", "Uint8Field", "255", uint8(255)},
		{"uint16 field", "Uint16Field", "65535", uint16(65535)},
		{"uint32 field", "Uint32Field", "4294967295", uint32(4294967295)},
		{"uint64 field", "Uint64Field", "18446744073709551615", uint64(18446744073709551615)},
		{"float32 field", "Float32Field", "3.14", float32(3.14)},
		{"float64 field", "Float64Field", "3.14159", 3.14159},
		{"slice field", "SliceField", "a,b,c", []string{"a", "b", "c"}},
		{"duration field", "DurationField", "1h30m", time.Hour + 30*time.Minute},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &TestStruct{}
			v := reflect.ValueOf(cfg).Elem()
			field := v.FieldByName(tt.fieldName)

			err := setFieldValue(field, tt.value)
			if err != nil {
				t.Errorf("setFieldValue() error = %v", err)
				return
			}

			got := field.Interface()
			switch expected := tt.expected.(type) {
			case []string:
				gotSlice := got.([]string)
				if len(gotSlice) != len(expected) {
					t.Errorf("slice length = %v, want %v", len(gotSlice), len(expected))
				}
				for i := range expected {
					if gotSlice[i] != expected[i] {
						t.Errorf("slice[%d] = %v, want %v", i, gotSlice[i], expected[i])
					}
				}
			default:
				if got != tt.expected {
					t.Errorf("field value = %v, want %v", got, tt.expected)
				}
			}
		})
	}

	// Test special types separately
	t.Run("location field", func(t *testing.T) {
		cfg := &TestStruct{}
		v := reflect.ValueOf(cfg).Elem()
		field := v.FieldByName("LocationField")

		err := setFieldValue(field, "America/New_York")
		if err != nil {
			t.Errorf("setFieldValue() error = %v", err)
			return
		}

		got := field.Interface().(*time.Location)
		if got.String() != "America/New_York" {
			t.Errorf("location = %v, want America/New_York", got.String())
		}
	})

	t.Run("url field", func(t *testing.T) {
		cfg := &TestStruct{}
		v := reflect.ValueOf(cfg).Elem()
		field := v.FieldByName("URLField")

		err := setFieldValue(field, "https://example.com/path?query=value")
		if err != nil {
			t.Errorf("setFieldValue() error = %v", err)
			return
		}

		got := field.Interface().(url.URL)
		if got.String() != "https://example.com/path?query=value" {
			t.Errorf("url = %v, want https://example.com/path?query=value", got.String())
		}
	})
}

func TestLoadNestedConfig(t *testing.T) {
	tests := []struct {
		name    string
		setup   func()
		cleanup func()
		want    *TestNestedConfig
		wantErr bool
	}{
		{
			name: "load nested config with all values",
			setup: func() {
				os.Setenv("APP_NAME", "MyApp")
				os.Setenv("SERVER_HOST", "0.0.0.0")
				os.Setenv("SERVER_PORT", "3000")
				os.Setenv("DATABASE_URL", "postgres://localhost/mydb")
				os.Setenv("DATABASE_MAX_CONNS", "25")
			},
			cleanup: func() {
				os.Unsetenv("APP_NAME")
				os.Unsetenv("SERVER_HOST")
				os.Unsetenv("SERVER_PORT")
				os.Unsetenv("DATABASE_URL")
				os.Unsetenv("DATABASE_MAX_CONNS")
			},
			want: &TestNestedConfig{
				AppName: "MyApp",
				Server: NestedServerConfig{
					Host: "0.0.0.0",
					Port: 3000,
				},
				Database: NestedDatabaseConfig{
					URL:      "postgres://localhost/mydb",
					MaxConns: 25,
				},
			},
			wantErr: false,
		},
		{
			name: "load nested config with defaults",
			setup: func() {
				os.Setenv("DATABASE_URL", "sqlite://test.db")
			},
			cleanup: func() {
				os.Unsetenv("DATABASE_URL")
			},
			want: &TestNestedConfig{
				AppName: "TestApp",
				Server: NestedServerConfig{
					Host: "localhost",
					Port: 8080,
				},
				Database: NestedDatabaseConfig{
					URL:      "sqlite://test.db",
					MaxConns: 10,
				},
			},
			wantErr: false,
		},
		{
			name: "missing required nested field",
			setup: func() {
				// DATABASE_URL is required but not set
			},
			cleanup: func() {},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			defer tt.cleanup()

			cfg := &TestNestedConfig{}
			err := Load(cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && cfg != nil {
				if cfg.AppName != tt.want.AppName {
					t.Errorf("AppName = %v, want %v", cfg.AppName, tt.want.AppName)
				}
				if cfg.Server.Host != tt.want.Server.Host {
					t.Errorf("Server.Host = %v, want %v", cfg.Server.Host, tt.want.Server.Host)
				}
				if cfg.Server.Port != tt.want.Server.Port {
					t.Errorf("Server.Port = %v, want %v", cfg.Server.Port, tt.want.Server.Port)
				}
				if cfg.Database.URL != tt.want.Database.URL {
					t.Errorf("Database.URL = %v, want %v", cfg.Database.URL, tt.want.Database.URL)
				}
				if cfg.Database.MaxConns != tt.want.Database.MaxConns {
					t.Errorf("Database.MaxConns = %v, want %v", cfg.Database.MaxConns, tt.want.Database.MaxConns)
				}
			}
		})
	}
}
