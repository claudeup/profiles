## Profile Submission

**Profile name:** `category/profile-name.json`

**Category:** [languages | platforms | stacks | tools | workflow]

**Profile type:** [standard | composable]

## Description

<!-- Describe what this profile is for and who should use it -->

## Use Case

<!-- Explain the specific development scenario this profile addresses -->

## What's Included

<!-- List the main components. For composable profiles, list the included profiles. -->

## **Plugins (or included profiles):**

**Scope:** [user | project]

## **Marketplaces:**

## **MCP Servers (if any):**

## Testing

- [ ] I have tested this profile with `claudeup profile apply` locally
- [ ] Profile passes validation (`./scripts/validate.sh`)
- [ ] All plugins work as expected at the specified scope
- [ ] No hardcoded credentials or API keys
- [ ] Secrets use proper secret sources (env, 1password, keychain)
- [ ] Project detection rules work correctly (if included)
- [ ] Composable profiles resolve correctly (if using `includes`)

## Additional Notes

<!-- Any special requirements, dependencies, or considerations -->

---

### Checklist

- [ ] Profile is in the correct category directory
- [ ] File name uses lowercase with hyphens
- [ ] Description is clear and detailed (10+ characters)
- [ ] Standard profiles have `marketplaces` field
- [ ] Composable profiles have non-empty `includes` field
- [ ] Plugins use correct `name@marketplace-ref` format
- [ ] Marketplaces use valid GitHub repo format (`owner/repo`)
- [ ] `perScope` assigns plugins to appropriate scope (user vs project)
- [ ] No hardcoded credentials or API keys
- [ ] Tested locally with claudeup

### For Maintainers

- [ ] Code review completed
- [ ] Security review completed (MCP commands, plugin sources, credentials)
- [ ] Category is appropriate
- [ ] Profile adds value to the community
