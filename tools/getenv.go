package tools

import "os"

// 获取系统环境变量
func GetEnv(envName string) (string, bool) {
	env := os.Getenv(envName)
	if env == "" {
		return env, false
	}
	return env, true
}
