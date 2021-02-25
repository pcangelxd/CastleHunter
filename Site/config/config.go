package config

import (
	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

type Config struct {
	Address		string		`toml:"address"`
	Email		string		`toml:"email"`
	Password	string		`toml:"password"`
	Host		string		`toml:"host"`
	Port		string		`toml:"port"`
	Cost7Day	string		`toml:"cost7Day"`
	Cost14Day	string		`toml:"cost14Day"`
	Cost30Day	string		`toml:"cost30Day"`
	BankDetails	string		`toml:"bankDetails"`
}

// Создание новой конфигурации
func NewConfig(configPath string) *Config {
	cfg := &Config{
	}

	_, err := toml.DecodeFile(configPath, cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg
}

// Создание новой конфигурации логгера
func NewConfigLogger() *zap.Logger {
	cfg := zap.Config{
		Encoding:         "console",
		Level:            zap.NewAtomicLevelAt(zapcore.DebugLevel),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey: "msg",

			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalLevelEncoder,

			TimeKey:    "time",
			EncodeTime: zapcore.RFC3339TimeEncoder,

			CallerKey:    "caller",
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}
	logger, _ := cfg.Build()

	defer logger.Sync()

	return logger
}