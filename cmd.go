package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type cmdConfig struct {
	// --out
	// --out-stdout
	OutStdout       bool
	OutStdoutFormat string

	// --out-json
	OutJson string

	// --out-github-pr
	OutGithubPr string

	// --out-curl
	OutCurl string
}

var rootCmd = &cobra.Command{
	Use:   "committer-coverage-stats <path-to-repo> [flags]",
	Short: "Committer Coverage Stats",
	Long:  `Committer Coverage Stats`,
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	fmt.Println("Running...")
}

func init() {
	parseOutputFlags(rootCmd)

	bindConfig(rootCmd)
}

func parseOutputFlags(cmd *cobra.Command) {
	parseOutputStdout(cmd)
	parseOutputJson(cmd)
	parseOutputGithubPr(cmd)
	parseOutputCurl(cmd)
}

// parseOutputStdout parses the output flags for standard output
func parseOutputStdout(cmd *cobra.Command) {
	cmd.Flags().BoolP("out-stdout", "", true, "Enable standard output (default)")
	// format
	cmd.Flags().StringP("out-stdout-format", "", "plain", "Set stdout format (default: plain, options: plain, colored, verbose)")

	viper.BindPFlag("out-stdout-format", cmd.Flags().Lookup("out-stdout-format"))
}

func parseOutputJson(cmd *cobra.Command) {
	cmd.Flags().StringP("out-json", "", "", "Enable JSON output")
}

func parseOutputGithubPr(cmd *cobra.Command) {
	cmd.Flags().StringP("out-github-pr", "", "", "Enable GitHub Pull Request comment output")
}

func parseOutputCurl(cmd *cobra.Command) {
	cmd.Flags().StringP("out-curl", "", "", "Enable CURL output to specified URL")
}

func bindConfig(cmd *cobra.Command) {
	err := viper.Unmarshal(&cmdConfig{})
	if err != nil {
		fmt.Println("Error binding config:", err)
		os.Exit(1)
	}
}
