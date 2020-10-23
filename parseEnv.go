package parseEnv

import (
	"github.com/spf13/viper"
	"os"
	"strings"
)

func ParseEnv() {
	for _, element := range os.Environ() {
		variable := strings.Split(element, "=")
		if strings.Contains(variable[0], "__") && !strings.HasPrefix(variable[0], "_") {
			viper.Set(strings.Replace(variable[0], "__", ".", -1), variable[1])
		}
	}
}
