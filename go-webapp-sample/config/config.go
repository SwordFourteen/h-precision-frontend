package config

import (
	"embed"
	"flag"
	"fmt"
	"os"

	"github.com/ybkuroki/go-webapp-sample/util"
	"gopkg.in/yaml.v3"
)

// Config represents the composition of yml settings.
// Config 结构体用于表示从 yml 文件中加载的配置信息，
// 包含数据库、Redis、日志、静态内容、安全性等多方面的设置。
type Config struct {
	Database struct {
		Dialect   string `default:"sqlite3"` // 数据库方言，默认为 sqlite3
		Host      string `default:"book.db"` // 数据库文件名或主机名，默认为 book.db
		Port      string // 数据库端口号
		Dbname    string // 数据库名称
		Username  string // 数据库用户名
		Password  string // 数据库密码
		Migration bool   `default:"false"` // 是否启用数据库迁移，默认为 false
	}
	Redis struct {
		Enabled            bool   `default:"false"`                          // 是否启用 Redis，默认为 false
		ConnectionPoolSize int    `yaml:"connection_pool_size" default:"10"` // Redis 连接池大小，默认为 10
		Host               string // Redis 主机名
		Port               string // Redis 端口号
	}
	Extension struct {
		MasterGenerator bool `yaml:"master_generator" default:"false"` // 是否启用主生成器功能
		CorsEnabled     bool `yaml:"cors_enabled" default:"false"`     // 是否启用 CORS（跨域资源共享）
		SecurityEnabled bool `yaml:"security_enabled" default:"false"` // 是否启用安全功能
	}
	Log struct {
		RequestLogFormat string `yaml:"request_log_format" default:"${remote_ip} ${account_name} ${uri} ${method} ${status}"` // 请求日志的格式
	}
	StaticContents struct {
		Enabled bool `default:"false"` // 是否启用静态内容服务
	}
	Swagger struct {
		Enabled bool   `default:"false"` // 是否启用 Swagger 文档
		Path    string // Swagger 文档的路径
	}
	Security struct {
		AuthPath    []string `yaml:"auth_path"`    // 需要认证的路径
		ExculdePath []string `yaml:"exclude_path"` // 排除认证的路径
		UserPath    []string `yaml:"user_path"`    // 用户权限路径
		AdminPath   []string `yaml:"admin_path"`   // 管理员权限路径
	}
}

const (
	// DEV represents development environment
	// 表示开发环境
	DEV = "develop"
	// PRD represents production environment
	// 表示生产环境
	PRD = "production"
	// DOC represents docker container
	// 表示 Docker 容器环境
	DOC = "docker"
)

// LoadAppConfig reads the settings written to the yml file
// LoadAppConfig 从嵌入式的 yml 文件中读取应用程序的配置，并根据环境变量选择不同的配置文件。
func LoadAppConfig(yamlFile embed.FS) (*Config, string) {
	var env *string

	// 从环境变量 WEB_APP_ENV 中读取当前环境，如果环境变量未设置，默认使用 "develop"
	if value := os.Getenv("WEB_APP_ENV"); value != "" {
		env = &value
	} else {
		// 如果未设置环境变量，则从命令行参数中读取 "env" 参数，默认为 "develop"
		env = flag.String("env", "develop", "To switch configurations.")
		flag.Parse() // 解析命令行参数
	}

	// 根据当前环境加载相应的 yml 配置文件
	file, err := yamlFile.ReadFile(fmt.Sprintf(AppConfigPath, *env))
	if err != nil {
		// 如果读取文件失败，输出错误信息并退出程序
		fmt.Printf("Failed to read application.%s.yml: %s", *env, err)
		os.Exit(ErrExitStatus)
	}

	config := &Config{}
	// 解析 yml 文件内容到 Config 结构体中
	if err := yaml.Unmarshal(file, config); err != nil {
		// 如果解析失败，输出错误信息并退出程序
		fmt.Printf("Failed to read application.%s.yml: %s", *env, err)
		os.Exit(ErrExitStatus)
	}

	// 返回解析后的配置和当前环境
	return config, *env
}

// LoadMessagesConfig loads the messages.properties.
// LoadMessagesConfig 从嵌入式的 properties 文件中加载消息配置。
func LoadMessagesConfig(propsFile embed.FS) map[string]string {
	// 读取并解析 messages.properties 文件
	messages := util.ReadPropertiesFile(propsFile, MessagesConfigPath)
	if messages == nil {
		// 如果加载失败，输出错误信息并退出程序
		fmt.Printf("Failed to load the messages.properties.")
		os.Exit(ErrExitStatus)
	}
	// 返回加载的消息配置
	return messages
}
