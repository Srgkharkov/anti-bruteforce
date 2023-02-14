package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/Srgkharkov/anti-bruteforce/internal/pkg/app"
	"github.com/Srgkharkov/anti-bruteforce/internal/pkg/grpc"
	"google.golang.org/grpc"
	//"google.golang.org/grpc/codes"
	//"google.golang.org/grpc/status"
	//"github.com/kisom/netallow"
)

// anti-bruteforce is used to implement helloworld.GreeterServer.
type Grpcserver struct {
	//grpcpb.UnimplementedAntiBruteforceServer
	grpcpb.AntiBruteforceServer
	port int
	app  *app.App
	//lis *bufconn.Listener
	//grpcpb.grpcpb
	//Wl *netallow.BasicNet
	//Bl *netallow.BasicNet
}

func New(port int, app *app.App) *Grpcserver {
	return &Grpcserver{port: port, app: app}
}

func (g *Grpcserver) Run() error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", g.port))
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	grpcpb.RegisterAntiBruteforceServer(s, g)

	log.Printf("anti-bruteforce listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}

// SayHello implements helloworld.GreeterServer
func (s *Grpcserver) AccessRequest(ctx context.Context, in *grpcpb.Request) (*grpcpb.Result, error) {
	//inWl := false
	//inBl := false
	s.app.AccessRequest(time.Now(), in.GetLogin(), in.GetPass(), in.GetIp())

	ip := net.ParseIP(in.GetIp())
	log.Printf("Bucket received: login %v, pass %v, IP %v", in.GetLogin(), in.GetPass(), ip)
	//if s.app.Firewall.Wl.Permitted(ip) {
	//	inWl = true
	//	log.Printf("The IP exist in WL")
	//}
	//if s.app.Firewall.Bl.Permitted(ip) {
	//	inBl = true
	//	log.Printf("The IP exist in BL")
	//}

	if s.app.Firewall.Wl.Permitted(ip) {
		if !s.app.Firewall.Bl.Permitted(ip) {
			return &grpcpb.Result{Status: grpcpb.Status_STATUS_TRUE}, nil
		}
	} else if s.app.Firewall.Bl.Permitted(ip) {
		return &grpcpb.Result{Status: grpcpb.Status_STATUS_FALSE}, nil
	}

	return &grpcpb.Result{Status: grpcpb.Status_STATUS_ERROR}, nil
}

func (s *Grpcserver) AddToWhileList(ctx context.Context, in *grpcpb.SubNet) (*grpcpb.Result, error) {
	_, curIPNet, err := net.ParseCIDR(in.GetSubnet())
	if err != nil {
		return &grpcpb.Result{Status: grpcpb.Status_STATUS_ERROR}, nil
	}
	s.app.Firewall.Wl.Add(curIPNet)
	log.Printf("Added the new IPNet to WhiteList: %v", curIPNet.String())
	return &grpcpb.Result{Status: grpcpb.Status_STATUS_TRUE}, nil
}

func (s *Grpcserver) DelFromWhileList(ctx context.Context, in *grpcpb.SubNet) (*grpcpb.Result, error) {
	_, curIPNet, err := net.ParseCIDR(in.GetSubnet())
	if err != nil {
		return &grpcpb.Result{Status: grpcpb.Status_STATUS_ERROR}, nil
	}
	s.app.Firewall.Wl.Remove(curIPNet)
	log.Printf("Removed the IPNet from WhiteList: %v", curIPNet.String())
	return &grpcpb.Result{Status: grpcpb.Status_STATUS_TRUE}, nil
}

func (s *Grpcserver) AddToBlackList(ctx context.Context, in *grpcpb.SubNet) (*grpcpb.Result, error) {
	_, curIPNet, err := net.ParseCIDR(in.GetSubnet())
	if err != nil {
		return &grpcpb.Result{Status: grpcpb.Status_STATUS_ERROR}, nil
	}
	s.app.Firewall.Bl.Add(curIPNet)
	log.Printf("Added the new IPNet to BlackList: %v", curIPNet.String())
	return &grpcpb.Result{Status: grpcpb.Status_STATUS_TRUE}, nil
}

func (s *Grpcserver) DelFromBlackList(ctx context.Context, in *grpcpb.SubNet) (*grpcpb.Result, error) {
	_, curIPNet, err := net.ParseCIDR(in.GetSubnet())
	if err != nil {
		return &grpcpb.Result{Status: grpcpb.Status_STATUS_ERROR}, nil
	}
	s.app.Firewall.Bl.Remove(curIPNet)
	log.Printf("Removed the IPNet from BlackList: %v", curIPNet.String())
	return &grpcpb.Result{Status: grpcpb.Status_STATUS_TRUE}, nil
}

func (s *Grpcserver) ClearBucket(ctx context.Context, in *grpcpb.Request) (*grpcpb.Result, error) {
	//_, curIPNet, err := net.ParseCIDR(in.GetSubnet())
	//if err != nil {
	//	return &grpcpb.Result{Status: grpcpb.Status_STATUS_ERROR}, nil
	//}
	//s.app.Firewall.Bl.Remove(curIPNet)
	//log.Printf("Removed the IPNet from BlackList: %v", curIPNet.String())
	return &grpcpb.Result{Status: grpcpb.Status_STATUS_TRUE}, nil
}
