package version

import (
	"github.com/aisk/logp"
)

// Version is lean-cli's version.
const Version = "0.16.3"

func PrintCurrentVersion() {
	logp.Info("当前命令行工具版本：", Version)
}
