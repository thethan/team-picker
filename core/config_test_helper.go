package core

func SetupTestConfig() *Config {
	config := Config{DataLocation: "../data/"}
	config.Logging = &Logging{
		LogLevel: uint(3),
	}

	return &config
}

func SetupTestDependencies() *Dependencies {
	dep := &Dependencies{}
	return dep
}