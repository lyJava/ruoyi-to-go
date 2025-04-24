package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"ry-go/business/dao/daoImpl"
	"ry-go/business/domain"
	"ry-go/configuation"
	"ry-go/utils"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

//goland:noinspection GoUnhandledErrorResult
func main() {
	configuation.GlobalStartTime = time.Now()
	if err := configuation.InitGlobalConfig(); err != nil {
		log.Println("全局配置文件初始化错误")
	}

	pgDb, err := configuation.InitGormPostgresConn()
	if err != nil {
		log.Println("Gorm Postgresql Init Failed===")
	}

	redisClient, _ := configuation.InitRedisClient()

	configuation.GlobalRedisClient = redisClient

	//dictTypeCacheInit(db, redisClient)

	// 命令行启动
	// todo 使用wire后，命名行启动使用 go run wire_gen.go main.go
	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	// 设置为中国时区
	time.Local = location

	redisStore := utils.NewRedisStore(redisClient,
		configuation.CacheConfig.Cache.Captcha.Prefix,
		configuation.CacheConfig.Cache.Captcha.Expiration)

	app, clean, err := WireApp(pgDb, redisClient, redisStore, sync.Map{})
	if err != nil {
		panic(err)
	}
	defer clean()
	defer func(ZapLog *zap.Logger) {
		if err = ZapLog.Sync(); err != nil {
			log.Printf("Zap Log Sync Error ===%+v", err)
		}
	}(configuation.ZapLog)

	// 修复点：在独立协程启动服务器
	serverErr := make(chan error)
	go func() {
		// echoRequestRouter(app)
		port := strconv.Itoa(configuation.ServerConfig.Server.Port)
		log.Printf("echo服务启动端口=====:%s", port)
		serverErr <- app.Start(":" + port)
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err = <-serverErr:
		log.Fatal("服务意外退出")
	case <-quit:
		log.Println("收到关闭信号")

		// 停止所有服务
		stopCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err = app.Shutdown(stopCtx); err != nil {
			log.Println("服务关闭失败")
		}
		log.Println("服务已优雅退出")
	}

	// 阻塞主程序直到收到信号
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Println("Main program exiting")

}

func dictTypeCacheInit(db *gorm.DB, client *redis.Client) {
	// 定义常量提升可维护性
	const (
		dictCacheKeyPrefix   = "sys_dict:"   // 缓存键前缀
		defaultCacheDuration = 1 * time.Hour // 默认缓存时间
		maxConcurrency       = 5             // 最大并发协程数
	)

	zap.L().Sugar().Infof("字典数据类型缓存初始化===")
	ctx := context.Background()

	sysDictDataDaoImpl := daoImpl.NewSysDictDataDaoImpl(db, client)
	sysDictTypeDaoImpl := daoImpl.NewSysDictTypeDaoImpl(db, client, sysDictDataDaoImpl)

	dictTypes, err := sysDictTypeDaoImpl.SelectAll(ctx)
	if err != nil {
		zap.L().Sugar().Errorf("字典数据类型缓存初始化错误===", zap.Error(err))
	}

	// 使用工作池控制并发
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxConcurrency) // 并发控制信号量

	for _, dictType := range dictTypes {

		if dictType == nil {
			continue
		}

		go func(dt *domain.SysDictType) {
			defer func() {
				<-sem
				wg.Done()
			}()

			// 获取字典数据
			dataCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			dataList, err := sysDictDataDaoImpl.SelectDictDataByType(dataCtx, dt.DictType)
			if err != nil {
				return
			}

			// 缓存数据到Redis
			cacheKey := fmt.Sprintf("%s%s", dictCacheKeyPrefix, dt.DictType)
			utils.Set(dataCtx, client, cacheKey, dataList, defaultCacheDuration)

		}(dictType)
	}

	wg.Wait()
}

type DeptData struct {
	Id       int64       `json:"id"`
	ParentId int64       `json:"parentId"`
	Name     string      `json:"name"`
	Children []*DeptData `json:"children,omitempty"`
}
