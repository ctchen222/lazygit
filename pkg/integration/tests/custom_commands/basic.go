package custom_commands

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var Basic = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Using a custom command to create a new file",
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
				Command: "touch myfile",
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
			Press("a").
			Lines(
				Contains("myfile"),
			)
	},
})
