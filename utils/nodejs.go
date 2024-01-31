package utils

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var genNode = &cobra.Command{
	Use:   "node-gen",
	Short: "Generate nodejs project",
	Long: `Generate nodejs project
	Example usage:
	nixjs node-gen -D ./projects/myapp -l "lib1 lib2" -d "lib3 lib4"`,
	Run: generateNodeJS,
}
var genRepo = &cobra.Command{
	Use:   "gen-repo",
	Short: "Generate CRUD node template",
	Long:  `Generate CRUD node tempalte`,
	Run:   generateGitHubRepo,
}

func init() {
	rootCmd.AddCommand(genNode)
	rootCmd.AddCommand(genRepo)
	genNode.Flags().StringP("directory", "D", "./myApp", "Output directory for the project")
	genNode.Flags().BoolP("yes", "y", false, "Generate default NodeJs package.json file")
	genNode.Flags().StringP("libs", "l", " ", "List of Node.js libraries to install")
	genNode.Flags().StringP("dev-libs", "d", " ", "List of Node.js libraries to install")
}

func generateNodeJS(cmd *cobra.Command, args []string) {
	dir, _ := cmd.Flags().GetString("directory")
	createDirectory(dir)
	checkNPMInstallation()

	changeDirectory(dir)

	initNodeProject(cmd)

	installLibraries(cmd, "libs")
	installLibraries(cmd, "dev-libs")
}

func generateGitHubRepo(cmd *cobra.Command, args []string) {
	// Assuming the GitHub repository is https://github.com/Faanilo/API-EXPRESS
	repoURL := "https://github.com/Faanilo/API-EXPRESS.git"
	dir, _ := cmd.Flags().GetString("directory")

	fmt.Println("Generating GitHub repository...")

	// Clone the repository
	cmdClone := exec.Command("git", "clone", repoURL, dir)
	cmdClone.Stdout = os.Stdout
	cmdClone.Stderr = os.Stderr
	err := cmdClone.Run()

	if err != nil {
		fmt.Println("Error cloning repository:", err)
		return
	}

	fmt.Println("GitHub repository generated successfully!")
}
