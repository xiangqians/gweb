// @author xiangqian
// @date 20:47 2023/06/10
package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"time"
)

// 数据库连接池
var db *sqlx.DB

// Init 初始化数据库连接池
func Init(config Config) (*sqlx.DB, error) {
	var err error

	// 进行连接验证，sql.Open 本身不会建立连接
	// 因为每一个连接都是惰性创建的，如何验证 sql.Open 调用之后，sql.DB对象可用呢？
	// 通常使用sql.Ping()函数初始化，调用完毕后会马上把连接返回给连接池。
	db, err = sqlx.Connect(config.Driver, config.DataSource)
	if err != nil {
		return nil, err
	}

	// 在 Golang 中，database/sql 包已经集成了连接池的功能

	// sql.DB 连接池是如何工作的呢？
	// sql.DB 连接池包含两种类型的连接：正在使用连接 和 空闲连接。
	// 当使用 "空闲" 连接执行数据库操作（例如 Exec，Query）时，该连接被标记为 "正在使用"，操作完成后，该连接被标记为 "空闲"。

	// 设置池中 "打开" 连接（"正在使用" 连接和 "空闲" 连接）数量的上限。默认情况下，打开的连接数是无限的。"打开" 连接 包括 "正在使用" 连接和 "空闲" 连接，不仅仅是 "正在使用" 连接。
	// 一般来说，MaxOpenConns 设置得越大，可以并发执行的数据库查询就越多，连接池本身成为应用程序中的瓶颈的风险就越低。
	// 但让它无限并不是最好的选择。默认情况下，PostgreSQL 最多 100 个打开连接的硬限制，如果达到这个限制的话，它将导致 PostgreSQL 驱动返回 "sorry, too many clients already" 错误。最大打开连接数限制可以在 postgresql.conf 文件中对 max_connections 设置来更改。
	// 为了避免这个错误，将池中打开的连接数量限制在 100 以下是有意义的，可以为其他需要使用 PostgreSQL 的应用程序或会话留下足够的空间。
	// 设置 MaxOpenConns 限制的另一个好处是，它充当一个非常基本的限流器，防止数据库同时被大量任务压垮。
	// 如果达到 MaxOpenConns 限制，并且所有连接都在使用中，那么任何新的数据库任务将被迫等待，直到有连接空闲。
	// 在我们的 API 上下文中，用户的 HTTP 请求可能在等待空闲连接时无限期地 "挂起"。因此，为了缓解这种情况，使用上下文为数据库任务设置超时是很重要的。
	db.SetMaxOpenConns(2)
	//db.SetMaxOpenConns(25)

	// 设置池中 "空闲" 连接数的上限。缺省情况下，最大空闲连接数为 2。
	// 理论上，在池中允许更多的空闲连接将增加性能。因为它减少了从头建立新连接发生概率，因此有助于节省资源。
	// 但要意识到保持空闲连接是有代价的。它占用了本来可以用于应用程序和数据库的内存，而且如果一个连接空闲时间过长，它也可能变得不可用。例如，默认情况下 MySQL 会自动关闭任何 8 小时未使用的连接。
	// 因此，与使用更小的空闲连接池相比，将 MaxIdleConns 设置得过高可能会导致更多的连接变得不可用，浪费资源。因此保持适量的空闲连接是必要的。理想情况下，你只希望保持一个连接空闲，可以快速使用。
	// 另一件要指出的事情是 MaxIdleConns 值应该总是小于或等于 MaxOpenConns。Go会强制保证这点，并在必要时自动减少 MaxIdleConns 值。
	db.SetMaxIdleConns(2)
	//db.SetMaxIdleConns(25)

	// 设置一个连接保持可用的最长时间。默认连接的存活时间没有限制，永久可用。
	// 如果设置 ConnMaxLifetime 的值为 1 小时，意味着所有的连接在创建后，经过一个小时就会被标记为失效连接，标志后就不可复用。
	// 但需要注意：
	// 1、这并不能保证一个连接将在池中存在一整个小时；有可能某个连接由于某种原因变得不可用，并在此之前自动关闭。
	// 2、连接在创建后一个多小时内仍然可以被使用，只是在这个时间之后它不能被重用。
	// 3、这不是一个空闲超时。连接将在创建后一小时过期，而不是在空闲后一小时过期。
	// 4、Go每秒运行一次后台清理操作，从池中删除过期的连接。
	// 理论上，ConnMaxLifetime 为无限大（或设置为很长生命周期）将提升性能，因为这样可以减少新建连接。但是在某些情况下，设置短期存活时间有用。比如：
	// 1、如果 SQL 数据库对连接强制设置最大存活时间，这时将 ConnMaxLifetime 设置稍短时间更合理。
	// 2、有助于数据库替换（优雅地交换数据库）
	// 如果决定对连接池设置 ConnMaxLifetime，那么一定要记住连接过期（然后重新创建）的频率。例如，如果连接池中有 100 个打开的连接，而 ConnMaxLifetime 为 1 分钟，那么应用程序平均每秒可以杀死并重新创建多达 1.67 个连接。频率太大而最终影响性能。
	db.SetConnMaxLifetime(5 * time.Minute)

	// SetConnMaxIdleTime() 函数在 Go 1.15 版本引入对 ConnMaxIdleTime 进行配置。
	// 其效果和 ConnMaxLifeTime 类似，但这里设置的是：在被标记为失效之前一个连接最长空闲时间。
	// 例如，如果我们将 ConnMaxIdleTime 设置为 1 小时，那么自上次使用以后在池中空闲了 1 小时的任何连接都将被标记为过期并被后台清理操作删除。
	// 这个配置非常有用，因为它意味着我们可以对池中空闲连接的数量设置相对较高的限制，但可以通过删除不再真正使用的空闲连接来周期性地释放资源。
	db.SetConnMaxIdleTime(5 * time.Minute)

	return db, nil
}

// Add 新增
func Add(sql string, args ...any) (rowsAffected int64, insertId int64, err error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		return
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return
	}

	insertId, err = result.LastInsertId()
	return
}

// Upd 更新
func Upd(sql string, args ...any) (rowsAffected int64, err error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		return
	}

	rowsAffected, err = result.RowsAffected()
	return
}

// Get 查询
func Get(dest any, sql string, args ...any) error {
	return db.Get(dest, sql, args...)
}

// List 查询列表
func List(dest any, sql string, args ...any) error {
	return db.Select(dest, sql, args...)
}

// Del 删除
func Del(sql string, args ...any) (rowsAffected int64, err error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		return
	}

	rowsAffected, err = result.RowsAffected()
	return
}

// Begin 开启事务
func Begin() (*Tx, error) {
	// sql.Begin() 调用完毕后将连接传递给 sql.Tx 类型对象
	// 当 Commit() 或 Rollback() 函数调用后释放连接
	tx, err := db.Beginx()
	if err != nil {
		return nil, err
	}
	return &Tx{tx: tx}, nil
}

type Tx struct {
	tx *sqlx.Tx
}

// Add 新增
func (tx *Tx) Add(sql string, args ...any) (rowsAffected int64, insertId int64, err error) {
	result, err := tx.tx.Exec(sql, args...)
	if err != nil {
		return
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		return
	}

	insertId, err = result.LastInsertId()
	return
}

// Upd 更新
func (tx *Tx) Upd(sql string, args ...any) (rowsAffected int64, err error) {
	result, err := tx.tx.Exec(sql, args...)
	if err != nil {
		return
	}

	rowsAffected, err = result.RowsAffected()
	return
}

// Get 查询
func (tx *Tx) Get(dest any, sql string, args ...any) error {
	return tx.tx.Get(dest, sql, args...)
}

// List 查询列表
func (tx *Tx) List(dest any, sql string, args ...any) error {
	return tx.tx.Select(dest, sql, args...)
}

// Del 删除
func (tx *Tx) Del(sql string, args ...any) (rowsAffected int64, err error) {
	result, err := tx.tx.Exec(sql, args...)
	if err != nil {
		return
	}

	rowsAffected, err = result.RowsAffected()
	return
}

// Commit 提交事务
func (tx *Tx) Commit() error {
	return tx.tx.Commit()
}

// Rollback 回滚事务
func (tx *Tx) Rollback() error {
	return tx.tx.Rollback()
}

// Stats 数据库状态
func Stats() string {
	stats := db.Stats()
	statsString := fmt.Sprint("\n\t最大连接数\t\t:", stats.MaxOpenConnections,
		"\n\t连接池状态",
		"\n\t\t当前连接数\t:", stats.OpenConnections, `（"正在使用"连接和"空闲"连接）`,
		"\n\t\t正在使用连接数\t:", stats.InUse,
		"\n\t\t空闲连接数\t:", stats.Idle,
		"\n\t统计",
		"\n\t\t等待连接数\t:", stats.WaitCount,
		"\n\t\t等待创建新连接时长（秒）:", stats.WaitDuration.Seconds(),
		"\n\t\t空闲超限关闭数\t:", stats.MaxIdleClosed, "（达到MaxIdleConns而关闭的连接数量）",
		"\n\t\t空闲超时关闭数\t:", stats.MaxIdleTimeClosed,
		"\n\t\t连接超时关闭数\t:", stats.MaxLifetimeClosed, "（达到ConnMaxLifetime而关闭的连接数量）")
	return statsString
}

type Config struct {
	Driver     string // 驱动名
	DataSource string // 数据源
}
