package chrome

import (
	"fmt"
	"github.com/netless-io/rebirth/internal/app/rebirth/cli"
	"github.com/netless-io/rebirth/internal/pkg/utils"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

// Path get chrome executable file path
func execPath() (string, bool) {
	var p []string
	switch runtime.GOOS {
	case "linux":
		p = []string{
			"google-chrome",
			"google-chrome-stable",
			"/usr/bin/google-chrome",
		}
	case "darwin":
		p = []string{
			"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome",
		}
	case "windows":
		p = []string{
			"chrome",
			"chrome.exe", // in case PATHEXT is misconfigured
			`C:\Program Files (x86)\Google\Chrome\Application\chrome.exe`,
			`C:\Program Files\Google\Chrome\Application\chrome.exe`,
			filepath.Join(os.Getenv("USERPROFILE"), `AppData\Local\Google\Chrome\Application\chrome.exe`),
		}
	}

	if found, ok := findExecPath(p); ok {
		return found, true
	}

	return "", false
}

// findExecPath Check whether the file is executable
func findExecPath(p []string) (string, bool) {
	for _, cp := range p {
		found, err := exec.LookPath(cp)

		if err == nil {
			return found, true
		}
	}

	return "", false
}

// execPathHandle Check whether the chrome path entered by the user or the built-in chrome path of the program is an executable file
func execPathHandle(chromePath *string) (string, error) {
	cp := *chromePath

	errorStatus := fmt.Sprint("chrome browser is not found in the current system")
	if cp != "" {
		errorStatus = fmt.Sprintf("chrome executable file does not exist: %s", cp)
	}

	var execNotExist = cli.StatusError{
		Status:     errorStatus,
		StatusCode: 126,
	}

	// The user did not specify --chrome-path flag, and the built-in chrome path does not exist
	if cp == "" {
		return "", execNotExist
	}

	if p, ok := findExecPath([]string{cp}); ok {
		return p, nil
	}

	return "", execNotExist
}

// findUserDataPath Obtain the corresponding `rebirth` temporary directory according to different systems
func findUserDataPath() (string, bool) {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		return "/tmp/rebirth", true
	}

	var (
		temp        = os.Getenv("TEMP")
		tmp         = os.Getenv("TMP")
		userProfile = os.Getenv("USERPROFILE")
	)

	if temp != "" {
		return temp, true
	}

	if tmp != "" {
		return tmp, true
	}

	if userProfile != "" {
		return path.Join(userProfile, "rebirth"), true
	}

	if d, err := os.UserHomeDir(); err == nil {
		return path.Join(d, "AppData", "Local", "Temp", "rebirth"), true
	}


	return "", false
}

// userDataPathHandle Whether to obtain the temporary directory of the corresponding system
func userDataPathHandle(d *string) (string, error) {
	if *d == "" {
		return "", cli.StatusError{
			Status:     fmt.Sprint("Please make sure that TEMP or TMP or USERPROFILE environment variables exist on your computer Or use --chrome-data-dir to explicitly set the path"),
			StatusCode: 126,
		}
	}

	return *d, nil
}

// findExtensionDir Get the chrome extension directory
// There may be a duplicate value, but it is irrelevant
func findExtensionDir() (string, bool) {
	var list []string

	// e.g:
	// $ rebirth -> @/usr/local/bin/rebirth -> ~/.rebirth/rebirth -> ~/.rebirth/extension/
	if d, ok := utils.ExecDir(); ok {
		list = append(list, path.Join(d, "extension"))
	}

	if p, err := os.UserHomeDir(); err == nil {
		list = append(list, path.Join(p, ".rebirth", "extension"))
	}

	if d, err := os.Getwd(); err == nil {
		// $ cd /usr/lib/
		// $ rebirth -> /usr/lib/extension
		list = append(list, path.Join(d, "extension"))

		// ----------------------------------------------------------------------------
		// The following path is mainly compatible with the development environment
		// ----------------------------------------------------------------------------

		// $ cd ./rebirth/
		// $ go run ./cmd/rebirth/ -> ./web/extension/dist
		list = append(list, path.Join(d, "web", "extension", "dist"))

		// $ cd ./rebirth/cmd
		// $ go run ./rebirth/ -> ../web/extension/dist
		list = append(list, path.Join(d, "..", "web", "extension", "dist"))

		// $ cd ./rebirth/cmd/rebirth
		// $ go run ./ -> ../../web/extension/dist
		list = append(list, path.Join(d, "..", "..", "web", "extension", "dist"))
	}

	for _, dir := range list {
		if d, err := extensionDirHandle(&dir); err == nil {
			return d, true
		}
	}

	return "", false
}

// extensionDirHandle Check whether the plugin directory meets the conditions
func extensionDirHandle(dir *string) (string, error) {
	d := *dir

	if d == "" {
		return "", cli.StatusError{
			Status:     fmt.Sprint("The chrome extension directory is not found, please use --extension-dir to display the specified"),
			StatusCode: 126,
		}
	}

	if !utils.IsDir(d) {
		return "", cli.StatusError{
			Status:     fmt.Sprintf("%s is not a chrome extension directory", d),
			StatusCode: 126,
		}
	}

	mustFile := []string{
		"key.pem",
		"manifest.json",
		"background.js",
		"content_script.js",
		"injected.js",
	}

	for _, f := range mustFile {
		if !utils.IsFile(path.Join(d, f)) {
			return "", cli.StatusError{
				Status:     fmt.Sprintf("Please make sure there is %s file in the %s directory", strings.Join(mustFile, ", "), d),
				StatusCode: 126,
			}
		}
	}

	return d, nil
}
