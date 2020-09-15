package logging

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

const (
	// DefaultLoggingLevel is the default logging level...
	DefaultLoggingLevel = zerolog.WarnLevel
)

// LogConfiguration is the logging configuration for both Zerolog and lumberjack
type LogConfiguration struct {
	ConsoleLoggingEnabled bool
	Directory             string
	EncodeLogsAsJSON      bool
	FileLoggingEnabled    bool
	Filename              string
	LoggingLevel          string
	MaxFileAge            int
	MaxFileBackups        int
	MaxFileSize           int
}

// PSLogger holds the default logger and configuration
type PSLogger struct {
	*zerolog.Logger
	Config *LogConfiguration
}

// Start returns a configured PSLogger struct
func Start(dir string, config string) (*PSLogger, error) {
	// Check for a logging configuration and create if it doesn't exists, create using the defaults
	_, err := os.Stat(path.Join(dir, config))
	if os.IsNotExist(err) {
		if e := createDefaultLoggingConfiguration(dir, config); e != nil {
			return nil, err
		}
	}

	// Open the logging configuration file
	f, err := os.Open(path.Join(dir, config))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Decode the logging configuration and configure the logger
	lc := LogConfiguration{}
	d := json.NewDecoder(f)
	if err = d.Decode(&lc); err != nil {
		return nil, err
	}

	lcl := LogConfiguration{
		ConsoleLoggingEnabled: lc.ConsoleLoggingEnabled,
		EncodeLogsAsJSON:      lc.EncodeLogsAsJSON,
		FileLoggingEnabled:    lc.FileLoggingEnabled,
		Directory:             lc.Directory,
		Filename:              lc.Filename,
		MaxFileAge:            lc.MaxFileAge,
		MaxFileBackups:        lc.MaxFileBackups,
		MaxFileSize:           lc.MaxFileSize,
		LoggingLevel:          lc.LoggingLevel,
	}

	log, err := lcl.configure()

	if err != nil {
		return nil, err
	}

	lvl, err := zerolog.ParseLevel(lc.LoggingLevel)
	if err != nil {
		// If level can't be parsed
		lvl = DefaultLoggingLevel
	}

	zerolog.SetGlobalLevel(lvl)

	return log, nil
}

func (lc *LogConfiguration) configure() (*PSLogger, error) {
	var writers []io.Writer

	if lc.ConsoleLoggingEnabled {
		writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr})
	}

	if lc.FileLoggingEnabled {
		rw, err := lc.createRollingFile()
		if err != nil {
			return nil, err
		}

		writers = append(writers, rw)
	}

	multipass := io.MultiWriter(writers...)
	logger := zerolog.New(multipass).With().Timestamp().Logger()

	return &PSLogger{
		Logger: &logger,
		Config: lc,
	}, nil
}

func (lc *LogConfiguration) createRollingFile() (io.Writer, error) {
	if err := os.MkdirAll(lc.Directory, 0755); err != nil {
		if os.IsPermission(err) {
			// if the binary doesn't have permissions to create in the requested directory
			// fall back to creating the log in the current running directory
			d, _ := os.Getwd()

			if e := os.MkdirAll(d, 0755); e != nil {
				return nil, fmt.Errorf("Failed to create the log in both '%s' and '%s'", lc.Directory, d)
			}

			// Set the output to the local directory
			lc.Directory = d
		}
	} else {
		// If not a permissions error, return the error
		return nil, err
	}

	// All good, return the logger
	return &lumberjack.Logger{
		Filename:   path.Join(lc.Directory, lc.Filename),
		MaxAge:     lc.MaxFileAge,
		MaxBackups: lc.MaxFileBackups,
		MaxSize:    lc.MaxFileSize,
	}, nil
}

// Creates a default logging configuration file
func createDefaultLoggingConfiguration(dir string, conf string) error {
	c := LogConfiguration{
		ConsoleLoggingEnabled: false,
		Directory:             "",
		EncodeLogsAsJSON:      true,
		FileLoggingEnabled:    true,
		Filename:              "logs/output.log",
		MaxFileAge:            30,
		MaxFileBackups:        10,
		MaxFileSize:           10,
		LoggingLevel:          "warn",
	}

	f, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		return err
	}

	// make sure the directories exist
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	if err := ioutil.WriteFile(path.Join(dir, conf), f, 0644); err != nil {
		return err
	}

	return nil
}
