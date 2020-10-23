package chrome

type Config struct {
	Path   string
	UserDataDir  string
	ExtensionDir string
}

func (c *Config) Init() {
	if p, ok := execPath(); ok {
		c.Path = p
	}

	if d, ok := findUserDataPath(); ok {
		c.UserDataDir = d
	}

	if d, ok := findExtensionDir(); ok {
		c.ExtensionDir = d
	}
}
