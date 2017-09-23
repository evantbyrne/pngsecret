# pngsecret

Hide messages in PNG images. Encodes data with subtle transparency changes.


## Install

Make sure that [Go](https://golang.org/) is installed and your PATH includes GOBIN. Then run the following:

    $ go get -u github.com/evantbyrne/pngsecret

**Note:** The `$` at the beginning of newlines in this document represents the bash shell prompt, and is not a part of the actual commands.


## Usage

    $ pngsecret
    usage: pngsecret [<flags>] <command> [<args> ...]

    Hide messages in PNG images.

    Flags:
      --help  Show context-sensitive help (also try --help-long and --help-man).

    Commands:
      help [<command>...]
        Show help.

      decode <file_in>
        Decode a message.

      encode <file_in> <file_out> <message>
        Encode a message.


Add a secret message to a PNG:

    $ pngsecret encode cat.png totally-normal-cat-picture.png "Hello, World"


Decode a PNG:

    $ pngsecret decode totally-normal-cat-picture.png
    Hello, World
