package logger

import (
	"fmt"
	"github.com/druidcaesa/gotool"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
	"vue-admin-system-golang/utils/fileFunc"
)

var Logger *logrus.Logger

func LoggerToFile1() gin.HandlerFunc {
	fileName := "./app2.log" //写入文件
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("err", err)
	}

	//实例化
	Logger = logrus.New()

	//设置输出
	Logger.Out = src

	//设置日志级别
	Logger.SetLevel(logrus.DebugLevel)

	//设置日志格式
	Logger.SetFormatter(&logrus.TextFormatter{})
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 从request.body 读取请求数据
		rawData, _ := c.GetRawData()

		// 日志格式
		Logger.Infof("| %3d | %13v | %15s | %s | %s | %s", statusCode, latencyTime, clientIP, reqMethod, reqUri, rawData)
	}

}
func LoggerToFile() gin.HandlerFunc {
	logFilePath := "./logFile"
	err := fileFunc.CreateDir(logFilePath)
	if err != nil {
		panic("创建日志文件夹失败")
	}
	logFileName := "golang"
	fileName := path.Join(logFilePath, logFileName)
	if !gotool.FileUtils.Exists(fileName) {
		create, err := os.Create(fileName)
		if err != nil {
			gotool.Logs.ErrorLog().Print(err)
		}
		defer create.Close()
	}
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println(err)
	}
	Logger = logrus.New()
	Logger.Out = src
	Logger.SetLevel(logrus.DebugLevel)
	logWriter, err := rotatelogs.New(
		fileName+".%Y%m%d.log",
		rotatelogs.WithLinkName(fileName),
		rotatelogs.WithMaxAge(365*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// 新增 Hook
	Logger.AddHook(lfHook)

	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()

		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		// 日志格式
		Logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}
