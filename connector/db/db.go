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
	"fmt"
	"time"

	"github.com/boltdb/bolt"
	"github.com/sirupsen/logrus"
)

const (
	DefaultDBFile = "connections.db"
)

// TODO: Handle regions
type Connection struct {
	// The Identifier of the connection.
	// The client is supposed to generate valid and unique
	// identifier.
	// TODO: Move this logic to DB.
	ID string `json:"id"`
	// The Source Gateway ID
	SourceID string `json:"source_id"`
	// The Source Cloud Provider (for example AWS)
	SourceProvider string `json:"source_provider"`
	// The region where Source resources will be stored.
	//
	// Region is considered in the context of specified provider.
	SourceRegion string `json:"source_region"`
	// The ID of VPC if the Source Gateway is associated with one.
	// Otherwise this field remains empty.
	SourceVPC string `json:"source_vpc"`
	// The CIDRs that should be reachable from the destination
	// machines.
	SourceCIDRs []string `json:"source_cidr"`
	// The Destination Gateway ID
	DestinationID string `json:"destination_id"`
	// The Destination Cloud Provider (for example GCP)
	DestinationProvider string `json:"destination_provider"`
	// The region where Destination resources will be stored.
	//
	// Region is considered in the context of specified provider.
	DestinationRegion string `json:"destination_region"`
	// The ID of VPC if the Destination Gateway is associated with one.
	// Otherwise this field remains empty.
	DestinationVPC string `json:"destination_vpc"`
	// The CIDRs that should be reachable from the source
	// machines.
	DestinationCIDRs []string `json:"destination_cidr"`
	// The state of the Connection.
	//
	// Describes if the Creation/Deletion is either in
	// progress, was incomplete or completed.
	State ConnectionState
	// If the State reports that not everything was completed
	// and it is no longer in progress - it means there was an
	// issue. Error contains information about the failure so
	// the user may manually fix the issue if it doesn't fix
	// itself alone.
	Error string
	// The time when the entry was created (which means when
	// the creation of the connection has started)
	//
	// TODO: Support it.
	Created time.Time `json:"created"`
	// The time when there was the last update of the given
	// Connection.
	//
	// TODO: Support it.
	LastUpdated time.Time `json:"last_updated"`
	// Tags are the way of providing additional data when writing
	// custom providers that need the exchange of additional
	// information when creating requested resources.
	//
	// Not supported yet.
	Tags map[string]string
}

// The state of the Connection
type ConnectionState string

const (
	// The connection was created properly and can be used as a
	// bridge between different Cloud Providers.
	StateActive ConnectionState = "ACTIVE"
	// The connection creation is still in progress.
	//
	// If the script responsible for creating the connection got
	// suddenly interrupted/killed this state may remain for the
	// connection forever and should be treated similarly to
	// StatePartiallyCreated. The biggest difference is that, the
	// StateCreationInProgress will most likely fix itself after
	// restarting the operation.
	StateCreationInProgress ConnectionState = "CREATION_IN_PROGRESS"
	// The connection deletion is still in progress.
	//
	// If the script responsible for deleting the connection got
	// suddenly interrupted/killed this state may remain for the
	// connection forever and should be treated similarly to
	// StatePartiallyDelete. The biggest difference is that, the
	// StateDeletionInProgress will most likely fix itself after
	// restarting the operation.
	StateDeletionInProgress ConnectionState = "DELETION_IN_PROGRESS"
	// StatePartiallyCreated let's us know that some resources could
	// have been created, but there was an error during handling
	// requests and it should be investigated.
	StatePartiallyCreated ConnectionState = "PARTIALLY_CREATED"
	// StatePartiallyDeleted let's us know that some resources could
	// have been deleted, but there was an error during handling
	// requests and it should be investigated.
	StatePartiallyDeleted ConnectionState = "PARTIALLY_DELETED"
)

const (
	BucketName = "Connections"
)

type Client struct {
	db     *bolt.DB
	logger *logrus.Entry
}

func NewClient(filePath string, logger *logrus.Entry) (*Client, error) {
	db, err := bolt.Open(filePath, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, fmt.Errorf(
			"failed to open BoltDB file for new Client from filepath '%s': %w", filePath, err,
		)
	}

	err = db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(BucketName))
		return err
	})
	if err != nil {
		return nil, fmt.Errorf(
			"failed to refresh BoltDB Buckets. BoltDB Filepath: '%s' "+
				"BoldDB bucket: %s, err: %w", filePath, BucketName, err,
		)
	}
	return &Client{db, logger}, nil
}

func (c *Client) Close() error {
	if c.db == nil {
		return nil
	}
	if err := c.db.Close(); err != nil {
		return fmt.Errorf("failed to close BoltDB Client: %w", err)
	}
	c.db = nil
	c.logger.Debugf("Successfully closed the DB Client")
	return nil
}

// TODO: Handle adding tags and making sure to not
// clear them.
func (c *Client) PutConnection(conn Connection) error {
	return c.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		encoded, err := json.Marshal(conn)
		c.logger.Debugf(
			"Connection %s Update: %s", conn.ID, encoded,
		)
		if err != nil {
			return err
		}
		return b.Put([]byte(conn.ID), encoded)
	})
}

func (c *Client) DeleteConnection(connID string) error {
	return c.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		c.logger.Debugf(
			"Connection %s Deletion", connID,
		)
		return b.Delete([]byte(connID))
	})
}

func (c *Client) GetConnection(connID string) (*Connection, error) {
	var conn Connection
	err := c.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		c.logger.Debugf(
			"Connection %s Get", connID,
		)
		v := b.Get([]byte(connID))
		if v == nil {
			return fmt.Errorf("Connection not found")
		}
		return json.Unmarshal(v, &conn)
	})
	if err != nil {
		return nil, err
	}
	return &conn, nil
}

func (c *Client) GetConnectionWithGateways(gatewayA, providerA, gatewayB, providerB string) (*Connection, error) {
	c.logger.Debugf(
		"Looking for connection between gateways %s-%s and their providers %s-%s",
		gatewayA, gatewayB, providerA, providerB,
	)
	connections, err := c.ListConnections()
	if err != nil {
		return nil, fmt.Errorf("cannot retrieve the list of connections: %w", err)
	}
	for i := range connections {
		if connections[i].SourceID == gatewayA &&
			connections[i].SourceProvider == providerA &&
			connections[i].DestinationID == gatewayB &&
			connections[i].DestinationProvider == providerB {
			c.logger.Debugf("Connection found: %s", connections[i].ID)
			return &connections[i], nil
		}
		if connections[i].SourceID == gatewayB &&
			connections[i].SourceProvider == providerB &&
			connections[i].DestinationID == gatewayA &&
			connections[i].DestinationProvider == providerA {
			c.logger.Debugf("Connection found: %s", connections[i].ID)
			return &connections[i], nil
		}
	}

	c.logger.Debug("Connection not found")
	return nil, nil
}

func (c *Client) ListConnections() ([]Connection, error) {
	var connections []Connection
	err := c.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		c.logger.Debugf(
			"Connection List",
		)
		return b.ForEach(func(k, v []byte) error {
			c.logger.Debugf(
				"Connection found: %s: %s", k, v,
			)
			var c Connection
			if err := json.Unmarshal(v, &c); err != nil {
				return err
			}
			connections = append(connections, c)
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return connections, nil
}
