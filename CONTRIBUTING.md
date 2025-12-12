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
   claudeup profile use my-profile
   ```

4. **Submit a pull request**
   - Use the PR template
   - Include a clear description of what the profile is for
   - Mention any special requirements or dependencies

## Profile Guidelines

### Required Fields

Every profile must include:

- **name**: Lowercase, alphanumeric with hyphens (e.g., `golang-api`)
- **description**: Clear, detailed description (minimum 10 characters)
- **marketplaces**: At least one marketplace (usually `anthropics/claude-code-plugins`)

### Optional but Recommended

- **plugins**: Specific plugins from the marketplaces
- **mcpServers**: MCP servers with proper secret handling
- **detect**: Auto-detection rules for `claudeup profile suggest`

### Quality Standards

✅ **DO:**
- Focus on a specific use case or technology stack
- Include only relevant plugins (aim for 3-7, max 10)
- Write detailed descriptions explaining what the profile includes
- Add project detection rules for auto-suggestion
- Test the profile yourself before submitting
- Include proper secret management for MCP servers

❌ **DON'T:**
- Create overly broad or generic profiles
- Include every plugin "just in case"
- Hardcode secrets or credentials
- Use suspicious or unverified MCP server commands
- Copy profiles without adding value

### Example Profile

```json
{
  "name": "golang-api",
  "description": "Go API development with REST, gRPC, and database tooling. Includes code generation, testing support, and common Go development patterns.",
  "plugins": [
    "superpowers@superpowers-marketplace",
    "commit-commands@claude-code-plugins"
  ],
  "mcpServers": [],
  "marketplaces": [
    {
      "source": "github",
      "repo": "anthropics/claude-code-plugins"
    },
    {
      "source": "github",
      "repo": "obra/superpowers-marketplace"
    }
  ],
  "detect": {
    "files": ["go.mod", "go.sum"],
    "contains": {
      "go.mod": "module "
    }
  }
}
```

## Secret Management

If your profile includes MCP servers that need API keys:

```json
{
  "mcpServers": [
    {
      "name": "my-service",
      "command": "npx",
      "args": ["-y", "@example/mcp"],
      "secrets": {
        "API_KEY": {
          "description": "API key for the service",
          "sources": [
            {"type": "env", "key": "SERVICE_API_KEY"},
            {"type": "1password", "ref": "op://Private/Service/credential"},
            {"type": "keychain", "service": "my-service", "account": "default"}
          ]
        }
      }
    }
  ]
}
```

Secret sources are tried in order. Use multiple sources to support different user setups.

## Categories

Choose the most appropriate category for your profile:

- **ai-development** - AI/ML development, LLM applications, agents
- **backend** - Server-side APIs, databases, backend frameworks
- **data-science** - Data analysis, scientific computing, notebooks
- **devops** - Infrastructure, containers, orchestration, CI/CD
- **frontend** - Web frameworks, UI development (beyond built-in profiles)
- **mobile** - React Native, Flutter, mobile app development
- **security** - Security testing, auditing, penetration testing

If your profile doesn't fit any category, suggest a new one in your PR.

## Validation

Profiles are automatically validated on PR submission. Common issues:

- **Invalid JSON** - Use `jq` or a JSON validator to check syntax
- **Missing required fields** - Ensure name, description, and marketplaces are present
- **Invalid marketplace format** - Must be `owner/repository`
- **Invalid name format** - Use lowercase letters, numbers, and hyphens only

Run validation locally before submitting:

```bash
./scripts/validate.sh
```

## Review Process

1. **Automated validation** - GitHub Actions checks JSON structure
2. **Manual review** - Maintainers review for:
   - Quality and usefulness
   - Security concerns (MCP commands, plugin sources)
   - Appropriate categorization
   - Clear documentation
3. **Merge** - Approved profiles are merged and immediately available

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

Thank you for making claudeup better! 🚀
