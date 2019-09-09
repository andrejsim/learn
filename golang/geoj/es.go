package main

import (
	"context"
	"github.com/apex/log"
	"github.com/kelseyhightower/envconfig"
	"github.com/olivere/elastic"
	"github.com/sirupsen/logrus"
)

type configuration struct {
	ElasticSearchDomain string `required:"true"`
	ElasticSearchIndex  string `default:"testgeo"`
	ElasticSearchRBAC   bool   `default:"true"`
	Region              string `required:"true"`
	LogLevel            string `default:"info"`
}

var cfg configuration
var esClient *elastic.Client

func run() {
	err := envconfig.Process("geo", &cfg)
	if err != nil {
		log.WithField("err", err).Fatal("Could not load required config")
	}
	lvl, err := log.ParseLevel(cfg.LogLevel)
	if err == nil {
		log.SetLevel(lvl)
	}
	log.Info(cfg.ElasticSearchDomain)
	log.Info(cfg.ElasticSearchIndex)

	// Create an Elasticsearch client
	opts := []elastic.ClientOptionFunc{
		elastic.SetURL(cfg.ElasticSearchDomain),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(false),
	}

	l := logrus.WithField("source", "es")
	if cfg.LogLevel == "debug" {
		opts = append(opts, elastic.SetTraceLog(l))
	}

	//if cfg.ElasticSearchRBAC {
	//	// Setup request signing for role based access
	//	log.Debug("Using V4 signing for accessing elasticsearch")
	//	transport := signer.NewTransport(session.New(&aws.Config{Region: &cfg.Region}), elasticsearchservice.ServiceName)
	//	httpClient := &http.Client{
	//		Transport: transport,
	//	}
	//	opts = append(opts, elastic.SetHttpClient(httpClient))
	//}


	esClient, err = elastic.NewClient(opts...)
	if err != nil {
		log.Fatal(err.Error())
	}

	bulkRequest := esClient.Bulk()

	log.Info(esClient.String())

	for i := 0; i < 10; i++  {
		meta := funcSyntheticGeojson()
		req := elastic.NewBulkIndexRequest().
			Index(cfg.ElasticSearchIndex).
			Type("geo_shape").
			Id(string(i)).
			Doc(meta)

		bulkRequest.Add(req)

		log.WithField("request",req.String()).Debug("bulk add req")
	}

	ctx := context.Background()
	log.WithField("ctx", ctx).Info("backgrond context")

	bulkResponse, err := bulkRequest.Do(ctx)
	if err != nil {
		log.WithField("error",err.Error()).Info("bulk do error")
	}

	log.WithField("bulkResponse",len(bulkResponse.Items)).Info("done")

}