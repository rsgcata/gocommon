package cli

import (
	"errors"
	"fmt"
	"io"
	"maps"
	"os"
	"slices"
	"strings"
)

type InputOptionDefinition struct {
	name        string
	description string
	required    bool
	defaultVal  string
}

type InputOption struct {
	InputOptionDefinition
	rawVal string
}

type InputOptionDefinitionMap map[string]InputOptionDefinition
type InputOptionsMap map[string]InputOption

type Command interface {
	Id() string
	Description() string
	Exec(options InputOptionsMap, stdWriter io.Writer) error
	InputDefinition() InputOptionDefinitionMap
}

func BuildOptionsFrom(
	rawOptions []string,
	cmd Command,
) (InputOptionsMap, []error) {
	options := InputOptionsMap{}
	var optionErrors []error
	for _, arg := range rawOptions {
		if !strings.HasPrefix(arg, "--") {
			continue
		}
		parts := strings.SplitN(strings.TrimLeft(arg, "--"), "=", 2)
		if len(parts) < 1 {
			continue
		}

		optionName := strings.TrimSpace(parts[0])
		if _, exists := options[optionName]; exists {
			optionErrors = append(
				optionErrors,
				fmt.Errorf("option '%s' is defined twice", optionName),
			)
		}

		optionValue := ""
		if len(parts) == 2 {
			optionValue = strings.TrimSpace(parts[1])
		}

		options[optionName] = InputOption{
			InputOptionDefinition: cmd.InputDefinition()[optionName],
			rawVal:                optionValue,
		}
	}

	for _, optionDef := range cmd.InputDefinition() {
		option, optionSet := options[optionDef.name]
		if optionDef.required && (!optionSet || option.rawVal == "") {
			optionErrors = append(
				optionErrors,
				fmt.Errorf("option '%s' is required", optionDef.name),
			)
		}
	}

	return options, optionErrors
}

func runCommand(cmd Command, rawOptions []string, stdOutWriter io.Writer) {
	optionsMap, errs := BuildOptionsFrom(rawOptions, cmd)
	if len(errs) > 0 {
		fmt.Printf(
			"Failed to execute command %s with error: %s\n",
			cmd.Id(),
			errors.Join(errs...).Error(),
		)
	}
	if cmdErr := cmd.Exec(optionsMap, stdOutWriter); cmdErr != nil {
		fmt.Printf(
			"Failed to execute command %s with error: %s\n",
			cmd.Id(),
			cmdErr.Error(),
		)
	}
}

func parseCmdInput(args []string) (cmdName string, rawOptions []string) {
	if len(args) > 1 {
		if args[0] == "--" {
			args = args[1:]
		}
	}

	if len(args) != 0 {
		cmdName = strings.TrimSpace(args[0])
		rawOptions = args[1:]
	}

	return
}

type CommandsRegistry struct {
	commands map[string]Command
}

func (registry *CommandsRegistry) Register(cmd Command) error {
	if _, exists := registry.commands[cmd.Id()]; exists {
		return fmt.Errorf("command '%s' is already registered", cmd.Id())
	}
	registry.commands[cmd.Id()] = cmd
	return nil
}

func (registry *CommandsRegistry) Commands() map[string]Command {
	cmdCopy := make(map[string]Command, len(registry.commands))
	for name, cmd := range registry.commands {
		cmdCopy[name] = cmd
	}
	return cmdCopy
}

// Bootstrap Will bootstrap everything needed for the user CLI request. Will process the
// user input and run the requested command
func Bootstrap(
	args []string,
	availableCommands CommandsRegistry,
) {
	_ = availableCommands.Register(
		&HelpCommand{slices.Collect(maps.Values(availableCommands.Commands()))},
	)
	cmdName, rawOptions := parseCmdInput(args)
	if cmdName == "" {
		cmdName = (&HelpCommand{}).Id()
	}

	stdOutWriter := os.Stdout
	for _, cmd := range availableCommands.Commands() {
		if cmdName == cmd.Id() {
			runCommand(cmd, rawOptions, stdOutWriter)
			return
		}
	}

	fmt.Printf("The command %s does not exist\n", cmdName)
}
