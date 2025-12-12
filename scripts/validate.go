// ABOUTME: Go validation script for profile JSON files
// ABOUTME: Used in CI/CD to validate profile structure and content
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Profile struct {
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Plugins      []string      `json:"plugins"`
	MCPServers   []MCPServer   `json:"mcpServers"`
	Marketplaces []Marketplace `json:"marketplaces"`
	Detect       *Detection    `json:"detect,omitempty"`
}

type MCPServer struct {
	Name    string                 `json:"name"`
	Command string                 `json:"command"`
	Args    []string               `json:"args"`
	Scope   string                 `json:"scope,omitempty"`
	Secrets map[string]interface{} `json:"secrets,omitempty"`
}

type Marketplace struct {
	Source string `json:"source"`
	Repo   string `json:"repo"`
}

type Detection struct {
	Files    []string          `json:"files,omitempty"`
	Contains map[string]string `json:"contains,omitempty"`
}

type ValidationError struct {
	File    string
	Message string
}

type ValidationWarning struct {
	File    string
	Message string
}

func main() {
	errors := []ValidationError{}
	warnings := []ValidationWarning{}
	checked := 0

	fmt.Println("Validating profiles...")
	fmt.Println()

	err := filepath.Walk("profiles", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(path, ".json") {
			return nil
		}

		checked++
		relPath := strings.TrimPrefix(path, "profiles/")
		fmt.Printf("Checking %s... ", relPath)

		data, err := os.ReadFile(path)
		if err != nil {
			fmt.Println("❌ FAILED")
			errors = append(errors, ValidationError{
				File:    relPath,
				Message: fmt.Sprintf("Failed to read file: %v", err),
			})
			return nil
		}

		var profile Profile
		if err := json.Unmarshal(data, &profile); err != nil {
			fmt.Println("❌ FAILED")
			errors = append(errors, ValidationError{
				File:    relPath,
				Message: fmt.Sprintf("Invalid JSON: %v", err),
			})
			return nil
		}

		// Validate required fields
		if profile.Name == "" {
			fmt.Println("❌ FAILED")
			errors = append(errors, ValidationError{
				File:    relPath,
				Message: "Missing required field: name",
			})
			return nil
		}

		if profile.Description == "" {
			fmt.Println("❌ FAILED")
			errors = append(errors, ValidationError{
				File:    relPath,
				Message: "Missing required field: description",
			})
			return nil
		}

		if len(profile.Marketplaces) == 0 {
			fmt.Println("❌ FAILED")
			errors = append(errors, ValidationError{
				File:    relPath,
				Message: "Missing required field: marketplaces (must have at least one)",
			})
			return nil
		}

		// Validate name format
		nameRegex := regexp.MustCompile(`^[a-z0-9-]+$`)
		if !nameRegex.MatchString(profile.Name) {
			warnings = append(warnings, ValidationWarning{
				File:    relPath,
				Message: "Name should be lowercase alphanumeric with hyphens only",
			})
		}

		// Validate description length
		if len(profile.Description) < 10 {
			warnings = append(warnings, ValidationWarning{
				File:    relPath,
				Message: "Description should be at least 10 characters",
			})
		}

		// Validate plugin count
		if len(profile.Plugins) > 10 {
			warnings = append(warnings, ValidationWarning{
				File:    relPath,
				Message: fmt.Sprintf("Profile has many plugins (%d). Consider splitting or removing unused ones.", len(profile.Plugins)),
			})
		}

		// Validate marketplace format
		repoRegex := regexp.MustCompile(`^[^/]+/[^/]+$`)
		for _, marketplace := range profile.Marketplaces {
			if marketplace.Source != "github" {
				errors = append(errors, ValidationError{
					File:    relPath,
					Message: fmt.Sprintf("Invalid marketplace source: %s (only 'github' is supported)", marketplace.Source),
				})
			}
			if !repoRegex.MatchString(marketplace.Repo) {
				fmt.Println("❌ FAILED")
				errors = append(errors, ValidationError{
					File:    relPath,
					Message: fmt.Sprintf("Invalid marketplace repo format: %s (should be owner/repo)", marketplace.Repo),
				})
				return nil
			}
		}

		fmt.Println("✅ PASSED")
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking profiles directory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("================================")
	fmt.Println("Validation Summary")
	fmt.Println("================================")
	fmt.Printf("Profiles checked: %d\n", checked)
	fmt.Printf("✅ Passed: %d\n", checked-len(errors))

	if len(warnings) > 0 {
		fmt.Printf("⚠️  Warnings: %d\n", len(warnings))
		fmt.Println()
		for _, w := range warnings {
			fmt.Printf("  %s: %s\n", w.File, w.Message)
		}
	}

	if len(errors) > 0 {
		fmt.Printf("❌ Failed: %d\n", len(errors))
		fmt.Println()
		for _, e := range errors {
			fmt.Printf("  %s: %s\n", e.File, e.Message)
		}
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("✅ All profiles are valid!")
}
