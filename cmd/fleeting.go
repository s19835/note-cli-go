/*
Copyright © 2025 s19835 <s19835@sci.pdn.ac.lk>
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

// fleetingCmd represents the fleeting command
var fleetingCmd = &cobra.Command{
	Use:   "fleeting [note]",
	Short: "Create a fleeting note",
	Long: `Purpose: Quick brain dumps, random thoughts, ideas, overheard quotes, sudden inspiration. Notes which are temporary and meant to be processed later.
	Title: YYYYMMDDHHMM.md, Location: notes/fleeting/`,
	Run: func(cmd *cobra.Command, args []string) {
		note := strings.Join(args, " ")
		now := time.Now()
		filename := now.Format("200601021504") + ".md"

		baseDir := "notes/fleeting"
		os.MkdirAll(baseDir, os.ModePerm)

		fullpath := filepath.Join(baseDir, filename)

		content := fmt.Sprintf(`# %s - %s`, now.Format("2006-01-02 15:04"), note)

		err := os.WriteFile(fullpath, []byte(content), 0644)
		if err != nil {
			fmt.Println("Fail to create Note:", err)
			return
		}
		fmt.Println("# Creates ", fullpath)
	},
}

func init() {
	rootCmd.AddCommand(fleetingCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fleetingCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fleetingCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
