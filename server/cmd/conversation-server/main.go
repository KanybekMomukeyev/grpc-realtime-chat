package main

import (
	"database/sql"
	"io/ioutil"
	"log"
	"net"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/KanybekMomukeyev/grpc-realtime-chat/server/chat"
	manager "github.com/KanybekMomukeyev/grpc-realtime-chat/server/chat/cassandra"
	pb "github.com/KanybekMomukeyev/grpc-realtime-chat/server/chat/chatpb"
	"flag"
)

var db *sql.DB

var (
	certFile   = flag.String("cert_file", "../certfiles/ssl.crt", "The TLS cert file")
	keyFile    = flag.String("key_file", "../certfiles/ssl.key", "The TLS key file")
	jwtPrivateKey = flag.String("jwt_key_file", "../certfiles/jwt-key.pem", "The TLS key file")
)

type config struct {
	DBCluster       []string
	DBKeyspace      string `default:"chat"`
	//TLSCert         string `default:"/etc/auth/cert.pem"`
	//TLSKey          string `default:"/etc/auth/key.pem"`
	//JWTPrivateKey   string `default:"/etc/auth/jwt-key.pem"`
	MaxAttempts     int    `default:"20"`
	ListenAddr      string `default:"127.0.0.1:7400"`
	DebugListenAddr string `default:"127.0.0.1:7401"`
}

func main() {
	var c config
	err := envconfig.Process("CONVERSATION", &c)
	if err != nil {
		log.Fatalf("[CRITICAL][conversation-server] Could not process the config enviromment: %v \n", err)
	}

	cluster := gocql.NewCluster(c.DBCluster...)
	cluster.Keyspace = c.DBKeyspace

	session, _ := cluster.CreateSession()
	defer session.Close()

	cM := manager.ConversationManager{session}

	ta, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	//ta, err := credentials.NewServerTLSFromFile(c.TLSCert, c.TLSKey)
	if err != nil {
		log.Fatalf("[CRITICAL][conversation-server] %v", err)
	}
	gs := grpc.NewServer(grpc.Creds(ta))

	//key, err := ioutil.ReadFile(c.JWTPrivateKey)
	key, err := ioutil.ReadFile(*jwtPrivateKey)

	if err != nil {
		log.Fatalf("[CRITICAL][conversation-server] Error reading the jwt private key: %s", err)
	}
	cS, err := chat.NewConversationService(key, cM)
	pb.RegisterConversationServiceServer(gs, cS)

	ln, err := net.Listen("tcp", c.ListenAddr)
	if err != nil {
		log.Fatalf("[CRITICAL][conversation-server] %v", err)
	}
	go gs.Serve(ln)

	log.Println("[INFO][conversation-server] Conversation service started successfully.")
	log.Fatal(http.ListenAndServe(c.DebugListenAddr, nil))
}
