package config

func init() {
	loadEnv()
	connectElastic()
}
