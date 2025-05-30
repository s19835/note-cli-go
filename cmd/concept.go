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

// conceptCmd represents the concept command
var conceptCmd = &cobra.Command{
	Use:   "concept [note]",
	Short: "Add a small concept or unorganized learning",
	Long: `Purpose: Mini-notes of stuff you learned but haven't fully processed into atomic permanent notes, Messy but still valuable.
	Title: concept-YYYY-MM.md, Location: notes/concepts/`,
	Run: func(cmd *cobra.Command, args []string) {
		note := strings.Join(args, " ")
		now := time.Now()

		filename := fmt.Sprintf("concepts-%d-%02d.md", now.Year(), now.Month())
		baseDir := "notes/concepts"
		fullpath := filepath.Join(baseDir, filename)

		os.MkdirAll(baseDir, os.ModePerm)

		content := fmt.Sprintf("\n## %s\n\n- %s\n", now.Format("2006-01-02 15:04"), note)

		f, err := os.OpenFile(fullpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println("Could not write to concept file: ", err)
			return
		}
		defer f.Close()

		if _, err := f.WriteString(content); err != nil {
			fmt.Println("Error writing to file: ", err)
		} else {
			fmt.Println("Creates ", fullpath)
		}
	},
}

func init() {
	rootCmd.AddCommand(conceptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// conceptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// conceptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
