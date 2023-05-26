package main

func LogInfo(module string, code string) {
	Log.Create.Info().Str("module", module).Str("code", code).Msg(Lang.Get(code))
}

func LogInfoWithMessage(module string, code string, message string) {
	Log.Create.Info().Str("module", module).Str("code", code).Msg(message)
}

func LogFatal(module string, code string) {
	Log.Create.Fatal().Str("module", module).Str("code", code).Msg(Lang.Get(code))
}

func LogFatalWithException(module string, code string, err error) {
	Log.Create.Fatal().Str("module", module).Str("code", code).Str("exception", err.Error()).Msg(Lang.Get(code))
}

