// 代码生成时间: 2025-10-12 21:43:49
package main

import (
    "fmt"
    "os"
    "path/filepath"
    "text/template"
    "golang.org/x/exp/slices"

    "github.com/spf13/cobra"
)

// Template for program generation
const progTemplateStr = `// +build ignore

package main

import (
    "fmt"
)

func main() {
    fmt.Println("{{.HelloMessage}}")
}
`

// HelloMessage is a struct that holds the message to be printed by the generated program
type HelloMessage struct {
    Message string
}

func main() {
    var cmd = &cobra.Command{
        Use:   "generate",
        Short: "Generates a program that prints a greeting message",
        Long:  `Generates a program that prints a greeting message.`,
        RunE:  generateProgramCmd,
    }

    err := cmd.Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v
", err)
        os.Exit(1)
    }
}

// generateProgramCmd is the command that generates a program based on a template
func generateProgramCmd(cmd *cobra.Command, args []string) error {
    // Default message
    msg := HelloMessage{Message: "Hello, World!"}

    // Check if a custom message is provided
    if len(args) > 0 {
        msg.Message = args[0]
    }

    // Create a new template from the string
    template, err := template.New("program").Parse(progTemplateStr)
    if err != nil {
        return err
    }

    // Create a file for the generated program
    f, err := os.Create(filepath.Join(os.TempDir(), "hello_program.go"))
    if err != nil {
        return err
    }
    defer f.Close()

    // Execute the template with the message
    return template.Execute(f, msg)
}