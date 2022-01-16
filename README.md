# PWC
![Tests](https://github.com/imanhodjaev/pwc/actions/workflows/run-tests.yml/badge.svg)

If you need to remember dozens of passwords without having to remember all of them
then with password cards you can have it.
Each password card is a set of grids with random letters and digits on it has 8 rows
which have different colors, columns mapped to different symbols.
Everything needed is to remember a combination of a symbol and a color or an index of row,
then compose the letters and digits of your passwords from there.

## Build

```sh
$ go build -o pwc
```

## Installation

```sh
$ go get -u github.com/imanhodjaev/pwc
```

## CLI Use

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

## Example card

<p>
<img src="https://raw.githubusercontent.com/imanhodjaev/pwc/main/example/password-card.jpg" width="400"/>
</p>

## How it works

```sh
$ ./pwc explain
Alphabet

               Header symbols: ■□▲△○●★☂☀☁☹☺♠♣♥♦♫€¥£$!?¡¿⊙◐◩�
                      Numbers: 0123456789
                 Alphanumeric: 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ
     Alphanumeric and symbols: 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ@#$%!&(MISSING)*<>?€+{}[]()/\

Algorithm

1. Render the first row with shuffled "■□▲△○●★☂☀☁☹☺♠♣♥♦♫€¥£$!?¡¿⊙◐◩�",
2. Iterate over the rest of the rows and for each randomly selected letter we shuffle the entire alphabet
    a. If the card should include symbols then
       use the alphanumeric and symbols alphabet for every even character
       and use the alphanumeric alphabet for the odd columns,
    b. If the card should have a digits area then
       use the numeric alphabet for the lower half of the rows 5-8.

Row colors

White #ffffff, Gray #c0c0c0, Red #ffc0c0, Green #c0ffc0, Yellow #ffffc0, Blue #c0c0ff, Magenta #ffc0ff, Cyan #c0ffff
```

P.S. it was inspired by https://www.passwordcard.org/en.


<p align="center">✨ 🚀 ✨</p>
