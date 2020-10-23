package chrome

import (
	"github.com/netless-io/rebirth/internal/app/rebirth/cli"
	"os/exec"
)

func Cmd(config *Config) (*exec.Cmd, error) {
	chromePath, err := execPathHandle(&config.Path)
	if err != nil {
		return nil, err
	}

	userDataPath, err := userDataPathHandle(&config.UserDataDir)
	if err != nil {
		return nil, err
	}

	extensionDir, err := extensionDirHandle(&config.ExtensionDir)
	if err != nil {
		return nil, err
	}

	return exec.Command(chromePath, launchFlag(&userDataPath, &extensionDir)...), nil
}

// launchFlag chrome launch flags
func launchFlag(userDataDir *string, extensionDir *string) []string {
	args := []string{
		// Users don't need to interact with the document for video or audio sources to start playing automatically.
		"--autoplay-policy=no-user-gesture-required",

		// Enable screen capturing support for MediaStream API.
		"--enable-usermedia-screen-capturing",

		// Allow non-secure origins to use the screen capture API and the desktopCapture extension API.
		"--allow-http-screen-capture",

		// Enables remote debug over HTTP on the specified port.
		"--remote-debugging-port=9222",

		// Adds the given extension ID to all the permission allowlists.
		"--whitelisted-extension-id=cnlcagjlokccajjeehpfccjlkgflmmgj",

		// The /dev/shm partition is too small in certain VM environments, causing Chrome to fail or crash (see http://crbug.com/715363).
		// Use this flag to work-around this issue (a temporary directory will always be used to create anonymous shared memory files).
		"--disable-dev-shm-usage",

		// Disable task throttling of timer tasks from background pages.
		"--disable-background-timer-throttling",

		// set chrome windows size
		"--window-size=1792,1097",

		// Treat given (insecure) origins as secure origins.
		//"--unsafely-treat-insecure-origin-as-secure=http://127.0.0.1",

		// Skip First Run tasks(hide ®welcome use chrome popup)
		"--no-first-run",

		// Enable indication that browser is controlled by automation.
		// The save password window will no longer pop up when a form with a password is submitted
		"--enable-automation",

		// Specifies plain text store encryption
		"--password-store=basic",

		// Use fake keychain (only mac)
		"--use-mock-keychain",

		// Force all monitors to be treated as though they have the srgb color profile
		"--force-color-profile=srgb",

		// Disables syncing browser data to a Google Account.
		"--disable-sync",

		// Prevent renderer process backgrounding when set.
		"--disable-renderer-backgrounding",

		// Normally when the user attempts to navigate to a page that was the result of a post we prompt to make sure they want to.
		// This switch may be used to disable that check. This switch is used during automated testing.
		"--disable-prompt-on-repost",

		// Disables pop-up blocking.
		"--disable-popup-blocking",

		// Disables the IPC flooding protection.
		"--disable-ipc-flooding-protection",

		// Disable several subsystems which run network requests in the background.
		"--disable-background-networking",

		// Suppresses hang monitor dialogs in renderer processes.
		// This may allow slow unload handlers on a page to prevent the tab from closing,
		// but the Task Manager can be used to terminate the offending process in this case.
		"--disable-hang-monitor",

		// Disables TranslateUI
		"--disable-features=TranslateUI",

		// Disables installation of default apps on first run.
		"--disable-default-apps",

		// Disable default component extensions with background pages.
		"--disable-component-extensions-with-background-pages",

		// Disables the client-side phishing detection feature.
		"--disable-client-side-phishing-detection",

		// Disables the crash reporting.
		"--disable-breakpad",

		// Disable backgrounding renders for occluded windows.
		"--disable-backgrounding-occluded-windows",

		// Enables the recording of metrics reports but disables reporting.
		"--metrics-recording-only",

		// load chrome extension
		"--load-extension=" + *extensionDir,

		// disable chrome extensions except rebirth
		"--disable-extensions-except=" + *extensionDir,

		// Directory where the browser stores the user profile.
		// Chrome's debug log is inside (if --debug is turned on)
		"--user-data-dir=" + *userDataDir,
	}

	// If debug is enabled, chrome debug log is recorded
	if cli.RootFlags.Debug {
		args = append(args, "--enable-logging=stderr --v=1")
	}

	return args
}
