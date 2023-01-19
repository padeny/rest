package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/rest-go/auth"
	"github.com/rest-go/rest/pkg/log"
	"github.com/rest-go/rest/pkg/server"
)

func parseConfig() *Config {
	addr := flag.String("addr", ":3000", "listen addr")
	url := flag.String("db.url", "", "db url")
	cfgPath := flag.String("config", "", "path to config file")
	flag.Parse()

	cfg := &Config{}
	if *cfgPath != "" {
		var err error
		cfg, err = NewConfig(*cfgPath)
		if err != nil {
			log.Fatal(err)
		}
	}
	if cfg.Addr == "" {
		cfg.Addr = *addr
	}
	if *url != "" {
		cfg.DB.URL = *url
	}
	return cfg
}

func main() {
	cfg := parseConfig()
	restServer := server.New(&cfg.DB, server.EnableAuth(cfg.Auth.Enabled))
	if cfg.Auth.Enabled {
		log.Info("auth is enabled")
		restAuth, err := auth.New(cfg.DB.URL, []byte(cfg.Auth.Secret))
		if err != nil {
			log.Fatal("initialize auth error ", err)
		}
		http.Handle("/auth/", restAuth)
		http.Handle("/", restAuth.Middleware(restServer))
	} else {
		http.Handle("/", restServer)
	}

	s := &http.Server{
		Addr:              cfg.Addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	log.Info("listen on addr: ", cfg.Addr)
	log.Fatal(s.ListenAndServe())
}
