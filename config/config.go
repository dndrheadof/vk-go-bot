package config

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/spf13/viper"
)

type Config struct {
	Token   string `mapstructure:"token"`
	GroupID int    `mapstructure:"group_id"`
}

const configPath = "config/config.yml"

func NewConfig() (Config, error) {
	cfg := Config{}

	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return cfg, fmt.Errorf("ошибка чтения конфига: %w", err)
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("ошибка конвертации конфига: %w", err)
	}

	err = ValidateConfig(cfg)
	if err != nil {
		return cfg, err
	}

	return cfg, nil
}

func ValidateConfig(s interface{}) (err error) {
	structType := reflect.TypeOf(s)
	if structType.Kind() != reflect.Struct {
		return errors.New("конфиг должен иметь тип struct")
	}

	structVal := reflect.ValueOf(s)
	fieldNum := structVal.NumField()

	for i := 0; i < fieldNum; i++ {
		field := structVal.Field(i)
		fieldName := structType.Field(i).Name
		isSet := field.IsValid() && !field.IsZero()

		if !isSet {
			err = errors.Join(fmt.Errorf("переменная %s не установлена в %s", fieldName, configPath), err)
		}

	}

	return err
}
