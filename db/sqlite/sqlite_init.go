// Copyright (c) 2025 Cisco Systems, Inc. and its affiliates
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

package sqlite

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

// Table constants (copied from parent db package to avoid import cycle)
const (
	vpcTable                = "vpcs"
	vpcIndexTable           = "vpc_index"
	regionTable             = "regions"
	instanceTable           = "instances"
	subnetTable             = "subnets"
	clusterTable            = "clusters"
	podTable                = "pods"
	kubernetesServiceTable  = "kubernetes_services"
	kubernetesNodeTable     = "kubernetes_nodes"
	namespaceTable          = "namespaces"
	accountTable            = "accounts"
	routeTableTable         = "route_tables"
	aclTable                = "acls"
	securityGroupTable      = "security_groups"
	ngTable                 = "nat_gateways"
	routerTable             = "routers"
	igwTable                = "igws"
	vpcEndpointTable        = "vpcEndpoints"
	publicIPTable           = "publicIPs"
	lbTable                 = "lbs"
	networkInterfaceTable   = "network_interfaces"
	syncTimeTable           = "sync_time"
	keyPairTable            = "keyPairs"
	vpnConcentratorTable    = "vpnConcentrators"
	vpcConnectionTable      = "vpcConnections"
	vpcConnectionGraphTable = "vpc_connection_graph"
)

// Add sqlite table names to this list
var tableNames = []string{
	vpcTable,
	vpcIndexTable,
	regionTable,
	instanceTable,
	subnetTable,
	clusterTable,
	podTable,
	kubernetesServiceTable,
	kubernetesNodeTable,
	namespaceTable,
	accountTable,
	routeTableTable,
	ngTable,
	routerTable,
	igwTable,
	vpcEndpointTable,
	aclTable,
	securityGroupTable,
	publicIPTable,
	lbTable,
	networkInterfaceTable,
	syncTimeTable,
	keyPairTable,
	vpnConcentratorTable,
	vpcConnectionTable,
	vpcConnectionGraphTable,
}

type DbObject interface {
	DbId() string
	GetProvider() string
	SetSyncTime(string)
}

type sqliteClient struct {
	db    *sql.DB
	mutex sync.RWMutex
}

func NewSQLiteClient() interface{} {
	return &sqliteClient{}
}

func (client *sqliteClient) Open(filename string) error {
	var err error

	// Enable WAL mode for better concurrency
	dsn := filename + "?_journal_mode=WAL&_busy_timeout=5000&_synchronous=NORMAL&_cache_size=1000"

	client.db, err = sql.Open("sqlite3", dsn)
	if err != nil {
		return fmt.Errorf("failed to open SQLite database: %v", err)
	}

	// Set connection pool settings for concurrency
	client.db.SetMaxOpenConns(25)
	client.db.SetMaxIdleConns(25)

	// Test the connection
	if err = client.db.Ping(); err != nil {
		return fmt.Errorf("failed to ping SQLite database: %v", err)
	}

	// Create tables
	return client.createTables()
}

func (client *sqliteClient) Close() error {
	return client.db.Close()
}

func (client *sqliteClient) DropDB() error {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	// Drop all tables and recreate them
	for _, tableName := range tableNames {
		_, err := client.db.Exec(fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName))
		if err != nil {
			return fmt.Errorf("failed to drop table %s: %v", tableName, err)
		}
	}

	return client.createTables()
}

func (client *sqliteClient) createTables() error {
	// Create a generic table structure for all resource types
	// Each table has: id (primary key), data (JSON), provider, account_id, region, sync_time
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS %s (
			id TEXT PRIMARY KEY,
			data TEXT NOT NULL,
			provider TEXT,
			account_id TEXT,
			region TEXT,
			sync_time TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
		CREATE INDEX IF NOT EXISTS idx_%s_provider ON %s(provider);
		CREATE INDEX IF NOT EXISTS idx_%s_account ON %s(account_id);
		CREATE INDEX IF NOT EXISTS idx_%s_region ON %s(region);
	`

	for _, tableName := range tableNames {
		sql := fmt.Sprintf(createTableSQL,
			tableName,
			tableName, tableName,
			tableName, tableName,
			tableName, tableName)

		if _, err := client.db.Exec(sql); err != nil {
			return fmt.Errorf("failed to create table %s: %v", tableName, err)
		}
	}

	return nil
}

// Generic put function that works with any DbObject
func (client *sqliteClient) putObject(obj interface{}, id, tableName string) error {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	// Marshal object to JSON
	data, err := json.Marshal(obj)
	if err != nil {
		return fmt.Errorf("failed to marshal object: %v", err)
	}

	// Extract metadata if the object implements DbObject
	var provider, accountID, region string
	if dbObj, ok := obj.(DbObject); ok {
		provider = dbObj.GetProvider()
	}

	// Use UPSERT (INSERT OR REPLACE) for atomic updates
	query := `
		INSERT OR REPLACE INTO ` + tableName + ` 
		(id, data, provider, account_id, region, sync_time, updated_at) 
		VALUES (?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
	`

	_, err = client.db.Exec(query, id, string(data), provider, accountID, region, "")
	if err != nil {
		return fmt.Errorf("failed to put object in %s: %v", tableName, err)
	}

	return nil
}

// Generic get function
func (client *sqliteClient) getObject(id, tableName string, result interface{}) error {
	client.mutex.RLock()
	defer client.mutex.RUnlock()

	query := "SELECT data FROM " + tableName + " WHERE id = ?"
	row := client.db.QueryRow(query, id)

	var data string
	if err := row.Scan(&data); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("object with id %s not found in %s", id, tableName)
		}
		return fmt.Errorf("failed to get object from %s: %v", tableName, err)
	}

	return json.Unmarshal([]byte(data), result)
}

// Generic delete function
func (client *sqliteClient) deleteObject(id, tableName string) error {
	client.mutex.Lock()
	defer client.mutex.Unlock()

	query := "DELETE FROM " + tableName + " WHERE id = ?"
	result, err := client.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete object from %s: %v", tableName, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("object with id %s not found in %s", id, tableName)
	}

	return nil
}

// Helper function to list objects from a table
func (client *sqliteClient) listObjects(tableName string, constructor func() interface{}) ([]interface{}, error) {
	client.mutex.RLock()
	defer client.mutex.RUnlock()

	query := "SELECT data FROM " + tableName
	rows, err := client.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to list objects from %s: %v", tableName, err)
	}
	defer rows.Close()

	var results []interface{}
	for rows.Next() {
		var data string
		if err := rows.Scan(&data); err != nil {
			return nil, fmt.Errorf("failed to scan row from %s: %v", tableName, err)
		}

		obj := constructor()
		if err := json.Unmarshal([]byte(data), obj); err != nil {
			return nil, fmt.Errorf("failed to unmarshal object from %s: %v", tableName, err)
		}

		results = append(results, obj)
	}

	return results, nil
}

// Helper functions (copied from parent package to avoid import cycle)

// extractPorts extracts individual ports from a port range string
func extractPorts(portRange string) []string {
	ports := []string{}

	// Handle single port
	if !strings.Contains(portRange, "-") {
		ports = append(ports, strings.TrimSpace(portRange))
		return ports
	}

	// Handle port range
	parts := strings.Split(portRange, "-")
	if len(parts) == 2 {
		start, err1 := strconv.Atoi(strings.TrimSpace(parts[0]))
		end, err2 := strconv.Atoi(strings.TrimSpace(parts[1]))

		if err1 == nil && err2 == nil {
			// For analysis, we'll focus on commonly known risky ports
			riskyPorts := []int{
				21, 22, 23, 25, 53, 69, 80, 110, 111, 135, 139, 143, 443, 445, 993, 995, 1433, 1521, 3306, 3389, 5432, 5984, 6379, 8080, 8443, 9200, 27017,
			}

			for _, port := range riskyPorts {
				if port >= start && port <= end {
					ports = append(ports, strconv.Itoa(port))
				}
			}

			// Also add the start and end ports
			ports = append(ports, strconv.Itoa(start))
			ports = append(ports, strconv.Itoa(end))
		}
	}

	return ports
}

// removeDuplicates removes duplicate strings from a slice
func removeDuplicates(slice []string) []string {
	keys := make(map[string]bool)
	var result []string

	for _, item := range slice {
		if !keys[item] {
			keys[item] = true
			result = append(result, item)
		}
	}

	return result
}

// contains checks if a string slice contains a specific string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// hasTransitGatewayRoute checks if a route table has a route via transit gateway
func hasTransitGatewayRoute(routeTable interface{}, tgwID string) bool {
	// This is a placeholder implementation
	// You'll need to implement this based on your specific route table structure
	return false
}
