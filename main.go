package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"github.com/asim/go-micro/plugins/server/grpc/v4"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"io/ioutil"
	"productCatalog/common"
	_ "productCatalog/docs"
	"productCatalog/ent"
	"productCatalog/global"
	grpcHandler "productCatalog/grpc"
	"productCatalog/handler"
	"productCatalog/middleware"
	pb "productCatalog/proto"
	productCatalogService "productCatalog/service"
	"productCatalog/subscriber"
	"strconv"
	"time"

	"github.com/IBM/sarama"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
	"go-micro.dev/v4/server"

	// Enable below line if you need to overwrite default settings for hystrix-go
	// hystrixGo "github.com/afex/hystrix-go/hystrix"
	//"github.com/asim/go-micro/plugins/registry/consul/v4"
	// circuit breaker example used for Client wrapper only
	//"github.com/asim/go-micro/plugins/wrapper/breaker/hystrix/v4"
	limiter "github.com/asim/go-micro/plugins/wrapper/ratelimiter/uber/v4"
	opentracing4 "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"go-micro.dev/v4/cache"
	log "go-micro.dev/v4/logger"
)

const (
	version     = "latest"
	address     = "0.0.0.0:8084"
	address_rpc = "0.0.0.0:8085"
	QPS         = 100 // QPS for rate limiter
)

// @title LIS Billing productCatalog Swagger API

// @description This Billing productCatalog service consists of four parts: Bundle, Item price and Promotion. For details https://vibrantamerica.atlassian.net/wiki/spaces/LIS/pages/328695927/Billing+Service
// @termsOfService http://swagger.io/terms/

// @contact.name LIS Support
// @version 1.0
// @host 192.168.10.212
// @BasePath /v1/productCatalog
// @schemes http
func main() {
	// Initial Logger
	log.DefaultLogger = common.InitLogger()
	// Initialize kubernetes environment variables
	global.InitServiceAddress()
	common.InitKuberEnv()

	serviceName := "go.micro.lis.service.productCatalog"
	if common.KuberEnv.RunEnv == "Staging" || common.KuberEnv.RunEnv == "Akstaging" {
		serviceName = "go.micro.lis.service.productCatalog.staging"
	}

	var err error
	var capiClient *api.Client
	if common.KuberEnv.RunEnv == "Aks" || common.KuberEnv.RunEnv == "Akstaging" {
		log.Info("init aks consul config")
		capiConfig := api.Config{
			Scheme:  "https",
			Address: common.KuberEnv.ConsulConfigAddr + ":8501",
			Token:   common.KuberEnv.ConsulToken,
			TLSConfig: api.TLSConfig{
				CAFile:             "etc/consul-ca-cert.pem",
				InsecureSkipVerify: true,
			},
		}
		capiClient, err = api.NewClient(&capiConfig)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Info("init production consul config")
		capiConfig := api.Config{
			Scheme:  "http",
			Address: common.KuberEnv.ConsulConfigAddr + ":8500",
			Token:   common.KuberEnv.ConsulToken,
		}
		capiClient, err = api.NewClient(&capiConfig)
		if err != nil {
			log.Fatal(err)
		}
	}

	var redisClient *redis.Client

	// ! NOTE: There's unknown bug now that if this is init later, then can't retrieve anything from Consul. Need to be done right after consleConfig init.
	// ! My guess is something to do with shared context would be overwritten in some other init stages.

	common.CapiGetSecrets(capiClient, common.KuberEnv.ConsulPrefix+"/secrets")

	// DB client
	var mysqlInfo *common.MySQLConfig
	if common.KuberEnv.RunEnv == "Aks" || common.KuberEnv.RunEnv == "Production" {
		mysqlInfo = common.CapiGetMySqlConfig(capiClient, common.KuberEnv.ConsulPrefix+"/mysql")
	} else {
		mysqlInfo = common.CapiGetMySqlConfig(capiClient, common.KuberEnv.ConsulPrefix+"/mysqlDev")
	}

	// Init Redis Cache
	redisClient = common.CapiGetRedisConfig(capiClient, common.KuberEnv.ConsulPrefix+"/redisSentinel")
	defer func(redisClient *redis.Client) {
		_ = redisClient.Close()
	}(redisClient)

	//Registry center
	//consulRegistry := consul.NewRegistry(
	//	consul.Config(&api.Config{
	//		Address: consulConfigAddr + ":8500",
	//		Token:   common.ConsulToken,
	//	}),
	//)

	if common.KuberEnv.RunEnv == "Aks" || common.KuberEnv.RunEnv == "Production" {
		var pem []byte
		//use ssl on the client side for database connection
		rootCertPool := x509.NewCertPool()
		pem, err = ioutil.ReadFile("etc/DigiCertGlobalRootCA.crt.pem")
		if err != nil {
			log.Fatal(err)
		}
		if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
			log.Fatal("Failed to append PEM.")
		}
		//clientCert := make([]tls.Certificate, 0, 1)
		//certs, err := tls.LoadX509KeyPair("/path/client-cert.pem",
		//	"/path/client-key.pem")
		//if err != nil {
		//	log.Fatal(err)
		//}

		//clientCert = append(clientCert, certs)
		mysql.RegisterTLSConfig("custom", &tls.Config{
			RootCAs: rootCertPool,
			//Certificates: clientCert,
			ServerName: "lisportalprod2.mysql.database.azure.com", // hostname
		})
	}

	var client *ent.Client
	if common.KuberEnv.RunEnv != "Akstaging" {
		if common.KuberEnv.RunEnv == "Aks" || common.KuberEnv.RunEnv == "Production" {
			client, err = ent.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host+")/"+mysqlInfo.Database+"?parseTime=true&tls=custom")
		} else {
			client, err = ent.Open("mysql", mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host+":"+strconv.Itoa(mysqlInfo.Port)+")/"+mysqlInfo.Database+"?parseTime=true")
		}
		if err != nil {
			log.Fatalf("failed opening connection to MySQL: %v", err)
		}
		defer func(client *ent.Client) {
			_ = client.Close()
		}(client)

		log.Info("MYSQL CMD:", mysqlInfo.User+":"+mysqlInfo.Pwd+"@tcp("+mysqlInfo.Host+":"+strconv.Itoa(mysqlInfo.Port)+")/"+mysqlInfo.Database)

		// Run the auto migration tool.
		if mysqlInfo.AutoMigration {
			err = client.Schema.Create(context.Background())
			if err != nil {
				log.Fatalf("failed creating schema resources: %v", err)
			}
		}
	}
	var secDbclient, thirDbclient *sql.DB
	if common.KuberEnv.RunEnv == "Staging" || common.KuberEnv.RunEnv == "Production" {
		// mysql db client for shipping fee service
		secDbclient, err = sql.Open("mysql", mysqlInfo.SecondaryUser+":"+mysqlInfo.SecondaryPwd+"@tcp("+mysqlInfo.SecondaryHost+":"+strconv.Itoa(mysqlInfo.SecondaryPort)+")/"+mysqlInfo.SecondaryDatabase+"?parseTime=true")
		if err != nil {
			log.Fatalf("Failed to open second db client: %v", err)
		}
		thirDbclient, err = sql.Open("mysql", mysqlInfo.ThirdUser+":"+mysqlInfo.ThirdPwd+"@tcp("+mysqlInfo.ThirdHost+":"+strconv.Itoa(mysqlInfo.ThirdPort)+")/"+mysqlInfo.ThirdDatabase+"?parseTime=true")
		if err != nil {
			log.Fatalf("Failed to open third db client: %v", err)
		}
		defer func(client1 *sql.DB, client2 *sql.DB) {
			_ = client1.Close()
			_ = client2.Close()
		}(secDbclient, thirDbclient)
		// See "Important settings" section.
		secDbclient.SetConnMaxLifetime(time.Minute * 3)
		secDbclient.SetMaxOpenConns(10)
		secDbclient.SetMaxIdleConns(10)
		thirDbclient.SetConnMaxLifetime(time.Minute * 3)
		thirDbclient.SetMaxOpenConns(10)
		thirDbclient.SetMaxIdleConns(10)
	}

	// Connect to kafka broker
	producerConfig := sarama.NewConfig()
	producerConfig.Producer.Return.Successes = true
	//producerConfig.Version = sarama.V2_0_0_0

	consumerConfig := sarama.NewConfig()
	consumerConfig.Consumer.Return.Errors = true
	//consumerConfig.Version = sarama.V2_0_0_0
	var brokers = []string{}
	if common.KuberEnv.RunEnv == "Aks" || common.KuberEnv.RunEnv == "Akstaging" {
		kafkaInfo := common.CapiGetKafkaConfig(capiClient, common.KuberEnv.ConsulPrefix+"/kafkaConfig")
		brokers = kafkaInfo.Brokers

		producerConfig.Net.SASL.Enable = true
		producerConfig.Net.SASL.Mechanism = "SCRAM-SHA-256"
		producerConfig.Net.SASL.User = kafkaInfo.User
		producerConfig.Net.SASL.Password = kafkaInfo.Pwd
		producerConfig.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
			return &subscriber.XDGSCRAMClient{HashGeneratorFcn: subscriber.SHA256}
		}

		consumerConfig.Net.SASL.Enable = true
		consumerConfig.Net.SASL.Mechanism = "SCRAM-SHA-256"
		consumerConfig.Net.SASL.User = kafkaInfo.User
		consumerConfig.Net.SASL.Password = kafkaInfo.Pwd
		consumerConfig.Net.SASL.SCRAMClientGeneratorFunc = func() sarama.SCRAMClient {
			return &subscriber.XDGSCRAMClient{HashGeneratorFcn: subscriber.SHA256}
		}
	} else if common.KuberEnv.RunEnv == "Production" || common.KuberEnv.RunEnv == "Staging" {
		brokers = []string{"192.168.60.9:9095", "192.168.60.10:9095"}
	}

	producer, err := sarama.NewSyncProducer(brokers, producerConfig)
	if err != nil {
		log.Fatalf("Failed to start Sarama producer: %v", err)
	}

	consumerGroup, err := sarama.NewConsumerGroup(brokers, "lis-product-catalog", consumerConfig)
	if err != nil {
		log.Fatalf("Failed to start consumer group: %v", err)
	}

	//for grpc healthcheck
	//http.HandleFunc("/healthcheck", func(http.ResponseWriter, *http.Request) {})
	//go log.Fatal(http.ListenAndServe(":8084", nil))

	// Create microservice
	service := micro.NewService(
		// enable for grpc server
		micro.Server(grpc.NewServer()),
		//http server
		//micro.Server(srv),
		micro.Name(serviceName),
		micro.Version(version),
		micro.Address(address_rpc),
		//micro.Registry(consulRegistry),
		// WrapHandler is used to wrap inbound requests, aka server-side wrapping,
		// the other one called WrapClient is used to wrap outbound requests, aka client-side wrapping
		micro.WrapHandler(opentracing4.NewHandlerWrapper(opentracing.GlobalTracer())),
		micro.WrapHandler(limiter.NewHandlerWrapper(QPS)),
		micro.WrapHandler(middleware.LogHandler),
		// An example of typical client wrapper hystrix for circuit breaker
		// micro.WrapClient(hystrix.NewClientWrapper()),
		//micro.Broker(kafka.NewBroker(broker.Addrs(BrokerConfig...))),
		micro.Cache(cache.NewCache()),
	)
	/* Example of overwriting default hystrix settings
	hystrixGo.DefaultMaxConcurrent = 100
	hystrixGo.DefaultTimeout = 200
	*/

	// Create tls certificate
	// serverCert, err := tls.LoadX509KeyPair("cert/go.ehr.automation.crt", "cert/go.ehr.automation.key")
	//if err != nil {
	//	log.Errorf(err.Error())
	//}
	//service.Init() // Parse CLI flags

	// Register subscriber
	// micro.RegisterSubscriber("audit_logs", srv, subscriber.SubEv, server.SubscriberQueue("go.micro.audit_logs"))

	//Register rpc handler
	gRPCserver := service.Server()
	err = gRPCserver.Init(server.Wait(nil))
	if err != nil {
		log.Fatal(err)
	}
	// server.TLSConfig(&tls.Config{
	// 	Certificates: []tls.Certificate{serverCert},
	// 	ClientAuth:   tls.NoClientCert,
	// }

	// Register subscriber
	//micro.RegisterSubscriber("lis_logs", gRPCserver, subscriber.SubEv, server.SubscriberQueue("go.micro.lis"))

	// Register handler
	err = pb.RegisterProductCatalogServiceHandler(gRPCserver, &grpcHandler.ProductCatalogRpcHandler{
		TaxService: productCatalogService.NewStripeService(client, redisClient, common.StripeKeyZb),
	})
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.Use(middleware.AuthMiddleware(true))
	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	router.Use(middleware.Ginzap(common.ZapLogger, "", true))
	// Logs all panic to error log
	//   - stack means whether output the stack info.
	router.Use(middleware.RecoveryWithZap(common.ZapLogger, true))

	productCatalogHandler := handler.ProductCatalog{
		MerchandiseService: productCatalogService.NewMerchandiseService(client, redisClient),
		Producer:           productCatalogService.NewProducerService(producer),
		Consumer:           handler.NewConsumer(consumerGroup),
	}
	productCatalogHandler.MerchandiseRegister(router.Group("/manage/merchandise"))

	testAuth := router.Group("/")
	testAuth.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pilipilipalapala",
		})
	})

	//log.Info("StartConsumer")
	go productCatalogHandler.StartConsumer()

	//Setup HTTP server using Gin in a separate goroutine
	go func() {
		if err = router.Run(address); err != nil {
			log.Fatalf("Failed to run HTTP server: %v", err)
		}
	}()

	// Run service
	if err = service.Run(); err != nil {
		log.Fatal(err)
	}

}
