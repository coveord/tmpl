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
  tmpl -t <templateFile> -v <variableFile> [flags]

Flags:
      --debug                 Wheter to have debug information
  -t, --templateFile string   The source template file
  -v, --variableFile string   The source of all variables (yaml formatted)
  ```