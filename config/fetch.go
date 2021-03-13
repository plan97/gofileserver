package config

import (
	"flag"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config contains the input for the fileserver.
type Config struct {
	// BaseDir to serve.
	BaseDir string

	// Addr to listen.
	Addr string

	// HTTPS enable.
	HTTPS bool

	// SSLCertFile refers to the location of SSL certificate.
	SSLCertFile string

	// SSLKeyFile refers to the location of SSL key.
	SSLKeyFile string

	// AllowCORS enable.
	AllowCors bool
}

// Fetch configuration from flags.
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
