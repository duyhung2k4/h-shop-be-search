package config

func init() {
	loadEnv()
	connectElastic()
	connectRedis()
}
