# Contributing to claudeup/profiles

Thank you for contributing to the claudeup community! This guide will help you submit high-quality profiles.

## Submission Process

1. **Fork the repository**

   ```bash
   gh repo fork claudeup/profiles --clone
   cd profiles
   ```

2. **Create your profile**
   - Copy `examples/template.json` as a starting point
   - Place it in the appropriate category directory
   - Name the file using lowercase with hyphens: `my-profile.json`

3. **Test your profile**

   ```bash
   # Validate JSON structure
   ./scripts/validate.sh

   # Test it with claudeup
   cp profiles/category/my-profile.json ~/.claudeup/profiles/
   claudeup profile apply my-profile
   ```

4. **Submit a pull request**
   - Use the PR template
   - Include a clear description of what the profile is for
   - Mention any special requirements or dependencies

## Profile Types

### Standard profiles

Define plugins and marketplaces directly. Must include `marketplaces` and typically use `perScope` to control where plugins are installed.

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

### Composable profiles

Combine other profiles using `includes`. No `marketplaces` field needed -- plugins and marketplaces are inherited from the included profiles.

```json
{
  "name": "fullstack-go",
  "description": "Go fullstack: essentials + Go, backend, frontend, testing",
  "includes": ["essentials", "go", "backend", "frontend", "testing"]
}
```

Composable profiles belong in the `stacks/` category.

## Profile Guidelines

### Required Fields

Every profile must include:

- **name**: Lowercase, alphanumeric with hyphens (e.g., `backend`)
- **description**: Clear, detailed description (minimum 10 characters)
- **marketplaces** or **includes**: At least one (standard profiles need marketplaces, composable profiles need includes)

### Optional Fields

| Field        | Description                                         |
| ------------ | --------------------------------------------------- |
| `perScope`   | Plugins scoped to `user` or `project` level         |
| `mcpServers` | MCP server configurations (inside `perScope`)       |
| `detect`     | Auto-detection rules for `claudeup profile suggest` |
| `localItems` | Local extensions (agents, commands, skills, hooks)  |

### Scope Guidelines

Use `perScope` to control where plugins are installed:

- **user** scope -- tools useful across all projects (memory, git workflow, code review)
- **project** scope -- tools specific to a technology or framework (language LSPs, testing frameworks)

### Quality Standards

DO:

- Focus on a specific use case or technology stack
- Include only relevant plugins
- Write detailed descriptions explaining what the profile includes
- Add project detection rules for auto-suggestion
- Test the profile yourself before submitting
- Use proper secret management for MCP servers

DON'T:

- Create overly broad or generic profiles
- Include every plugin "just in case"
- Hardcode secrets, credentials, or API keys
- Use suspicious or unverified MCP server commands
- Copy profiles without adding value

## Secret Management

If your profile includes MCP servers that need API keys, use the `secrets` field with multiple sources:

```json
{
  "perScope": {
    "user": {
      "mcpServers": [
        {
          "name": "my-service",
          "command": "npx",
          "args": ["-y", "@example/mcp-server"],
          "secrets": {
            "API_KEY": {
              "description": "API key for the service",
              "sources": [
                { "type": "env", "key": "SERVICE_API_KEY" },
                {
                  "type": "1password",
                  "ref": "op://Private/Service/credential"
                },
                {
                  "type": "keychain",
                  "service": "my-service",
                  "account": "default"
                }
              ]
            }
          }
        }
      ]
    }
  }
}
```

Secret sources are tried in order. Use multiple sources to support different user setups. Never hardcode credentials in profile files.

## Categories

Choose the most appropriate category for your profile:

- **languages/** -- Language-specific tools and LSP integrations
- **platforms/** -- Platform-specific development (backend, frontend, mobile)
- **stacks/** -- Composable profiles that combine others into complete environments
- **tools/** -- Cross-cutting tools (memory, AI, plugin development)
- **workflow/** -- Development workflow automation (testing, git, security, docs)

If your profile doesn't fit any category, suggest a new one in your PR.

## Validation

Profiles are automatically validated on PR submission. Common issues:

- **Invalid JSON** -- Use `jq` or a JSON validator to check syntax
- **Missing required fields** -- Standard profiles need name, description, and marketplaces. Composable profiles need name, description, and includes.
- **Invalid marketplace format** -- Must be `owner/repository`
- **Invalid name format** -- Use lowercase letters, numbers, and hyphens only
- **Empty includes** -- Composable profiles must include at least one profile

Run validation locally before submitting:

```bash
./scripts/validate.sh
```

## Review Process

1. **Automated validation** -- GitHub Actions checks JSON structure and scans for hardcoded secrets
2. **Manual review** -- Maintainers review for:
   - Quality and usefulness
   - Security concerns (MCP commands, plugin sources, credentials)
   - Appropriate categorization
   - Clear documentation
3. **Merge** -- Approved profiles are merged and immediately available

## Maintenance

After your profile is merged:

- Update it if plugins change or become deprecated
- Respond to issues or questions about your profile
- Consider becoming a maintainer if you contribute regularly

## Code of Conduct

- Be respectful and constructive
- Focus on helping users solve real problems
- Share knowledge and best practices
- Report security issues privately to the maintainers

## Questions?

- Open an issue for general questions
- Tag maintainers in your PR for specific feedback
- Join the discussion in existing PRs to learn from others
