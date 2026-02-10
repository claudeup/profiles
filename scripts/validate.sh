#!/bin/bash
# ABOUTME: Validates all profile JSON files in the repository
# ABOUTME: Checks for valid JSON syntax and required fields

set -e

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

PROFILES_DIR="profiles"
ERRORS=0
WARNINGS=0
CHECKED=0

echo "Validating profiles..."
echo ""

# Find all JSON files in profiles directory
while IFS= read -r -d '' profile; do
    CHECKED=$((CHECKED + 1))
    profile_name=$(basename "$profile")
    category=$(basename "$(dirname "$profile")")

    echo -n "Checking ${category}/${profile_name}... "

    # Check valid JSON
    if ! jq empty "$profile" 2>/dev/null; then
        echo -e "${RED}FAILED${NC}"
        echo "  ❌ Invalid JSON syntax"
        ERRORS=$((ERRORS + 1))
        continue
    fi

    # Check required fields
    has_error=false

    if ! jq -e '.name' "$profile" >/dev/null 2>&1; then
        echo -e "${RED}FAILED${NC}"
        echo "  ❌ Missing required field: name"
        has_error=true
    fi

    if ! jq -e '.description' "$profile" >/dev/null 2>&1; then
        echo -e "${RED}FAILED${NC}"
        echo "  ❌ Missing required field: description"
        has_error=true
    fi

    # Check if profile uses includes (composable profile)
    has_includes=$(jq -e '.includes' "$profile" >/dev/null 2>&1 && echo "true" || echo "false")

    # Marketplaces are required unless profile uses includes
    if [ "$has_includes" = "false" ]; then
        if ! jq -e '.marketplaces' "$profile" >/dev/null 2>&1; then
            echo -e "${RED}FAILED${NC}"
            echo "  ❌ Missing required field: marketplaces (or includes)"
            has_error=true
        fi
    fi

    # Validate includes format if present
    if [ "$has_includes" = "true" ]; then
        if ! jq -e '.includes | type == "array"' "$profile" >/dev/null 2>&1; then
            echo -e "${RED}FAILED${NC}"
            echo "  ❌ includes must be an array"
            has_error=true
        fi

        include_count=$(jq -r '.includes | length' "$profile" 2>/dev/null || echo "0")
        if [ "$include_count" -eq 0 ]; then
            echo -e "${RED}FAILED${NC}"
            echo "  ❌ includes array cannot be empty"
            has_error=true
        fi
    fi

    if [ "$has_error" = true ]; then
        ERRORS=$((ERRORS + 1))
        continue
    fi

    # Check name format (lowercase, alphanumeric + hyphens)
    name=$(jq -r '.name' "$profile")
    if ! echo "$name" | grep -qE '^[a-z0-9-]+$'; then
        echo -e "${YELLOW}WARNING${NC}"
        echo "  ⚠️  Name should be lowercase alphanumeric with hyphens only: $name"
        WARNINGS=$((WARNINGS + 1))
    fi

    # Check description length
    desc_len=$(jq -r '.description | length' "$profile")
    if [ "$desc_len" -lt 10 ]; then
        echo -e "${YELLOW}WARNING${NC}"
        echo "  ⚠️  Description should be at least 10 characters"
        WARNINGS=$((WARNINGS + 1))
    fi

    # Check plugin count
    plugin_count=$(jq -r '.plugins | length' "$profile" 2>/dev/null || echo "0")
    if [ "$plugin_count" -gt 10 ]; then
        echo -e "${YELLOW}WARNING${NC}"
        echo "  ⚠️  Profile has many plugins ($plugin_count). Consider splitting or removing unused ones."
        WARNINGS=$((WARNINGS + 1))
    fi

    # Check marketplace format (skip for composable profiles)
    if [ "$has_includes" = "false" ]; then
        invalid_marketplace=false
        while IFS= read -r marketplace; do
            if ! echo "$marketplace" | grep -qE '^[^/]+/[^/]+$'; then
                echo -e "${RED}FAILED${NC}"
                echo "  ❌ Invalid marketplace repo format: $marketplace (should be owner/repo)"
                invalid_marketplace=true
            fi
        done < <(jq -r '.marketplaces[]?.repo // empty' "$profile")

        if [ "$invalid_marketplace" = true ]; then
            ERRORS=$((ERRORS + 1))
            continue
        fi
    fi

    echo -e "${GREEN}PASSED${NC}"

done < <(find "$PROFILES_DIR" -name "*.json" -type f -print0)

echo ""
echo "================================"
echo "Validation Summary"
echo "================================"
echo "Profiles checked: $CHECKED"
echo -e "${GREEN}Passed: $((CHECKED - ERRORS))${NC}"
if [ $WARNINGS -gt 0 ]; then
    echo -e "${YELLOW}Warnings: $WARNINGS${NC}"
fi
if [ $ERRORS -gt 0 ]; then
    echo -e "${RED}Failed: $ERRORS${NC}"
    exit 1
fi

echo ""
echo "✅ All profiles are valid!"
exit 0
