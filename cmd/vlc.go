package cmd

import (
	"errors"
	"github.com/spf13/cobra"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var vlcCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Pack file using variable-length code",
	Run:   pack,
}

const packedFileExtension = "vlc"

var ErrEmptyPath = errors.New("path is not specified")

func pack(_ *cobra.Command, args []string) {
	if len(args) == 0 || args[0] == "" {
		handleError(ErrEmptyPath)
	}

	filePath := args[0]
	r, err := os.Open(filePath)
	if err != nil {
		handleError(err)
	}
	data, err := io.ReadAll(r)
	if err != nil {
		handleError(err)
	}

	//packed := Encode(data)
	packed := "" + string(data)

	err = os.WriteFile(packedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleError(err)
	}

	defer func(r *os.File) {
		err := r.Close()
		if err != nil {
			handleError(err)
		}
	}(r)

}

func packedFileName(path string) string {
	fileName := filepath.Base(path)
	return strings.TrimSuffix(fileName, filepath.Ext(fileName)) + "." + packedFileExtension
}

func init() {
	packCmd.AddCommand(vlcCmd)
}
