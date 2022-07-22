package arena

import (
	"os"
	"os/exec"
	"strings"

	"github.com/diambra/cli/pkg/diambra"
	"github.com/diambra/cli/pkg/log"
	"github.com/diambra/cli/pkg/pyarena"
	"github.com/go-kit/log/level"
	"github.com/spf13/cobra"
)

var romScripts = map[string]string{
	"check roms": pyarena.CheckRoms,
	"list roms":  pyarena.ListRoms,
}

func findPython() string {
	for _, name := range []string{
		"python",
		"python3",
	} {
		_, err := exec.LookPath(name)
		if err == nil {
			return name
		}
	}
	return "python"
}

func NewRomCmds(logger *log.Logger) ([]*cobra.Command, error) {
	c, err := diambra.NewConfig()
	if err != nil {
		level.Error(logger).Log("msg", err.Error())
		os.Exit(1)
	}
	var (
		pythonPath string
	)
	cmds := []*cobra.Command{}
	for name, script := range romScripts {
		level.Debug(logger).Log("msg", "hello") // "name", name, "script", script)
		cmd := &cobra.Command{
			Use:   strings.ReplaceAll(name, " ", "-"),
			Short: name,
			Long:  "This command runs the " + name + " rom utility",
			Run: func(_ *cobra.Command, args []string) {
				cmd := exec.Command(pythonPath, "-c", script)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				cmd.Args = append(cmd.Args, args...)
				cmd.Env = append(os.Environ(),
					"DIAMBRAROMSPATH="+c.RomsPath,
				)
				if err := cmd.Run(); err != nil {
					level.Error(logger).Log("msg", "command failed", "err", err.Error())
					os.Exit(1)
				}
			},
		}
		c.AddRomsPathFlag(cmd.Flags())
		cmd.Flags().StringVar(&pythonPath, "python", findPython(), "Path to python executable")
		cmds = append(cmds, cmd)
	}
	return cmds, nil
}
