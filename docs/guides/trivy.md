# Moving towards configuration scanning with Vul
Overtime we've taken [vul][vul] to be the go-to scanning tool for a vareity of things. This also includes terraform scanning.

This section describes some differences between Vul and tfsecurity.

| Feature              | Vul                                                  | tfsecurity                |
|----------------------|--------------------------------------------------------|----------------------|
| Policy Distribution | Embedded and Updated via Registry                      | Embedded             |
| Custom Policies      | Rego                                                   | Rego, JSON, and YAML |
| Supported Formats    | Dockerfile, JSON, YAML, Terraform, CloudFormation etc. | Terraform  Only      |


# Comparison with examples
## Simple scan
### With Vul
```shell
$ vul config <dir>
```
### With tfsecurity
```shell
$ tfsecurity <dir>
```

## Passing tfvars
### With Vul
```shell
$ vul --tf-vars <vars.tf> <dir>
```
### With tfsecurity
```shell
$ tfsecurity <dir> --tf-vars-file <vars.tf>
```

## Report formats
### With Vul
```shell
$ vul config --format <format-type> <dir>
```

### With tfsecurity
```shell
$ tfsecurity <dir> --format <format-type>
```

We welcome any feedback if you find features that today are not available with Vul misconfigration scanning that are available in tfsecurity. 

[vul]: https://github.com/khulnasoft-lab/vul