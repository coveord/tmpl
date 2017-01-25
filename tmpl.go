package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"text/template"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/kubernetes/pkg/util/uuid"
)

var debug = false

func main() {
	RootCmd := &cobra.Command{
		Use:   "tmpl -t <templateFile> -v <variableFile>",
		Short: "template parser",
		Long: `A really fast template parser
 from a guy who had much free time.
 
 You can override variable or define new ones using environment variables
prefixed with TMPL_ ex: TMPL_XVAR`,
		Run: runCmd,
	}

	RootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Wheter to have debug information")
	RootCmd.Flags().StringP("vars", "v", "", "The main variables source (yaml)")
	RootCmd.Flags().StringP("template", "t", "", "The template (dir or file)")
	RootCmd.MarkFlagRequired("template")
	RootCmd.Flags().StringP("output", "o", "", "Output stdout by default for a single file, (required for dir template)")

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func runCmd(cmd *cobra.Command, args []string) {
	varFilePath, _ := cmd.Flags().GetString("vars")
	templateFilePath, _ := cmd.Flags().GetString("template")
	outputFilePath, _ := cmd.Flags().GetString("output")

	// Init Variable Source
	viper := viper.New()
	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("TMPL")
	viper.AutomaticEnv()

	if varFilePath != "" {
		f, err := os.Open(varFilePath)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		if viper.ReadConfig(f) != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
	}

	// verify parameter validity
	if templateFilePath == "" {
		fmt.Fprintln(os.Stderr, "--template <templateFile> flag required")
		os.Exit(-1)
	}

	templateFile, err := os.Open(templateFilePath)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
	s, err := templateFile.Stat()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	if s.IsDir() {
		if outputFilePath == "" {
			fmt.Fprintf(os.Stderr, "%s is a Directory, -o flag is required\n", templateFile.Name())
			os.Exit(-1)
		}

		os.MkdirAll(outputFilePath, 0711)
		names, _ := templateFile.Readdirnames(-1)
		sort.Strings(names)
		for _, n := range names {
			outFilePath := filepath.Join(outputFilePath, n)
			inFilePath := filepath.Join(templateFilePath, n)
			out, _ := os.Create(outFilePath)
			in, _ := os.Open(inFilePath)

			fmt.Println(outFilePath)
			processFile(viper, in, out)
			out.Close()
		}
		// process all files individually
	} else {
		// TODO: when template is not a dir but a file
	}
}

func processFile(v *viper.Viper, in io.Reader, out io.Writer) {
	b, _ := ioutil.ReadAll(in)
	tmpl, _ := template.New(string(uuid.NewUUID())).Parse(string(b))

	tmpl.Execute(out, v.AllSettings())
}
