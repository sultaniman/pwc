# PWC

## Build

```sh
$ go build -o pwc
```

## Use

```sh
$ ./pwc classic -h
Generate classic password card

Usage:
  cmd classic [flags]

Flags:
  -e, --encrypted string   When given will encrypt generated card and write to file (default "card.aes")
  -h, --help               help for classic
  -d, --include-digits     Rows 5-8 will be digits only
  -s, --include-symbols    With regular a-zA-Z include @#$%&*<>?€+{}[]()/\
  -o, --output string      Output file (default "card.jpg")
  -p, --print-passphrase   Prints passphrase in the console
```

## Misc
```
Alphabets

First row: ■□▲△○●★☂☀☁☹☺♠♣♥♦♫€¥£$!?¡¿⊙◐◩�
Alphanumeric: 0123456789abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ
Alphanumeric and symbols: 0123456789abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ@#$%&*<>?€+{}[]()/\
Numeric: 0123456789

Row colors:
#ffffff White
#c0c0c0 Gray
#ffc0c0 Red
#c0ffc0 Green
#ffffc0 Yellow
#c0c0ff Blue
#ffc0ff Magenta
#c0ffff Cyan
```
