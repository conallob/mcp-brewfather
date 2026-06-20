# mcp-brewfather

A [Model Context Protocol (MCP)](https://modelcontextprotocol.io) server for the [Brewfather](https://brewfather.app) homebrewing API. Exposes your Brewfather data — batches, recipes, and inventory — as MCP tools so AI assistants can help you brew.

## Prerequisites

- A [Brewfather](https://brewfather.app) account with API access enabled
- Brewfather User ID and API Key (Settings → API)

## Installation

### Homebrew

```sh
brew tap conallob/taps
brew install mcp-brewfather
```

### Docker

```sh
docker pull ghcr.io/conallob/mcp-brewfather:latest
```

### From source

```sh
go install github.com/conallob/mcp-brewfather/cmd/mcp-brewfather@latest
```

## Configuration

Set the required environment variables:

```sh
export BREWFATHER_USER_ID=your_user_id
export BREWFATHER_API_KEY=your_api_key
```

## Usage with Claude Desktop

Add to your Claude Desktop config (`~/Library/Application Support/Claude/claude_desktop_config.json`):

```json
{
  "mcpServers": {
    "brewfather": {
      "command": "mcp-brewfather",
      "env": {
        "BREWFATHER_USER_ID": "your_user_id",
        "BREWFATHER_API_KEY": "your_api_key"
      }
    }
  }
}
```

### Docker variant

```json
{
  "mcpServers": {
    "brewfather": {
      "command": "docker",
      "args": ["run", "--rm", "-i",
        "-e", "BREWFATHER_USER_ID",
        "-e", "BREWFATHER_API_KEY",
        "ghcr.io/conallob/mcp-brewfather:latest"
      ],
      "env": {
        "BREWFATHER_USER_ID": "your_user_id",
        "BREWFATHER_API_KEY": "your_api_key"
      }
    }
  }
}
```

## Available Tools

| Tool | Description |
|---|---|
| `list_batches` | List brewing batches, optionally filtered by status |
| `get_batch` | Get full details of a batch by ID |
| `get_batch_readings` | Get fermentation sensor readings for a batch |
| `list_recipes` | List brewing recipes |
| `get_recipe` | Get full details of a recipe by ID |
| `list_fermentables` | List fermentable ingredients in inventory |
| `list_hops` | List hop ingredients in inventory |
| `list_yeasts` | List yeast ingredients in inventory |
| `list_miscs` | List miscellaneous ingredients in inventory |

### Batch statuses

`Planning`, `Brewing`, `Fermenting`, `Conditioning`, `Completed`, `Archived`

## Release process

Releases are driven by pushing a `v*` tag:

```sh
git tag v1.0.0
git push origin v1.0.0
```

GitHub Actions will:
1. Build binaries for Linux (amd64/arm64), macOS (amd64/arm64), and Windows (amd64)
2. Publish multi-arch container images to `ghcr.io/conallob/mcp-brewfather`
3. Update the Homebrew formula in [conallob/homebrew-taps](https://github.com/conallob/homebrew-taps)

The `HOMEBREW_TAP_GITHUB_TOKEN` secret must be set in the repository with `repo` write access to `conallob/homebrew-taps`.

## Development

```sh
go test ./...
go build ./cmd/mcp-brewfather
```

## License

MIT
