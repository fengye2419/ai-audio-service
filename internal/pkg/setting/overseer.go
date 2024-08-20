package setting

var (
	Overseer = struct {
		FetcherFilePath string
		HTTPPort        string
		HTTPAddress     string
		DebugMode       bool
	}{
		FetcherFilePath: "",
		HTTPAddress:     "",
	}
)

func InitOverseerConfig() {
	sec := Cfg.Section("overseer")
	Overseer.FetcherFilePath = sec.Key("FETCHER_FILE_PATH").MustString("")
	Overseer.HTTPPort = sec.Key("HTTP_PORT").MustString("3000")
	Overseer.HTTPAddress = ":" + Overseer.HTTPPort
	Overseer.DebugMode = sec.Key("DEBUG_MODE").MustBool(true)
}
