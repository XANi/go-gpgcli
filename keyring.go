package gpgcli
import (
	"strings"
	"fmt"
)

func (g *Gpg) GetFingerprintById( short string) (s string, err error) {
	stdout,stderr, err := g.cmd("--list-keys","--fingerprint","--with-colons",short)
	if err != nil {
		return s, fmt.Errorf("Cant find key %s: %s\nerror %s", short, stderr, err)
	}
	for _, line := range strings.Split(stdout,"\n") {
		if strings.HasPrefix(line, `fpr:`) {
			tmp := strings.Split(line, `:`)
			if len(tmp[9]) > 0 && strings.Contains(tmp[9], s) {
				return tmp[9], err
			}
		}
	}
	return s, fmt.Errorf("coudn't find fingerprint")
}
