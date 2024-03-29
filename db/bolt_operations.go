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
