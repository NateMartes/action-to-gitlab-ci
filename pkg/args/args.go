package args

import (
	"github.com/spf13/pflag"
	"os"
	"fmt"
	"bytes"
)

// Get the usage message for action-to-gitlab-ci
func UsageMsg() string {
	var buf bytes.Buffer
    pflag.CommandLine.SetOutput(&buf)
    fmt.Fprintf(&buf, "usage: action-to-gitlab-ci [-dh] <file>\n")
    fmt.Fprintf(&buf, "Convert Github Workflows to Gitlab CI\n")
    fmt.Fprintf(&buf, "Options:\n")
    pflag.PrintDefaults()
    pflag.CommandLine.SetOutput(os.Stderr)
    return buf.String()
}

// Parse arguments from the command line
func ParseArgs() {
	
	// Args
	pflag.BoolP("debug", "d", false, "Turn on debug messaging")
	pflag.BoolP("help", "h", false, "Display help message")
	
	pflag.Parse()
}

// Get the value of the help argument
func HelpVal() (bool, error) {
    v, err := pflag.CommandLine.GetBool("help")
    return v, err
}

// Get the value of the debug argument
func DebugVal() (bool, error) {
    v, err := pflag.CommandLine.GetBool("debug")
    return v, err
}