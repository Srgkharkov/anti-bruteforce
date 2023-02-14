package main

import (
	"flag"
	"log"

	grpcserver "github.com/Srgkharkov/anti-bruteforce/internal/app/grpcserver"
	app2 "github.com/Srgkharkov/anti-bruteforce/internal/pkg/app"
)

var (
	grpcport = flag.Int("port", 50051, "The anti-bruteforce port")
	n        = flag.Int("n", 10, "Try per minute for unique login")
	m        = flag.Int("m", 100, "Try per minute for unique pass")
	k        = flag.Int("k", 1000, "Try per minute for unique ip")
)

//var wl = netallow.NewBasicNet()
//var bl = netallow.NewBasicNet()

//type bucket struct {
//	time  time.Time
//	login string
//	pass  string
//	ip    string
//}

//type whiteListIPNets struct {
//	IPNets     []net.IPNet
//	mapByIPNet map[string]uint64
//}

func main() {
	flag.Parse()
	app, err := app2.New(*n, *m, *k)
	if err != nil {
		log.Fatal(err)
	}
	grpcserver := grpcserver.New(*grpcport, app)
	err = grpcserver.Run()
	if err != nil {
		log.Fatal(err)
	}
	//s := grpc.NewServer()
	//srv := handler.CreateBruteforceServer(*port, *n, *m, *k)
	//grpcpb.RegisterAntiBruteforceServer(s, &srv)
	//grpcpb.RegisterAntiBruteforceServer(s, grpcserver)
	//log.Printf("anti-bruteforce listening at %v", lis.Addr())
	//if err := s.Serve(lis); err != nil {
	//	log.Fatalf("failed to serve: %v", err)
	//}
	//grpcserver.
}
