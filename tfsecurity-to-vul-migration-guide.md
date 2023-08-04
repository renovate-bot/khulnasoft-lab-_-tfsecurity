# Migrating from tfsecurity to Vul
Overtime we've taken [Vul][vul] to be the go-to scanning tool for a vareity of things. This also includes terraform scanning. For further information, have a look at the announcement ["tfsecurity is joining the Vul family".](https://github.com/khulnasoft-labs/tfsecurity/discussions/1994)

### Main differences between Vul and tfsecurity

Vul's design keeps misconfiguration up to date automatically. New misconfiguration are updated in Vul by pulling from the Container Registry. The embedded misconfiguration in Vul are only used if Vul cannot pull from the remote registry. See the [following documentation](https://khulnasoft-labs.github.io/vul/v0.41/docs/scanner/misconfiguration/policy/builtin/#policy-distribution) for further details.

## Comparison with examples
### Simple scan
#### With Vul
```shell
$ vul config <dir>
```
#### With tfsecurity
```shell
$ tfsecurity <dir>
```

The documentation can be found in Vul under the [following link.](https://khulnasoft-labs.github.io/vul/latest/docs/scanner/misconfiguration/)

### Passing tfvars
#### With Vul
```shell
$ vul --tf-vars <vars.tf> <dir>
```
#### With tfsecurity
```shell
$ tfsecurity <dir> --tf-vars-file <vars.tf>
```

The documentation can be found in Vul under the [following link.](https://khulnasoft-labs.github.io/vul/v0.41/docs/scanner/misconfiguration/#terraform-value-overrides)

### Report formats
#### With Vul
```shell
$ vul config --format <format-type> <dir>
```

#### With tfsecurity
```shell
$ tfsecurity <dir> --format <format-type>
```

The documentation can be found in Vul under the [following link.](https://khulnasoft-labs.github.io/vul/v0.41/docs/configuration/reporting/)

## FAQs

**Does Vul support junit?**

Yes, Vul supports different report templates. These can either be set, loaded through a file or by providing a default template such as for JUnit. 

For more information, please [the documentation.](https://khulnasoft-labs.github.io/vul/v0.41/docs/configuration/reporting/#junit)

**Does Vul support multiple outputs?**

Currently, the following outputs are supported by Vul:

* Table
* JSON
* SARIF
* Template
* SBOM

e.g.
```
vul config --output report.json --format json ./bad_iac/docker
```
This will saver the json report into a `report.json` file.

[Documentation](https://khulnasoft-labs.github.io/vul/v0.41/docs/configuration/reporting/)

Note that one report can be generated per scan. However, if you require multiple different reports, the same scan would pull the information from the cache to generate a new report format.

**Can Vul skip files?**

Yes, you can specify that Vul should skip a directory, using the following flag `--skip-dirs`.

[Documentation](https://khulnasoft-labs.github.io/vul/v0.41/docs/configuration/others/)

Alternatively, it is possible to skip files, using this flag `--skip-files`.

[Documentation](https://khulnasoft-labs.github.io/vul/v0.41/docs/configuration/others/#skip-files)

## Feedback

We welcome any feedback if you find features that today are not available with Vul misconfigration scanning that are available in tfsecurity. 

For further information on scanning terraform with Vul, do have a look at the [Vul Terraform Guide](https://khulnasoft-labs.github.io/vul/latest/tutorials/terraform/scannig/).

[vul]: https://github.com/khulnasoft-labs/vul
