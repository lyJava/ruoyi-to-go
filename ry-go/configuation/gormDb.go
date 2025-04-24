package configuation

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"ry-go/utils"

	"github.com/fatih/color"
	"github.com/natefinch/lumberjack"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// customLogger 自定义日志记录器
type customLogger struct {
	logger.Interface
	ConsoleOut                io.Writer     // 控制台输出
	FileOut                   io.Writer     // 文件输出
	IgnoreRecordNotFoundError bool          // 是否忽略记录未找到错误
	Colorful                  bool          // 是否启用彩色输出
	SlowThreshold             time.Duration // 慢sql阈值
	ZapLogger                 *zap.Logger   // Zap 日志记录器
}

// LogMode 实现日志输出的 Format 方法
//
//goland:noinspection GoUnusedParameter
func (cusLog *customLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *cusLog
	return &newLogger
}

func (cusLog *customLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	cusLog.log(ctx, "INFO", msg, data...)
}

func (cusLog *customLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	cusLog.log(ctx, "WARN", msg, data...)
}

func (cusLog *customLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	// 检查是否为记录未找到错误
	for _, arg := range data {
		if err, ok := arg.(error); ok {
			if errors.Is(err, gorm.ErrRecordNotFound) && cusLog.IgnoreRecordNotFoundError {
				// 忽略记录未找到错误
				return
			}
		}
	}
	// 记录其他错误
	cusLog.log(ctx, "ERROR", msg, data...)
}

//goland:noinspection GoUnusedParameter
func (cusLog *customLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	// 获取 SQL 查询语句和执行的行数
	execSql, rows := fc()
	// 获取文件名和行号
	_, file, line, ok := runtime.Caller(2)
	if ok {
		// 计算查询的执行时间
		duration := time.Since(begin)
		// 判断是否为慢查询
		if duration > cusLog.SlowThreshold {
			cusLog.log(ctx, "WARN", fmt.Sprintf("SLOW SQL >= %v\n%s:%d\n%v\n[rows:%d]\nSQL: %s\r\n[elapsed: %v]", cusLog.SlowThreshold, file, line, duration, rows, execSql, duration))
		} else {
			cusLog.log(ctx, "DEBUG", fmt.Sprintf("%s:%d\n%v\n[rows:%d]\nSQL: %s\r\n[elapsed: %v]", file, line, duration, rows, execSql, duration))
		}
	} else {
		cusLog.log(ctx, "TRACE", fmt.Sprintf("%s:%d\ngorm获取执行信息失败\n[rows:%d]\nSQL: %s\r\n", file, line, rows, execSql))
	}
}

//goland:noinspection GoUnhandledErrorResult
func (cusLog *customLogger) log(ctx context.Context, level string, format string, args ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	message := fmt.Sprintf(format, args...)

	requestId := utils.GetValForContext(ctx, "requestId")

	// 根据日志级别设置颜色
	var logMessage string
	// 控制台输出
	if cusLog.ConsoleOut != nil {
		if cusLog.Colorful {
			switch level {
			case "INFO":
				logMessage = color.GreenString("%s [request_id:%s] [%s] %s\n", timestamp, requestId, level, message)
			case "WARN":
				logMessage = color.YellowString("%s [request_id:%s] [%s] %s\n", timestamp, requestId, level, message)
			case "ERROR":
				logMessage = color.HiRedString("%s [request_id:%s] [%s] %s\n", timestamp, requestId, level, message)
			case "TRACE":
				logMessage = color.HiBlueString("%s [request_id:%s] [%s] %s\n", timestamp, requestId, level, message)
			case "DEBUG":
				logMessage = color.HiMagentaString("%s [request_id:%s] [%s] %s\n", timestamp, requestId, level, message)
			default:
				logMessage = fmt.Sprintf("%s [request_id:%s] [%s] %s\n", timestamp, requestId, level, message)
			}
		} else {
			logMessage = fmt.Sprintf("%s [request_id:%s] [%s] %s\n", timestamp, requestId, level, message)
		}
		cusLog.ConsoleOut.Write([]byte(logMessage))
	}

	if cusLog.FileOut != nil {
		cusLog.FileOut.Write([]byte(fmt.Sprintf("%s [request_id:%s] [%s] %s\n", timestamp, requestId, level, message)))
	}

	// 设置zap日志
	sprintZapLog := fmt.Sprintf("\n%s [request_id:%s] %s %s\n", timestamp, requestId, level, message)
	switch level {
	case "INFO":
		cusLog.ZapLogger.Info(sprintZapLog)
	case "WARN":
		cusLog.ZapLogger.Warn(sprintZapLog)
	case "ERROR":
		cusLog.ZapLogger.Error(sprintZapLog)
	case "DEBUG":
		cusLog.ZapLogger.Debug(sprintZapLog)
	case "TRACE":
		cusLog.ZapLogger.Debug(sprintZapLog)
	default:
		cusLog.ZapLogger.Info(sprintZapLog)
	}
}

// initLogger 初始化日志配置
func initLogger(logName, logPath string) *lumberjack.Logger {
	logDir := filepath.Join(logPath)
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Printf("创建日志目录错误===%s", fmt.Errorf("创建日志目录失败: %w", err))
	}

	logFileName := filepath.Join(logDir, logName)

	// 配置日志切割
	return &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    10,
		MaxAge:     30,
		MaxBackups: 30,
		Compress:   false,
		LocalTime:  true,
	}
}

// InitGormPostgresConn gorm初始化postgresql连接
//
//goland:noinspection GoUnhandledErrorResult
func InitGormPostgresConn() (*gorm.DB, error) {

	lumberjackLogger := initLogger("gorm_pg_sql_zero.log", "logs/postgres")

	// 创建自定义的 logger 实例
	customGormLogger := &customLogger{
		ConsoleOut:                os.Stdout,
		FileOut:                   lumberjackLogger,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
		SlowThreshold:             3 * time.Second, // 设置慢sql阈值为3秒
		ZapLogger:                 ZapLog,          // 增加zap日志记录
	}

	postgresqlItem := PostgresConfig.Postgresql

	// 构建连接
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		postgresqlItem.Host, postgresqlItem.Port, postgresqlItem.User, postgresqlItem.Password, postgresqlItem.DbName)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		CreateBatchSize: 2000,
		PrepareStmt:     true,
		Logger:          customGormLogger,
	})

	if err != nil {
		zap.L().Sugar().Errorf("Failed to connect to Postgresql database %+v", err)
		return nil, err
	}
	zap.L().Sugar().Infof("Gorm Create Postgresql Database connection successful")

	sqlDb, err := db.DB()
	if err != nil {
		zap.L().Sugar().Errorf("获取 postgres 底层连接失败: %+v", err)
		return nil, err
	}
	setConnPoolProps("Postgresql", sqlDb)

	if err = sqlDb.Ping(); err != nil {
		zap.L().Sugar().Errorf("postgres 连接健康检查失败: %+v", err)
		return nil, err
	}

	return db, nil
}

func setConnPoolProps(dbType string, db *sql.DB) {
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(50)
	db.SetConnMaxLifetime(time.Hour)
	zap.L().Sugar().Info("Gorm Database Connection Pool Setting success")
}
