package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/kelseyhightower/envconfig"
	"github.com/nats-io/nats"
	natsp "github.com/nats-io/nats/encoders/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/KanybekMomukeyev/grpc-realtime-chat/server/chat"
	manager "github.com/KanybekMomukeyev/grpc-realtime-chat/server/chat/cassandra"
	pb "github.com/KanybekMomukeyev/grpc-realtime-chat/server/chat/chatpb"
	"flag"
)

var (
	certFile   = flag.String("cert_file", "../certfiles/ssl.crt", "The TLS cert file")
	keyFile    = flag.String("key_file", "../certfiles/ssl.key", "The TLS key file")
	jwtPrivateKey = flag.String("jwt_key_file", "../certfiles/jwt.pem", "The TLS key file")
)

type config struct {
	DBCluster       []string `default:"127.0.0.1"`
	DBKeyspace      string `default:"chat"`
	NatsAddress     string `default:"nats://nats:4222"`
	//TLSCert         string `default:"/etc/auth/cert.pem"`
	//TLSKey          string `default:"/etc/auth/key.pem"`
	//JWTPrivateKey   string `default:"/etc/auth/jwt-key.pem"`
	MaxAttempts     int    `default:"20"`
	ListenAddr      string `default:"127.0.0.1:7400"`
	DebugListenAddr string `default:"127.0.0.1:7401"`
}

func main() {
	var c config
	err := envconfig.Process("MESSAGE", &c)
	if err != nil {
		log.Fatalf("[CRITICAL][message-server] Could not process the config enviromment: %v \n", err)
	}

	cluster := gocql.NewCluster(c.DBCluster...)
	cluster.Keyspace = c.DBKeyspace

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatalf("[CRITICAL][message-server] Error connecting to cassandra: %s", err)
	}
	defer session.Close()

	cM := manager.ConversationManager{Session: session}
	mM := manager.MessageManager{Session: session}

	natsConnect, _ := nats.Connect(c.NatsAddress)
	encodedConn, _ := nats.NewEncodedConn(natsConnect, natsp.PROTOBUF_ENCODER)
	defer ec.Close()

	ta, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	//ta, err := credentials.NewServerTLSFromFile(c.TLSCert, c.TLSKey)
	if err != nil {
		log.Fatalf("[CRITICAL][message-server] %v", err)
	}
	gs := grpc.NewServer(grpc.Creds(ta))

	key, err := ioutil.ReadFile(*jwtPrivateKey)
	//key, err := ioutil.ReadFile(c.JWTPrivateKey)
	if err != nil {
		log.Fatalf("[CRITICAL][message-server] Error reading the jwt private key: %s", err)
	}
	cS, err := chat.NewMessageService(key, mM, cM, encodedConn)
	if err != nil {
		log.Fatalf("[CRITICAL][message-server] %v", err)
	}
	pb.RegisterMessageServiceServer(gs, cS)

	ln, err := net.Listen("tcp", c.ListenAddr)
	if err != nil {
		log.Fatalf("[CRITICAL][message-server] %v", err)
	}
	go gs.Serve(ln)

	log.Println("[INFO][message-server] Conversation service started successfully.")
	log.Fatal(http.ListenAndServe(c.DebugListenAddr, nil))
}
