// @author xiangqian
// @date 2025/07/20 13:09
package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const KeepTTL = -1

var cli *redis.Client

var ctx context.Context

func Init(config Config) (*redis.Client, string, error) {
	// 创建 Redis 客户端配置
	cli = redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%d", config.Host, config.Port), // 服务器地址
		Password:     config.Passwd,                                  // 密码
		DB:           int(config.Db),                                 // 数据库
		PoolSize:     5,                                              // 连接池大小
		MinIdleConns: 1,                                              // 最小空闲连接
		MaxRetries:   3,                                              // 最大重试次数
		DialTimeout:  5 * time.Second,                                // 连接超时
	})

	// 创建上下文
	ctx = context.Background()

	// 测试连接
	pong, err := cli.Ping(ctx).Result()
	return cli, pong, err
}

// Redis 数据类型之 String（字符串）
// 存储文本、数字（整数或浮点数，并支持数值进行加法、减法等操作）、二进制数据，最大可以存储 512MB 的内容。
// 使用场景：缓存、短信验证码、计数器、分布式 session。

// Set 设置键值
func Set(key string, value string, exp time.Duration) error {
	return cli.Set(ctx, key, value, exp).Err()
}

// Get 获取键值
func Get(key string) (string, error) {
	return cli.Get(ctx, key).Result()
}

// Redis 数据类型之 List（列表）
// 有序的字符串列表，支持头部和尾部的插入、删除操作，可以实现队列、栈等数据结构。
// 使用场景：发布订阅等

// Redis 数据类型之 Set（集合）
// 无序的字符串集合，元素不重复，支持集合间的交集、并集、差集等操作。
// 使用场景：共同好友、点赞或点踩等

// SAdd 向集合中添加一个或多个元素
func SAdd(key string, values ...any) error {
	return cli.SAdd(ctx, key, values...).Err()
}

// SGet 随机获取集合中的一个元素
func SGet(key string) (string, error) {
	return cli.SRandMember(ctx, key).Result()
}

// SGets 获取集合中的所有元素
func SGets(key string) ([]string, error) {
	return cli.SMembers(ctx, key).Result()
}

// SCount 获取 Set 元素个数
func SCount(key string) (int64, error) {
	return cli.SCard(ctx, key).Result()
}

// SHas 检查集合是否包含指定元素
func SHas(key string, value any) (bool, error) {
	return cli.SIsMember(ctx, key, value).Result()
}

// SDel 从集合中移除一个或多个元素
func SDel(key string, values ...any) error {
	return cli.SRem(ctx, key, values...).Err()
}

// Redis 数据类型之 Hash（哈希）
// 存储键值对的集合，适合存储对象的属性信息，支持对单个字段的读写操作。
// 使用场景：存储对象

// Expire 设置 key 过期时间
func Expire(key string, exp time.Duration) error {
	return cli.Expire(ctx, key, exp).Err()
}

// Del 删除 key
func Del(keys ...string) error {
	return cli.Del(ctx, keys...).Err()
}

type Config struct {
	Host   string
	Port   uint16
	Passwd string
	Db     uint8
}
