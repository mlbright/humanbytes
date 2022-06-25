# humanize

> Format byte sizes into more humanly readable values

This program was inspired by a [post][wandering-thoughts] by Chris Siebenmann.
I wondered if there was an equivalent (and easy) way to perform the functions of his Python script in Golang.

There is!
This is one possible solution, using the regular expression function substitution he talks about.

## Build

    go build

## Run

    ./humanbytes < test.txt

[wandering-thoughts]: https://utcc.utoronto.ca/~cks/space/blog/python/RegexpFunctionSubstitutionWin?showcomments#comments
