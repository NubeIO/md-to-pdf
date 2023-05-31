package main

import (
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type application struct {
	pandocPath     string
	userHome       string
	pandocDataDir  string
	commandTimeout time.Duration
}

var logger = logrus.New()
var permWrite = os.FileMode(0600)
var permDir = os.FileMode(0755)

func main() {
	app := &application{}

	var host string
	var wait time.Duration
	var debugOutput bool
	flag.StringVar(&app.userHome,
		"home",
		lookupEnvOrString("PANDOC_HOME", "/home"),
		"set your home dir to save a file not in tmp dir")
	flag.StringVar(&host,
		"host",
		lookupEnvOrString("PANDOC_HOST", ":8080"),
		"IP and Port to bind to. Can also be set through the PANDOC_HOST environment variable.")
	flag.BoolVar(&debugOutput,
		"debug",
		lookupEnvOrBool("PANDOC_DEBUG", false),
		"Enable DEBUG mode. Can also be set through the PANDOC_DEBUG environment variable.")
	flag.DurationVar(&wait,
		"graceful-timeout",
		lookupEnvOrDuration("PANDOC_GRACEFUL_TIMEOUT", 5*time.Second),
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m. Can also be set through the PANDOC_GRACEFUL_TIMEOUT environment variable.")
	flag.DurationVar(&app.commandTimeout,
		"command-timeout",
		lookupEnvOrDuration("PANDOC_COMMAND_TIMEOUT", 1*time.Minute),
		"the timeout for the conversion command. Can also be set through the PANDOC_COMMAND_TIMEOUT environment variable.")
	flag.StringVar(&app.pandocPath,
		"pandoc-path",
		lookupEnvOrString("PANDOC_PATH", "/usr/share/pandoc"),
		"The path of the pandoc binary. Can also be set through the PANDOC_PATH environment variable.")
	flag.StringVar(&app.pandocDataDir,
		"pandoc-data-dir",
		lookupEnvOrString("PANDOC_DATA_DIR", "/.pandoc"),
		"The pandoc data dir containing the templates. Can also be set through the PANDOC_DATA_DIR environment variable.")
	flag.Parse()

	gin.SetMode(gin.ReleaseMode)

	logger.SetOutput(os.Stdout)
	logger.SetLevel(logrus.InfoLevel)
	if debugOutput {
		gin.SetMode(gin.DebugMode)
		logger.SetLevel(logrus.DebugLevel)
		logger.Debug("DEBUG mode enabled")
	}

	logger.Infof("host: %s", host)
	logger.Infof("debug: %t", debugOutput)
	logger.Infof("graceful timeout: %s", wait)
	logger.Infof("command timeout: %s", app.commandTimeout)
	logger.Infof("pandoc path: %s", app.pandocPath)
	logger.Infof("pandoc data dir: %s", app.pandocDataDir)
	logger.Infof("homedir: %s", app.userHome)

	srv := &http.Server{
		Addr:    host,
		Handler: app.routes(),
	}
	logger.Infof("Starting server on %s", host)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logger.Error(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(err)
	}
	logger.Info("shutting down")
	os.Exit(0)
}
