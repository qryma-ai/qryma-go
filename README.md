# Qryma Go SDK

A Go SDK for the Qryma Search API, providing a simple and intuitive interface for accessing Qryma's powerful search capabilities.

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Usage Examples](#usage-examples)
- [API Reference](#api-reference)
- [Configuration](#configuration)
- [Error Handling](#error-handling)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Installation

You can install the Qryma Go SDK using go get:

```bash
go get github.com/qryma-ai/qryma-go
```

## Quick Start

```go
package main

import (
	"fmt"
	"log"

	"github.com/qryma-ai/qryma-go"
)

func main() {
	// To install: go get github.com/qryma-ai/qryma-go
	client, err := qryma.Qryma(qryma.ClientConfig{
		APIKey: "ak-********************",
	})
	if err != nil {
		log.Fatal(err)
	}

	client.Search("artificial intelligence", qryma.SearchOptions{
		Lang: "en",
	})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("%+v\n", response)
}
```

## Usage Examples

### Basic Search

```go
package main

import (
	"fmt"
	"log"

	"github.com/qryma-ai/qryma-go"
)

func main() {
	client, err := qryma.Qryma(qryma.ClientConfig{
		APIKey: "ak-********************",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Simple search with just query
	response, err := client.Search("python programming")
	if err != nil {
		log.Fatal(err)
	}

	// Access the organic results
	if organic, ok := response["organic"].([]interface{}); ok {
		for _, result := range organic {
			res := result.(map[string]interface{})
			fmt.Println(res["title"])
			fmt.Println(res["link"])
			fmt.Println(res["snippet"])
			fmt.Println()
		}
	}
}
```

### Search with All Parameters

```go
package main

import (
	"fmt"
	"log"

	"github.com/qryma-ai/qryma-go"
)

func main() {
	client, err := qryma.Qryma(qryma.ClientConfig{
		APIKey: "ak-********************",
	})
	if err != nil {
		log.Fatal(err)
	}

	// Search with all available parameters
	response, err := client.Search("machine learning tutorials", qryma.SearchOptions{
		Lang:   "en",
		Start:  0,
		Safe:   false,
		Detail: false,
	})
	if err != nil {
		log.Fatal(err)
	}

	if organic, ok := response["organic"].([]interface{}); ok {
		fmt.Printf("Found %d results\n", len(organic))
	}
}
```

### Custom Configuration

You can specify additional configuration options:

```go
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/qryma-ai/qryma-go"
)

func main() {
	client, err := qryma.Qryma(qryma.ClientConfig{
		APIKey:  "ak-********************",
		BaseURL: "https://custom.qryma.com",
		Timeout: 60 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Search("custom search")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", response)
}
```

### API Response Format

The `Search()` method returns the raw API response in the following format:

```go
map[string]interface{}{
	"organic": []interface{}{
		map[string]interface{}{
			"title":      "Result Title",
			"date":       "",
			"link":       "https://example.com",
			"position":   1,
			"site_name":  "Example.com",
			"snippet":    "Description text...",
		},
	},
}
```

**Field descriptions:**
- `title`: Search result title
- `date`: Publication date (if available)
- `link`: URL of the search result
- `position`: Position in the results list
- `site_name`: Name of the website
- `snippet`: Brief description or excerpt from the page

## API Reference

### Qryma(config ClientConfig) (*QrymaClient, error)

Factory function to create a Qryma client instance.

**Parameters:**
- `config.APIKey`: Your Qryma API key (required)
- `config.BaseURL`: Base URL for the API (optional, default: `https://search.qryma.com`)
- `config.Timeout`: Request timeout (optional, default: 30 seconds)

**Returns:**
- `*QrymaClient`: Qryma client instance
- `error`: Any error that occurred

### QrymaClient.Search(query string, options ...SearchOptions) (QrymaResponse, error)

Perform a search with the given query and return the raw API response.

**Parameters:**
- `query`: Search query string (required)
- `options`: Search options (optional)
  - `Lang`: Language code for search results (e.g., 'am' for Amharic, 'en' for English)
  - `Start`: Starting position of results (default: 0)
  - `Safe`: Safe search mode: true or false (default: false)
  - `Detail`: Include detailed results (default: false)

**Returns:**
- `QrymaResponse`: Raw API response (map[string]interface{})
- `error`: Any error that occurred

### Alternative: Using QrymaClient directly

If you prefer, you can still use the `NewClient` function directly:

```go
package main

import (
	"fmt"
	"log"

	"github.com/qryma-ai/qryma-go"
)

func main() {
	client, err := qryma.NewClient("ak-********************")
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Search("artificial intelligence")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", response)
}
```

## Configuration

### Environment Variables

You can configure the API key using environment variables:

```bash
export QRYMA_API_KEY="your-api-key"
```

Then in your code:

```go
package main

import (
	"os"

	"github.com/qryma-ai/qryma-go"
)

func main() {
	client, err := qryma.Qryma(qryma.ClientConfig{
		APIKey: os.Getenv("QRYMA_API_KEY"),
	})
	if err != nil {
		panic(err)
	}
}
```

## Error Handling

The SDK returns errors for API errors:

```go
package main

import (
	"log"
	"strings"

	"github.com/qryma-ai/qryma-go"
)

func main() {
	client, err := qryma.Qryma(qryma.ClientConfig{
		APIKey: "ak-********************",
	})
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Search("test query")
	if err != nil {
		if strings.Contains(err.Error(), "timed out") {
			log.Fatal("Network timeout error")
		} else if strings.Contains(err.Error(), "API request failed") {
			log.Fatal("API error")
		} else {
			log.Fatalf("General error: %v", err)
		}
	}

	// Process results...
}
```

Common error conditions:
- Invalid API key
- Rate limiting
- Network timeouts
- Invalid parameters

## Testing

The SDK includes a simple test file. To run the test:

1. First, replace the placeholder API key in `client/client_test.go` with your actual API key
2. Then run the test:

```bash
go test -v ./...
```

## Contributing

Contributions are welcome! Please see our contributing guide for more information.

## License

MIT License - see the [LICENSE](LICENSE) file for details.

## Support

If you encounter any issues or have questions, please:

1. Check the [documentation](https://qryma.com/documentation.html)
2. Open an issue on GitHub
3. Contact support at support@qryma.com

## Changelog

### 0.1.0
- Basic search functionality
- Simple `Qryma()` factory function for easy initialization
- Advanced search with SearchOptions
- Result extraction methods
- API status check
- Error handling
- Comprehensive data models
