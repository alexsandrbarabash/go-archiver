package cmd

import (
	"archiver/lib/vlc"
	"io"
	"os"
	"github.com/spf13/cobra"
)

var vlcUnpackCmd = &cobra.Command{
	Use:   "vlc",
	Short: "Unpack file using variable-length code test",
	Run:   unpack,
}

func unpack(_ *cobra.Command, args []string) {

	if len(args) == 0 || args[0] == "" {
		handleErr(ErrEmptyPath)
	}

	filePath := args[0]

	r, err := os.Open(filePath)
	if err != nil {
		handleErr(err)
	}
	defer r.Close()

	data, err := io.ReadAll(r)

	if err != nil {
		handleErr(err)
	}

	packed := vlc.Decode(data)

	err = os.WriteFile(unpackedFileName(filePath), []byte(packed), 0644)
	if err != nil {
		handleErr(err)
	}
}

func init() {
	unpackCmd.AddCommand(vlcUnpackCmd)
}
