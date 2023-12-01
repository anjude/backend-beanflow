package enum

type EnvKey string

const (
	EkConfigPath EnvKey = "CONFIG_PATH"
	EkEnv        EnvKey = "ENV"
)

func (e EnvKey) ToString() string {
	return string(e)
}

type Env string

const (
	LOCAL Env = "LOCAL"
	DEV   Env = "DEV"
	LIVE  Env = "LIVE"
)

func (e Env) ToString() string {
	return string(e)
}
