package tools

import (
	"httpserver/logger"
	"os"
)

// 获取系统环境变量
func GetEnv(envName string) (string, bool) {
	env := os.Getenv(envName)
	if env == "" {
		logger.Errorf("获取系统环境变量失败，请检查变量名")
		return env, false
	}
	return env, true
}
