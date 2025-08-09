package db

import (
	"context"
	"fmt"

	"github.com/app-net-interface/awi-infra-guard/grpc/go/infrapb"
	"github.com/app-net-interface/awi-infra-guard/types"
)

func (p *providerWithDB) GetVpcConnectionGraph(ctx context.Context, params *infrapb.GetVpcConnectionGraphRequest) (*types.VpcConnectionGraph, error) {
	// Get VPC connection from the database
	conn, err := p.dbClient.GetVpcConnection(params.VpcId)
	if err != nil || conn == nil {
		p.logger.Errorf("Failed to get VPC connection for VPC %s: %v", params.VpcId, err)
		return nil, fmt.Errorf("failed to get VPC connection: %w", err)
	}
	// Create base graph structure
	graph := &types.VpcConnectionGraph{
		Id:       conn.ID,
		Provider: params.Provider,
		Accounts: []string{conn.FromVpcAccountId, conn.ToVpcAccountId},
		Regions:  []string{conn.FromVpcRegion, conn.ToVpcRegion},
	}

	// Create VPC nodes
	sourceNode := &types.VpcConnectionGraphNode{
		Id:       conn.FromVpcId,
		Name:     "Source VPC",
		NodeType: "vpc",
		Provider: params.Provider,
	}

	destNode := &types.VpcConnectionGraphNode{
		Id:       conn.ToVpcId,
		Name:     "Target VPC",
		NodeType: "vpc",
		Provider: params.Provider,
	}
	graph.Nodes = []*types.VpcConnectionGraphNode{sourceNode, destNode}

	// Create connection edge
	edge := &types.VpcConnectionGraphEdge{
		Id:             conn.ID,
		SourceNodeId:   conn.FromVpcId,
		TargetNodeId:   conn.ToVpcId,
		ConnectionType: conn.ConnectionType,
		Status:         conn.Status,
		AccountId:      conn.Account, // Set the account that owns the connection
	}
	graph.Edges = []*types.VpcConnectionGraphEdge{edge}

	// Get source VPC internal graph
	sourceGraphReq := &infrapb.GetVpcConnectivityGraphRequest{
		Provider:  params.Provider,
		AccountId: conn.FromVpcAccountId,
		VpcId:     conn.FromVpcId,
		Region:    conn.FromVpcRegion,
		Creds:     params.Creds,
	}
	sourceGraph, err := p.GetVpcConnectivityGraph2(ctx, sourceGraphReq)
	if err != nil {
		p.logger.Warnf("Failed to get source VPC graph: %v", err)
	} else {
		graph.SrcVpcGraph = sourceGraph
	}

	// Get destination VPC internal graph
	destGraphReq := &infrapb.GetVpcConnectivityGraphRequest{
		Provider:  params.Provider,
		AccountId: conn.ToVpcAccountId,
		VpcId:     conn.ToVpcId,
		Region:    conn.ToVpcRegion,
		Creds:     params.Creds,
	}
	destGraph, err := p.GetVpcConnectivityGraph2(ctx, destGraphReq)
	if err != nil {
		p.logger.Warnf("Failed to get destination VPC graph: %v", err)
	} else {
		graph.DestVpcGraph = destGraph
	}

	p.logger.Infof("Built VPC connection graph with %d nodes and %d edges", len(graph.Nodes), len(graph.Edges))
	return graph, nil
}
