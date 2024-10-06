# crd

``` shell
❯ crd
text2midi

Usage:
  crd [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  gen         generate codes and files
  help        Help about any command
  info        show definitions
  key         key info
  midi        midi util
  text        text processor
  write       write midi

Flags:
      --attr strings    additional attribute definitions
      --chord strings   additional chord definitions
      --debug           enable debug
  -h, --help            help for crd
  -o, --output string   output file

Use "crd [command] --help" for more information about a command.
```

## Examples

``` shell
❯ crd text conv --help
convert text to yaml

Spaces, newlines, any text between ; and a newline are ignored.
Syntax details: input/ast/chords.y.

Examples:
# C triad with 1 beat
C[1]
# Bb minor triad with 2 beat
Bbm[2]
# D A7/E E Rest (D major)
# '_' is required when chord symbol is a number
echo 'D[1] A_7/E[1] E[2] R[1]' | crd text conv syllable --key D | crd write midi -o out.midi
```

## Build

``` shell
./task build
```
