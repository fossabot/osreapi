package windows_remote_executor

import (
	"fmt"
	"runtime"

	"github.com/common-nighthawk/go-figure"
)

var (
	Version   string
	GitUrl    string
	GitBranch string
	GitCommit string
	BuildTime string
	UserName  string
	UserEmail string
	title     = figure.NewFigure("Remote Executor", "doom", true).String()
)

func VersionInfo() string {
	var info = fmt.Sprintf("Go Version: %s\nOS: %s\nArch: %s\n",
		runtime.Version(), runtime.GOOS, runtime.GOARCH)
	if Version != "" {
		info += fmt.Sprintf("Version: %s\n", Version)
	}
	if GitUrl != "" {
		info += fmt.Sprintf("Git URL: %s\n", GitUrl)
	}
	if GitBranch != "" {
		info += fmt.Sprintf("Git Branch: %s\n", GitBranch)
	}
	if GitBranch != "" {
		info += fmt.Sprintf("Git Commit: %s\n", GitCommit)
	}
	if GitBranch != "" {
		info += fmt.Sprintf("Build Time: %s\n", BuildTime)
	}
	return info
}

func PrintHeadInfo() {
	fmt.Println(title)
	fmt.Print(VersionInfo())
}
