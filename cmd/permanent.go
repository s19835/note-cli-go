/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// permanentCmd represents the permanent command
var permanentCmd = &cobra.Command{
	Use:   "permanent [note]",
	Short: "Create Permanent note with a structured format",
	Long: `Fully processed ideas or concepts, backed by understanding, usually written in you own words. These are atomic, connected, and resuable.
	Title: Short descriptive title, Location: notes/permanent/`,
	Run: func(cmd *cobra.Command, args []string) {
		title := strings.Join(args, " ")
		slug := slugify(title)
		filename := slug + ".md"

		baseDir := "notes/permanent"
		fullpath := filepath.Join(baseDir, filename)

		os.MkdirAll(baseDir, os.ModePerm)

		created := time.Now().Format("2006-01-02 15:04")
		content := fmt.Sprintf(`# %s

**created:** %s

## Summary

> One-linear explanation.

## Context

Explain where you learned it, why it matters.

## Details

Explain the concept, give examples, diagrams, etc.

## Related Notes

relevent permanent notes

## Tags

[tag1] [tag2] [concept]
`, title, created)

		err := os.WriteFile(fullpath, []byte(content), 0644)
		if err != nil {
			fmt.Println("could not created permanant note")
			return
		}

		fmt.Println("Creates ", fullpath)
	},
}

func init() {
	rootCmd.AddCommand(permanentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// permanentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// permanentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "_")
	s = strings.ReplaceAll(s, "/", "_")
	return s
}
