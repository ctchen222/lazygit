package custom_commands

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var MultiplePrompts = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Using a custom command with multiple prompts",
	ExtraCmdArgs: "",
	Skip:         false,
	SetupRepo: func(shell *Shell) {
		shell.EmptyCommit("blah")
	},
	SetupConfig: func(cfg *config.AppConfig) {
		cfg.UserConfig.CustomCommands = []config.CustomCommand{
			{
				Key:     "a",
				Context: "files",
				Command: `echo "{{index .PromptResponses 1}}" > {{index .PromptResponses 0}}`,
				Prompts: []config.CustomCommandPrompt{
					{
						Type:  "input",
						Title: "Enter a file name",
					},
					{
						Type:  "menu",
						Title: "Choose file content",
						Options: []config.CustomCommandMenuOption{
							{
								Name:        "foo",
								Description: "Foo",
								Value:       "FOO",
							},
							{
								Name:        "bar",
								Description: "Bar",
								Value:       "BAR",
							},
							{
								Name:        "baz",
								Description: "Baz",
								Value:       "BAZ",
							},
						},
					},
					{
						Type:  "confirm",
						Title: "Are you sure?",
						Body:  "Are you REALLY sure you want to make this file? Up to you buddy.",
					},
				},
			},
		}
	},
	Run: func(
		shell *Shell,
		input *Input,
		keys config.KeybindingConfig,
	) {
		input.Model().WorkingTreeFileCount(0)

		input.Views().Files().
			IsFocused().
			Press("a")

		input.ExpectPrompt().Title(Equals("Enter a file name")).Type("myfile").Confirm()

		input.ExpectMenu().Title(Equals("Choose file content")).Select(Contains("bar")).Confirm()

		input.ExpectConfirmation().
			Title(Equals("Are you sure?")).
			Content(Equals("Are you REALLY sure you want to make this file? Up to you buddy.")).
			Confirm()

		input.Model().WorkingTreeFileCount(1)
		input.Views().Files().SelectedLine(Contains("myfile"))
		input.Views().Main().Content(Contains("BAR"))
	},
})
