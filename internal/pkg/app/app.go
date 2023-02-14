package app

import (
	"net"
	"time"

	"github.com/Srgkharkov/anti-bruteforce/internal/pkg/bucket"
	"github.com/Srgkharkov/anti-bruteforce/internal/pkg/firewall"
)

type App struct {
	Port, N, M, K int
	Buckets       *bucket.Buckets
	Firewall      *firewall.Firewall
}

func New(N, M, K int) (*App, error) {
	return &App{
		N:        N,
		M:        M,
		K:        K,
		Firewall: firewall.New(),
		Buckets:  bucket.NewBuckets(),
	}, nil
}

func (app *App) AccessRequest(time time.Time, login string, pass string, ipstr string) bool {
	ip := net.ParseIP(ipstr)
	if app.Firewall.Wl.Permitted(ip) {
		if !app.Firewall.Bl.Permitted(ip) {
			return true
		}
	} else if app.Firewall.Bl.Permitted(ip) {
		return false
	}

	//todo
	return true
}

func (app *App) AddToWhileList(subnet string) error {
	_, curIPNet, err := net.ParseCIDR(subnet)
	if err != nil {
		return err
	}

	app.Firewall.Wl.Add(curIPNet)
	return nil
}

func (app *App) DelFromWhileList(subnet string) error {
	_, curIPNet, err := net.ParseCIDR(subnet)
	if err != nil {
		return err
	}

	app.Firewall.Wl.Remove(curIPNet)
	return nil
}

func (app *App) AddToBlackList(subnet string) error {
	_, curIPNet, err := net.ParseCIDR(subnet)
	if err != nil {
		return err
	}

	app.Firewall.Bl.Add(curIPNet)
	return nil
}

func (app *App) DelFromBlackList(subnet string) error {
	_, curIPNet, err := net.ParseCIDR(subnet)
	if err != nil {
		return err
	}

	app.Firewall.Bl.Remove(curIPNet)
	return nil
}

func (app *App) ClearBucket(login string, pass string, ipstr string) {
	//todo
	return
}
