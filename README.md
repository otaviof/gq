# `gq`

Is a swissarmy knife text processor to apply [Go Templates][gotmpl] against structure file formats,
like YAML and JSON, heavily inspired on [`jq`][jq] and [`yq`][yq].

For instance:

```sh
gq '{{ index . "current-context" }}' ~/.kube/config
cat ~/.kube/config |gq '{{ index . "current-context" }}'
gq --type=json '{{ range .HttpHeaders }}{{ printf "%s\n" . }}{{ end }}' ~/.docker/config.json
```

## Installing

You can use `gq` as a [container-image][containerimage], or get it locally via:

```sh
go get -u github.com/otaviof/gq
```

When having a local clone of this repository, you may:

```sh
make install
```

[containerimage]: https://quay.io/otaviof/gq
[gotmpl]: https://golang.org/pkg/text/template/
[jq]: https://stedolan.github.io/jq/
[yq]: https://github.com/mikefarah/yq