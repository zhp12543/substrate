package client

import (
	"github.com/zhp12543/substrate/wsrpc/jsonrpc"
	"github.com/zhp12543/substrate/wsrpc/types/rpccall"
	"github.com/zhp12543/substrate/wsrpc/utils/decoder"
	"errors"
)

func createSystem(p *jsonrpc.WsProvider) *system {
	s := system{}
	s.provider = p
	s.section = "system"
	return &s
}

type system struct {
	rpcBase
}

func (s *system) Name() (string, error) {
	result, err := s.call("name", nil)
	return result.(string), err
}

func (s *system) Version() (string, error) {
	result, err := s.call("version", nil)
	return result.(string), err
}

func (s *system) Chain() (string, error) {
	result, err := s.call("chain", nil)
	return result.(string), err
}

func (s *system) Health() (*rpccall.Health, error) {
	result, err := s.call("health", nil)
	if err != nil {
		return nil, err
	}
	health := rpccall.Health{}
	err = decoder.MapDecode(result, &health)
	return &health, err
}

func (s *system) Peers() ([]rpccall.PeerInfo, error) {
	result, err := s.call("peers", nil)
	var peers []rpccall.PeerInfo
	err = decoder.MapDecode(result, &peers)
	return peers, err
}

func (s *system) NetworkState() (interface{}, error) {
	result, err := s.call("networkState", nil)
	return result, err
}

func (s *system) Properties() (*rpccall.ChainProperties, error) {
	result, err := s.call("properties", nil)
	if err != nil {
		return nil, err
	}
	properties := rpccall.ChainProperties{}
	err = decoder.MapDecode(result, &properties)
	return &properties, err
}

func (s *system) Nonce(address string) (int64, error) {
	var params []interface{} = nil
	if address != "" {
		params = []interface{}{address}
	}
	result, err := s.call("accountNextIndex", params)
	if err != nil {
		return 0, err
	}
	if result == nil {
		return 0, errors.New("get nonce is nil")
	}
	nonceF := result.(float64)

	return int64(nonceF), err
}
