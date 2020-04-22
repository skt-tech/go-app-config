package main

// ServerConfig parameters
type ServerConfig struct {
	Port int
}

// DatabaseConfig parameters
type DatabaseConfig struct {
	HostName   string
	DBUser     string
	DBPassword string
}

// Configurations - parameters
type Configurations struct {
	Server       ServerConfig
	Database     DatabaseConfig
	MaxWorkers   int
	MaxQueues    int
	SensitiveVal string
}
