package bootstrap

import (
	"fmt"
	"log"
	"net/http"
	"os"
	
	"github.com/olivere/elastic/v7"
	
	"github.com/devhg/es/config"
	"github.com/devhg/es/httpapi"
	"github.com/devhg/es/resource"
)

func MustInit(conf *config.ServerConfig) http.Handler {
	
	resource.EsClient = NewEsClient(conf)
	
	router := httpapi.NewHTTPRouter()
	
	return router
}

func NewEsClient(conf *config.ServerConfig) *elastic.Client {
	url := fmt.Sprintf("http://%s:%d", conf.Elastic.Host, conf.Elastic.Port)
	client, err := elastic.NewClient(
		elastic.SetSniff(false), // docker需要 https://blog.csdn.net/finghting321/article/details/105991741
		elastic.SetURL(url),     // elastic 服务地址
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ERR ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "ELASTIC INFO ", log.LstdFlags)),
	)
	
	if err != nil {
		log.Fatalln("Failed to create elastic client")
	}
	return client
}
