// @author xiangqian
// @date 2025/08/02 12:55
package main

import (
	pkg_ini "gopkg.in/ini.v1"
	"gweb/pkg/db"
	"gweb/pkg/minio"
	"gweb/pkg/redis"
	"time"
)

func LoadConfig() (Config, error) {
	var config Config
	file, err := pkg_ini.Load("config.ini")
	if err != nil {
		return config, err
	}

	// 数据库配置
	section, err := file.GetSection("db")
	if err != nil {
		return config, err
	}
	config.Db = db.Config{
		Driver:     section.Key("driver").String(),
		DataSource: section.Key("data-source").String(),
	}

	// Redis 配置
	section, err = file.GetSection("redis")
	if err != nil {
		return config, err
	}
	config.Redis = redis.Config{
		Host:   section.Key("redis-host").String(),
		Port:   uint16(section.Key("port").MustUint()),
		Passwd: section.Key("passwd").String(),
		Db:     uint8(section.Key("db").MustUint()),
	}

	// MinIO 配置
	section, err = file.GetSection("minio")
	if err != nil {
		return config, err
	}
	config.Minio = minio.Config{
		Endpoint:  section.Key("endpoint").String(),
		AccessKey: section.Key("access-key").String(),
		SecretKey: section.Key("secret-key").String(),
	}

	// JWT 配置
	section, err = file.GetSection("jwt")
	if err != nil {
		return config, err
	}
	config.Jwt = Jwt{
		Key:        section.Key("key").String(),
		ExpireTime: section.Key("expire-time").MustDuration(),
	}

	// HTTP 配置
	section, err = file.GetSection("http")
	if err != nil {
		return config, err
	}
	config.Http = Http{
		Port: uint16(section.Key("port").MustUint()),
	}

	return config, nil
}

// Config 配置
type Config struct {
	Db    db.Config    // 数据库配置
	Redis redis.Config // Redis 配置
	Minio minio.Config // MinIO 配置
	Jwt   Jwt          // JWT 配置
	Http  Http         // HTTP 配置
}

// Jwt JWT 配置
type Jwt struct {
	Key        string        // 密钥
	ExpireTime time.Duration // 过期时间
}

// Http HTTP 配置
type Http struct {
	Port uint16 // 监听端口
}
