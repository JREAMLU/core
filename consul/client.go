package consul

import (
	"fmt"
	"net"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/consul/api"
)

const (
	ErrorKeyNotExist = "Key(%s) does not exist"
	ErrorDirNotExist = "Dir(%s) does not exist"
)

type Client struct {
	config       *api.Config
	consulClient *api.Client
	register     *api.AgentServiceRegistration
}

func NewClient(opts ...ClientOptionFunc) (*Client, error) {
	client := &Client{
		config: api.DefaultConfig(),
	}

	for _, opt := range opts {
		opt(client)
	}

	consulClient, err := api.NewClient(client.config)
	if err != nil {
		return nil, err
	}
	client.consulClient = consulClient

	consulClient.Agent().ServiceRegister(client.register)

	return client, nil
}

type ClientOptionFunc func(*Client) error

func SetRegister(reg *api.AgentServiceRegistration) ClientOptionFunc {
	return func(client *Client) error {
		client.register = reg
		return nil
	}
}

func SetAddress(address string) ClientOptionFunc {
	return func(client *Client) error {
		if address != "" {
			client.config.Address = address
		}
		return nil
	}
}

func SetScheme(scheme string) ClientOptionFunc {
	return func(client *Client) error {
		if scheme != "" {
			client.config.Scheme = scheme
		}
		return nil
	}
}

func SetDatacenter(datacenter string) ClientOptionFunc {
	return func(client *Client) error {
		if datacenter != "" {
			client.config.Datacenter = datacenter
		}
		return nil
	}
}

func SetHttpBasicAuth(userName, password string) ClientOptionFunc {
	return func(client *Client) error {
		if userName != "" {
			client.config.HttpAuth = &api.HttpBasicAuth{
				Username: userName,
				Password: password,
			}
		}
		return nil
	}
}

func SetWaitTime(waitTime time.Duration) ClientOptionFunc {
	return func(client *Client) error {
		if waitTime > 0 {
			client.config.WaitTime = waitTime
		}
		return nil
	}
}

func SetToken(token string) ClientOptionFunc {
	return func(client *Client) error {
		if token != "" {
			client.config.Token = token
		}
		return nil
	}
}

func (client *Client) KV() *api.KV {
	return client.consulClient.KV()
}

func (client *Client) Deregister() error {
	return client.consulClient.Agent().ServiceDeregister("api.srv")
}

func (client *Client) Put(key, value string) error {
	pair := &api.KVPair{
		Key:   key,
		Value: []byte(value),
	}
	_, err := client.KV().Put(pair, nil)
	return err
}

func (client *Client) Get(key string) (string, error) {
	kvPair, _, err := client.KV().Get(key, nil)
	if err != nil {
		return "", err
	}
	if kvPair == nil {
		return "", fmt.Errorf(ErrorKeyNotExist, key)
	}
	return string(kvPair.Value), nil
}

// GetOrDefault 获取指定key的值，没有就返回缺省值
func (client Client) GetOrDefault(key, defaultValue string) string {
	value, err := client.Get(key)
	if err != nil {
		return defaultValue
	}

	return value
}

func (client *Client) GetInt(key string) (int, error) {
	value, err := client.Get(key)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(value)
}

// GetInt64 获取指定key的int64值
func (client Client) GetInt64(key string) (int64, error) {
	value, err := client.Get(key)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(value, 10, 64)
}

// GetOrDefaultInt64 获取指定key的int64值，没有就返回缺省值
func (client Client) GetOrDefaultInt64(key string, defaultValue int64) int64 {
	value, err := client.GetInt64(key)
	if err != nil {
		return defaultValue
	}

	return value
}

func (client *Client) GetFloat64(key string) (float64, error) {
	value, err := client.Get(key)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(value, 64)
}

func (client *Client) GetHostPort(key string) (string, string, error) {
	value, err := client.Get(key)
	if err != nil {
		return "", "", err
	}
	return net.SplitHostPort(value)
}

//GetChildKeys 获取下一级所有key
func (client *Client) GetChildKeys(keyPrefix string) ([]string, error) {
	if !strings.HasSuffix(keyPrefix, "/") {
		keyPrefix += "/"
	}

	keys, _, err := client.KV().Keys(keyPrefix, "/", nil)
	if err != nil {
		return nil, err
	}

	keyPrefixIndex := -1
	for i := range keys {
		if keys[i] == keyPrefix {
			keyPrefixIndex = i
			break
		}
	}

	if keyPrefixIndex == -1 {
		return keys, nil
	}

	if keyPrefixIndex+1 == len(keys) {
		return keys[:keyPrefixIndex], nil
	}

	if keyPrefixIndex == 0 {
		return keys[1:], nil
	}

	return append(keys[:keyPrefixIndex], keys[keyPrefixIndex+1:]...), nil
}

//GetChildValues 获取下一级所有key的值
//TODO 如果下一级key也是目录？
func (client *Client) GetChildValues(keyPrefix string) (api.KVPairs, error) {
	keys, err := client.GetChildKeys(keyPrefix)
	if err != nil {
		return nil, err
	}
	return client.GetValues(keys)
}

//GetArray
//  例如
//              key                         value
//      conn/mongodb/goimhistory/1  172.16.9.221:27017
//      conn/mongodb/goimhistory/2  172.16.9.222:27017
//  keyPrefix=conn/mongodb/goimhistory/
//  return ["172.16.9.221:27017","172.16.9.222:27017"]
func (client *Client) GetArray(keyPrefix string) ([]string, error) {
	kvPairs, err := client.GetChildValues(keyPrefix)
	if err != nil {
		return nil, err
	}

	result := make([]string, len(kvPairs))
	for i := range kvPairs {
		result[i] = string(kvPairs[i].Value)
	}
	return result, nil
}

//GetMap
// 例如
//              key                         value
//  conn/redis/config/master/1/db           0
//  conn/redis/config/master/1/ip           172.16.9.221
//  conn/redis/config/master/1/poolsize     1
//  conn/redis/config/master/1/port         6379
//  keyPrefix=conn/redis/config/master/1/
//  return map[db:0 ip:172.16.9.221 poolsize:1 port:6379]
func (client *Client) GetMap(keyPrefix string) (map[string]string, error) {
	kvPairs, err := client.GetChildValues(keyPrefix)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for i := range kvPairs {
		result[path.Base(kvPairs[i].Key)] = string(kvPairs[i].Value)
	}
	return result, nil
}

const FunctionIDsTmpl = "service/go/%s/functionids"

func (client *Client) GetValues(keys []string) (api.KVPairs, error) {
	txn := make(api.KVTxnOps, len(keys))
	for i := range keys {
		txn[i] = &api.KVTxnOp{
			Verb: api.KVGet,
			Key:  keys[i],
		}
	}
	success, resp, _, err := client.KV().Txn(txn, nil)
	if err != nil {
		return nil, err
	}

	if len(resp.Errors) != 0 {
		return nil, fmt.Errorf("%v", resp)
	}

	if !success {
		//TODO
		return nil, nil
	}
	return resp.Results, nil
}

//GetFunctionIDs 获取服务所有权限function id和对应的uri
func (client *Client) GetUriAndFunctionIDs(serviceName string) (map[string]string, error) {
	keyPrefix := fmt.Sprintf(FunctionIDsTmpl, serviceName)
	keys, _, err := client.KV().Keys(keyPrefix, "", &api.QueryOptions{})
	if err != nil {
		return nil, err
	}
	functionIDs := make(map[string]string)
	kvPairs, err := client.GetValues(keys)
	for i := range kvPairs {
		functionIDs[strings.Replace(kvPairs[i].Key, keyPrefix, "", 1)] = string(kvPairs[i].Value)
	}
	return functionIDs, nil
}
