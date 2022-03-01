package types

// ClaveProdServ01 01010101 No existe en el catálogo.
const ClaveProdServ01 = "01010101"

const (
	UnidadServicio  = "E48"
	UnidadActividad = "ACT"
	UnidadPieza     = "H87"
	UnidadKilogramo = "KGM"
	UnidadUno       = "C62"
	UnidadLitro     = "LTR"
	UnidadMetro     = "MTR"
)

// May be one of AGU, BCN, BCS, CAM, CHP, CHH, COA, COL, DIF, CMX, DUR, GUA, GRO, HID, JAL, MEX, MIC, MOR, NAY, NLE, OAX, PUE, QUE, ROO, SLP, SIN, SON, TAB, TAM, TLA, VER, YUC, ZAC, AL, AK, AZ, AR, CA, NC, SC, CO, CT, ND, SD, DE, FL, GA, HI, ID, IL, IN, IA, KS, KY, LA, ME, MD, MA, MI, MN, MS, MO, MT, NE, NV, NJ, NY, NH, NM, OH, OK, OR, PA, RI, TN, TX, UT, VT, VA, WV, WA, WI, WY, ON, QC, NS, NB, MB, BC, PE, SK, AB, NL, NT, YT, UN
type Estado string

// May be one of 01, 02, 03
type Exportacion string

const (
	// Exportacion01 No aplica.
	Exportacion01 Exportacion = "01"
	// Exportacion02 Definitiva.
	Exportacion02 Exportacion = "02"
	// Exportacion03 Provisional.
	Exportacion03 Exportacion = "03"
)

func (e Exportacion) Desc() string {
	switch e {
	case Exportacion01:
		return "No aplica"
	case Exportacion02:
		return "Definitiva"
	case Exportacion03:
		return "Provisional"
	}
	return ""
}

// May be one of 01, 02, 03, 04, 05, 06, 08, 12, 13, 14, 15, 17, 23, 24, 25, 26, 27, 28, 29, 30, 31, 99
type FormaPago string

const (
	// FormaPago01 Efectivo.
	FormaPago01 = "01"
	// FormaPago02 Cheque nominativo.
	FormaPago02 = "02"
	// FormaPago03 Transferencia electrónica de fondos.
	FormaPago03 = "03"
	// FormaPago04 Tarjeta de crédito.
	FormaPago04 = "04"
	// FormaPago05 Monedero electrónico.
	FormaPago05 = "05"
	// FormaPago06 Dinero electrónico.
	FormaPago06 = "06"
	// FormaPago08 Vales de despensa.
	FormaPago08 = "08"
	// FormaPago12 Dación en pago.
	FormaPago12 = "12"
	// FormaPago13 Pago por subrogación.
	FormaPago13 = "13"
	// FormaPago14 Pago por consignación.
	FormaPago14 = "14"
	// FormaPago15 Condonación.
	FormaPago15 = "15"
	// FormaPago17 Compensación.
	FormaPago17 = "17"
	// FormaPago23 Novación.
	FormaPago23 = "23"
	// FormaPago24 Confusión.
	FormaPago24 = "24"
	// FormaPago25 Remisión de deuda.
	FormaPago25 = "25"
	// FormaPago26 Prescripción o caducidad.
	FormaPago26 = "26"
	// FormaPago27 A satisfacción del acreedor.
	FormaPago27 = "27"
	// FormaPago28 Tarjeta de débito.
	FormaPago28 = "28"
	// FormaPago29 Tarjeta de servicios.
	FormaPago29 = "29"
	// FormaPago30 Aplicación de anticipos.
	FormaPago30 = "30"
	// FormaPago31 Intermediario pagos.
	FormaPago31 = "31"
	// FormaPago99 Por definir.
	FormaPago99 = "99"
)

func (f FormaPago) Desc() string {
	switch f {
	case FormaPago01:
		return "Efectivo"
	case FormaPago02:
		return "Cheque nominativo"
	case FormaPago03:
		return "Transferencia electrónica de fondos"
	case FormaPago04:
		return "Tarjeta de crédito"
	case FormaPago05:
		return "Monedero electrónico"
	case FormaPago06:
		return "Dinero electrónico"
	case FormaPago08:
		return "Vales de despensa"
	case FormaPago12:
		return "Dación en pago"
	case FormaPago13:
		return "Pago por subrogación"
	case FormaPago14:
		return "Pago por consignación"
	case FormaPago15:
		return "Condonación"
	case FormaPago17:
		return "Compensación"
	case FormaPago23:
		return "Novación"
	case FormaPago24:
		return "Confusión"
	case FormaPago25:
		return "Remisión de deuda"
	case FormaPago26:
		return "Prescripción o caducidad"
	case FormaPago27:
		return "A satisfacción del acreedor"
	case FormaPago28:
		return "Tarjeta de débito"
	case FormaPago29:
		return "Tarjeta de servicios"
	case FormaPago30:
		return "Aplicación de anticipos"
	case FormaPago31:
		return "Intermediario pagos"
	case FormaPago99:
		return "Por definir"
	}
	return ""
}

// May be one of 001, 002, 003
type Impuesto string

const (
	//ImpuestoISR 001.
	ImpuestoISR Impuesto = "001"
	//ImpuestoIVA 002.
	ImpuestoIVA Impuesto = "002"
	//ImpuestoIEPS 003.
	ImpuestoIEPS Impuesto = "003"
)

func (i Impuesto) Desc() string {
	switch i {
	case ImpuestoISR:
		return "ISR"
	case ImpuestoIVA:
		return "IVA"
	case ImpuestoIEPS:
		return "IEPS"
	}
	return ""
}

// May be one of 01, 02, 03, 04, 05, 06, 07, 08, 09, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 66, 67, 68, 69
type Localidad string

// May be one of PUE, PPD
type MetodoPago string

const (
	MetodoPagoPUE MetodoPago = "PUE"
	MetodoPagoPPD MetodoPago = "PPD"
)

func (m MetodoPago) Desc() string {
	switch m {
	case MetodoPagoPUE:
		return "Pago en una sola exhibición"
	case MetodoPagoPPD:
		return "Pago en parcialidades o diferido"
	}
	return ""
}

// May be one of AED, AFN, ALL, AMD, ANG, AOA, ARS, AUD, AWG, AZN, BAM, BBD, BDT, BGN, BHD, BIF, BMD, BND, BOB, BOV, BRL, BSD, BTN, BWP, BYR, BZD, CAD, CDF, CHE, CHF, CHW, CLF, CLP, CNY, COP, COU, CRC, CUC, CUP, CVE, CZK, DJF, DKK, DOP, DZD, EGP, ERN, ETB, EUR, FJD, FKP, GBP, GEL, GHS, GIP, GMD, GNF, GTQ, GYD, HKD, HNL, HRK, HTG, HUF, IDR, ILS, INR, IQD, IRR, ISK, JMD, JOD, JPY, KES, KGS, KHR, KMF, KPW, KRW, KWD, KYD, KZT, LAK, LBP, LKR, LRD, LSL, LYD, MAD, MDL, MGA, MKD, MMK, MNT, MOP, MRO, MUR, MVR, MWK, MXN, MXV, MYR, MZN, NAD, NGN, NIO, NOK, NPR, NZD, OMR, PAB, PEN, PGK, PHP, PKR, PLN, PYG, QAR, RON, RSD, RUB, RWF, SAR, SBD, SCR, SDG, SEK, SGD, SHP, SLL, SOS, SRD, SSP, STD, SVC, SYP, SZL, THB, TJS, TMT, TND, TOP, TRY, TTD, TWD, TZS, UAH, UGX, USD, USN, UYI, UYU, UZS, VEF, VND, VUV, WST, XAF, XAG, XAU, XBA, XBB, XBC, XBD, XCD, XDR, XOF, XPD, XPF, XPT, XSU, XTS, XUA, XXX, YER, ZAR, ZMW, ZWL
type Moneda string

const (
	MonedaMXN Moneda = "MXN"
	MonedaUSD Moneda = "USD"
	MonedaXXX Moneda = "XXX"
)

// May be one of 01, 02, 03
type ObjetoImp string

const (
	// ObjetoImp01 No objeto de impuesto.
	ObjetoImp01 ObjetoImp = "01"
	// ObjetoImp02 Sí objeto de impuesto.
	ObjetoImp02 ObjetoImp = "02"
	// ObjetoImp03 Sí objeto del impuesto y no obligado al desglose.
	ObjetoImp03 ObjetoImp = "03"
)

func (o ObjetoImp) Desc() string {
	switch o {
	case ObjetoImp01:
		return "No objeto de impuesto"
	case ObjetoImp02:
		return "Sí objeto de impuesto"
	case ObjetoImp03:
		return "Sí objeto del impuesto y no obligado al desglose"
	}
	return ""
}

// May be one of AFG, ALA, ALB, DEU, AND, AGO, AIA, ATA, ATG, SAU, DZA, ARG, ARM, ABW, AUS, AUT, AZE, BHS, BGD, BRB, BHR, BEL, BLZ, BEN, BMU, BLR, MMR, BOL, BIH, BWA, BRA, BRN, BGR, BFA, BDI, BTN, CPV, KHM, CMR, CAN, QAT, BES, TCD, CHL, CHN, CYP, COL, COM, PRK, KOR, CIV, CRI, HRV, CUB, CUW, DNK, DMA, ECU, EGY, SLV, ARE, ERI, SVK, SVN, ESP, USA, EST, ETH, PHL, FIN, FJI, FRA, GAB, GMB, GEO, GHA, GIB, GRD, GRC, GRL, GLP, GUM, GTM, GUF, GGY, GIN, GNB, GNQ, GUY, HTI, HND, HKG, HUN, IND, IDN, IRQ, IRN, IRL, BVT, IMN, CXR, NFK, ISL, CYM, CCK, COK, FRO, SGS, HMD, FLK, MNP, MHL, PCN, SLB, TCA, UMI, VGB, VIR, ISR, ITA, JAM, JPN, JEY, JOR, KAZ, KEN, KGZ, KIR, KWT, LAO, LSO, LVA, LBN, LBR, LBY, LIE, LTU, LUX, MAC, MDG, MYS, MWI, MDV, MLI, MLT, MAR, MTQ, MUS, MRT, MYT, MEX, FSM, MDA, MCO, MNG, MNE, MSR, MOZ, NAM, NRU, NPL, NIC, NER, NGA, NIU, NOR, NCL, NZL, OMN, NLD, PAK, PLW, PSE, PAN, PNG, PRY, PER, PYF, POL, PRT, PRI, GBR, CAF, CZE, MKD, COG, COD, DOM, REU, RWA, ROU, RUS, ESH, WSM, ASM, BLM, KNA, SMR, MAF, SPM, VCT, SHN, LCA, STP, SEN, SRB, SYC, SLE, SGP, SXM, SYR, SOM, LKA, SWZ, ZAF, SDN, SSD, SWE, CHE, SUR, SJM, THA, TWN, TZA, TJK, IOT, ATF, TLS, TGO, TKL, TON, TTO, TUN, TKM, TUR, TUV, UKR, UGA, URY, UZB, VUT, VAT, VEN, VNM, WLF, YEM, DJI, ZMB, ZWE, ZZZ
type Pais string

const (
	PaisMEX Pais = "MEX"
	PaisUSA Pais = "USA"
	PaisCAN Pais = "CAN"
	PaisCHL Pais = "CHL"
	PaisCOL Pais = "COL"
	PaisBRA Pais = "BRA"
	PaisCHN Pais = "CHN"
)

// May be one of 601, 603, 605, 606, 607, 608, 609, 610, 611, 612, 614, 615, 616, 620, 621, 622, 623, 624, 625, 626, 628, 629, 630
type RegimenFiscal string

const (
	// RegimenFiscal601 General de Ley Personas Morales.
	RegimenFiscal601 RegimenFiscal = "601"
	// RegimenFiscal603 Personas Morales con Fines no Lucrativos.
	RegimenFiscal603 RegimenFiscal = "603"
	// RegimenFiscal605 Sueldos y Salarios e Ingresos Asimilados a Salarios.
	RegimenFiscal605 RegimenFiscal = "605"
	// RegimenFiscal606 Arrendamiento.
	RegimenFiscal606 RegimenFiscal = "606"
	// RegimenFiscal607 Régimen de Enajenación o Adquisición de Bienes.
	RegimenFiscal607 RegimenFiscal = "607"
	// RegimenFiscal608 Demás ingresos.
	RegimenFiscal608 RegimenFiscal = "608"
	// RegimenFiscal610 Residentes en el Extranjero sin Establecimiento Permanente en México.
	RegimenFiscal610 RegimenFiscal = "610"
	// RegimenFiscal611 Ingresos por Dividendos (socios y accionistas).
	RegimenFiscal611 RegimenFiscal = "611"
	// RegimenFiscal612 Personas Físicas con Actividades Empresariales y Profesionales.
	RegimenFiscal612 RegimenFiscal = "612"
	// RegimenFiscal614 Ingresos por intereses.
	RegimenFiscal614 RegimenFiscal = "614"
	// RegimenFiscal615 Régimen de los ingresos por obtención de premios.
	RegimenFiscal615 RegimenFiscal = "615"
	// RegimenFiscal616 Sin obligaciones fiscales.
	RegimenFiscal616 RegimenFiscal = "616"
	// RegimenFiscal620 Sociedades Cooperativas de Producción que optan por diferir sus ingresos.
	RegimenFiscal620 RegimenFiscal = "620"
	// RegimenFiscal621 Incorporación Fiscal.
	RegimenFiscal621 RegimenFiscal = "621"
	// RegimenFiscal622 Actividades Agrícolas, Ganaderas, Silvícolas y Pesqueras.
	RegimenFiscal622 RegimenFiscal = "622"
	// RegimenFiscal623 Opcional para Grupos de Sociedades.
	RegimenFiscal623 RegimenFiscal = "623"
	// RegimenFiscal624 Coordinados.
	RegimenFiscal624 RegimenFiscal = "624"
	// RegimenFiscal625 Régimen de las Actividades Empresariales con ingresos a través de Plataformas Tecnológicas.
	RegimenFiscal625 RegimenFiscal = "625"
	// RegimenFiscal626 Régimen Simplificado de Confianza.
	RegimenFiscal626 RegimenFiscal = "626"
)

func (r RegimenFiscal) Desc() string {
	switch r {
	case RegimenFiscal601:
		return "General de Ley Personas Morales"
	case RegimenFiscal603:
		return "Personas Morales con Fines no Lucrativos"
	case RegimenFiscal605:
		return "Sueldos y Salarios e Ingresos Asimilados a Salarios"
	case RegimenFiscal606:
		return "Arrendamiento"
	case RegimenFiscal607:
		return "Régimen de Enajenación o Adquisición de Bienes"
	case RegimenFiscal608:
		return "Demás ingresos"
	case RegimenFiscal610:
		return "Residentes en el Extranjero sin Establecimiento Permanente en México"
	case RegimenFiscal611:
		return "Ingresos por Dividendos (socios y accionistas)"
	case RegimenFiscal612:
		return "Personas Físicas con Actividades Empresariales y Profesionales"
	case RegimenFiscal614:
		return "Ingresos por intereses"
	case RegimenFiscal615:
		return "Régimen de los ingresos por obtención de premios"
	case RegimenFiscal616:
		return "Sin obligaciones fiscales"
	case RegimenFiscal620:
		return "Sociedades Cooperativas de Producción que optan por diferir sus ingresos"
	case RegimenFiscal621:
		return "Incorporación Fiscal"
	case RegimenFiscal622:
		return "Actividades Agrícolas, Ganaderas, Silvícolas y Pesqueras"
	case RegimenFiscal623:
		return "Opcional para Grupos de Sociedades"
	case RegimenFiscal624:
		return "Coordinados"
	case RegimenFiscal625:
		return "Régimen de las Actividades Empresariales con ingresos a través de Plataformas Tecnológicas"
	case RegimenFiscal626:
		return "Régimen Simplificado de Confianza"
	}
	return ""
}

// May be one of I, E, T, N, P
type TipoDeComprobante string

const (
	ComprobanteI TipoDeComprobante = "I"
	ComprobanteE TipoDeComprobante = "E"
	ComprobanteT TipoDeComprobante = "T"
	ComprobanteN TipoDeComprobante = "N"
	ComprobanteP TipoDeComprobante = "P"
)

func (t TipoDeComprobante) Desc() string {
	switch t {
	case ComprobanteI:
		return "Ingreso"
	case ComprobanteE:
		return "Egreso"
	case ComprobanteT:
		return "Traslado"
	case ComprobanteN:
		return "Nomina"
	case ComprobanteP:
		return "Pago"
	}
	return ""
}

// May be one of Tasa, Cuota, Exento
type TipoFactor string

const (
	TipoFactorTasa   TipoFactor = "Tasa"
	TipoFactorCuota  TipoFactor = "Cuota"
	TipoFactorExento TipoFactor = "Exento"
)

// May be one of 01, 02, 03, 04, 05, 06, 07, 08, 09
type TipoRelacion string

const (
	//TipoRelacion01 Nota de crédito de los documentos relacionados.
	TipoRelacion01 TipoRelacion = "01"
	//TipoRelacion02 Nota de débito de los documentos relacionados.
	TipoRelacion02 TipoRelacion = "02"
	//TipoRelacion03 Devolución de mercancía sobre facturas o traslados previos.
	TipoRelacion03 TipoRelacion = "03"
	//TipoRelacion04 Sustitución de los CFDI previos.
	TipoRelacion04 TipoRelacion = "04"
	//TipoRelacion05 Traslados de mercancias facturados previamente.
	TipoRelacion05 TipoRelacion = "05"
	//TipoRelacion06 Factura generada por los traslados previos.
	TipoRelacion06 TipoRelacion = "06"
	//TipoRelacion07 CFDI por aplicación de anticipo.
	TipoRelacion07 TipoRelacion = "07"
)

func (t TipoRelacion) Desc() string {
	switch t {
	case TipoRelacion01:
		return "Nota de crédito de los documentos relacionados"
	case TipoRelacion02:
		return "Nota de débito de los documentos relacionados"
	case TipoRelacion03:
		return "Devolución de mercancía sobre facturas o traslados previos"
	case TipoRelacion04:
		return "Sustitución de los CFDI previos"
	case TipoRelacion05:
		return "Traslados de mercancias facturados previamente"
	case TipoRelacion06:
		return "Factura generada por los traslados previos"
	case TipoRelacion07:
		return "CFDI por aplicación de anticipo"
	}
	return ""
}

// May be one of G01, G02, G03, I01, I02, I03, I04, I05, I06, I07, I08, D01, D02, D03, D04, D05, D06, D07, D08, D09, D10, P01, S01, CP01, CN01
type UsoCFDI string

const (
	// UsoCFDIG01	Adquisición de mercancias.
	UsoCFDIG01 UsoCFDI = "G01"
	// UsoCFDIG02	Devoluciones, descuentos o bonificacione
	UsoCFDIG02 UsoCFDI = "G02"
	// UsoCFDIG03	Gastos en general.
	UsoCFDIG03 UsoCFDI = "G03"
	// UsoCFDII01	Construcciones.
	UsoCFDII01 UsoCFDI = "I01"
	// UsoCFDII02	Mobilario y equipo de oficina por invers
	UsoCFDII02 UsoCFDI = "I02"
	// UsoCFDII03	Equipo de transporte.
	UsoCFDII03 UsoCFDI = "I03"
	// UsoCFDII04	Equipo de computo y accesorios.
	UsoCFDII04 UsoCFDI = "I04"
	// UsoCFDII05	Dados, troqueles, moldes, matrices y her
	UsoCFDII05 UsoCFDI = "I05"
	// UsoCFDII06	Comunicaciones telefónicas.
	UsoCFDII06 UsoCFDI = "I06"
	// UsoCFDII07	Comunicaciones satelitales.
	UsoCFDII07 UsoCFDI = "I07"
	// UsoCFDII08	Otra maquinaria y equipo.
	UsoCFDII08 UsoCFDI = "I08"
	// UsoCFDID01	Honorarios médicos, dentales y gastos ho
	UsoCFDID01 UsoCFDI = "D01"
	// UsoCFDID02	Gastos médicos por incapacidad o discapa
	UsoCFDID02 UsoCFDI = "D02"
	// UsoCFDID03	Gastos funerales.
	UsoCFDID03 UsoCFDI = "D03"
	// UsoCFDID04	Donativos.
	UsoCFDID04 UsoCFDI = "D04"
	// UsoCFDID05	Intereses reales efectivamente pagados p
	UsoCFDID05 UsoCFDI = "D05"
	// UsoCFDID06	Aportaciones voluntarias al SAR.
	UsoCFDID06 UsoCFDI = "D06"
	// UsoCFDID07	Primas por seguros de gastos médicos.
	UsoCFDID07 UsoCFDI = "D07"
	// UsoCFDID08	Gastos de transportación escolar obligat
	UsoCFDID08 UsoCFDI = "D08"
	// UsoCFDID09	Depósitos en cuentas para el ahorro, pri
	UsoCFDID09 UsoCFDI = "D09"
	// UsoCFDID10	Pagos por servicios educativos (colegiat
	UsoCFDID10 UsoCFDI = "D10"
	// UsoCFDIS01	Sin efectos fiscales.
	UsoCFDIS01 UsoCFDI = "S01"
)

// May be one 01, 02, 03, 04, 05
type Periodicidad string

const (
	PeriodicidadDiario    Periodicidad = "01"
	PeriodicidadSemanal   Periodicidad = "02"
	PeriodicidadQuincenal Periodicidad = "03"
	PeriodicidadMensual   Periodicidad = "04"
	PeriodicidadBimestral Periodicidad = "05"
)

func (p Periodicidad) Desc() string {
	switch p {
	case PeriodicidadDiario:
		return "Diario"
	case PeriodicidadSemanal:
		return "Semanal"
	case PeriodicidadQuincenal:
		return "Quincenal"
	case PeriodicidadMensual:
		return "Mensual"
	case PeriodicidadBimestral:
		return "Bimestral"
	}
	return ""
}

// Banco es un elemento del Catálogo de bancos.
//
// Catálogo de bancos se utiliza cuando los contribuyentes realicen operaciones con diferentes bancos
// nacionales, para el registro contable de cada póliza debe apoyarse en el catálogo de bancos.
type Banco string

// CodAgrup es un elemento del código agrupador de cuentas del sat.
//
// El código agrupador del SAT tiene el objetivo de que la información sea presentada de manera
// uniforme, para lo cual es necesario que los contribuyentes asocien las cuentas de su catálogo de
// cuentas al código agrupador por naturaleza y preponderancia de la cuenta.
type CodAgrup string

// CodAgrupH es un elemento del catálogo de códigos agrupadores hidrocarburos.
type CodAgrupH string
