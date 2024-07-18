package cartaporte30

import (
	"encoding/xml"

	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/types"
)

func Unmarshal(b []byte) (*CartaPorte30, error) {
	carta := &CartaPorte30{}
	if err := xml.Unmarshal(b, carta); err != nil {
		return nil, err
	}
	return carta, nil
}

// CartaPorte30 Complemento para incorporar al Comprobante Fiscal Digital por Internet (CFDI), la información relacionada a los bienes y/o mercancías, ubicaciones de origen, puntos intermedios y destinos, así como lo referente al medio por el que se transportan; que circulen por vía terrestre, férrea, aérea o naveguen por vía marítima; además de incluir el traslado de hidrocarburos y petrolíferos.
type CartaPorte30 struct {
	// Ubicaciones Nodo requerido para registrar las distintas ubicaciones que sirven para indicar el domicilio del origen y/o destino que tienen los bienes y/o mercancías que se trasladan a través de los distintos medios de transporte.
	Ubicaciones Ubicaciones `xml:"Ubicaciones"`
	// Mercancias Nodo requerido para registrar la información de los bienes y/o mercancías que se trasladan en los distintos medios de transporte.
	Mercancias *Mercancias `xml:"Mercancias"`
	// FiguraTransporte Nodo condicional para indicar los datos de la(s) figura(s) del transporte que interviene(n) en el traslado de los bienes y/o mercancías realizado a través de los distintos medios de transporte dentro del territorio nacional, cuando el dueño de dicho medio sea diferente del emisor del comprobante con el complemento Carta Porte.
	FiguraTransporte *FiguraTransporte `xml:"FiguraTransporte,omitempty"`
	// Version Atributo requerido con valor prefijado en el cual se indica la versión del complemento Carta Porte.
	Version string `xml:"Version,attr"`
	// IdCCP Atributo requerido para expresar los 36 caracteres del folio del complemento Carta Porte (IdCCP) de la transacción de timbrado conforme al estándar RFC 4122, para la identificación del CFDI con complemento Carta Porte.
	IdCCP string `xml:"IdCCP,attr"`
	// TranspInternac Atributo requerido para expresar si los bienes y/o mercancías que son transportadas ingresan o salen del territorio nacional.
	TranspInternac string `xml:"TranspInternac,attr"`
	// RegimenAduanero Atributo condicional para expresar el tipo de régimen que se encuentra asociado con el traslado de los bienes y/o mercancías de procedencia extranjera.
	RegimenAduanero RegimenAduanero `xml:"RegimenAduanero,attr,omitempty"`
	// EntradaSalidaMerc Atributo condicional para precisar si los bienes y/o mercancías ingresan o salen del territorio nacional.
	EntradaSalidaMerc string `xml:"EntradaSalidaMerc,attr,omitempty"`
	// PaisOrigenDestino Atributo condicional para registrar la clave del país de origen o destino de los bienes y/o mercancías que se trasladan a través de los distintos medios de transporte.
	PaisOrigenDestino types.Pais `xml:"PaisOrigenDestino,attr,omitempty"`
	// ViaEntradaSalida Atributo condicional para registrar la vía de ingreso o salida de los bienes y/o mercancías en territorio nacional.
	ViaEntradaSalida string `xml:"ViaEntradaSalida,attr,omitempty"`
	// TotalDistRec Atributo condicional para indicar en kilómetros, la suma de las distancias recorridas, registradas en el atributo “DistanciaRecorrida”, para el traslado de los bienes y/o mercancías.
	TotalDistRec decimal.Decimal `xml:"TotalDistRec,attr,omitempty"`
	// RegistroISTMO Atributo opcional para registrar las regiones, sí el traslado de los bienes y/o mercancías se realiza al interior de los Polos de Desarrollo para el Bienestar del istmo de Tehuantepec.
	RegistroISTMO string `xml:"RegistroISTMO,attr,omitempty"`
	// UbicacionPoloOrigen Atributo condicional para registrar la región en donde inicia el traslado de los bienes y/o mercancias al interior de los Polos de Desarrollo para el Bienestar del istmo de Tehuantepec.
	UbicacionPoloOrigen string `xml:"UbicacionPoloOrigen,attr,omitempty"`
	// UbicacionPoloDestino Atributo condicional para registrar la región en donde termina el traslado de los bienes y/o mercancias al interior de los Polos de Desarrollo para el Bienestar del istmo de Tehuantepec.
	UbicacionPoloDestino string `xml:"UbicacionPoloDestino,attr,omitempty"`
}

// Ubicaciones Nodo requerido para registrar las distintas ubicaciones que sirven para indicar el domicilio del origen y/o destino que tienen los bienes y/o mercancías que se trasladan a través de los distintos medios de transporte.
type Ubicaciones []*Ubicacion

func (u *Ubicaciones) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ubics struct {
		Slice []*Ubicacion `xml:"Ubicacion"`
	}

	if err := d.DecodeElement(&ubics, &start); err != nil {
		return err
	}
	*u = ubics.Slice
	return nil
}

// Ubicacion Nodo requerido para registrar la ubicación que sirve para indicar el domicilio del origen y/o destino parcial o final, que tienen los bienes y/o mercancías que se trasladan a través de los distintos medios de transporte.
type Ubicacion struct {
	// Domicilio Nodo condicional para registrar información del domicilio de origen y/o destino de los bienes y/o mercancías que se trasladan a través de los distintos medios de transporte.
	Domicilio *Domicilio `xml:"Domicilio,omitempty"`
	// TipoUbicacion Atributo requerido para precisar si el tipo de ubicación corresponde al origen o destino de las ubicaciones para el traslado de los bienes y/o mercancías en los distintos medios de transporte.
	TipoUbicacion string `xml:"TipoUbicacion,attr"`
	// IDUbicacion Atributo condicional para registrar una clave que sirva para identificar el punto de salida o entrada de los bienes y/o mercancías que se trasladan a través de los distintos medios de transporte, la cual estará integrada de la siguiente forma: para origen el acrónimo “OR” o para destino el acrónimo “DE” seguido de 6 dígitos numéricos asignados por el contribuyente que emite el comprobante para su identificación.
	IDUbicacion string `xml:"IDUbicacion,attr,omitempty"`
	// RFCRemitenteDestinatario Atributo requerido para registrar el RFC del remitente o destinatario de los bienes y/o mercancías que se trasladan a través de los distintos medios de transporte.
	RFCRemitenteDestinatario string `xml:"RFCRemitenteDestinatario,attr"`
	// NombreRemitenteDestinatario Atributo opcional para registrar el nombre del remitente o destinatario de los bienes y/o mercancías que se trasladan a través de los distintos medios de transporte.
	NombreRemitenteDestinatario string `xml:"NombreRemitenteDestinatario,attr,omitempty"`
	// NumRegIdTrib Atributo condicional para registrar el número de identificación o registro fiscal del país de residencia, para los efectos fiscales del remitente o destinatario de los bienes y/o mercancías que se trasladan cuando se trate de residentes en el extranjero.
	NumRegIdTrib string `xml:"NumRegIdTrib,attr,omitempty"`
	// ResidenciaFiscal Atributo condicional para registrar la clave del país de residencia para efectos fiscales del remitente o destinatario de los bienes y/o mercancías, conforme el catálogo de CFDI c_Pais publicado en el portal del SAT en Internet de acuerdo a la especificación ISO 3166-1.
	ResidenciaFiscal types.Pais `xml:"ResidenciaFiscal,attr,omitempty"`
	// NumEstacion Atributo condicional para registrar la clave de la estación de origen o destino para el traslado de los bienes y/o mercancías que se realiza a través de los distintos medios de transporte, esto de acuerdo al valor de la columna “Clave identificación” del catálogo c_Estaciones del complemento Carta Porte que permita asociarla al tipo de transporte.
	NumEstacion string `xml:"NumEstacion,attr,omitempty"`
	// NombreEstacion Atributo condicional para registrar el nombre de la estación de origen o destino por la que se pasa para efectuar el traslado de los bienes y/o mercancías a través de los distintos medios de transporte, conforme al catálogo c_Estaciones del complemento Carta Porte.
	NombreEstacion string `xml:"NombreEstacion,attr,omitempty"`
	// NavegacionTrafico Atributo condicional para registrar el tipo de puerto de origen o destino en el cual se documentan los bienes y/o mercancías que se trasladan vía marítima.
	NavegacionTrafico string `xml:"NavegacionTrafico,attr,omitempty"`
	// FechaHoraSalidaLlegada Atributo requerido para registrar la fecha y hora estimada en la que salen o llegan los bienes y/o mercancías de origen o al destino, respectivamente. Se expresa en la forma AAAA-MM-DDThh:mm:ss.
	FechaHoraSalidaLlegada types.FechaH `xml:"FechaHoraSalidaLlegada,attr"`
	// TipoEstacion Atributo condicional para registrar el tipo de estación por el que pasan los bienes y/o mercancías durante su traslado a través de los distintos medios de transporte.
	TipoEstacion string `xml:"TipoEstacion,attr,omitempty"`
	// DistanciaRecorrida Atributo condicional para registrar en kilómetros la distancia recorrida entre la ubicación de origen y la de destino parcial o final, por los distintos medios de transporte que trasladan los bienes y/o mercancías.
	DistanciaRecorrida decimal.Decimal `xml:"DistanciaRecorrida,attr,omitempty"`
}

// Mercancias Nodo requerido para registrar la información de los bienes y/o mercancías que se trasladan en los distintos medios de transporte.
type Mercancias struct {
	// Mercancia Nodo requerido para registrar detalladamente la información de los bienes y/o mercancías que se trasladan en los distintos medios de transporte.
	Mercancia []*Mercancia `xml:"Mercancia"`
	// Autotransporte Nodo condicional para registrar la información que permita la identificación del autotransporte de carga, por medio del cual se trasladan los bienes y/o mercancías, que transitan a través de las carreteras del territorio nacional.
	Autotransporte *Autotransporte `xml:"Autotransporte,omitempty"`
	// TransporteMaritimo Nodo condicional para registrar la información que permita la identificación de la embarcación a través de la cual se trasladan los bienes y/o mercancías por vía marítima.
	TransporteMaritimo *TransporteMaritimo `xml:"TransporteMaritimo,omitempty"`
	// TransporteAereo Nodo condicional para registrar la información que permita la identificación del transporte aéreo por medio del cual se trasladan los bienes y/o mercancías.
	TransporteAereo *TransporteAereo `xml:"TransporteAereo,omitempty"`
	// TransporteFerroviario Nodo condicional para registrar la información que permita la identificación del carro o contenedor en el que se trasladan los bienes y/o mercancías por vía férrea.
	TransporteFerroviario *TransporteFerroviario `xml:"TransporteFerroviario,omitempty"`
	// PesoBrutoTotal Atributo requerido para registrar la suma del peso bruto total estimado de los bienes y/o mercancías que se trasladan en los distintos medios de transporte.
	PesoBrutoTotal decimal.Decimal `xml:"PesoBrutoTotal,attr"`
	// UnidadPeso Atributo requerido para registrar la clave de la unidad de medida estandarizada del peso de los bienes y/o mercancías que se trasladan a través de los distintos medios de transporte.
	UnidadPeso string `xml:"UnidadPeso,attr"`
	// PesoNetoTotal Atributo condicional para registrar la suma de los valores indicados en el atributo “PesoNeto” del nodo “DetalleMercancia”.
	PesoNetoTotal decimal.Decimal `xml:"PesoNetoTotal,attr,omitempty"`
	// NumTotalMercancias Atributo requerido para registrar el número total de los bienes y/o mercancías que se trasladan en los distintos medios de transporte, identificándose por cada nodo "Mercancia" registrado en el complemento.
	NumTotalMercancias int `xml:"NumTotalMercancias,attr"`
	// CargoPorTasacion Atributo opcional para expresar el monto del importe pagado por la tasación de los bienes y/o mercancías que se trasladan vía aérea.
	CargoPorTasacion decimal.Decimal `xml:"CargoPorTasacion,attr,omitempty"`
	// LogisticaInversaRecoleccionDevolucion Atributo condicional para expresar si se hace uso de alguno de los servicios de logística inversa, recolección o devolución para el traslado de los bienes y/o mercancías.
	LogisticaInversaRecoleccionDevolucion string `xml:"LogisticaInversaRecoleccionDevolucion,attr,omitempty"`
}

// Mercancia Nodo requerido para registrar detalladamente la información de los bienes y/o mercancías que se trasladan en los distintos medios de transporte.
type Mercancia struct {
	// DocumentacionAduanera Nodo condicional para registrar la información del(los) documento(s) aduanero(s) que se encuentra(n) asociado(s) al traslado de los bienes y/o mercancías por los distintos medios de transporte de procedencia extranjera para acreditar la legal estancia o tenencia durante su traslado en territorio nacional.
	DocumentacionAduanera []*DocumentacionAduanera `xml:"DocumentacionAduanera,omitempty"`
	// GuiasIdentificacion Nodo condicional para registrar la información del(los) número(s) de guía(s) que se encuentre(n) asociado(s) al(los) paquete(s) que se traslada(n) dentro del territorio nacional.
	GuiasIdentificacion []*GuiasIdentificacion `xml:"GuiasIdentificacion,omitempty"`
	// CantidadTransporta Nodo opcional para registrar la cantidad de los bienes y/o mercancías que se trasladan a través de los distintos medios de transporte, que será captada o distribuida en distintos puntos, a fin de identificar el punto de origen y destino correspondiente.
	CantidadTransporta []*CantidadTransporta `xml:"CantidadTransporta,omitempty"`
	// DetalleMercancia Nodo condicional para registrar especificaciones de los bienes y/o mercancías que se trasladan a través de los distintos medios de transporte.
	DetalleMercancia *DetalleMercancia `xml:"DetalleMercancia,omitempty"`
	// BienesTransp Atributo requerido para registrar la clave de producto de los bienes y/o mercancías que se trasladan en los distintos medios de transporte.
	BienesTransp string `xml:"BienesTransp,attr"`
	// ClaveSTCC Atributo opcional para expresar la clave de producto de la STCC (por sus siglas en inglés, Standard Transportation Commodity Code), cuando el medio de transporte utilizado para el traslado de los bienes y/o mercancías sea ferroviario.
	ClaveSTCC string `xml:"ClaveSTCC,attr,omitempty"`
	// Descripcion Atributo requerido para detallar las características de los bienes y/o mercancías que se trasladan en los distintos medios de transporte.
	Descripcion string `xml:"Descripcion,attr"`
	// Cantidad Atributo requerido para expresar la cantidad total de los bienes y/o mercancías que se trasladan a través de los distintos medios de transporte.
	Cantidad decimal.Decimal `xml:"Cantidad,attr"`
	// ClaveUnidad Atributo requerido para registrar la clave de la unidad de medida estandarizada aplicable para la cantidad de los bienes y/o mercancías que se trasladan en los distintos medios de transporte. La unidad debe corresponder con la descripción de los bienes y/o mercancías registrados.
	ClaveUnidad string `xml:"ClaveUnidad,attr"`
	// Unidad Atributo opcional para registrar la unidad de medida propia para la cantidad de los bienes y/o mercancías que se trasladan a través de los distintos medios de transporte. La unidad debe corresponder con la descripción de los bienes y/o mercancías.
	Unidad string `xml:"Unidad,attr,omitempty"`
	// Dimensiones Atributo opcional para expresar las medidas del empaque de los bienes y/o mercancías que se trasladan en los distintos medios de transporte. Se debe registrar la longitud, la altura y la anchura en centímetros o en pulgadas, separados dichos valores con una diagonal, i.e. 30/40/30cm.
	Dimensiones string `xml:"Dimensiones,attr,omitempty"`
	// MaterialPeligroso Atributo condicional para precisar que los bienes y/o mercancías que se trasladan son considerados o clasificados como material peligroso.
	MaterialPeligroso string `xml:"MaterialPeligroso,attr,omitempty"`
	// CveMaterialPeligroso Atributo condicional para indicar la clave del tipo de material peligroso que se transporta de acuerdo a la NOM-002-SCT/2011.
	CveMaterialPeligroso string `xml:"CveMaterialPeligroso,attr,omitempty"`
	// Embalaje Atributo condicional para precisar la clave del tipo de embalaje que se requiere para transportar el material o residuo peligroso.
	Embalaje string `xml:"Embalaje,attr,omitempty"`
	// DescripEmbalaje Atributo opcional para expresar la descripción del embalaje de los bienes y/o mercancías que se trasladan y que se consideran material o residuo peligroso.
	DescripEmbalaje string `xml:"DescripEmbalaje,attr,omitempty"`
	// SectorCOFEPRIS Atributo opcional para expresar la clasificación del producto que se traslada a través de los distintos medios de transporte y que debe contar con autorización por la autoridad correspondiente.
	SectorCOFEPRIS SectorCOFEPRIS `xml:"SectorCOFEPRIS,attr,omitempty"`
	// NombreIngredienteActivo Atributo condicional para expresar el nombre común del ingrediente activo de los precursores, químicos de uso dual, plaguicidas o fertilizantes que se trasladan a través de los distintos medios de transporte.
	NombreIngredienteActivo string `xml:"NombreIngredienteActivo,attr,omitempty"`
	// NomQuimico Atributo condicional para expresar el nombre de la sustancia activa de los precursores, químicos de uso dual o sustancias tóxicas que se traslada a través de los distintos medios de transporte.
	NomQuimico string `xml:"NomQuimico,attr,omitempty"`
	// DenominacionGenericaProd Atributo condicional para expresar el fármaco o la sustancia activa del medicamento, psicotrópico o estupefaciente que se traslada a través de los distintos medios de transporte.
	DenominacionGenericaProd string `xml:"DenominacionGenericaProd,attr,omitempty"`
	// DenominacionDistintivaProd Atributo condicional para expresar la marca con la que se comercializa el producto o nombre que le asigna el laboratorio o fabricante a sus especialidades farmacéuticas con el fin de distinguirlas de otras similares del medicamento, psicotrópico o estupefaciente que se traslada a través de los distintos medios de transporte.
	DenominacionDistintivaProd string `xml:"DenominacionDistintivaProd,attr,omitempty"`
	// Fabricante Atributo condicional para expresar el nombre o razón social del establecimiento que realiza la fabricación o manufactura del medicamento, precursor, químico de uso dual, psicotrópico o estupefaciente que se traslada a través de los distintos medios de transporte.
	Fabricante string `xml:"Fabricante,attr,omitempty"`
	// FechaCaducidad Atributo condicional para registrar la fecha de caducidad del medicamento, psicotrópico o estupefaciente; o para expresar la fecha de reanálisis del precursor o químico de uso dual que se traslada a través de los distintos medios de transporte. Se expresa en la forma AAAA-MM-DD.
	FechaCaducidad string `xml:"FechaCaducidad,attr,omitempty"`
	// LoteMedicamento Atributo condicional para expresar la denominación que identifica y confiere trazabilidad del medicamento, precursor, químico de uso dual, psicotrópico o estupefaciente elaborado en un ciclo de producción, bajo condiciones equivalentes de operación y durante un periodo.
	LoteMedicamento string `xml:"LoteMedicamento,attr,omitempty"`
	// FormaFarmaceutica Atributo condicional para expresar la forma farmacéutica o mezcla del medicamento, precursor, químico de uso dual, psicotrópico o estupefaciente que presenta ciertas características físicas para su adecuada dosificación, conservación y administración.
	FormaFarmaceutica FormaFarmaceutica `xml:"FormaFarmaceutica,attr,omitempty"`
	// CondicionesEspTransp Atributo condicional para expresar la condición en la cual es necesario mantener el medicamento, precursor, químico de uso dual, psicotrópicos o estupefacientes durante el traslado y almacenamiento.
	CondicionesEspTransp CondicionesEspTransp `xml:"CondicionesEspTransp,attr,omitempty"`
	// RegistroSanitarioFolioAutorizacion Atributo condicional para expresar el registro sanitario o folio de autorización con el que cuenta la empresa para el traslado del medicamento, psicotrópico o estupefaciente.
	RegistroSanitarioFolioAutorizacion string `xml:"RegistroSanitarioFolioAutorizacion,attr,omitempty"`
	// PermisoImportacion Atributo condicional para registrar el folio del permiso de importación con el que cuenta el medicamento, precursor, químico de uso dual, psicotrópico o estupefaciente.
	PermisoImportacion string `xml:"PermisoImportacion,attr,omitempty"`
	// FolioImpoVUCEM Atributo condicional para registrar el número de folio de importación VUCEM para la identificación del documento, para el traslado de medicamentos, precursores o químicos de uso dual, sustancias tóxicas, plaguicidas o fertizantes.
	FolioImpoVUCEM string `xml:"FolioImpoVUCEM,attr,omitempty"`
	// NumCAS Atributo condicional para expresar el número Chemical Abstracts Service (CAS) con el que se identifica el compuesto químico de la sustancia tóxica.
	NumCAS string `xml:"NumCAS,attr,omitempty"`
	// RazonSocialEmpImp Atributo condicional para expresar el nombre o razón social de la empresa importadora de las sustancias tóxicas.
	RazonSocialEmpImp string `xml:"RazonSocialEmpImp,attr,omitempty"`
	// NumRegSanPlagCOFEPRIS Atributo condicional para expresar el número de registro sanitario para plaguicidas o fertilizantes cuya importación, comercialización y uso están permitidos en México, mismo que emite la Comisión Intersecretarial para el Control del Proceso y Uso de Plaguicidas, Fertilizantes y Sustancias Tóxicas (CICLOPLAFEST).
	NumRegSanPlagCOFEPRIS string `xml:"NumRegSanPlagCOFEPRIS,attr,omitempty"`
	// DatosFabricante Atributo condicional para registrar el país y nombre o razón social de quien produce o fabrica el ingrediente activo del plaguicida o fertilizante.
	DatosFabricante string `xml:"DatosFabricante,attr,omitempty"`
	// DatosFormulador Atributo condicional para registrar el país y nombre o razón social de quien formula el ingrediente activo del plaguicida o fertilizante.
	DatosFormulador string `xml:"DatosFormulador,attr,omitempty"`
	// DatosMaquilador Atributo condicional para registrar el país y nombre o razón social de quien maquila el ingrediente activo del plaguicida o fertilizante.
	DatosMaquilador string `xml:"DatosMaquilador,attr,omitempty"`
	// UsoAutorizado Atributo condicional para registrar el uso autorizado del plaguicida o fertilizante de acuerdo a la regulación del país.
	UsoAutorizado string `xml:"UsoAutorizado,attr,omitempty"`
	// PesoEnKg Atributo requerido para indicar en kilogramos el peso estimado de los bienes y/o mercancías que se trasladan en los distintos medios de transporte.
	PesoEnKg decimal.Decimal `xml:"PesoEnKg,attr"`
	// ValorMercancia Atributo condicional para expresar el monto del valor de los bienes y/o mercancías que se trasladan en los distintos medios de transporte, de acuerdo al valor mercado, al valor pactado en la contraprestación o bien al valor estimado que determine el contribuyente.
	ValorMercancia decimal.Decimal `xml:"ValorMercancia,attr,omitempty"`
	// Moneda Atributo condicional para identificar la clave de la moneda utilizada para expresar el valor de los bienes y/o mercancías que se trasladan en los distintos medios de transporte. Cuando se usa moneda nacional se registra MXN, de acuerdo a la especificación ISO 4217.
	Moneda types.Moneda `xml:"Moneda,attr,omitempty"`
	// FraccionArancelaria Atributo condicional que sirve para expresar la clave de la fracción arancelaria que corresponde con la descripción de los bienes y/o mercancías que se trasladan en los distintos medios de transporte.
	FraccionArancelaria string `xml:"FraccionArancelaria,attr,omitempty"`
	// UUIDComercioExt Atributo opcional para expresar el folio fiscal (UUID) del comprobante de comercio exterior que se relaciona.
	UUIDComercioExt string `xml:"UUIDComercioExt,attr,omitempty"`
	// TipoMateria Atributo condicional para expresar el estado de la materia o producto al realizar una operación de comercio exterior a través de los distintos medios de transporte.
	TipoMateria string `xml:"TipoMateria,attr,omitempty"`
	// DescripcionMateria Atributo condicional para expresar la descripción del estado de la materia o producto al realizar una operación de comercio exterior a través de los distintos medios de transporte.
	DescripcionMateria string `xml:"DescripcionMateria,attr,omitempty"`
}

// DocumentacionAduanera Nodo condicional para registrar la información del(los) documento(s) aduanero(s) que se encuentra(n) asociado(s) al traslado de los bienes y/o mercancías por los distintos medios de transporte de procedencia extranjera para acreditar la legal estancia o tenencia durante su traslado en territorio nacional.
type DocumentacionAduanera struct {
	// TipoDocumento Atributo requerido para expresar el tipo de documento aduanero que se encuentra asociado al traslado de los bienes y/o mercancías de procedencia extranjera durante su traslado en territorio nacional.
	TipoDocumento DocumentoAduanero `xml:"TipoDocumento,attr"`
	// NumPedimento Atributo condicional para expresar el número de pedimento de importación que se encuentra asociado con el traslado de los bienes y/o mercancías de procedencia extranjera para acreditar la legal estancia y tenencia durante su traslado en territorio nacional.
	NumPedimento string `xml:"NumPedimento,attr,omitempty"`
	// IdentDocAduanero Atributo condicional para expresar el identificador o folio del documento aduanero que se encuentra asociado al traslado de los bienes y/o mercancías de procedencia extranjera para acreditar la legal estancia o tenencia durante su traslado en territorio nacional.
	IdentDocAduanero string `xml:"IdentDocAduanero,attr,omitempty"`
	// RFCImpo Atributo condicional para expresar el RFC del importador de los bienes y/o mercancías que fue registrado en la documentación aduanera correspondiente y este se encuentre en la lista de RFC inscritos no cancelados del SAT (l_RFC).
	RFCImpo string `xml:"RFCImpo,attr,omitempty"`
}

// GuiasIdentificacion Nodo condicional para registrar la información del(los) número(s) de guía(s) que se encuentre(n) asociado(s) al(los) paquete(s) que se traslada(n) dentro del territorio nacional.
type GuiasIdentificacion struct {
	// NumeroGuiaIdentificacion Atributo requerido para expresar el número de guía de cada paquete que se encuentra asociado con el traslado de los bienes y/o mercancías en territorio nacional.
	NumeroGuiaIdentificacion string `xml:"NumeroGuiaIdentificacion,attr"`
	// DescripGuiaIdentificacion Atributo requerido para expresar la descripción del contenido del paquete o carga registrada en la guía, o en el número de identificación, que se encuentra asociado con el traslado de los bienes y/o mercancías dentro del territorio nacional.
	DescripGuiaIdentificacion string `xml:"DescripGuiaIdentificacion,attr"`
	// PesoGuiaIdentificacion Atributo requerido para indicar en kilogramos, el peso del paquete o carga que se está trasladando en territorio nacional  y que se encuentra registrado en la guía o el número de identificación correspondiente.
	PesoGuiaIdentificacion decimal.Decimal `xml:"PesoGuiaIdentificacion,attr"`
}

// CantidadTransporta Nodo opcional para registrar la cantidad de los bienes y/o mercancías que se trasladan a través de los distintos medios de transporte, que será captada o distribuida en distintos puntos, a fin de identificar el punto de origen y destino correspondiente.
type CantidadTransporta struct {
	// Cantidad Atributo requerido para expresar el número de bienes y/o mercancías que se trasladan en los distintos medios de transporte.
	Cantidad decimal.Decimal `xml:"Cantidad,attr"`
	// IDOrigen Atributo requerido para expresar la clave del identificador del origen de los bienes y/o mercancías que se trasladan por los distintos medios de transporte, de acuerdo al valor registrado en el atributo “IDUbicacion”, del nodo “Ubicacion”.
	IDOrigen string `xml:"IDOrigen,attr"`
	// IDDestino Atributo requerido para registrar la clave del identificador del destino de los bienes y/o mercancías que se trasladan a través de los distintos medios de transporte, de acuerdo al valor registrado en el atributo “IDUbicacion”, del nodo “Ubicacion”.
	IDDestino string `xml:"IDDestino,attr"`
	// CvesTransporte Atributo condicional para indicar la clave a través de la cual se identifica el medio por el que se transportan los bienes y/o mercancías.
	CvesTransporte string `xml:"CvesTransporte,attr,omitempty"`
}

// DetalleMercancia Nodo condicional para registrar especificaciones de los bienes y/o mercancías que se trasladan a través de los distintos medios de transporte.
type DetalleMercancia struct {
	// UnidadPesoMerc Atributo requerido para registrar la clave de la unidad de medida estandarizada del peso de los bienes y/o mercancías que se trasladan en los distintos medios de transporte.
	UnidadPesoMerc string `xml:"UnidadPesoMerc,attr"`
	// PesoBruto Atributo requerido para registrar el peso bruto total de los bienes y/o mercancías que se trasladan a través de los diferentes medios de transporte.
	PesoBruto decimal.Decimal `xml:"PesoBruto,attr"`
	// PesoNeto Atributo requerido para registrar el peso neto total de los bienes y/o mercancías que se trasladan en los distintos  medios de transporte.
	PesoNeto decimal.Decimal `xml:"PesoNeto,attr"`
	// PesoTara Atributo requerido para registrar el peso bruto, menos el peso neto de los bienes y/o mercancías que se trasladan a través de los distintos medios de transporte.
	PesoTara decimal.Decimal `xml:"PesoTara,attr"`
	// NumPiezas Atributo opcional para registrar el número de piezas de los bienes y/o mercancías que se trasladan en los distintos medios de transporte.
	NumPiezas int `xml:"NumPiezas,attr,omitempty"`
}

// Autotransporte Nodo condicional para registrar la información que permita la identificación del autotransporte de carga, por medio del cual se trasladan los bienes y/o mercancías, que transitan a través de las carreteras del territorio nacional.
type Autotransporte struct {
	// IdentificacionVehicular Nodo requerido para registrar los datos de identificación del autotransporte en el que se trasladan los bienes y/o mercancías.
	IdentificacionVehicular *IdentificacionVehicular `xml:"IdentificacionVehicular"`
	// Seguros Nodo requerido para registrar los datos de las pólizas de seguro que cubren los riesgos en el traslado de los bienes y/o mercancías.
	Seguros *Seguros `xml:"Seguros"`
	// Remolques Nodo condicional para registrar los datos del(los) remolque(s) o semirremolque(s) que se adaptan al autotransporte para realizar el traslado de los bienes y/o mercancías.
	Remolques Remolques `xml:"Remolques,omitempty"`
	// PermSCT Atributo requerido para registrar la clave del tipo de permiso proporcionado por la Secretaría de Infraestructura, Comunicaciones y Transportes (SICT) o la autoridad análoga, el cual debe corresponder con el tipo de autotransporte utilizado para el traslado de los bienes y/o mercancías de acuerdo al catálogo correspondiente.
	PermSCT string `xml:"PermSCT,attr"`
	// NumPermisoSCT Atributo requerido para registrar el número del permiso otorgado por la Secretaría de Infraestructura, Comunicaciones y Transportes (SICT) o la autoridad correspondiente, al autotransporte utilizado para el traslado de los bienes y/o mercancías.
	NumPermisoSCT string `xml:"NumPermisoSCT,attr"`
}

// IdentificacionVehicular Nodo requerido para registrar los datos de identificación del autotransporte en el que se trasladan los bienes y/o mercancías.
type IdentificacionVehicular struct {
	// ConfigVehicular Atributo requerido para expresar la clave de nomenclatura del autotransporte que es utilizado para transportar los bienes y/o mercancías.
	ConfigVehicular string `xml:"ConfigVehicular,attr"`
	// PesoBrutoVehicular Atributo requerido para indicar en toneladas el peso bruto vehicular permitido del autotransporte de acuerdo a la NOM-SCT-012-2017 que es utilizado para realizar el traslado de los bienes y/o mercancías.
	PesoBrutoVehicular decimal.Decimal `xml:"PesoBrutoVehicular,attr"`
	// PlacaVM Atributo requerido para registrar solo los caracteres alfanuméricos, sin guiones ni espacios de la placa vehicular del autotransporte que es utilizado para transportar los bienes y/o mercancías.
	PlacaVM string `xml:"PlacaVM,attr"`
	// AnioModeloVM Atributo requerido para registrar el año del autotransporte que es utilizado para transportar los bienes y/o mercancías.
	AnioModeloVM string `xml:"AnioModeloVM,attr"`
}

// Seguros Nodo requerido para registrar los datos de las pólizas de seguro que cubren los riesgos en el traslado de los bienes y/o mercancías.
type Seguros struct {
	// AseguraRespCivil Atributo requerido para registrar el nombre de la aseguradora que cubre los riesgos por responsabilidad civil del autotransporte utilizado para el traslado de los bienes y/o mercancías.
	AseguraRespCivil string `xml:"AseguraRespCivil,attr"`
	// PolizaRespCivil Atributo requerido para registrar el número de póliza asignado por la aseguradora, que cubre los riesgos por responsabilidad civil del autotransporte utilizado para el traslado de los bienes y/o mercancías.
	PolizaRespCivil string `xml:"PolizaRespCivil,attr"`
	// AseguraMedAmbiente Atributo condicional para registrar el nombre de la aseguradora, que cubre los posibles daños al medio ambiente cuando exista al menos una mercancía tipificada como material peligroso se debe registrar la información del atributo “AseguraMedAmbiente” (aplicable para los transportistas de materiales, residuos o remanentes y desechos peligrosos.
	AseguraMedAmbiente string `xml:"AseguraMedAmbiente,attr,omitempty"`
	// Atributo condicional para registrar el número de póliza asignado por la aseguradora, que cubre los posibles daños al medio ambiente cuando exista al menos una mercancía tipificada como material peligroso se debe registrar la información del atributo “AseguraMedAmbiente” (aplicable para los transportistas de materiales, residuos o remanentes y desechos peligrosos).
	PolizaMedAmbiente string `xml:"PolizaMedAmbiente,attr,omitempty"`
	// AseguraCarga Atributo opcional para registrar el nombre de la aseguradora que cubre los riesgos de la carga (bienes y/o mercancías) del autotransporte utilizado para el traslado.
	AseguraCarga string `xml:"AseguraCarga,attr,omitempty"`
	// PolizaCarga Atributo opcional para expresar el número de póliza asignado por la aseguradora que cubre los riesgos de la carga (bienes y/o mercancías) del autotransporte utilizado para el traslado.
	PolizaCarga string `xml:"PolizaCarga,attr,omitempty"`
	// PrimaSeguro Atributo opcional para registrar el valor del importe por el cargo adicional convenido entre el transportista y el cliente, el cual será igual al valor de la prima del seguro contratado, conforme a lo establecido en la cláusula novena del Acuerdo por el que se homologa la Carta de Porte regulada por la Ley de Caminos, Puentes y Autotransporte Federal, con el complemento Carta Porte que debe acompañar al Comprobante Fiscal Digital por Internet (CFDI).
	PrimaSeguro decimal.Decimal `xml:"PrimaSeguro,attr,omitempty"`
}

// Remolques Nodo condicional para registrar los datos del(los) remolque(s) o semirremolque(s) que se adaptan al autotransporte para realizar el traslado de los bienes y/o mercancías.
type Remolques []*Remolque

func (rem *Remolques) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var rem2 struct {
		Remolque []*Remolque `xml:"Remolque"`
	}
	if err := d.DecodeElement(&rem2, &start); err != nil {
		return err
	}
	*rem = rem2.Remolque
	return nil
}

// Remolque Nodo requerido para expresar la información del(los) remolque(s) o semirremolque(s) que se adapta(n) al autotransporte para realizar el traslado de los bienes y/o mercancías.
type Remolque struct {
	// SubTipoRem Atributo requerido para expresar la clave del subtipo de remolque o semirremolques que se emplean con el autotransporte para el traslado de los bienes y/o mercancías.
	SubTipoRem string `xml:"SubTipoRem,attr"`
	// Placa Atributo requerido para registrar los caracteres alfanuméricos, sin guiones ni espacios de la placa vehicular del remolque o semirremolque que es utilizado para transportar los bienes y/o mercancías.
	Placa string `xml:"Placa,attr"`
}

// TransporteMaritimo Nodo condicional para registrar la información que permita la identificación de la embarcación a través de la cual se trasladan los bienes y/o mercancías por vía marítima.
type TransporteMaritimo struct {
	// Contenedor Nodo opcional para registrar los datos del contenedor en el que se transportan los bienes y/o mercancías.
	Contenedor []ContenedorMaritimo `xml:"Contenedor"`
	// PermSCT Atributo opcional para registrar la clave del permiso proporcionado por la SCT, la cual debe corresponder con la embarcación que se está utilizando para el traslado de los bienes y/o mercancías, de acuerdo al catálogo correspondiente.
	PermSCT string `xml:"PermSCT,attr,omitempty"`
	// NumPermisoSCT Atributo opcional para registrar el número del permiso otorgado por la SCT a la embarcación utilizada para el traslado de los bienes y/o mercancías.
	NumPermisoSCT string `xml:"NumPermisoSCT,attr,omitempty"`
	// NombreAseg Atributo opcional para registrar el nombre de la aseguradora que cubre la protección e indemnización por responsabilidad civil de la embarcación en el traslado de los bienes y/o mercancías.
	NombreAseg string `xml:"NombreAseg,attr,omitempty"`
	// NumPolizaSeguro Atributo opcional para registrar el número de póliza asignada por la aseguradora que cubre la protección e indemnización por responsabilidad civil de la embarcación en el traslado de los bienes y/o mercancías.
	NumPolizaSeguro string `xml:"NumPolizaSeguro,attr,omitempty"`
	// TipoEmbarcacion Atributo requerido para registrar la clave de identificación del tipo de embarcación que es utilizado para trasladar los bienes y/o mercancías.
	TipoEmbarcacion string `xml:"TipoEmbarcacion,attr"`
	// Matricula Atributo requerido para registrar el número de la matrícula o registro de la embarcación que es utilizada para transportar los bienes y/o mercancías.
	Matricula string `xml:"Matricula,attr"`
	// NumeroOMI Atributo requerido para registrar el número de identificación asignado por la Organización Marítima Internacional, a la embarcación encargada de transportar los bienes y/o mercancías.
	NumeroOMI string `xml:"NumeroOMI,attr"`
	// AnioEmbarcacion Atributo opcional para registrar el año de la embarcación en la que se transportan los bienes y/o mercancías.
	AnioEmbarcacion int `xml:"AnioEmbarcacion,attr,omitempty"`
	// NombreEmbarc Atributo opcional para registrar el nombre de la embarcación en la que se realiza el traslado de los bienes y/o mercancías.
	NombreEmbarc string `xml:"NombreEmbarc,attr,omitempty"`
	// NacionalidadEmbarc Atributo requerido para registrar la clave del país correspondiente a la nacionalidad de la embarcación que transporta los bienes y/o mercancías.
	NacionalidadEmbarc types.Pais `xml:"NacionalidadEmbarc,attr"`
	// UnidadesDeArqBruto Atributo requerido para registrar el valor de las unidades de arqueo bruto conforme a las medidas internacionales definidas por el ITC para cada tipo de buque o embarcación en la que se transportan los bienes y/o mercancías.
	UnidadesDeArqBruto decimal.Decimal `xml:"UnidadesDeArqBruto,attr"`
	// TipoCarga Atributo requerido para especificar el tipo de carga en el cual se clasifican los bienes y/o mercancías que se transportan en la embarcación.
	TipoCarga string `xml:"TipoCarga,attr"`
	// NumCertITC Atributo requerido para registrar el número del certificado emitido por la ITC para la embarcación o buque que transporta los bienes y/o mercancías.
	NumCertITC string `xml:"NumCertITC,attr"`
	// Eslora Atributo opcional para registrar la longitud de eslora, definida en pies, con la que cuenta la embarcación o el buque en el que se transportan los bienes y/o mercancías.
	Eslora decimal.Decimal `xml:"Eslora,attr,omitempty"`
	// Manga Atributo opcional para registrar la longitud de manga, definida en pies, con la que cuenta la embarcación o el buque en el que se transportan los bienes y/o mercancías.
	Manga decimal.Decimal `xml:"Manga,attr,omitempty"`
	// Calado Atributo opcional para registrar la longitud del calado, definida en pies, con la que cuenta la embarcación o el buque en el que se transportan los bienes y/o mercancías.
	Calado decimal.Decimal `xml:"Calado,attr,omitempty"`
	// LineaNaviera Atributo opcional para registrar el nombre de la línea naviera autorizada de gestionar el traslado de los bienes y/o mercancías por vía marítima.
	LineaNaviera string `xml:"LineaNaviera,attr,omitempty"`
	// NombreAgenteNaviero Atributo requerido para registrar el nombre del agente naviero consignatario autorizado para gestionar el traslado de los bienes y/o mercancías por vía marítima.
	NombreAgenteNaviero string `xml:"NombreAgenteNaviero,attr"`
	// NumAutorizacionNaviero Atributo requerido para expresar el número de la autorización como agente naviero consignatario emitida por la SCT.
	NumAutorizacionNaviero string `xml:"NumAutorizacionNaviero,attr"`
	// NumViaje Atributo opcional para registrar el número del viaje con el que se identifica el traslado de los bienes y/o mercancías en el buque o la embarcación.
	NumViaje string `xml:"NumViaje,attr,omitempty"`
	// NumConocEmbarc Atributo opcional para registrar el número de conocimiento de embarque con el que se identifica el traslado de los bienes y/o mercancías.
	NumConocEmbarc string `xml:"NumConocEmbarc,attr,omitempty"`
}

// ContenedorMaritimo Nodo requerido para registrar los datos del contenedor en el que se transportan los bienes y/o mercancías.
type ContenedorMaritimo struct {
	// TipoContenedor Atributo requerido para registrar la clave de identificación correspondiente con el tipo de contenedor marítimo en el que se transportan los bienes y/o mercancías.
	TipoContenedor string `xml:"TipoContenedor,attr"`
	// MatriculaContenedor Atributo condicional para registrar la matrícula o el número de identificación del contenedor marítimo en el que se transportan los bienes y/o mercancías, el cual está integrado por el código del propietario, el número de serie y el dígito de control correspondiente.
	MatriculaContenedor string `xml:"MatriculaContenedor,attr,omitempty"`
	// NumPrecinto Atributo condicional para registrar el número del sello o precinto de los contenedores marítimos que son utilizados para trasladar los bienes y/o mercancías.
	NumPrecinto string `xml:"NumPrecinto,attr,omitempty"`
	// IdCCPRelacionado Atributo condicional para registrar el identificador del complemento Carta Porte (IdCCP) de un CFDI previamente certificado para el traslado de bienes o mercancías mediante autotransporte, únicamente aplica para traslados mediante ferri.
	IdCCPRelacionado string `xml:"IdCCPRelacionado,attr,omitempty"`
	// PlacaVMCCP Atributo condicional para registrar los caracteres alfanuméricos, sin guiones ni espacios de la placa vehicular del autotransporte registrado en el CFDI con complemento Carta Porte del autotransporte, únicamente aplica para traslado mediante ferri.
	PlacaVMCCP string `xml:"PlacaVMCCP,attr,omitempty"`
	// FechaCertificacionCCP Atributo condicional para registrar la fecha y hora de certificación del CDFI con complemento Carta Porte del autotransporte, únicamente aplica para traslado mediante ferri.
	FechaCertificacionCCP string `xml:"FechaCertificacionCCP,attr,omitempty"`
	// RemolquesCCP Nodo condicional para registrar los datos del(los) remolque(s) o semirremolque(s) que se adaptan al autotransporte que realizó el traslado de los bienes y/o mercancías registrado en el CFDI con complemento Carta Porte de autotransporte, únicamente aplica para traslado mediante ferri.
	RemolquesCCP *RemolquesCCP `xml:"RemolquesCCP,omitempty"`
}

// RemolquesCCP Nodo condicional para registrar los datos del(los) remolque(s) o semirremolque(s) que se adaptan al autotransporte que realizó el traslado de los bienes y/o mercancías registrado en el CFDI con complemento Carta Porte de autotransporte, únicamente aplica para traslado mediante ferri.
type RemolquesCCP struct {
	RemolqueCCP []RemolqueCCP `xml:"RemolqueCCP"`
}

// RemolqueCCP Nodo requerido para expresar la información del(los) remolque(s) o semirremolque(s) que se adapta(n) al autotransporte que realizó el traslado de los bienes y/o mercancías registrado en el CFDI con complemento Carta Porte, únicamente aplica para traslado mediante ferri.
type RemolqueCCP struct {
	SubTipoRemCCP string `xml:"SubTipoRemCCP,attr"`
	PlacaCCP      string `xml:"PlacaCCP,attr"`
}

// TransporteAereo Nodo condicional para registrar la información que permita la identificación del transporte aéreo por medio del cual se trasladan los bienes y/o mercancías.
type TransporteAereo struct {
	// PermSCT Atributo requerido para registrar la clave del permiso proporcionado por la SCT o la autoridad análoga, la cual debe corresponder con la aeronave que se está utilizando para realizar el traslado de los bienes y/o mercancías por vía aérea.
	PermSCT string `xml:"PermSCT,attr"`
	// NumPermisoSCT Atributo requerido para registrar el número de permiso o valor análogo proporcionado por la SCT o la autoridad análoga, según corresponda, para el transporte de bienes y/o mercancías por vía aérea.
	NumPermisoSCT string `xml:"NumPermisoSCT,attr"`
	// MatriculaAeronave Atributo opcional para registrar el número de la matrícula de la aeronave con la que se realiza el traslado de los bienes y/o mercancías en territorio nacional el cual tiene una longitud de 10 posiciones y se compone de valores alfanuméricos, más el carácter especial denominado guion medio “-“.
	MatriculaAeronave string `xml:"MatriculaAeronave,attr,omitempty"`
	// NombreAseg Atributo opcional para registrar el nombre de la aseguradora que cubre los riesgos de la aeronave con la que transportan los bienes y/o mercancías.
	NombreAseg string `xml:"NombreAseg,attr,omitempty"`
	// NumPolizaSeguro Atributo opcional para registrar el número de póliza asignado por la aseguradora que cubre la protección e indemnización por responsabilidad civil de la aeronave que transporta los bienes y/o mercancías.
	NumPolizaSeguro string `xml:"NumPolizaSeguro,attr,omitempty"`
	// NumeroGuia Atributo requerido para registrar el número de guía aérea con el que se trasladan los bienes y/o mercancías.
	NumeroGuia string `xml:"NumeroGuia,attr"`
	// LugarContrato Atributo opcional para registrar el lugar, entidad, región, localidad o análogo, donde se celebró el contrato para realizar el traslado de los bienes y/o mercancías.
	LugarContrato string `xml:"LugarContrato,attr,omitempty"`
	// CodigoTransportista Atributo requerido para registrar el valor del código que tiene asignado el transportista el cual debe contener alguna de las claves contenidas en el catálogo correspondiente.
	CodigoTransportista string `xml:"CodigoTransportista,attr"`
	// RFCEmbarcador Atributo opcional para registrar el RFC del embarcador de los bienes y/o mercancías que se trasladan.
	RFCEmbarcador string `xml:"RFCEmbarcador,attr,omitempty"`
	// NumRegIdTribEmbarc Atributo condicional para incorporar el número de identificación o registro fiscal del país de residencia cuando el embarcador sea residente en el extranjero para los efectos fiscales correspondientes de los bienes y/o mercancías que se trasladan.
	NumRegIdTribEmbarc string `xml:"NumRegIdTribEmbarc,attr,omitempty"`
	// ResidenciaFiscalEmbarc Atributo condicional para registrar la clave del país de residencia para efectos fiscales del embarcador de los bienes y/o mercancías.
	ResidenciaFiscalEmbarc types.Pais `xml:"ResidenciaFiscalEmbarc,attr,omitempty"`
	// NombreEmbarcador Atributo opcional para registrar el nombre del embarcador de los bienes y/o mercancías que se trasladan, ya sea nacional o extranjero.
	NombreEmbarcador string `xml:"NombreEmbarcador,attr,omitempty"`
}

// TransporteFerroviario Nodo condicional para registrar la información que permita la identificación del carro o contenedor en el que se trasladan los bienes y/o mercancías por vía férrea.
type TransporteFerroviario struct {
	// DerechosDePaso Nodo opcional para registrar los tipos de derechos de paso cubiertos por el transportista en las vías férreas de las cuales no es concesionario o asignatario, así como la distancia establecida en kilómetros.
	DerechosDePaso []DerechosDePaso `xml:"DerechosDePaso,omitempty"`
	// Carro Nodo requerido para registrar la información que permite identificar el (los) carro(s) en el (los) que se trasladan los bienes y/o mercancías por vía férrea.
	Carro []TransporteFerroviarioCarro `xml:"Carro"`
	// TipoDeServicio Atributo requerido para registrar la clave del tipo de servicio utilizado para el traslado de los bienes y/o mercancías por vía férrea.
	TipoDeServicio string `xml:"TipoDeServicio,attr"`
	// TipoDeTrafico Atributo requerido para registrar la clave del tipo de tráfico (interrelación entre concesionarios) para realizar el traslado de los bienes y/o mercancías por vía férrea dentro del territorio nacional.
	TipoDeTrafico string `xml:"TipoDeTrafico,attr"`
	// NombreAseg Atributo opcional para registrar el nombre de la aseguradora que cubre los riesgos para el traslado de los bienes y/o mercancías por vía férrea.
	NombreAseg string `xml:"NombreAseg,attr,omitempty"`
	// NumPolizaSeguro Atributo opcional para registrar el número de póliza asignada por la aseguradora para la protección e indemnización por responsabilidad civil en el traslado de los bienes y/o mercancías que se realiza por vía férrea.
	NumPolizaSeguro string `xml:"NumPolizaSeguro,attr,omitempty"`
}

// DerechosDePaso Nodo opcional para registrar los tipos de derechos de paso cubiertos por el transportista en las vías férreas de las cuales no es concesionario o asignatario, así como la distancia establecida en kilómetros.
type DerechosDePaso struct {
	// TipoDerechoDePaso Atributo requerido para registrar la clave del derecho de paso pagado por el transportista en las vías férreas de las cuales no es concesionario o asignatario.
	TipoDerechoDePaso string `xml:"TipoDerechoDePaso,attr"`
	// KilometrajePagado Atributo requerido para registrar el total de kilómetros pagados por el transportista en las vías férreas de las cuales no es concesionario o asignatario con el derecho de paso.
	KilometrajePagado decimal.Decimal `xml:"KilometrajePagado,attr"`
}

// TransporteFerroviarioCarro Nodo requerido para registrar la información que permite identificar el (los) carro(s) en el (los) que se trasladan los bienes y/o mercancías por vía férrea.
type TransporteFerroviarioCarro struct {
	// Contenedor Nodo condicional para especificar el tipo de contenedor o vagón en el que se trasladan los bienes y/o mercancías por vía férrea.
	Contenedor *TransporteFerroviarioCarroContenedor `xml:"Contenedor,omitempty"`
	// TipoCarro Atributo requerido para registrar la clave del tipo de carro utilizado para el traslado de los bienes y/o mercancías por vía férrea.
	TipoCarro string `xml:"TipoCarro,attr"`
	// MatriculaCarro Atributo requerido para registrar el número de contenedor, carro de ferrocarril o número económico del vehículo en el que se trasladan los bienes y/o mercancías por vía férrea.
	MatriculaCarro string `xml:"MatriculaCarro,attr"`
	// GuiaCarro Atributo requerido para registrar el número de guía asignado al contenedor, carro de ferrocarril o vehículo, en el que se trasladan los bienes y/o mercancías por vía férrea.
	GuiaCarro string `xml:"GuiaCarro,attr"`
	// ToneladasNetasCarro Atributo requerido para registrar la cantidad de las toneladas netas depositadas en el contenedor, carro de ferrocarril o vehículo en el que se trasladan los bienes y/o mercancías por vía férrea.
	ToneladasNetasCarro decimal.Decimal `xml:"ToneladasNetasCarro,attr"`
}

// TransporteFerroviarioCarroContenedor Nodo condicional para especificar el tipo de contenedor o vagón en el que se trasladan los bienes y/o mercancías por vía férrea.
type TransporteFerroviarioCarroContenedor struct {
	// TipoContenedor Atributo requerido para registrar la clave con la que se identifica al tipo de contenedor o el vagón en el que se realiza el traslado de los bienes y/o mercancías.
	TipoContenedor string `xml:"TipoContenedor,attr"`
	// PesoContenedorVacio Atributo requerido para registrar en kilogramos, el peso del contenedor vacío en el que se trasladan los bienes y/o mercancías.
	PesoContenedorVacio decimal.Decimal `xml:"PesoContenedorVacio,attr"`
	// PesoNetoMercancia Atributo requerido para registrar en kilogramos el peso neto de los bienes y/o mercancías que son trasladados en el contenedor.
	PesoNetoMercancia decimal.Decimal `xml:"PesoNetoMercancia,attr"`
}

// FiguraTransporte Nodo condicional para indicar los datos de la(s) figura(s) del transporte que interviene(n) en el traslado de los bienes y/o mercancías realizado a través de los distintos medios de transporte dentro del territorio nacional, cuando el dueño de dicho medio sea diferente del emisor del comprobante con el complemento Carta Porte.
type FiguraTransporte struct {
	// TiposFigura Nodo condicional para indicar los datos del(los) tipo(s) de figura(s) que participan en el traslado de los bienes y/o mercancías en los distintos medios de transporte.
	TiposFigura []*TiposFigura `xml:"TiposFigura"`
}

// TiposFigura Nodo condicional para indicar los datos del(los) tipo(s) de figura(s) que participan en el traslado de los bienes y/o mercancías en los distintos medios de transporte.
type TiposFigura struct {
	// PartesTransporte Nodo condicional para indicar los datos de las partes del transporte de las cuales el emisor del comprobante es distinto al dueño de las mismas, por ejemplo: vehículos, máquinas, contenedores, plataformas, etc; mismos que son utilizados para el traslado de los bienes y/o mercancías.
	PartesTransporte []*PartesTransporte `xml:"PartesTransporte,omitempty"`
	// Domicilio Nodo opcional para registrar información del domicilio del(los) tipo(s) de figura transporte que intervenga(n) en el traslado de los bienes y/o mercancías.
	Domicilio *Domicilio `xml:"Domicilio,omitempty"`
	// TipoFigura Atributo requerido para registrar la clave de la figura de transporte que interviene en el traslado de los bienes y/o mercancías.
	TipoFigura string `xml:"TipoFigura,attr"`
	// RFCFigura Atributo condicional para registrar el RFC de la figura de transporte que interviene en el traslado de los bienes y/o mercancías.
	RFCFigura string `xml:"RFCFigura,attr,omitempty"`
	// NumLicencia Atributo condicional para expresar el número de la licencia o el permiso otorgado al operador del autotransporte de carga en el que realiza el traslado de los bienes y/o mercancías.
	NumLicencia string `xml:"NumLicencia,attr,omitempty"`
	// NombreFigura Atributo requerido para registrar el nombre de la figura de transporte que interviene en el traslado de los bienes y/o mercancías.
	NombreFigura string `xml:"NombreFigura,attr,omitempty"`
	// NumRegIdTribFigura Atributo condicional para registrar el número de identificación o registro fiscal del país de residencia de la figura de transporte que interviene en el traslado de los bienes y/o mercancías, cuando se trate de residentes en el extranjero para los efectos fiscales correspondientes.
	NumRegIdTribFigura string `xml:"NumRegIdTribFigura,attr,omitempty"`
	// ResidenciaFiscalFigura Atributo condicional para registrar la clave del país de residencia de la figura de transporte que interviene en el traslado de los bienes y/o mercancías para los efectos fiscales correspondientes.
	ResidenciaFiscalFigura types.Pais `xml:"ResidenciaFiscalFigura,attr,omitempty"`
}

// PartesTransporte Nodo condicional para indicar los datos de las partes del transporte de las cuales el emisor del comprobante es distinto al dueño de las mismas, por ejemplo: vehículos, máquinas, contenedores, plataformas, etc; mismos que son utilizados para el traslado de los bienes y/o mercancías.
type PartesTransporte struct {
	// ParteTransporte Atributo requerido para registrar información de la parte del transporte de la cual el emisor del comprobante es distinto al dueño de la misma, por ejemplo: vehículos, máquinas, contenedores, plataformas, etc; que se utilicen para el traslado de los bienes y/o mercancías.
	ParteTransporte string `xml:"ParteTransporte,attr"`
}

// Domicilio Nodo opcional para registrar información del domicilio.
type Domicilio struct {
	// Calle Atributo opcional que sirve para registrar la calle en la que está ubicado el domicilio.
	Calle string `xml:"Calle,attr,omitempty"`
	// NumeroExterior Atributo opcional que sirve para registrar el número exterior en donde se ubica el domicilio.
	NumeroExterior string `xml:"NumeroExterior,attr,omitempty"`
	// NumeroInterior Atributo opcional que sirve para registrar el número interior, en caso de existir, en donde se ubica el domicilio.
	NumeroInterior string `xml:"NumeroInterior,attr,omitempty"`
	// Colonia Atributo opcional que sirve para expresar la clave de la colonia o dato análogo en donde se ubica el domicilio.
	Colonia string `xml:"Colonia,attr,omitempty"`
	// Localidad Atributo opcional para registrar la clave de la ciudad, población, distrito o dato análogo de donde se encuentra ubicado el domicilio.
	Localidad string `xml:"Localidad,attr,omitempty"`
	// Referencia Atributo opcional para registrar una referencia geográfica adicional que permita una fácil o precisa ubicación del domicilio; por ejemplo, las coordenadas del GPS.
	Referencia string `xml:"Referencia,attr,omitempty"`
	// Municipio Atributo opcional para registrar la clave del municipio, delegación o alcaldía, condado o dato análogo en donde se encuentra ubicado el domicilio.
	Municipio string `xml:"Municipio,attr,omitempty"`
	// Estado Atributo requerido para registrar el estado, entidad, región, comunidad, o dato análogo en donde se encuentra ubicado el domicilio.
	Estado string `xml:"Estado,attr"`
	// Pais Atributo requerido que sirve para registrar la clave del país en donde se encuentra ubicado el domicilio, conforme al catálogo c_Pais del CFDI publicado en el portal del SAT en Internet de acuerdo a la especificación ISO 3166-1.
	Pais types.Pais `xml:"Pais,attr"`
	// CodigoPostal Atributo requerido para registrar el código postal en donde se encuentra ubicado el domicilio.
	CodigoPostal string `xml:"CodigoPostal,attr"`
}
