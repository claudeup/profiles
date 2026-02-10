# claudeup profiles

Community profiles for [claudeup](https://github.com/claudeup/claudeup) -- pre-configured Claude Code setups you can apply in one command.

## What are profiles?

Profiles bundle plugins, MCP servers, marketplaces, and local extensions for specific development scenarios. Instead of manually installing plugins one by one, apply a complete profile designed for your workflow.

Profiles come in two types:

- **Standard profiles** define plugins, marketplaces, and MCP servers directly
- **Composable profiles** use `includes` to combine other profiles into curated stacks

## Browse Profiles

### Base

- **[base](profiles/base.json)** -- Foundation profile with core marketplaces, plugins, and local extensions

### Languages

- **[go](profiles/languages/go.json)** -- Go development with gopls LSP
- **[javascript-typescript](profiles/languages/javascript-typescript.json)** -- JavaScript and TypeScript development
- **[python](profiles/languages/python.json)** -- Python development

### Platforms

- **[backend](profiles/platforms/backend.json)** -- Backend development and API security
- **[frontend](profiles/platforms/frontend.json)** -- Frontend development: Next.js, Tailwind, shadcn, Vercel
- **[mobile](profiles/platforms/mobile.json)** -- Mobile development and security

### Stacks (composable)

Stacks combine multiple profiles using `includes` for complete development environments.

- **[essentials](profiles/stacks/essentials.json)** -- Core tools: memory, superpowers, git workflow, code quality
- **[fullstack-go](profiles/stacks/fullstack-go.json)** -- Go fullstack: essentials + Go, backend, frontend, testing
- **[fullstack-js](profiles/stacks/fullstack-js.json)** -- JavaScript/TypeScript fullstack: essentials + JS/TS, frontend, backend, testing
- **[fullstack-python](profiles/stacks/fullstack-python.json)** -- Python fullstack: essentials + Python, backend, testing
- **[mobile-dev](profiles/stacks/mobile-dev.json)** -- Mobile development: essentials + mobile, frontend, testing
- **[secure-dev](profiles/stacks/secure-dev.json)** -- Security-focused development: essentials + security, testing, debugging

### Tools

- **[ai-tools](profiles/tools/ai-tools.json)** -- AI chat, voice, and workflow tools from claude-code-tools
- **[memory](profiles/tools/memory.json)** -- Memory and context persistence: episodic memory, claude-mem
- **[plugin-dev](profiles/tools/plugin-dev.json)** -- Plugin and skill development for Claude Code
- **[superpowers](profiles/tools/superpowers.json)** -- Superpowers suite: enhanced capabilities and workflows

### Workflow

- **[code-quality](profiles/workflow/code-quality.json)** -- Code review, refactoring, and style enforcement
- **[debugging](profiles/workflow/debugging.json)** -- Debugging and error analysis
- **[documentation](profiles/workflow/documentation.json)** -- Code documentation and generation
- **[git](profiles/workflow/git.json)** -- Git workflow: commits, PRs, and branch management
- **[productivity](profiles/workflow/productivity.json)** -- Developer productivity, collaboration, and shell scripting
- **[security](profiles/workflow/security.json)** -- Security scanning, compliance, and safety hooks
- **[testing](profiles/workflow/testing.json)** -- Testing, TDD, and performance testing

## Installation

### Using claudeup CLI

```bash
# Copy a profile to your profiles directory
cp profiles/stacks/fullstack-go.json ~/.claudeup/profiles/

# Apply the profile
claudeup profile apply fullstack-go
```

### Manual Installation

1. Download the profile JSON file:

   ```bash
   curl -O https://raw.githubusercontent.com/claudeup/profiles/main/profiles/stacks/fullstack-go.json
   ```

2. Copy it to your claudeup profiles directory:

   ```bash
   mkdir -p ~/.claudeup/profiles
   cp fullstack-go.json ~/.claudeup/profiles/
   ```

3. Apply the profile:
   ```bash
   claudeup profile apply fullstack-go
   ```

## Profile Structure

### Standard profile

Defines plugins and marketplaces directly, optionally scoped to user or project level:

```json
{
  "name": "backend",
  "description": "Backend development and API security",
  "marketplaces": [{ "source": "github", "repo": "wshobson/agents" }],
  "perScope": {
    "project": {
      "plugins": [
        "backend-api-security@claude-code-workflows",
        "backend-development@claude-code-workflows"
      ]
    }
  }
}
```

### Composable profile

Combines other profiles using `includes`:

```json
{
  "name": "fullstack-go",
  "description": "Go fullstack: essentials + Go, backend, frontend, testing",
  "includes": ["essentials", "go", "backend", "frontend", "testing"]
}
```

When applied, claudeup resolves all included profiles recursively, merging their plugins and marketplaces.

### Profile fields

| Field          | Required | Description                                         |
| -------------- | -------- | --------------------------------------------------- |
| `name`         | Yes      | Lowercase, alphanumeric with hyphens                |
| `description`  | Yes      | What the profile provides (10+ characters)          |
| `marketplaces` | Yes\*    | Plugin marketplace sources (`owner/repo` format)    |
| `includes`     | Yes\*    | List of other profiles to compose                   |
| `perScope`     | No       | Plugins scoped to `user` or `project` level         |
| `mcpServers`   | No       | MCP server configurations                           |
| `detect`       | No       | Auto-detection rules for `claudeup profile suggest` |
| `localItems`   | No       | Local extensions (agents, commands, skills, hooks)  |

\* A profile must have either `marketplaces` or `includes` (or both).

### Scope layering

Profiles use `perScope` to control where plugins are installed:

- **user** scope -- plugins available across all projects
- **project** scope -- plugins active only in the current project

See [examples/template.json](examples/template.json) for a complete annotated example.

## Profile Categories

- **languages/** -- Language-specific tools and LSP integrations
- **platforms/** -- Platform-specific development (backend, frontend, mobile)
- **stacks/** -- Composable profiles that combine others into complete environments
- **tools/** -- Cross-cutting tools (memory, AI, plugin development)
- **workflow/** -- Development workflow automation (testing, git, security, docs)

## Community Marketplaces

Profiles reference plugins from these community marketplaces:

| Marketplace                                                                                 | Description                               |
| ------------------------------------------------------------------------------------------- | ----------------------------------------- |
| [anthropics/claude-plugins-official](https://github.com/anthropics/claude-plugins-official) | Official Claude Code plugins              |
| [obra/superpowers-marketplace](https://github.com/obra/superpowers-marketplace)             | Community productivity and workflow tools |
| [wshobson/agents](https://github.com/wshobson/agents)                                       | Backend, frontend, and workflow agents    |
| [pchalasani/claude-code-tools](https://github.com/pchalasani/claude-code-tools)             | AI chat, voice, and workflow tools        |
| [thedotmack/claude-mem](https://github.com/thedotmack/claude-mem)                           | Persistent memory across sessions         |
| [davila7/claude-code-templates](https://github.com/davila7/claude-code-templates)           | Next.js, testing, and framework templates |
| [jarrodwatts/claude-hud](https://github.com/jarrodwatts/claude-hud)                         | Status line HUD for Claude Code           |

## Validation

Before submitting a profile, validate it:

```bash
./scripts/validate.sh
```

## Contributing

We welcome community contributions! See [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

### Quick Start

1. Fork this repository
2. Copy `examples/template.json` as a starting point
3. Add your profile to the appropriate category directory
4. Test with the validation script: `./scripts/validate.sh`
5. Submit a pull request

## License

MIT License -- see [LICENSE](LICENSE) for details.

## Related Projects

- [claudeup/claudeup](https://github.com/claudeup/claudeup) -- Main claudeup CLI tool
