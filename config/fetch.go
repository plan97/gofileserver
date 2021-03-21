// Package config provides the necessary configurations for gofileserver.
//
// The configurations can be retrieved from command line flags if needed.
package config

import (
	"flag"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config for the fileserver.
type Config struct {
	// BaseDir to serve.
	BaseDir string

	// Addr to listen.
	Addr string

	// HTTPS flag to enable HTTPS using SSLCertFile and SSLKeyFile.
	HTTPS bool

	// SSLCertFile refers to the location of SSL certificate.
	SSLCertFile string

	// SSLKeyFile refers to the location of SSL key.
	SSLKeyFile string

	// AllowCORS flag to allow requests from all origin.
	AllowCors bool
}

// Fetch configuration from command line flags.
func (conf *Config) Fetch() (err error) {
	config := viper.New()

	flag.String("dir", "", "Directory to serve")
	flag.String("addr", ":80", "Listening address")
	flag.Bool("https", false, "Listen in HTTPS")
	flag.String("cert", "", "SSL certificate file")
	flag.String("key", "", "SSL key file")
	flag.Bool("cors", false, "Allow CORS")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err = config.BindPFlags(pflag.CommandLine)
	if err != nil {
		return
	}

	conf.BaseDir = config.GetString("dir")
	conf.Addr = config.GetString("addr")
	conf.HTTPS = config.GetBool("https")
	conf.SSLCertFile = config.GetString("cert")
	conf.SSLKeyFile = config.GetString("key")
	conf.AllowCors = config.GetBool("cors")

	conf.BaseDir, err = filepath.Abs(conf.BaseDir)
	return
}

// New returns a Config instance.
func New() *Config {
	return &Config{}
}
