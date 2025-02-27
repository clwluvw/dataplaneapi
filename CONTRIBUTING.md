# ![HAProxy](assets/images/haproxy-weblogo-210x49.png "HAProxy")

## Contributing guide

This document is a short guide for contributing to this project.

## API Specification - Development guide

Data Plane API is generated using [go-swagger](https://github.com/go-swagger/go-swagger) from the swagger spec found [here](https://github.com/haproxytech/client-native/blob/master/specification/build/haproxy_spec.yaml) using the following command.

```
make generate
```
or if you prefer to run it directly (not in docker)
```
make generate-native
```

`make generate-native` must also be used if you are using local client-native on disk (you have `replace github.com/haproxytech/client-native/v3 => ../client-native` in go.mod file)

This command generates some of the files in this project, which are marked with // Code generated by go-swagger; DO NOT EDIT.
comments at the top of the files. These are not to be edited, as they are overwritten when specification is changed and the above-mentioned command is run. If you want to change those files, please change the specification where necessary and then generate them again.

## Commit Messages and General Style

For commit messages and general style please follow the haproxy project's [CONTRIBUTING guide](https://github.com/haproxy/haproxy/blob/master/CONTRIBUTING) and use that where applicable.
