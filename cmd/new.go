package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/supreet321/GPTerminal/core"
)

func promptForMessage(promptMessage string) {
	prompt := promptui.Prompt{
		Label: promptui.Styler(promptui.FGRed)(promptMessage),
	}

	promptResult, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	if promptResult == "exit" {
		return
	}

	response := core.CreateNewChat(promptResult)
	fmt.Println(promptui.Styler(promptui.FGCyan)(response))
	promptForMessage("How else may I assist you today? Or you can type 'exit' to quit.")
}

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Starts a new chat",
	Long:  `Starts a new chat with the OpenAI API.`,
	Run: func(cmd *cobra.Command, args []string) {
		promptForMessage("How may I assist you today? Or you can type 'exit' to quit.")
		//
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
