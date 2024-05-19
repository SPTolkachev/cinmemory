package logger

type Config struct {
	DebugFn LogFn
	InfoFn  LogFn
	WarnFn  LogFn
	ErrorFn LogFn
}

func configure(logger *Logger, args ...Configurators) {
	config := Config{}
	for _, arg := range args {
		arg(&config)
	}

	logger.debugFn = config.DebugFn
	logger.infoFn = config.InfoFn
	logger.warnFn = config.WarnFn
	logger.errorFn = config.ErrorFn
}
