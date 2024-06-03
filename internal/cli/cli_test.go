package cli

import "testing"

func TestNewCLI(t *testing.T) {
	tests := []struct {
		name          string
		wantUse       string
		wantShort     string
		wantLong      string
		subCmdArgs    []string
		wantSubCmdUse string
	}{
		{
			name:          "Test NewCli",
			wantUse:       "logparser",
			wantShort:     "LogParser CLI application",
			wantLong:      "Logparser CLI application\nUse this CLI for all your log parsing needs.",
			subCmdArgs:    []string{"log"},
			wantSubCmdUse: "log -- ip unique - ip active $(IP_COUNT) - url top $(URL_COUNT)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd := NewCLI()

			// Verify the root command properties
			if rootCmd.Use != tt.wantUse {
				t.Errorf("Use: got %s, want %s", rootCmd.Use, tt.wantUse)
			}
			if rootCmd.Short != tt.wantShort {
				t.Errorf("Short: got %s, want %s", rootCmd.Short, tt.wantShort)
			}
			if rootCmd.Long != tt.wantLong {
				t.Errorf("Long: got %s, want %s", rootCmd.Long, tt.wantLong)
			}

			// Verify subcommand properties
			subCmd, _, err := rootCmd.Find(tt.subCmdArgs)
			if err != nil {
				t.Errorf("Unexpected error finding subcommand: %v", err)
			}
			if subCmd == nil {
				t.Error("Subcommand not found")
			} else {
				if subCmd.Use != tt.wantSubCmdUse {
					t.Errorf("Subcommand Use: got %s, want %s", subCmd.Use, tt.wantSubCmdUse)
				}
			}
		})
	}
}
