# envy

> [envsubst](https://linux.die.net/man/1/envsubst) variant, using [golang templates](https://golang.org/pkg/text/template/) and [Masterminds/sprig](https://github.com/Masterminds/sprig)


## Usage

```bash
$ envy
Usage: ./envy [opts] <template>...
  -input string
        <file> to render, defaults to first <template>
  -output string
        write output to <file> instead of stdout
```


## Example

### [`hosts.gotmpl`](hosts.gotmpl)

```
{{ range $host := splitList "," .hosts }}
{{- $host}}	{{ index $ (print "hosts_" $host) }}
{{end}}
```


### Execution

```bash
export hosts="front,back,db"
export hosts_front="192.168.42.1"
export hosts_back="192.168.42.2"
export hosts_db="192.168.42.3"


envy --output output.txt hosts.gotmpl
# or
envy --output output.txt --input hosts.gotmpl *.gotmpl

cat output.txt
# =>
front   192.168.42.1
back    192.168.42.2
db      192.168.42.3
```


## Related

* [Masterminds/sprig](https://github.com/Masterminds/sprig) – Useful template functions for Go templates.


## License

MIT © [Michael Mayer](http://schnittstabil.de)
