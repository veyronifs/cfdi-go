package catcuentas

import (
	"github.com/veyronifs/cfdi-go/types"
)

// Catalogo Estándar de catálogo de cuentas que se entrega como parte de la contabilidad electrónica.
//
// El catálogo de cuentas es el documento en el que se detalla de forma ordenada todos los números
// (claves) y nombres de las cuentas de activo, pasivo, capital, ingresos, costos, gastos y cuentas de orden,
// aplicables en la contabilidad de un ente económico, y contendrá los siguientes datos:
type Catalogo struct {
	// Ctas Nodo obligatorio para expresar el detalle de cada cuenta y subcuenta del catálogo.
	Ctas []*Cta `xml:"Ctas"`
	// Version Atributo requerido para expresar la versión del formato.
	//
	// fixed="1.3"
	Version string `xml:"Version,attr"`
	// RFC Atributo requerido para expresar el RFC del contribuyente que envía los datos.
	RFC string `xml:"RFC,attr"`
	// Mes Atributo requerido para expresar el mes en que inicia la vigencia del catálogo para la
	// balanza.
	Mes int `xml:"Mes,attr"`
	// Anio Atributo requerido para expresar el año en que inicia la vigencia del catálogo para la
	// balanza.
	Anio int `xml:"Anio,attr"`
	// Sello Atributo opcional para contener el sello digital del archivo de contabilidad
	// electrónica. El sello deberá ser expresado cómo una cadena de texto en formato Base 64.
	Sello string `xml:"Sello,attr,omitempty"`
	// NoCertificado Atributo opcional para expresar el número de serie del certificado de sello
	// digital que ampara el archivo de contabilidad electrónica, de acuerdo al acuse
	// correspondiente a 20 posiciones otorgado por el sistema del SAT.
	NoCertificado string `xml:"noCertificado,attr,omitempty"`
	// Certificado Atributo opcional que sirve para expresar el certificado de sello digital que
	// ampara al archivo de contabilidad electrónica como texto, en formato base 64.
	Certificado string `xml:"Certificado,attr,omitempty"`
}

// Cta Nodo obligatorio para expresar el detalle de cada cuenta y subcuenta del catálogo.
type Cta struct {
	// CodAgrup Atributo requerido para expresar el código asociador de cuentas y subcuentas
	// conforme al catálogo publicado en la página de internet del SAT. Se debe asociar cada
	// cuenta y subcuenta que sea más apropiado de acuerdo con la naturaleza y preponderancia de la cuenta o subcuenta.
	CodAgrup types.CodAgrup `xml:"CodAgrup,attr"`
	// NumCta Atributo requerido, es la clave con que se distingue la cuenta o subcuenta en la
	// contabilidad.
	NumCta string `xml:"NumCta,attr"`
	// Desc Atributo requerido para expresar el nombre de la cuenta o subcuenta.
	Desc string `xml:"Desc,attr"`
	// SubCtaDe Atributo opcional en el caso de subcuentas. Sirve para expresar la clave de la
	// cuenta a la que pertenece dicha subcuenta. Se convierte en requerido cuando se cuente con
	// la información.
	SubCtaDe string `xml:"SubCtaDe,attr,omitempty"`
	// Nivel Atributo requerido para expresar el nivel en el que se encuentra la cuenta o
	// subcuenta en el catálogo.
	Nivel int `xml:"Nivel,attr"`
	// Natur  Atributo requerido para expresar la naturaleza de la cuenta o subcuenta.
	//
	//	* D: Deudora
	//	* A: Acreedora
	//
	//	* Activo: D
	//	* Pasivo: A
	//	* Capital: A
	//	* Ingreso: A
	//	* Costo: D
	//	* Gasto: D
	//	* Resultado Integral de Financiamient: D,A
	//	* Cuentas de orden: D,A
	//
	// Existen cuentas de Activo, Pasivo y Capital que por  su naturaleza pueden presentarse de manera Deudora o Acreedora
	Natur Natur `xml:"Natur,attr"`
}

// Natur naturaleza de la cuenta o subcuenta.
type Natur string

const (
	NaturD Natur = "D"
	NaturA Natur = "A"
)
