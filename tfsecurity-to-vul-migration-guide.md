# Migrating from tfsecurity to Tunnel
Overtime we've taken [Tunnel][tunnel] to be the go-to scanning tool for a vareity of things. This also includes terraform scanning. For further information, have a look at the announcement ["tfsecurity is joining the Tunnel family".](https://github.com/khulnasoft-lab/tfsecurity/discussions/56)

### Main differences between Tunnel and tfsecurity

Tunnel's design keeps misconfiguration up to date automatically. New misconfiguration are updated in Tunnel by pulling from the Container Registry. The embedded misconfiguration in Tunnel are only used if Tunnel cannot pull from the remote registry. See the [following documentation](https://khulnasoft-lab.github.io/vul/v0.41/docs/scanner/misconfiguration/policy/builtin/#policy-distribution) for further details.

## Comparison with examples
### Simple scan
#### With Tunnel
```shell
$ vul config <dir>
```
#### With tfsecurity
```shell
$ tfsecurity <dir>
```

The documentation can be found in Tunnel under the [following link.](https://khulnasoft-lab.github.io/vul/latest/docs/scanner/misconfiguration/)

### Passing tfvars
#### With Tunnel
```shell
$ vul --tf-vars <vars.tf> <dir>
```
#### With tfsecurity
```shell
$ tfsecurity <dir> --tf-vars-file <vars.tf>
```

The documentation can be found in Tunnel under the [following link.](https://khulnasoft-lab.github.io/vul/v0.41/docs/scanner/misconfiguration/#terraform-value-overrides)

### Report formats
#### With Tunnel
```shell
$ vul config --format <format-type> <dir>
```

#### With tfsecurity
```shell
$ tfsecurity <dir> --format <format-type>
```

The documentation can be found in Tunnel under the [following link.](https://khulnasoft-lab.github.io/vul/v0.41/docs/configuration/reporting/)

## FAQs

**Does Tunnel support junit?**

Yes, Tunnel supports different report templates. These can either be set, loaded through a file or by providing a default template such as for JUnit. 

For more information, please [the documentation.](https://khulnasoft-lab.github.io/vul/v0.41/docs/configuration/reporting/#junit)

**Does Tunnel support multiple outputs?**

Currently, the following outputs are supported by Tunnel:

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

[Documentation](https://khulnasoft-lab.github.io/vul/v0.41/docs/configuration/reporting/)

Note that one report can be generated per scan. However, if you require multiple different reports, the same scan would pull the information from the cache to generate a new report format.

**Can Tunnel skip files?**

Yes, you can specify that Tunnel should skip a directory, using the following flag `--skip-dirs`.

[Documentation](https://khulnasoft-lab.github.io/vul/v0.41/docs/configuration/others/)

Alternatively, it is possible to skip files, using this flag `--skip-files`.

[Documentation](https://khulnasoft-lab.github.io/vul/v0.41/docs/configuration/others/#skip-files)

## Feedback

We welcome any feedback if you find features that today are not available with Tunnel misconfigration scanning that are available in tfsecurity. 

For further information on scanning terraform with Tunnel, do have a look at the [Tunnel Terraform Guide](https://khulnasoft-lab.github.io/vul/latest/tutorials/terraform/scannig/).

[tunnel]: https://github.com/khulnasoft/tunnel
