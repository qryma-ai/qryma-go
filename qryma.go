// Package qryma is the main package for the Qryma Go SDK.
//
// A Go SDK for the Qryma Search API, providing a simple and intuitive
// interface for accessing Qryma's powerful search capabilities.
//
// Example usage:
//
//	package main
//
//	import (
//		"fmt"
//		"log"
//
//		"github.com/qryma-ai/qryma-go"
//	)
//
//	func main() {
//		client, err := qryma.Qryma(qryma.ClientConfig{
//			APIKey: "ak-********************",
//		})
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		resp, err := client.Search("ces", qryma.SearchOptions{Lang: "zh-CN"})
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		fmt.Printf("%+v\n", resp)
//	}
package qryma

import (
	"github.com/qryma-ai/qryma-go/client"
	"github.com/qryma-ai/qryma-go/version"
)

// Re-export types from client package for convenience
type (
	SearchOptions = client.SearchOptions
	QrymaResponse = client.QrymaResponse
	QrymaClient   = client.QrymaClient
	ClientConfig  = client.ClientConfig
	ClientOption  = client.ClientOption
)

// Qryma creates a Qryma client instance.
// This is a convenience function that wraps client.Qryma.
func Qryma(config ClientConfig) (*QrymaClient, error) {
	return client.Qryma(config)
}

// NewClient creates a new QrymaClient
//
// This is a convenience function that wraps client.NewQrymaClient.
// For more configuration options, use client.NewQrymaClient directly.
func NewClient(apiKey string, opts ...ClientOption) (*QrymaClient, error) {
	return client.NewQrymaClient(apiKey, opts...)
}

// Version returns the SDK version
func Version() string {
	return version.Version
}
