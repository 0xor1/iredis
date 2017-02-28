package iredis

import (
	"github.com/garyburd/redigo/redis"
	"github.com/stretchr/testify/mock"
)



type MockPool struct {
	mock.Mock
}

func (m *MockPool) Get() redis.Conn {
	args := m.Called()
	conn := args.Get(0)
	if conn == nil {
		return nil
	}
	return conn.(redis.Conn)
}

type MockConn struct {
	mock.Mock
}

func (m *MockConn) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockConn) Err() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockConn) Do(commandName string, args ...interface{}) (interface{}, error) {
	_args := make([]interface{}, 0, len(args)+1)
	_args = append(_args, commandName)
	_args = append(_args, args...)
	a := m.Called(_args...)
	return a.Get(0), a.Error(1)
}

func (m *MockConn) Send(commandName string, args ...interface{}) error {
	_args := make([]interface{}, 0, len(args)+1)
	_args = append(_args, commandName)
	_args = append(_args, args...)
	a := m.Called(_args...)
	return a.Error(0)
}

func (m *MockConn) Flush() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockConn) Receive() (interface{}, error) {
	args := m.Called()
	return args.Get(0), args.Error(1)
}