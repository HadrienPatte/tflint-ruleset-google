# WORKING IN PROGRESS AND MAINTAINER WANTED!

This project was started to support terraform-provider-google resources in TFLint. Everything is working in progress and not available as a plugin. If you are interested in this project, consider the following:

- Open an issue for the rule you want
  - Provider-specific knowledge is very important. Sharing the problems you faced will make it clear what we need to develop.
- Write a rule and open a pull request
  - If you are already familiar with tflint rulesets, Writing rules to solve open issues helps all users of TFLint. If you are interested but are not familiar, please let us know. Support you as much as possible.
- Make automatic rule generation system
  - In reality, it is impossible to manage rules for so many resources by hand. tflint has a mechanism to automatically generate AWS validation rules from aws-sdk (See [here](https://github.com/terraform-linters/tflint/blob/master/docs/DEVELOPING.md#sdk-based)). Building such a system can bring a lot of value at low cost.

See [terraform-lintes/tflint#606](https://github.com/terraform-linters/tflint/issues/606) for more details.

# TFLint Ruleset for terraform-provider-google

TFLint ruleset plugin for Terraform Google Cloud Platform provider

## Requirements

- TFLint v0.14+
- Go v1.13

## Installation

Download the plugin and place it in `~/.tflint.d/plugins/tflint-ruleset-google` (or `./.tflint.d/plugins/tflint-ruleset-google`). When using the plugin, configure as follows in `.tflint.hcl`:

```hcl
plugin "google" {
    enabled = true
}
```

## Rules

|Name|Description|Severity|Enabled|Link|
| --- | --- | --- | --- | --- |
|google_compute_instance_example_type|Show machine type|ERROR|✔||

## Building the plugin

Clone the repository locally and run the following command:

```
$ make
```

You can easily install the built plugin with the following:

```
$ make install
```
