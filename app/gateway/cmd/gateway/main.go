package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/http"
	"kratos-blog/app/gateway/internal/conf"
	"kratos-blog/pkg/server"
	"os"
	"path/filepath"
	"time"

	_ "go.uber.org/automaxprocs"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = server.GateWay
	// Version is the version of the compiled software.
	Version = "1.0"
	// flagconf is the config flag.
	flagconf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagconf, "conf", server.GateWayConfPath, "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, r registry.Registrar) *kratos.App {

	return kratos.New(
		kratos.ID(id+"-gateway"),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
		),
		kratos.Registrar(r),
	)
}

func main() {
	flag.Parse()
	var logger log.Logger
	initLog(&logger)
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	var rc conf.Registry
	if err := c.Scan(&rc); err != nil {
		panic(err)
	}
	listenConfigChange(c, "domain", &bc)

	app, cleanup, err := wireApp(&bc, &rc, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}

func initLog(logger *log.Logger) {
	logFileName := fmt.Sprintf("%s%s", time.Now().Format("2006-01-02"), "-gateway.log")
	_, e := os.Stat("log")
	if e != nil && os.IsNotExist(e) {
		os.MkdirAll("log", os.ModePerm)
	}
	logPath := filepath.Join("log", logFileName)
	f, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return
	}
	*logger = log.With(log.NewStdLogger(f),
		"ts", log.DefaultTimestamp,
		"caller", log.DefaultCaller,
		"service.id", id+"-gateway",
		"service.name", Name,
		"service.version", Version,
		"trace.id", tracing.TraceID(),
		"span.id", tracing.SpanID(),
	)
}

func listenConfigChange(c config.Config, serverName string,
	change *conf.Bootstrap) {
	if err := c.Watch(serverName, func(key string, value config.Value) {
		fmt.Printf("config changed: %s = %v\n", key, value)
		var confC map[string]interface{}
		value.Scan(&confC)
		change.Domain.Open = confC["open"].(bool)
		list := confC["origin"].([]interface{})
		ParseJSONToStruct(list, &change.Domain.Origin)
	}); err != nil {
		log.Error(err)
	}
}

func ParseJSONToList(jsonStr string, list *[]interface{}) bool {
	if len(jsonStr) == 0 {
		return false
	}
	var memoryList []interface{}
	if err := json.Unmarshal([]byte(jsonStr), &memoryList); err != nil {
		fmt.Errorf("error parsing JSON: %s", err)
		return false
	}
	*list = memoryList
	return true
}

func ParseJSONToStruct(jsonStr interface{}, resultStruct interface{}) error {
	f, _ := json.Marshal(jsonStr)
	err := json.Unmarshal(f, resultStruct)
	if err != nil {
		return fmt.Errorf("error parsing JSON: %s", err)
	}
	return nil
}
