# PWC

![Tests](https://github.com/imanhodjaev/pwc/actions/workflows/run-tests.yml/badge.svg)
[![CodeQL](https://github.com/sultaniman/pwc/actions/workflows/codeql-analysis.yml/badge.svg)](https://github.com/sultaniman/pwc/actions/workflows/codeql-analysis.yml)

If you need to remember dozens of passwords without having to remember all of them
then with password cards you can have it.
Each password card is a set of grids with random letters and digits on it has 8 rows
which have different colors, columns mapped to different symbols.
Everything needed is to remember a combination of a symbol and a color or an index of row,
then compose the letters and digits of your passwords from there.

## 🔨 Build

```sh
$ go build -o pwc
```

## 💾 Installation

```sh
$ go get -u github.com/sultaniman/pwc
```

## 🖥️ CLI Use

```sh
$ ./pwc classic -h
```

## Example card

<p>
<img src="https://raw.githubusercontent.com/sultaniman/pwc/main/example/password-card.jpg" width="400"/>
</p>

## ❓ How it works

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

### Row colors

1. White `#ffffff`,
2. Gray `#c0c0c0`,
3. Red `#ffc0c0`,
4. Green `#c0ffc0`,
5. Yellow `#ffffc0`,
6. Blue `#c0c0ff`,
7. Magenta `#ffc0ff`,
8. Cyan `#c0ffff`.


| Color | HEX |
|------|-----|
| $\color{rgb(255,255,255)}{\textsf{White}}$ | `#ffffff` |
| $\color{rgb(192,192,192)}{\textsf{Gray}}$ | `#c0c0c0` |
| $\color{rgb(255,192,192)}{\textsf{Red}}$ | `#ffc0c0` |
| $\color{rgb(192,255,192)}{\textsf{Green}}$ | `#c0ffc0` |
| $\color{rgb(255,255,192)}{\textsf{Yellow}}$ | `#ffffc0` |
| $\color{rgb(192,192,255)}{\textsf{Blue}}$ | `#c0c0ff` |
| $\color{rgb(255,192,255)}{\textsf{Magenta}}$ | `#ffc0ff` |
| $\color{rgb(192,255,255)}{\textsf{Cyan}}$ | `#c0ffff` |

P.S. it was inspired by https://www.passwordcard.org/en.

<p align="center">✨ 🚀 ✨</p>
