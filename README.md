# tokengen

Generate random strings for tokens and passwords

Generated tokens are random bytes of characters, consisting only of upper and
lowercase alphanumeric characters. These tokens can be used as auth tokens,
passwords, or other secrets that do not impose any additional constraints like
the inclusion of special characters.

## Installation

The `tokengen` tool can be installed using `go`:
```sh
go install github.com/rgcottrell/tokengen@latest
```

The binary can be found in the `go` binary directory, which is typically
`$HOME/go/bin`. Make sure that this directory is included in your path.

## Usage

Without any additional parameters, `tokengen` will return a single 32-character
token:

```
$ tokengen
WSFoOoOFYTr3jqWrs44fjkZO0dRR1u9X
```

The returned values can be customized be specifying the total number of tokens
desired and/or the length of each token in characters:

```sh
$ tokengen -n 5 -l 40
yXI2XlsKlAN6lr2Kct3Vhnef5x8ykFYVpC3dZWsd
QzFPSy2MFMEqiAbqNhsdzhWIWyYHWP0IaVXIIo23
ij02hC6m8zaVfSAeXSzNMkm6z3ZjKsX8VyBCh2fl
pseUhxVcXRodgwZzKlWuhsuEdVLIYXF8Uy6Q1olP
FMKEXzylnZT9hvaz5Fyhjp0fXjnSwxERxfxABAuj
``````
