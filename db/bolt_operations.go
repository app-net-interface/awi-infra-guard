// Copyright (c) 2024 Cisco Systems, Inc. and its affiliates
// All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http:www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: Apache-2.0

package db

import (
	"encoding/json"

	"github.com/boltdb/bolt"
)

func delete_(client *boltClient, id, tableName string) error {
	return client.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(tableName))
		return bucket.Delete([]byte(id))
	})
}

func update[T any](client *boltClient, t T, id, tableName string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return client.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(tableName))

		if err := bucket.Put([]byte(id), data); err != nil {
			return err
		}
		return nil
	})
}

func get[T any](client *boltClient, id, tableName string) (*T, error) {
	var data []byte
	if err := client.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(tableName))
		data = bucket.Get([]byte(id))
		return nil
	}); err != nil {
		return nil, err
	}
	if data == nil {
		return nil, nil
	}
	var t T
	if err := json.Unmarshal(data, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

func list[T any](client *boltClient, tableName string) ([]*T, error) {
	var ts []*T
	if err := client.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(tableName))
		return bucket.ForEach(func(k, v []byte) error {
			var t T
			if err := json.Unmarshal(v, &t); err != nil {
				return err
			}
			ts = append(ts, &t)
			return nil
		})
	}); err != nil {
		return nil, err
	}

	return ts, nil
}
