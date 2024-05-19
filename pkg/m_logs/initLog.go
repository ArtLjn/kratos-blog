package m_logs

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/robfig/cron/v3"
	"kratos-blog/pkg/server"
	"os"
	"path/filepath"
	"time"
)

func InitLog(logger *log.Logger, id, name, version, prefix string) {
	c := cron.New(cron.WithSeconds())

	// 每天创建新的日志文件
	_, err := c.AddFunc("0 0 0 * * *", func() {
		InitDailyLogFile(logger, id, name, version, prefix)
	})

	// 每7天清理旧的日志文件
	_, err = c.AddFunc("0 0 0 */7 * *", func() {
		CleanOldFile(0, server.LogOutStreamPath)
	})

	// 开始定时任务
	c.Start()
	InitDailyLogFile(logger, id, name, version, prefix)
	if err != nil {
		log.Fatalf("Failed to set up log file creation and cleaning: %v", err)
	}
}

// CleanOldFile 清理指定天数之前的日志文件
func CleanOldFile(retainDays int, path string) {
	// 获取当前时间
	now := time.Now()

	// 计算保留的截止日期
	cutoff := now.AddDate(0, 0, -retainDays)

	// 读取日志目录中的所有文件
	files, err := os.ReadDir(path)
	if err != nil {
		log.Error("Error reading log directory: %v", err)
		return
	}

	for _, file := range files {
		// 检查文件是否是日志文件
		if !file.IsDir() {
			// 获取文件的修改时间
			fileInfo, _ := file.Info()
			modTime := fileInfo.ModTime()
			// 如果文件的修改时间早于截止日期，则删除该文件
			if modTime.Before(cutoff) {
				filePath := filepath.Join(path, file.Name())
				err = os.Remove(filePath)
				if err != nil {
					log.Error("Error removing old log file %s: %v", filePath, err)
				} else {
					fmt.Printf("Removed old log file: %s\n", filePath)
				}
			}
		}
	}
}

func InitDailyLogFile(logger *log.Logger, id, name, version, prefix string) {
	logFileName := fmt.Sprintf("%s-%s.log", time.Now().Format("2006-01-02"), prefix)
	logPath := filepath.Join(server.LogOutStreamPath, logFileName)

	// 创建日志目录，如果不存在的话
	if _, err := os.Stat(server.LogOutStreamPath); os.IsNotExist(err) {
		err := os.MkdirAll(server.LogOutStreamPath, os.ModePerm)
		if err != nil {
			log.Fatalf("Failed to create log directory: %v", err)
		}
	}

	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	*logger = log.With(log.NewStdLogger(f),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id+"-"+prefix,
		"service.name", name,
		"service.version", version,
		// 以下字段需要根据你的 tracing 包和实现来调整
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
}
