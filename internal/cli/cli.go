package cli

import (
	"errors"
	"flag"
	"os"
	"strings"

	"github.com/mraraneda/mrlogger"
)

//cli.FlagHandler(&sellerdni, &folio, &order)

// FlagHandler captura los flags declarados y los maneja
func FlagHandler(configfile *string) {
	flag.StringVar(configfile, "config", "", "archivo de configuración de la aplicacoión")

	flag.Parse()

	args := strings.Builder{}

	for _, v := range os.Args {
		args.WriteString(v)
		args.WriteString(" ")
	}

	mrlogger.Debug("CLI call:", args.String())

	// En este map se registran los flags obligatorios
	required := []string{"config"}

	// Este bucle evalúa si se ingresaron los parámetros obligatorios
	seen := make(map[string]bool)
	flag.VisitAll(func(f *flag.Flag) {
		if f.Value.String() != "" {
			seen[f.Name] = true
		}
	})
	for _, req := range required {
		if !seen[req] {
			err := errors.New("Falta ingresar el parámetro obligatorio: \"" + req + "\"")
			mrlogger.Error("Falta ingresar el parámetro obligatorio: \"" + req + "\"")
			mrlogger.Check(err, mrlogger.InThisPoint())
		}
	}

}
