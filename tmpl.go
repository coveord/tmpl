package main

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/k0kubun/pp"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	debug := false
	varFileName := ""
	templateFileName := ""

	var RootCmd = &cobra.Command{
		Use:   "tmpl -t <templateFile> -v <variableFile>",
		Short: "template parser",
		Long: `A really fast template parser
 from a guy who had much free time.
 
 You can override variable or define new ones using environment variables
  prefixed with TMPL_ ex: TMPL_XVAR`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(varFileName) <= 0 {
				log.Fatalln("-v <variableFile> flag required")
			}
			if len(templateFileName) <= 0 {
				log.Fatalln("-t <templateFile> flag required")
			}

			yaml := viper.New()
			yaml.SetConfigType("yaml")
			yaml.SetEnvPrefix("TMPL")
			yaml.AutomaticEnv()

			varFile, err := os.Open(varFileName)
			if err != nil {
				log.Fatalln(err)
			}
			yaml.ReadConfig(varFile)

			if debug {
				pp.Fprintln(os.Stderr, yaml.AllSettings())
			}

			tmpl, err := template.New(templateFileName).ParseFiles(templateFileName)
			if err != nil {
				log.Fatalln(err)
			}
			tmpl.Execute(os.Stdout, yaml.AllSettings())
		},
	}

	RootCmd.Flags().StringVarP(&varFileName, "variableFile", "v", varFileName, "The source of all variables (yaml formatted)")
	RootCmd.Flags().StringVarP(&templateFileName, "templateFile", "t", templateFileName, "The source template file")
	RootCmd.PersistentFlags().BoolVar(&debug, "debug", debug, "Wheter to have debug information")

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

}
