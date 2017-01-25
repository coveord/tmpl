# tmpl

a mix of [cobra](github.com/spf13/cobra), [viper](github.com/spf13/viper) and
`text/template` and you get something really powerfull from a simple way.

## Usage

```
A really fast template parser
 from a guy who had much free time.

 You can override variable or define new ones using environment variables
prefixed with TMPL_ ex: TMPL_XVAR

Usage:
  tmpl -t=<templateDirectory> -v=<variableFile> -o=<outputDirectory>

Flags:
      --debug              Wheter to have debug information
  -o, --output string      Output stdout by default for a single file, (required for dir template)
  -t, --template string   The template (dir or file)
  -v, --vars string        The main variables source (yaml)
  ```