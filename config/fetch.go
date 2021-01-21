package config

import (
	"flag"
	"path/filepath"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// Config contains the input for the fileserver.
type Config struct {
	// Dir to serve.
	Dir string

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

// Fetch configuration.
func (c *Config) Fetch() error {
	config := viper.New()

	flag.String("dir", "", "Directory to serve")
	flag.String("addr", ":80", "Listening address")
	flag.Bool("https", false, "Listen in HTTPS")
	flag.String("cert", "", "SSL certificate file")
	flag.String("key", "", "SSL key file")
	flag.Bool("cors", false, "Allow CORS")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	err := config.BindPFlags(pflag.CommandLine)
	if err != nil {
		return err
	}

	c.Dir = config.GetString("dir")
	c.Addr = config.GetString("addr")
	c.HTTPS = config.GetBool("https")
	c.SSLCertFile = config.GetString("cert")
	c.SSLKeyFile = config.GetString("key")
	c.AllowCors = config.GetBool("cors")

	c.Dir, err = filepath.Abs(c.Dir)
	return err
}

// New returns a Config instance.
func New() *Config {
	return &Config{}
}
