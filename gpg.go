package gpgcli

import (
	"os"
	"strings"
	"fmt"
)

var version string

type Gpg struct {
	filteredEnv []string
	gpgOptions  []string
}
type GpgConfig struct{
	HomeDir string
}


func New(arg ...*GpgConfig) (g *Gpg, err error) {
	var gpg Gpg
	if len(arg) > 0 {
		opts := arg[0]
		if len(opts.HomeDir) > 0 {
			gpg.gpgOptions = append (gpg.gpgOptions, "--homedir", opts.HomeDir)
		}
	} else if len(arg) > 1 {
 		return g, fmt.Errorf("only one config struct is allowed")
	}


	for _, v := range os.Environ() {
		// filter locale sillines so command output is always english
		if !strings.HasPrefix(v, "LANG=") && !strings.HasPrefix(v, "LC")  {
			gpg.filteredEnv = append(gpg.filteredEnv,v)
		}
	}
	return &gpg, err
}
