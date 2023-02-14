package firewall

import (
	"errors"
	"net"

	netallow "github.com/kisom/netallow"
)

var (
	ErrCanNotParseIP = errors.New("can not parse IP")
	//ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

type Firewall struct {
	Wl, Bl *netallow.BasicNet
}

func New() *Firewall {
	return &Firewall{
		Wl: NewBasicNet(),
		Bl: NewBasicNet(),
	}
}

type Status int

const (
	StatusUnspecified Status = iota
	StatusNoIpInLists
	StatusExistInBl
	StatusExistInWl
	StatusExistInBlAndWl
)

// String - Creating common behavior - give the type a String function
func (s Status) String() string {
	return [...]string{"STATUS_UNSPECIFIED", "STATUS_NO_IP_IN_LISTS", "STATUS_EXIST_IN_BL", "STATUS_EXIST_IN_WL", "STATUS_EXIST_IN_BL_AND_WL", "Friday", "Saturday"}[s]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (s Status) EnumIndex() int {
	return int(s)
}

func NewBasicNet() *netallow.BasicNet {
	return netallow.NewBasicNet()
}

func AddByStr(list netallow.BasicNet, strIPNet string) error {
	_, curIPNet, err := net.ParseCIDR(strIPNet)
	if err != nil {
		return err
	}
	list.Add(curIPNet)
	return nil
}

func DelByStr(list netallow.BasicNet, strIPNet string) error {
	_, curIPNet, err := net.ParseCIDR(strIPNet)
	if err != nil {
		return err
	}
	list.Remove(curIPNet)
	return nil
}

func (fw *Firewall) CheckIP(strIP string) (status Status, ip net.IP, err error) {
	curIP := net.ParseIP(strIP)
	if curIP == nil {
		return StatusUnspecified, nil, ErrCanNotParseIP
	}
	inbl := fw.Bl.Permitted(curIP)
	inwl := fw.Wl.Permitted(curIP)
	if inbl {
		if inwl {
			return StatusExistInBlAndWl, ip, nil
		} else {
			return StatusExistInBl, ip, nil
		}
	} else {
		if inwl {
			return StatusExistInWl, ip, nil
		} else {
			return StatusNoIpInLists, ip, nil
		}
	}
}
