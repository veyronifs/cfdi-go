package contabilidad13

import "fmt"

type Archivo struct {
	RFC  string
	Anio int
	Mes  int
	Tipo string
}

// FileName crea el nombre del archivo de acuerdo a la especificación del SAT.
func (a Archivo) FileName() string {
	return fmt.Sprintf("%s%d%02d%s", a.RFC, a.Anio, a.Mes, a.Tipo)
}

// FileNameClear crea el nombre del archivo con un formato legible para el usuario.
func (a Archivo) FileNameClear() string {
	return fmt.Sprintf("%s_%d_%02d_%s", a.RFC, a.Anio, a.Mes, a.Tipo)
}
