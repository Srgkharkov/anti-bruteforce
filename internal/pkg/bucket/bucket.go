package bucket

import (
	"errors"
	"net"
	"sync"
	"time"
)

var (
	ErrCanNotParseIP = errors.New("can not parse IP")
	//ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

type Buckets struct {
	Buckets        *ProtectedBucketSlice
	BucketsByLogin *ProtectedBucketMap
	BucketsByPass  *ProtectedBucketMap
	BucketsByIP    *ProtectedBucketMap
}

type ProtectedBucketMap struct {
	mx sync.RWMutex
	m  map[string][]*Bucket
}

type ProtectedBucketSlice struct {
	mx sync.RWMutex
	s  []*Bucket
}

type Bucket struct {
	//id    uint64
	time  time.Time
	login string
	pass  string
	ip    net.IP
}

func NewBuckets() *Buckets {
	return &Buckets{
		BucketsByLogin: NewProtectedBucketMap(),
		BucketsByPass:  NewProtectedBucketMap(),
		BucketsByIP:    NewProtectedBucketMap(),
		Buckets:        NewProtectedBucketSlice(),
	}
}

func NewProtectedBucketMap() *ProtectedBucketMap {
	return &ProtectedBucketMap{m: make(map[string][]*Bucket)}
}

func NewProtectedBucketSlice() *ProtectedBucketSlice {
	return &ProtectedBucketSlice{}
}

func (s *ProtectedBucketSlice) Append(curBacket *Bucket) {
	s.mx.Lock()
	s.s = append(s.s, curBacket)
	s.mx.Unlock()

	return
}

func (m *ProtectedBucketMap) Load(key string) (val []*Bucket, ok bool) {
	m.mx.RLock()
	val, ok = m.m[key]
	m.mx.RUnlock()

	return
}

func (m *ProtectedBucketMap) Store(key string, value []*Bucket) {
	m.mx.Lock()
	m.m[key] = value
	m.mx.Unlock()
}

func (m *ProtectedBucketMap) Clear() {
	m.mx.Lock()
	//todo
	m.mx.Unlock()
}

func CreateBucket(time time.Time, login string, pass string, ipstr string) (*Bucket, error) {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return nil, ErrCanNotParseIP
	}
	return &Bucket{time: time, login: login, pass: pass, ip: ip}, nil
}

func (b *Bucket) GetLogin() string {
	return b.login
}

func (b *Bucket) GetPass() string {
	return b.pass
}

func (b *Bucket) GetIP() net.IP {
	return b.ip
}

func (b *Buckets) AppendBucket(curBacket *Bucket) *Bucket {
	//curBacket := Bucket{time: time.Now(), login: login, pass: pass, ip: ip}
	b.Buckets.Append(curBacket)
	//b.muBuckets.Lock()
	//b.Buckets = append(b.Buckets, curBacket)
	//b.muBuckets.Unlock()

	//b.muBucketsByLogin.Lock()
	s, _ := b.BucketsByLogin.Load(curBacket.login) //[curBacket.login]
	b.BucketsByLogin.Store(curBacket.login, append(s, curBacket))
	//if ok {
	//	//s = append(s, &curBacket)
	//	//b.BucketsByLogin[curBacket.login] = append(s, curBacket)
	//	b.BucketsByLogin.Store(curBacket.login, append(s, curBacket))
	//} else {
	//	//s = make([]*Bucket, 0, 1)
	//	//s = append(s, &curBacket)
	//	//b.BucketsByLogin[login] = append(make([]*Bucket, 0, 1), &curBacket)
	//	//b.BucketsByLogin[curBacket.login] = []*Bucket{curBacket}
	//	b.BucketsByLogin.Store(curBacket.login, append(s, curBacket))
	//	b.BucketsByLogin.Store(curBacket.login, []*Bucket{curBacket})
	//}
	//b.muBucketsByLogin.Unlock()

	s, _ = b.BucketsByPass.Load(curBacket.pass) //[curBacket.login]
	b.BucketsByPass.Store(curBacket.pass, append(s, curBacket))

	//b.muBucketsByPass.Lock()
	//s, ok = b.BucketsByPass[curBacket.pass]
	//if ok {
	//	//s = append(s, &curBacket)
	//	b.BucketsByPass[curBacket.pass] = append(s, curBacket)
	//} else {
	//	//s = make([]*Bucket, 0, 1)
	//	//s = append(s, &curBacket)
	//	b.BucketsByPass[curBacket.pass] = []*Bucket{curBacket}
	//}
	//b.muBucketsByPass.Unlock()

	ipstr := string(curBacket.ip.To16())
	s, _ = b.BucketsByIP.Load(ipstr) //[curBacket.login]
	b.BucketsByIP.Store(ipstr, append(s, curBacket))

	//b.muBucketsByIP.Lock()
	//s, ok = b.BucketsByIP[ipstr]
	//if ok {
	//	//s = append(s, &curBacket)
	//	b.BucketsByIP[ipstr] = append(s, curBacket)
	//} else {
	//	//s = make([]*Bucket, 0, 1)
	//	//s = append(s, &curBacket)
	//	b.BucketsByIP[ipstr] = []*Bucket{curBacket}
	//}
	//b.muBucketsByIP.Unlock()
	return curBacket
}

func (b *Buckets) GetBucketsByLogin(login string, since time.Time) []*Bucket {
	buckets, found := b.BucketsByLogin.Load(login)
	//	b.muBucketsByLogin.RLock()
	//defer b.muBucketsByLogin.RUnlock()
	//buckets, found := b.BucketsByLogin[login]
	result := make([]*Bucket, 0, 1)
	if found {
		for _, bucket := range buckets {
			if bucket.time.After(since) {
				result = append(result, bucket)
			}
		}
	}
	return result
}

func (b *Buckets) GetBucketsByPass(pass string, since time.Time) []*Bucket {
	buckets, found := b.BucketsByPass.Load(pass)
	//b.muBucketsByPass.RLock()
	//defer b.muBucketsByPass.RUnlock()
	//buckets, found := b.BucketsByPass[pass]
	result := make([]*Bucket, 0, 1)
	if found {
		for _, bucket := range buckets {
			if bucket.time.After(since) {
				result = append(result, bucket)
			}
		}
	}
	return result
}

func (b *Buckets) GetBucketsByIP(ip net.IP, since time.Time) []*Bucket {
	buckets, found := b.BucketsByIP.Load(string(ip.To16()))
	//b.muBucketsByIP.RLock()
	//defer b.muBucketsByIP.RUnlock()
	//buckets, found := b.BucketsByIP[string(ip.To16())]
	result := make([]*Bucket, 0, 1)
	if found {
		for _, bucket := range buckets {
			if bucket.time.After(since) {
				result = append(result, bucket)
			}
		}
	}
	return result
}
