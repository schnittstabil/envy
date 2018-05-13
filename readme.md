# envy

> [envsubst](https://linux.die.net/man/1/envsubst) variant, using [golang templates](https://golang.org/pkg/text/template/)


## Usage

```bash
$ envy
Usage: envy <template>... <dest>
```


## Example


### [`templates/hello.gotmpl`](templates/hello.gotmpl)

```
Hello {{.name}},

{{template "credentials.gotmpl" .}}
```


### [`templates/credentials.gotmpl`](templates/credentials.gotmpl)

```
your credentials:
Login: {{.login}}
Password: {{.password}}
```


### Execution

```bash
export name="John Doe"
export login="john"
export password="doe"

envy templates/hello.gotmpl templates/credentials.gotmpl output.txt
# or
envy --name hello.gotmpl templates/* output.txt

cat output.txt
# Hello John Doe,
#
# your credentials:
# Login: john
# Password: doe
```


## License

MIT © [Michael Mayer](http://schnittstabil.de)
