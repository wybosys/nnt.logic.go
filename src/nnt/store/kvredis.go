package store

import (
	"github.com/go-redis/redis"
	"nnt/core/kernel"
	"strconv"
	"strings"
)

const (
	REDIS_DEFAULT_PORT = 6379
)

type redisHandle struct {
	cluster *redis.ClusterClient
	normal  *redis.Client
	prefix  string // 驱动没有实现prefix的处理
}

type KvRedis struct {
	AbstractKv

	Dbid    int
	Host    string
	Port    int
	Passwd  string
	Timeout int
	Retry   float64
	Cluster bool
	Prefix  string
	hdl     redisHandle
}

func KvNewRedis() *KvRedis {
	return new(KvRedis{
		Dbid:    0,
		Timeout: 1,
		Retry:   0.1,
		Port:    REDIS_DEFAULT_PORT,
	})
}

func (self *KvRedis) Config(cfg *kernel.JsonObject) bool {
	if !self.AbstractKv.Config(cfg) {
		return false
	}
	host, err := cfg.GetString("host")
	if err {
		return false
	}
	if v, _ := cfg.GetBool("cluster"); v {
		self.Dbid = 0
		self.Cluster = true
	} else if v, ok := cfg.GetInt("dbid"); ok {
		self.Dbid = v
	}
	if v, ok := cfg.GetString("prefix"); ok {
		self.Prefix = v
	}

	hs := strings.Split(host, ":")
	self.Host = hs[0]
	if len(hs) == 2 {
		v, _ := strconv.Atoi(hs[1])
		self.Port = v
	}

	self.Passwd, _ = cfg.GetString("password")
	if v, ok := cfg.GetInt("timeout"); ok {
		self.Timeout = v
	}
	if v, ok := cfg.GetFloat("retry"); ok {
		self.Retry = v
	}

	return true
}

func (self *KvRedis) Clone() *KvRedis {
	r := KvNewRedis()
	r.Dbid = self.Dbid
	r.Host = self.Host
	r.Port = self.Port
	r.Passwd = self.Passwd
	r.Timeout = self.Timeout
	r.Retry = self.Retry
	r.Prefix = self.Prefix
	r.Cluster = self.Cluster
	return r
}

func (self *KvRedis) Open() {
	self.hdl.open(self)
}

func (self *KvRedis) Close() {
	self.hdl.close()
}

func (self *KvRedis) TestOpen() {
	if !self.hdl.ping() {
		self.Close()
		self.Open()
	}
}

func (self *KvRedis) Get(key string) *kernel.Variant {
	return self.hdl.get(key)
}

// ------------------------------------------实现

func (self *redisHandle) open(kv *KvRedis) {
	self.prefix = kv.Prefix
	if kv.Cluster {
		self.normal = redis.NewClient(&redis.Options{
			Addr:     kv.Host + ":" + string(kv.Port),
			Password: kv.Passwd,
			DB:       kv.Dbid,
		})
	} else {
		self.cluster = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    []string{kv.Host + ":" + string(kv.Port)},
			Password: kv.Passwd,
		})
	}
}

func (self *redisHandle) key(k string) string {
	return self.prefix + k
}

func (self *redisHandle) close() {
	if self.normal != nil {
		self.normal.Close()
		self.normal = nil
	} else if self.cluster != nil {
		self.cluster.Close()
		self.cluster = nil
	}
}

func (self *redisHandle) ping() bool {
	if self.normal != nil {
		_, err := self.normal.Ping().Result()
		if err != nil {
			return true
		}
	} else if self.cluster != nil {
		_, err := self.cluster.Ping().Result()
		if err != nil {
			return true
		}
	}
	return false
}

func (self *redisHandle) get(key string) *kernel.Variant {
	k := self.key(key)
	if self.normal != nil {
		v, _ := self.normal.Get(k).Bytes()
		return kernel.VariantFromString(v)
	}
	v, _ := self.cluster.Get(k).Bytes()
	return kernel.VariantFromString(v)
}
