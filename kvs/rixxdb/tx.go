// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rixxdb

import (
	"github.com/abcum/rixxdb"
	"github.com/abcum/surreal/kvs"
)

type TX struct {
	pntr *rixxdb.TX
}

func one(all kvs.KV, err error) (kvs.KV, error) {

	switch err {
	case nil:
		break
	default:
		return nil, &kvs.DBError{}
	case rixxdb.ErrTxNotExpectedValue:
		return nil, &kvs.KVError{}
	}

	return all, err

}

func many(all []kvs.KV, err error) ([]kvs.KV, error) {

	switch err {
	case nil:
		break
	default:
		return nil, &kvs.DBError{}
	case rixxdb.ErrTxNotExpectedValue:
		return nil, &kvs.KVError{}
	}

	return all, err

}

func (tx *TX) Closed() bool {
	return tx.pntr.Closed()
}

func (tx *TX) Cancel() error {
	return tx.pntr.Cancel()
}

func (tx *TX) Commit() error {
	return tx.pntr.Commit()
}

func (tx *TX) Get(ver int64, key []byte) (kvs.KV, error) {
	all, err := tx.pntr.Get(ver, key)
	return one(all, err)
}

func (tx *TX) GetL(ver int64, key []byte) ([]kvs.KV, error) {
	all, err := tx.pntr.GetL(ver, key)
	out := make([]kvs.KV, len(all))
	for i, v := range all {
		out[i] = v
	}
	return many(out, err)
}

func (tx *TX) GetP(ver int64, key []byte, max uint64) ([]kvs.KV, error) {
	all, err := tx.pntr.GetP(ver, key, max)
	out := make([]kvs.KV, len(all))
	for i, v := range all {
		out[i] = v
	}
	return many(out, err)
}

func (tx *TX) GetR(ver int64, beg []byte, end []byte, max uint64) ([]kvs.KV, error) {
	all, err := tx.pntr.GetR(ver, beg, end, max)
	out := make([]kvs.KV, len(all))
	for i, v := range all {
		out[i] = v
	}
	return many(out, err)
}

func (tx *TX) Del(ver int64, key []byte) (kvs.KV, error) {
	all, err := tx.pntr.Del(ver, key)
	return one(all, err)
}

func (tx *TX) DelC(ver int64, key []byte, exp []byte) (kvs.KV, error) {
	return tx.pntr.DelC(ver, key, exp)
}

func (tx *TX) DelL(ver int64, key []byte) ([]kvs.KV, error) {
	all, err := tx.pntr.DelL(ver, key)
	out := make([]kvs.KV, len(all))
	for i, v := range all {
		out[i] = v
	}
	return many(out, err)
}

func (tx *TX) DelP(ver int64, key []byte, max uint64) ([]kvs.KV, error) {
	all, err := tx.pntr.DelP(ver, key, max)
	out := make([]kvs.KV, len(all))
	for i, v := range all {
		out[i] = v
	}
	return many(out, err)
}

func (tx *TX) DelR(ver int64, beg []byte, end []byte, max uint64) ([]kvs.KV, error) {
	all, err := tx.pntr.DelR(ver, beg, end, max)
	out := make([]kvs.KV, len(all))
	for i, v := range all {
		out[i] = v
	}
	return many(out, err)
}

func (tx *TX) Put(ver int64, key []byte, val []byte) (kvs.KV, error) {
	all, err := tx.pntr.Put(ver, key, val)
	return one(all, err)
}

func (tx *TX) PutC(ver int64, key []byte, val []byte, exp []byte) (kvs.KV, error) {
	all, err := tx.pntr.PutC(ver, key, val, exp)
	return one(all, err)
}

func (tx *TX) PutL(ver int64, key []byte, val []byte) ([]kvs.KV, error) {
	all, err := tx.pntr.PutL(ver, key, val)
	out := make([]kvs.KV, len(all))
	for i, v := range all {
		out[i] = v
	}
	return many(out, err)
}

func (tx *TX) PutP(ver int64, key []byte, val []byte, max uint64) ([]kvs.KV, error) {
	all, err := tx.pntr.PutP(ver, key, val, max)
	out := make([]kvs.KV, len(all))
	for i, v := range all {
		out[i] = v
	}
	return many(out, err)
}

func (tx *TX) PutR(ver int64, key []byte, val []byte, exp []byte, max uint64) ([]kvs.KV, error) {
	all, err := tx.pntr.PutR(ver, key, val, exp, max)
	out := make([]kvs.KV, len(all))
	for i, v := range all {
		out[i] = v
	}
	return many(out, err)
}