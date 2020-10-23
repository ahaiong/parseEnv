package parseEnv

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

func ReplaceForViper(x string, y string) {
	for _, element := range os.Environ() {
		variable := strings.Split(element, "=")
		if strings.Contains(variable[0], x) && !strings.HasPrefix(variable[0], "_") {
			viper.Set(strings.Replace(variable[0], x, y, -1), variable[1])
		}
	}
}
