package config

import (
	"os"
	"strings"
	"sync"

	"github.com/rs/zerolog"
	"github.com/spf13/viper"
)

type Config struct {
	logMode    zerolog.Level
	rjsApiUrl  string
	rjsApiKey  string
	rjsBalance uint64
	rjsPeriod  uint
	disHook    string
	disBotName string
	disBotAva  string
	disBotMsg  string
}

var (
	cfg  *Config
	once sync.Once
)

func (c Config) ZerologLevel() zerolog.Level {
	return c.logMode
}

func (c Config) RajaSMSApiURL() string {
	return c.rjsApiUrl
}

func (c Config) RajaSMSApiKey() string {
	return c.rjsApiKey
}

func (c Config) RajaSMSLowBalance() uint64 {
	return c.rjsBalance
}

func (c Config) RajaSMSGraceDays() uint {
	return c.rjsPeriod
}

func (c Config) DishookURL() string {
	return c.disHook
}

func (c Config) DishookBotName() string {
	return c.disBotName
}

func (c Config) DishookBotAvatarURL() string {
	return c.disBotAva
}

func (c Config) DishookBotMessage() string {
	return c.disBotMsg
}

func Get() *Config {
	once.Do(func() {
		cfg = read()
	})
	return cfg
}

func read() *Config {
	fang := viper.New()

	fang.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	fang.AutomaticEnv()

	fang.SetConfigName("rajasms-account-monitor")
	fang.SetConfigType("yml")
	fang.AddConfigPath(".")

	value, available := os.LookupEnv("CONFIGDIR_PATH")
	if available {
		fang.AddConfigPath(value)
	}

	_ = fang.ReadInConfig()

	var balance uint64 = uint64(fang.GetUint("rajasms.lowbalance"))
	if balance == 0 {
		balance = 100000
	}

	period := fang.GetUint("rajasms.graceperiod")
	if period == 0 {
		period = 7
	}

	botMsg := fang.GetString("discord.bot.message")
	if len(strings.TrimSpace(botMsg)) == 0 {
		botMsg = "Reminder akun RajaSMS"
	}

	var logmode zerolog.Level
	switch fang.GetString("logmode") {
	case "debug":
		logmode = zerolog.DebugLevel
	case "info":
		logmode = zerolog.InfoLevel
	case "warn":
		logmode = zerolog.WarnLevel
	case "error":
		logmode = zerolog.ErrorLevel
	default:
		logmode = zerolog.Disabled
	}

	return &Config{
		logMode:    logmode,
		rjsApiUrl:  setDefaultString(fang.GetString("rajasms.api.url"), "", true),
		rjsApiKey:  setDefaultString(fang.GetString("rajasms.api.key"), "", true),
		rjsBalance: balance,
		rjsPeriod:  period,
		disHook:    setDefaultString(fang.GetString("discord.webhookurl"), "", true),
		disBotName: setDefaultString(fang.GetString("discord.bot.name"), "AWS Lambda - RajaSMS Monitor", true),
		disBotAva:  setDefaultString(fang.GetString("discord.bot.avatarurl"), "https://adsmedia.co.id/wp-content/uploads/2017/07/cropped-adsmedia-LOGO-100x100-270x270.jpg", true),
		disBotMsg:  botMsg,
	}
}

func setDefaultString(value, fallback string, trimSpace bool) string {
	if trimSpace {
		value = strings.TrimSpace(value)
	}
	if len(value) <= 0 {
		return fallback
	}
	return value
}
