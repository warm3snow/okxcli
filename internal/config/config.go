package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var (
	config *Config
	once   sync.Once
)

// Init load and parse config
func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.cexcli")
	viper.AddConfigPath("/etc/cexcli")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	config = &Config{}
	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	if err := config.ValidateConfig(); err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	once.Do(func() {
		if config == nil {
			Init()
		}
	})
	return config
}

// CEXConfig CEX 配置结构
type CEXConfig struct {
	BaseURL string       `yaml:"base_url" mapstructure:"base_url"`
	API     CEXAPIConfig `yaml:"api" mapstructure:"api"`
}

type CEXAPIConfig struct {
	APIKey      string `yaml:"api_key" mapstructure:"api_key"`
	SecretKey   string `yaml:"secret_key" mapstructure:"secret_key"`
	Passphrase  string `yaml:"passphrase" mapstructure:"passphrase"`
	IsSimulated bool   `yaml:"is_simulated" mapstructure:"is_simulated"`
}

type FeishuConfig struct {
	WebhookURL string `yaml:"webhook_url" mapstructure:"webhook_url"`
}

type NotificationsConfig struct {
	EnableMacOS  bool   `yaml:"enable_macos" mapstructure:"enable_macos"`
	EnableFeishu bool   `yaml:"enable_feishu" mapstructure:"enable_feishu"`
	AppName      string `yaml:"app_name" mapstructure:"app_name"`
}

type SchedulerConfig struct {
	Interval      int `yaml:"interval" mapstructure:"interval"` // 统一的任务间隔 (秒)
	RetryInterval int `yaml:"retry_interval" mapstructure:"retry_interval"`
	MaxRetries    int `yaml:"max_retries" mapstructure:"max_retries"`
}

type LoggingConfig struct {
	Level  string `yaml:"level" mapstructure:"level"`
	Format string `yaml:"format" mapstructure:"format"`
}

type Config struct {
	CEX           CEXConfig           `yaml:"cex" mapstructure:"cex"`
	Feishu        FeishuConfig        `yaml:"feishu" mapstructure:"feishu"`
	Notifications NotificationsConfig `yaml:"notifications" mapstructure:"notifications"`
	Scheduler     SchedulerConfig     `yaml:"scheduler" mapstructure:"scheduler"`
	Logging       LoggingConfig       `yaml:"logging" mapstructure:"logging"`
	Traders       map[string]string   `yaml:"traders" mapstructure:"traders"` // 交易员配置
}

// validateConfig 验证配置
func (config Config) ValidateConfig() error {

	// 检查通知器配置
	if !config.Notifications.EnableMacOS && !config.Notifications.EnableFeishu {
		return fmt.Errorf("no notifiers enabled, please enable at least one notifier")
	}

	// 检查飞书配置
	if config.Notifications.EnableFeishu && config.Feishu.WebhookURL == "" {
		return fmt.Errorf("feishu notifications enabled but webhook URL not configured")
	}

	return nil
}
