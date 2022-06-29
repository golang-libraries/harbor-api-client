package harbor

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/go-logr/logr"
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/golang-libraries/harbor-api-client/pkg/harbor/project"
	"github.com/golang-libraries/harbor-api-client/pkg/harbor/robot"
	"github.com/golang-libraries/harbor-api-client/pkg/logger"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"
)

const (
	DefaultHost             = "localhost"
	DefaultBasePath         = "/api/v2.0"
	DefaultLoggerScope      = "HarborClient"
	DefaultCABundleFilePath = ""
	DefaultInsecureSSL      = false
	DefaultUsername         = ""
	DefaultPassword         = ""
)

var (
	schemes = []string{"http", "https"}
)

type ClientConfig struct {
	envCfg  *viper.Viper
	schemes []string
}

func (cfg *ClientConfig) Host() string {
	return cfg.envCfg.GetString("host")
}
func (cfg *ClientConfig) Url() string {
	return fmt.Sprintf("%s/%s", cfg.Host(), cfg.BasePath())
}
func (cfg *ClientConfig) BasePath() string {
	return cfg.envCfg.GetString("basepath")
}
func (cfg *ClientConfig) Schemes() []string {
	return cfg.schemes
}
func (cfg *ClientConfig) Username() string {
	return cfg.envCfg.GetString("username")
}
func (cfg *ClientConfig) Password() string {
	return cfg.envCfg.GetString("password")
}

func (cfg *ClientConfig) InsecureSSL() bool {
	return cfg.envCfg.GetBool("insecure_ssl")
}

func (cfg *ClientConfig) CAPath() string {
	return cfg.envCfg.GetString("ca_path")
}

func NewClientConfig() (ClientConfig, error) {
	var prefix = "HARBOR"
	env := viper.New()
	env.SetDefault("host", DefaultHost)
	env.SetDefault("basepath", DefaultBasePath)
	env.SetDefault("username", DefaultUsername)
	env.SetDefault("password", DefaultPassword)
	env.SetDefault("ca_path", DefaultCABundleFilePath)
	env.SetDefault("insecure_ssl", DefaultInsecureSSL)
	env.AllowEmptyEnv(false)
	env.SetEnvPrefix(prefix)
	env.AutomaticEnv()
	cfg := ClientConfig{
		envCfg:  env,
		schemes: schemes,
	}
	return cfg, nil
}

type Client struct {
	transport *httptransport.Runtime
	authInfo  runtime.ClientAuthInfoWriter
	log       logr.Logger
	Project   project.Service
	Robot     robot.Service
}

func NewRootCABundle(caBundlePath string) (*x509.CertPool, error) {
	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}
	bundleFile, err := os.ReadFile(caBundlePath)
	if err != nil {
		return nil, err
	}
	certPool.AppendCertsFromPEM(bundleFile)
	return certPool, nil
}

func NewHTTPClient(timeoutSec int, caBundlePath string, insecureSSL bool) (*http.Client, error) {
	if timeoutSec == 0 {
		timeoutSec = 5
	}
	var rootCAs *x509.CertPool
	if caBundlePath != "" {
		var err error
		rootCAs, err = NewRootCABundle(caBundlePath)
		if err != nil {
			return nil, err
		}
	}
	tlsConfig := &tls.Config{
		InsecureSkipVerify: insecureSSL,
		RootCAs:            rootCAs,
	}
	transport := &http.Transport{
		TLSClientConfig: tlsConfig,
	}
	client := &http.Client{
		Timeout:   time.Second * time.Duration(timeoutSec),
		Transport: transport,
	}
	return client, nil
}

func NewTransport() (*httptransport.Runtime, error) {
	cfg, err := NewClientConfig()
	if err != nil {
		return nil, err
	}
	httpClient, err := NewHTTPClient(10, cfg.CAPath(), cfg.InsecureSSL())
	if err != nil {
		return nil, err
	}
	transport := httptransport.NewWithClient(cfg.Host(), cfg.BasePath(), cfg.Schemes(), httpClient)
	authInfo := httptransport.BasicAuth(cfg.Username(), cfg.Password())
	transport.DefaultAuthentication = authInfo
	return transport, nil
}

func NewClient() (*Client, error) {
	log, err := logger.NewLogger(DefaultLoggerScope, false)
	if err != nil {
		return nil, err
	}
	transport, err := NewTransport()
	if err != nil {
		return nil, err
	}
	return &Client{
		transport: transport,
		log:       log,
		Project:   project.NewProjectSvc(transport, log),
		Robot:     robot.NewRobotSvc(transport, log),
	}, nil
}
