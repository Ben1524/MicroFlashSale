package conf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config 主配置结构
type Config struct {
	MySQLConnections []MySQLConnection `json:"mysql_connections"`
	RedisConnections []RedisConnection `json:"redis_connections"`
	MicroServices    []MicroService    `json:"micro_services"`
	LogConfig        LogConfig         `json:"log_config"`
	ConsulConfig     ConsulConfig      `json:"cosult_config"`
}

// MySQLConnection MySQL连接配置
type MySQLConnection struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Username string `json:"username"`
	Database string `json:"database"`
}

// RedisConnection Redis连接配置
type RedisConnection struct {
	Name        string `json:"name"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Password    string `json:"password"`
	DB          int    `json:"db"`
	PoolSize    int    `json:"pool_size"`
	PoolTimeout int    `json:"pool_timeout"`
}

// MicroService 微服务配置
type MicroService struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
	Version  string `json:"version"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level  string `json:"level"`
	Output string `json:"output"`
	Format string `json:"format"`
}

// ConsulConfig Consul配置
type ConsulConfig struct {
	Host       string `json:"host"`
	Port       int    `json:"port"`
	Datacenter string `json:"datacenter"`
}

var CONFIG *Config

func init() {
	// 读取配置文件
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Printf("读取配置文件失败: %v\n", err)
		return
	}

	err = json.Unmarshal(data, &CONFIG)
	if err != nil {
		fmt.Printf("解析配置文件失败: %v\n", err)
		return
	}
	//
	// 打印解析结果
	fmt.Println("MySQL连接配置:")
	for _, conn := range CONFIG.MySQLConnections {
		fmt.Printf("  名称: %s, 主机: %s:%d, 用户: %s, 数据库: %s\n",
			conn.Name, conn.Host, conn.Port, conn.Username, conn.Database)
	}

	fmt.Println("\nRedis连接配置:")
	for _, conn := range CONFIG.RedisConnections {
		fmt.Printf("  名称: %s, 主机: %s:%d, DB: %d, 连接池大小: %d\n",
			conn.Name, conn.Host, conn.Port, conn.DB, conn.PoolSize)
	}

	fmt.Println("\n微服务配置:")
	for _, srv := range CONFIG.MicroServices {
		fmt.Printf("  名称: %s, 地址: %s:%d, 协议: %s\n",
			srv.Name, srv.Host, srv.Port, srv.Protocol)
	}

	fmt.Println("\n日志配置:")
	fmt.Printf("  级别: %s, 输出: %s, 格式: %s\n",
		CONFIG.LogConfig.Level, CONFIG.LogConfig.Output, CONFIG.LogConfig.Format)

	fmt.Println("\nConsul配置:")
	fmt.Printf("  主机: %s:%d, 数据中心: %s\n",
		CONFIG.ConsulConfig.Host, CONFIG.ConsulConfig.Port, CONFIG.ConsulConfig.Datacenter)
}
