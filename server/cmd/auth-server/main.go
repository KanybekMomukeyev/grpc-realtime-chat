package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/KanybekMomukeyev/grpc-realtime-chat/server/user"
	manager "github.com/KanybekMomukeyev/grpc-realtime-chat/server/user/mysql"
	pb "github.com/KanybekMomukeyev/grpc-realtime-chat/server/user/userpb"
	"flag"
)

var db *sql.DB

var (
	certFile   = flag.String("cert_file", "../certfiles/ssl.crt", "The TLS cert file")
	keyFile    = flag.String("key_file", "../certfiles/ssl.key", "The TLS key file")
	jwtPrivateKey = flag.String("jwt_key_file", "../certfiles/jwt-key.pem", "The TLS key file")
)

type config struct {
	DBHost          string
	DBPort          string `default:"3306"`
	DBUsername      string `default:"user-service"`
	DBPassword      string `default:"nazgulum"`
	DBName          string `default:"chat"`
	TLSCert         string `default:"/etc/auth/cert.pem"`
	TLSKey          string `default:"/etc/auth/key.pem"`
	JWTPrivateKey   string `default:"/etc/auth/jwt-key.pem"`
	MaxAttempts     int    `default:"20"`
	ListenAddr      string `default:"127.0.0.1:7300"`
	DebugListenAddr string `default:"127.0.0.1:7301"`
}

func main() {
	var myConfig config
	err := envconfig.Process("AUTH", &myConfig)
	if err != nil {
		log.Fatalf("[CRITICAL][auth-server] Could not process the config enviromment: %v", err)
	}
	log.Println("[INFO][auth-server] Auth service starting...")

	var dbError error
	// Connect to database.
	fmt.Print(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", myConfig.DBUsername, myConfig.DBPassword, myConfig.DBHost, myConfig.DBPort, myConfig.DBName))

	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", myConfig.DBUsername, myConfig.DBPassword, myConfig.DBHost, myConfig.DBPort, myConfig.DBName))
	if err != nil {
		log.Printf("[ERROR][auth-server] %v", err)
	}
	for attempts := 1; attempts < myConfig.MaxAttempts; attempts++ {
		dbError = db.Ping()
		if dbError == nil {
			break
		}
		log.Printf("[ERROR][auth-server] %v", dbError)
		time.Sleep(time.Duration(attempts) * time.Second)
	}
	if dbError != nil {
		log.Fatalf("[CRITICAL][auth-server] Could not Connect to dababase: %v \n", dbError)
	}
	uM := manager.UserManager{db}


	ta, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
	//ta, err := credentials.NewServerTLSFromFile(myConfig.TLSCert, myConfig.TLSKey)
	if err != nil {
		log.Fatalf("[CRITICAL][auth-server] %v", err)
	}
	gs := grpc.NewServer(grpc.Creds(ta))

	//key, err := ioutil.ReadFile(myConfig.JWTPrivateKey)
	key, err := ioutil.ReadFile(*jwtPrivateKey)
	if err != nil {
		log.Fatal(fmt.Errorf("Error reading the jwt private key: %s", err))
	}
	as, err := user.NewAuthServer(key, uM)
	if err != nil {
		log.Fatalf("[CRITICAL][auth-server] %v", err)
	}
	pb.RegisterAuthServiceServer(gs, as)

	ln, err := net.Listen("tcp", myConfig.ListenAddr)
	if err != nil {
		log.Fatalf("[CRITICAL][auth-server] %v", err)
	}
	go gs.Serve(ln)

	log.Println("Auth service started successfully.")
	log.Fatal(http.ListenAndServe(myConfig.DebugListenAddr, nil))
}
