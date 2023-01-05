// Copyright 2022 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package types

import (
	"time"
)

// Config stores the system configuration.
type Config struct {
	// InstanceID specifis the ID of the gitness instance.
	// NOTE: If the value is not provided the hostname of the machine is used.
	InstanceID string `envconfig:"GITNESS_INSTANCE_ID"`
	Debug      bool   `envconfig:"GITNESS_DEBUG"`
	Trace      bool   `envconfig:"GITNESS_TRACE"`

	// Git defines the git configuration parameters
	Git struct {
		BaseURL        string `envconfig:"GITNESS_GIT_BASE_URL" default:"http://localhost:3000"` // clone url
		Root           string `envconfig:"GITNESS_GIT_ROOT"`
		TmpDir         string `envconfig:"GITNESS_GIT_TMP_DIR"`          // directory for temporary data (repo clone)
		ServerHookPath string `envconfig:"GITNESS_GIT_SERVER_HOOK_PATH"` // path to binary used as git server hook
		DefaultBranch  string `envconfig:"GITNESS_GIT_DEFAULTBRANCH" default:"main"`
	}

	// Server defines the server configuration parameters.
	Server struct {
		// HTTP defines the http configuration parameters
		HTTP struct {
			Bind  string `envconfig:"GITNESS_HTTP_BIND" default:":3000"`
			Proto string `envconfig:"GITNESS_HTTP_PROTO"`
			Host  string `envconfig:"GITNESS_HTTP_HOST"`
		}

		// GRPC defines the grpc configuration parameters
		GRPC struct {
			Bind string `envconfig:"GITNESS_GRPC_BIND" default:":3001"`
		}

		// Acme defines Acme configuration parameters.
		Acme struct {
			Enabled bool   `envconfig:"GITNESS_ACME_ENABLED"`
			Endpont string `envconfig:"GITNESS_ACME_ENDPOINT"`
			Email   bool   `envconfig:"GITNESS_ACME_EMAIL"`
		}
	}

	// Database defines the database configuration parameters.
	Database struct {
		Driver     string `envconfig:"GITNESS_DATABASE_DRIVER" default:"sqlite3"`
		Datasource string `envconfig:"GITNESS_DATABASE_DATASOURCE" default:"database.sqlite3"`
	}

	// Token defines token configuration parameters.
	Token struct {
		Expire time.Duration `envconfig:"GITNESS_TOKEN_EXPIRE" default:"720h"`
	}

	// Cors defines http cors parameters
	Cors struct {
		AllowedOrigins   []string `envconfig:"GITNESS_CORS_ALLOWED_ORIGINS"   default:"*"`
		AllowedMethods   []string `envconfig:"GITNESS_CORS_ALLOWED_METHODS"   default:"GET,POST,PATCH,PUT,DELETE,OPTIONS"`
		AllowedHeaders   []string `envconfig:"GITNESS_CORS_ALLOWED_HEADERS"   default:"Origin,Accept,Accept-Language,Authorization,Content-Type,Content-Language,X-Requested-With,X-Request-Id"` //nolint:lll // struct tags can't be multiline
		ExposedHeaders   []string `envconfig:"GITNESS_CORS_EXPOSED_HEADERS"   default:"Link"`
		AllowCredentials bool     `envconfig:"GITNESS_CORS_ALLOW_CREDENTIALS" default:"true"`
		MaxAge           int      `envconfig:"GITNESS_CORS_MAX_AGE"           default:"300"`
	}

	// Secure defines http security parameters.
	Secure struct {
		AllowedHosts          []string          `envconfig:"GITNESS_HTTP_ALLOWED_HOSTS"`
		HostsProxyHeaders     []string          `envconfig:"GITNESS_HTTP_PROXY_HEADERS"`
		SSLRedirect           bool              `envconfig:"GITNESS_HTTP_SSL_REDIRECT"`
		SSLTemporaryRedirect  bool              `envconfig:"GITNESS_HTTP_SSL_TEMPORARY_REDIRECT"`
		SSLHost               string            `envconfig:"GITNESS_HTTP_SSL_HOST"`
		SSLProxyHeaders       map[string]string `envconfig:"GITNESS_HTTP_SSL_PROXY_HEADERS"`
		STSSeconds            int64             `envconfig:"GITNESS_HTTP_STS_SECONDS"`
		STSIncludeSubdomains  bool              `envconfig:"GITNESS_HTTP_STS_INCLUDE_SUBDOMAINS"`
		STSPreload            bool              `envconfig:"GITNESS_HTTP_STS_PRELOAD"`
		ForceSTSHeader        bool              `envconfig:"GITNESS_HTTP_STS_FORCE_HEADER"`
		BrowserXSSFilter      bool              `envconfig:"GITNESS_HTTP_BROWSER_XSS_FILTER"    default:"true"`
		FrameDeny             bool              `envconfig:"GITNESS_HTTP_FRAME_DENY"            default:"true"`
		ContentTypeNosniff    bool              `envconfig:"GITNESS_HTTP_CONTENT_TYPE_NO_SNIFF"`
		ContentSecurityPolicy string            `envconfig:"GITNESS_HTTP_CONTENT_SECURITY_POLICY"`
		ReferrerPolicy        string            `envconfig:"GITNESS_HTTP_REFERRER_POLICY"`
	}

	// Admin defines admin user params (no admin setup if either is empty)
	Admin struct {
		DisplayName string `envconfig:"GITNESS_ADMIN_DISPLAY_NAME"`
		Email       string `envconfig:"GITNESS_ADMIN_EMAIL"`
		Password    string `envconfig:"GITNESS_ADMIN_PASSWORD"`
	}

	Redis struct {
		Endpoint           string `envconfig:"GITNESS_REDIS_ENDPOINT"             default:"localhost:6379"`
		MaxRetries         int    `envconfig:"GITNESS_REDIS_MAX_RETRIES"          default:"3"`
		MinIdleConnections int    `envconfig:"GITNESS_REDIS_MIN_IDLE_CONNECTIONS" default:"0"`
		Password           string `envconfig:"GITNESS_REDIS_PASSWORD"`
	}

	Events struct {
		Mode                  string `envconfig:"GITNESS_EVENTS_MODE"                     default:"inmemory"`
		Namespace             string `envconfig:"GITNESS_EVENTS_NAMESPACE"                default:"gitness"`
		MaxStreamLength       int64  `envconfig:"GITNESS_EVENTS_MAX_STREAM_LENGTH"        default:"1000"`
		ApproxMaxStreamLength bool   `envconfig:"GITNESS_EVENTS_APPROX_MAX_STREAM_LENGTH" default:"true"`
	}

	Webhook struct {
		MaxRetryCount       int64 `envconfig:"GITNESS_WEBHOOK_MAX_RETRY_COUNT" default:"3"`
		Concurrency         int   `envconfig:"GITNESS_WEBHOOK_CONCURRENCY" default:"4"`
		AllowLoopback       bool  `envconfig:"GITNESS_WEBHOOK_ALLOW_LOOPBACK" default:"false"`
		AllowPrivateNetwork bool  `envconfig:"GITNESS_WEBHOOK_ALLOW_PRIVATE_NETWORK" default:"false"`
	}
}
