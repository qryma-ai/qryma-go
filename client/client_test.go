package client

import (
	"testing"
)

func TestNewQrymaClient(t *testing.T) {
	// 使用示例 API 密钥
	client, err := NewQrymaClient("ak-3eb4db7c06354782926ab35106d4b461")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if client == nil {
		t.Fatal("Expected client to not be nil")
	}
}

func TestNewQrymaClient_EmptyAPIKey(t *testing.T) {
	client, err := NewQrymaClient("")
	if err == nil {
		t.Fatal("Expected error for empty API key, got nil")
	}
	if client != nil {
		t.Fatal("Expected client to be nil")
	}
}

func TestWithBaseURL(t *testing.T) {
	client, err := NewQrymaClient("ak-3eb4db7c06354782926ab35106d4b461", WithBaseURL("https://custom.qryma.com"))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if client == nil {
		t.Fatal("Expected client to not be nil")
	}
}
