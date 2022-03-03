package logging

import (
	"os"
	"testing"
)

func TestLoadConfiguration(t *testing.T) {
	const (
		logLevel = 4
		filePath = "/var/log/lumerin.log"
	)
	config, err := loadConfiguration("config.json")
	if err != nil {
		t.Error("loadConfiguration() returned an error")
	}
	if config.Level != logLevel {
		t.Errorf("Expected level to be %v but received %v", logLevel, config.Level)
	}
	if config.FilePath != filePath {
		t.Errorf("Expected level to be %s but received %s", filePath, config.FilePath)
	}
}

func TestInitStandardLoggerInitialized(t *testing.T) {
	Init(true)
	_, err := GetLogger()
	if err != nil {
		t.Error("StandardLogger not initialized")
	}
}

func TestInitLogFileCreated(t *testing.T) {
	createTestDirectory(t)

	Init(true)
	_, err := os.Stat("/var/log/test/lumerin.log")
	if err != nil {
		t.Error("Config file not found")
	}

	deleteTestDirectoryAndContents(t)
}

func TestStandardLoggerLogsToFile(t *testing.T) {
	createTestDirectory(t)

	Init(true)
	sl, _ := GetLogger()
	sl.InfoEvent("TestEvent:", "this is a test log")
	logFile, _ := os.ReadFile("/var/log/test/lumerin.log")
	logEntry := string(logFile)
	if logEntry == "" {
		t.Error("Log entry not created")
	}

	deleteTestDirectoryAndContents(t)
}

// Helper functions
func createTestDirectory(t *testing.T) {
	err := os.MkdirAll("/var/log/test", 0666)
	if err != nil {
		t.Errorf("Test directory not created")
	}
}

func deleteTestDirectoryAndContents(t *testing.T) {
	Cleanup()
	err := os.RemoveAll("/var/log/test")
	if err != nil {
		t.Error("Could not delete test directory and contents")
	}
}
