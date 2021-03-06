package consul_common

import (
	"github.com/spf13/pflag"
// 	"errors"
	"fmt"
	"os"
	"strings"
)

const (
	k_ENV_CONSUL = "CONSUL_HOST"
	k_ENV_DC     = "CONSUL_DC"

	k_CONSUL_URL = "localhost:8500"
	k_CONSUL_DC  = "dc1"
)


func SetupConsulFlags(ff *pflag.FlagSet, consulConf *ClientConfig, vv *int) error {
	
	ff.CountVarP(vv, "verbose", "v", "Enable (and/or increase) verbosity level")
	
	ff.StringVar(&consulConf.URL, "consul", "", "Consul HTTP API endpoint to use (default: $CONSUL_HOST)")
	ff.StringVar(&consulConf.Datacenter, "dc", "", "Consul datacenter identifier to use")

	var url, dc string
	if url = os.Getenv(k_ENV_CONSUL); "" == url {
		url = k_CONSUL_URL
	}
	if dc = os.Getenv(k_ENV_DC); "" == dc {
		dc = k_CONSUL_DC
	}
	consulConf.URL = url
	consulConf.Datacenter = dc
	
	// Extended flags: HTTP basic auth & SSL
// 	ff.StringVar(&consulConf.AuthUser, "http-user", "", "Consul HTTP basic auth Username")
// 	ff.StringVar(&consulConf.AuthPass, "http-pass", "", "Consul HTTP basic auth Password")

// 	ff.BoolVarP(&consulConf.directAPI, "passthrough", "z", false, "Enable direct operation against a Consul server, bypassing the (normally local) agent")

	return nil
}

func PrintConfig(consulConf *ClientConfig) {
	
	fmt.Println("direct->", consulConf.directAPI)
	fmt.Println("consul->", consulConf.URL)
	fmt.Println("datacenter->", consulConf.Datacenter)	
}

func NormalizePrefix(pp *string, delim string) {
	
	var w string = *pp
	var n int = strings.LastIndex(w,delim)
	for n > 1 {		// do-while( lastIndex > 1 )
		w = w[0:n]
		n = strings.LastIndex(w,delim)
	}
	*pp = w
	// do nothing :)
}
