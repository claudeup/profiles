# claudeup/profiles

Community profiles for [claudeup](https://github.com/claudeup/claudeup) - discover and share Claude Code configurations.

## What are profiles?

Profiles are pre-configured setups for claudeup that bundle plugins, MCP servers, and marketplaces for specific development scenarios. Instead of manually installing plugins one by one, you can apply a complete profile designed for your workflow.

## Browse Profiles

### AI Development
- **[langchain](profiles/ai-development/langchain.json)** - LangChain development for AI agents and LLM applications

### Backend
- **[golang-api](profiles/backend/golang-api.json)** - Go API development with REST, gRPC, and database tooling
- **[python-fastapi](profiles/backend/python-fastapi.json)** - FastAPI development with async support and SQLAlchemy

### Data Science
- **[jupyter-ml](profiles/data-science/jupyter-ml.json)** - Data science and ML with Jupyter notebooks

### DevOps
- **[kubernetes](profiles/devops/kubernetes.json)** - Kubernetes manifest development and cluster management

## Installation

### Using claudeup CLI (Coming Soon)

```bash
# Install a community profile
claudeup profile install backend/golang-api

# Browse available profiles
claudeup profile browse

# Search for profiles
claudeup profile search kubernetes
```

### Manual Installation

1. Download the profile JSON file you want:
   ```bash
   curl -O https://raw.githubusercontent.com/claudeup/profiles/main/profiles/backend/golang-api.json
   ```

2. Copy it to your claudeup profiles directory:
   ```bash
   cp golang-api.json ~/.claudeup/profiles/
   ```

3. Apply the profile:
   ```bash
   claudeup profile use golang-api
   ```

## Contributing

We welcome community contributions! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on submitting new profiles.

### Quick Start

1. Fork this repository
2. Copy `examples/template.json` as a starting point
3. Add your profile to the appropriate category directory
4. Test with the validation script: `./scripts/validate.sh`
5. Submit a pull request

## Profile Categories

- **ai-development/** - AI, ML, and LLM application development
- **backend/** - Server-side API development (Go, Python, Node, etc.)
- **data-science/** - Data analysis, Jupyter, scientific computing
- **devops/** - Infrastructure, containers, CI/CD
- **frontend/** - Web development frameworks beyond Next.js
- **mobile/** - React Native, Flutter, mobile development
- **security/** - Security testing and auditing tools

## Profile Structure

Each profile is a JSON file containing:

```json
{
  "name": "my-profile",
  "description": "What this profile does",
  "plugins": [
    "plugin-name@marketplace-name"
  ],
  "mcpServers": [
    {
      "name": "server-name",
      "command": "npx",
      "args": ["-y", "@package/mcp-server"]
    }
  ],
  "marketplaces": [
    {
      "source": "github",
      "repo": "owner/repository"
    }
  ],
  "detect": {
    "files": ["config.json"],
    "contains": {"package.json": "framework-name"}
  }
}
```

See [examples/template.json](examples/template.json) for a complete annotated example.

## Validation

Before submitting a profile, validate it:

```bash
# Using bash script
./scripts/validate.sh

# Using Go (used in CI)
go run scripts/validate.go
```

## License

MIT License - see [LICENSE](LICENSE) for details.

## Related Projects

- [claudeup/claudeup](https://github.com/claudeup/claudeup) - Main claudeup CLI tool
- [anthropics/claude-code-plugins](https://github.com/anthropics/claude-code-plugins) - Official Claude Code plugins
- [obra/superpowers-marketplace](https://github.com/obra/superpowers-marketplace) - Community productivity tools
