package db

import "errors"

func PoolIsNil() string {
	return "ERR Pool is nil "
}
func ErrPoolIsNil() error {
	return errors.New(PoolIsNil())
}
