package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func lookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultVal
}

func lookupEnvOrBool(key string, defaultVal bool) bool {
	if val, ok := os.LookupEnv(key); ok {
		v, err := strconv.ParseBool(val)
		if err != nil {
			return defaultVal
		}
		return v
	}
	return defaultVal
}

func lookupEnvOrDuration(key string, defaultVal time.Duration) time.Duration {
	if val, ok := os.LookupEnv(key); ok {
		v, err := time.ParseDuration(val)
		if err != nil {
			return defaultVal
		}
		return v
	}
	return defaultVal
}

func errorJson(errorText string) gin.H {
	return gin.H{
		"error": errorText,
	}
}

func (app *application) killProcessIfRunning(cmd *exec.Cmd) {
	if cmd.Process == nil {
		return
	}
	if err := cmd.Process.Release(); err != nil {
		return
	}
	if err := cmd.Process.Kill(); err != nil {
		return
	}
}
