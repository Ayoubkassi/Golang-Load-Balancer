package main

import (
	"flag"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	port = flag.Int("port", 8080, "where to start farely")
)

type Service struct {
	Name     string
	Replicas []string
}

//config is a representation of the configuration
//given to farely from a config source
type Config struct {
	Services []Service

	//Name of the strategy to be used in load balncing between instances
	Strategy string
}

type Farely struct {
	Config *Config
}

func (f *Farely) ServerHtpp(res http.ResponseWriter, req *http.Request) {

}

func main() {
	flag.Parse()

	conf := &Config{}

	farely := &Farely{
		Config: conf,
	}

	server := http.Server{
		Addr:    "",
		Handler: farely,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
