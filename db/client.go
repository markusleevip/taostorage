package db

import (
	"encoding/json"
	"github.com/markusleevip/taodb/client"
	"github.com/markusleevip/taodb/resp"
	"net"
)

var (
	db *TaoDb
)

func GetDb() *TaoDb {
	return db
}

type TaoDb struct {
	pool *client.Pool
}

func New(addr string) (t *TaoDb, err error) {
	pool, err := client.New(func() (net.Conn, error) {
		return net.Dial("tcp", addr)
	})
	if err != nil {
		return &TaoDb{}, err
	}
	db = &TaoDb{pool: pool}
	return db, nil
}

func (db *TaoDb) Set(key string, value []byte) error {
	if db == nil || db.pool == nil {
		return ErrPoolIsNil()
	}
	cn, err := db.pool.Get()
	if err != nil {
		return err
	}
	defer db.pool.Put(cn)

	cn.WriteCmd("SET", []byte(key), value)

	if err := cn.Flush(); err != nil {
		cn.MarkFailed()
		return err
	}

	t, err := cn.PeekType()
	if err != nil {
		return err
	}
	switch t {
	case resp.TypeError:
		_, err = cn.ReadError()
		return err
	}
	return nil
}
func (db *TaoDb) Get(key string) ([]byte, error) {
	if db == nil || db.pool == nil {
		return nil, ErrPoolIsNil()
	}
	cn, err := db.pool.Get()
	if err != nil {
		return nil, err
	}
	defer db.pool.Put(cn)
	cn.WriteCmdString("GET", key)

	if err := cn.Flush(); err != nil {
		cn.MarkFailed()
		return nil, err
	}

	t, err := cn.PeekType()
	if err != nil {
		return nil, err
	}
	switch t {
	case resp.TypeBulk:
		s, _ := cn.ReadBulk(nil)
		return s, nil
	case resp.TypeError:
		_, err = cn.ReadError()
		return nil, err
	}

	return nil, nil
}
func (db *TaoDb) Del(key string) error {
	if db == nil || db.pool == nil {
		return ErrPoolIsNil()
	}
	cn, err := db.pool.Get()
	if err != nil {
		return err
	}
	defer db.pool.Put(cn)

	cn.WriteCmdString("DEL", key)

	if err := cn.Flush(); err != nil {
		cn.MarkFailed()
		return err
	}

	t, err := cn.PeekType()
	if err != nil {
		return err
	}
	switch t {
	case resp.TypeError:
		_, err = cn.ReadError()
		return err
	}
	return nil

}

func (db *TaoDb) State(string) (string, error) {
	return "", nil
}
func (db *TaoDb) Iterator(prefix string) (map[string]string, error) {

	if db == nil || db.pool == nil {
		return nil, ErrPoolIsNil()
	}
	cn, err := db.pool.Get()
	if err != nil {
		return nil, err
	}
	defer db.pool.Put(cn)
	cn.WriteCmdString("ITERATOR", prefix)

	if err := cn.Flush(); err != nil {
		cn.MarkFailed()
		return nil, err
	}

	t, err := cn.PeekType()
	if err != nil {
		return nil, err
	}
	switch t {
	case resp.TypeBulk:
		data, _ := cn.ReadBulk(nil)
		var mapResult map[string]string
		json.Unmarshal(data,&mapResult)
		return mapResult, nil
	case resp.TypeError:
		_, err = cn.ReadError()
		return nil, err
	}

	return nil, nil

}
func (db *TaoDb) IteratorOnlyKey(prefix string) ([]string, error) {
	return nil, nil
}
