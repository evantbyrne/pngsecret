package main

import (
	"github.com/alecthomas/kingpin"
	lib "github.com/evantbyrne/pngsecret/lib"
	"os"
)

var (
	app = kingpin.New("pngsecret", "Hide messages in PNG images.")

	appDecode       = app.Command("decode", "Decode a message.")
	appDecodeFileIn = appDecode.Arg("file_in", "File to read from.").Required().String()

	appEncode        = app.Command("encode", "Encode a message.")
	appEncodeFileIn  = appEncode.Arg("file_in", "File to read from.").Required().String()
	appEncodeFileOut = appEncode.Arg("file_out", "File to write to.").Required().String()
	appEncodeMessage = appEncode.Arg("message", "Message to encode.").Required().String()
)

func main() {
	var (
		kp = kingpin.MustParse(app.Parse(os.Args[1:]))
	)

	switch kp {

	case appDecode.FullCommand():
		lib.CommandDecode(*appDecodeFileIn)

	case appEncode.FullCommand():
		lib.CommandEncode(*appEncodeFileIn, *appEncodeFileOut, *appEncodeMessage)

	}
}
