package version

import "fmt"

var commitHash string
var branchName string
var binaryName string

const (
	major string = "1"
	minor string = "0"
)

func String() string {
	return fmt.Sprintf("%s.%s (%s/%s)", major, minor, branchName, commitHash)
}

func BinaryName() string {
	return binaryName
}
