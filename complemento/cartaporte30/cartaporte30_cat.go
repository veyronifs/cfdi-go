package cartaporte30

// Expresa el tipo de documento aduanero que se encuentra asociado al traslado de los bienes y/o mercancías de procedencia extranjera durante su traslado en territorio nacional.
type DocumentoAduanero string

const (
	DocumentoAduanero01 DocumentoAduanero = "01"
	DocumentoAduanero02 DocumentoAduanero = "02"
	DocumentoAduanero03 DocumentoAduanero = "03"
	DocumentoAduanero04 DocumentoAduanero = "04"
	DocumentoAduanero05 DocumentoAduanero = "05"
	DocumentoAduanero06 DocumentoAduanero = "06"
	DocumentoAduanero07 DocumentoAduanero = "07"
	DocumentoAduanero08 DocumentoAduanero = "08"
	DocumentoAduanero09 DocumentoAduanero = "09"
	DocumentoAduanero10 DocumentoAduanero = "10"
	DocumentoAduanero11 DocumentoAduanero = "11"
	DocumentoAduanero12 DocumentoAduanero = "12"
	DocumentoAduanero13 DocumentoAduanero = "13"
	DocumentoAduanero14 DocumentoAduanero = "14"
	DocumentoAduanero15 DocumentoAduanero = "15"
	DocumentoAduanero16 DocumentoAduanero = "16"
	DocumentoAduanero17 DocumentoAduanero = "17"
	DocumentoAduanero18 DocumentoAduanero = "18"
	DocumentoAduanero19 DocumentoAduanero = "19"
	DocumentoAduanero20 DocumentoAduanero = "20"
)

// Desc regresa la descripción del documento aduanero.
func (d DocumentoAduanero) Desc() string {
	switch d {
	case DocumentoAduanero01:
		return "Pedimento"
	case DocumentoAduanero02:
		return "Autorización de importación temporal"
	case DocumentoAduanero03:
		return "Autorización de importación temporal de embarcaciones"
	case DocumentoAduanero04:
		return "Autorización de importación temporal de mercancías, destinadas al mantenimiento y reparación de las mercancías importadas temporalmente"
	case DocumentoAduanero05:
		return "Autorización para la importación de vehículos especialmente construidos o transformados, equipados con dispositivos o aparatos diversos para cumplir con contrato derivado de licitación pública"
	case DocumentoAduanero06:
		return "Aviso de exportación temporal"
	case DocumentoAduanero07:
		return "Aviso de traslado de mercancías de empresas con Programa IMMEX, RFE u Operador Económico Autorizado"
	case DocumentoAduanero08:
		return "Aviso para el traslado de autopartes ubicadas en la franja o región fronteriza a la industria terminal automotriz o manufacturera de vehículos de autotransporte en el resto del territorio nacional"
	case DocumentoAduanero09:
		return "Constancia de importación temporal, retorno o transferencia de contenedores"
	case DocumentoAduanero10:
		return "Constancia de transferencia de mercancías"
	case DocumentoAduanero11:
		return "Autorización de donación de mercancías al Fisco Federal que se encuentren en el extranjero"
	case DocumentoAduanero12:
		return "Cuaderno ATA"
	case DocumentoAduanero13:
		return "Listas de intercambio"
	case DocumentoAduanero14:
		return "Permiso de Importación Temporal"
	case DocumentoAduanero15:
		return "Permiso de importación temporal de casa rodante"
	case DocumentoAduanero16:
		return "Permiso de importación temporal de embarcaciones"
	case DocumentoAduanero17:
		return "Solicitud de donación de mercancías en casos de emergencias o desastres naturales"
	case DocumentoAduanero18:
		return "Aviso de consolidado"
	case DocumentoAduanero19:
		return "Aviso de cruce de mercancias"
	case DocumentoAduanero20:
		return "Otro"
	}
	return ""
}

// SectorCOFEPRIS Expresar la clasificación del producto que se traslada a través de los distintos medios de transporte y que debe contar con autorización por la autoridad correspondiente.
type SectorCOFEPRIS string

const (
	SectorCOFEPRIS01 SectorCOFEPRIS = "01"
	SectorCOFEPRIS02 SectorCOFEPRIS = "02"
	SectorCOFEPRIS03 SectorCOFEPRIS = "03"
	SectorCOFEPRIS04 SectorCOFEPRIS = "04"
	SectorCOFEPRIS05 SectorCOFEPRIS = "05"
)

// Desc regresa la descripción del sector COFEPRIS.
func (s SectorCOFEPRIS) Desc() string {
	switch s {
	case SectorCOFEPRIS01:
		return "Medicamento"
	case SectorCOFEPRIS02:
		return "Precursores y químicos de uso dual"
	case SectorCOFEPRIS03:
		return "Psicotrópicos y estupefacientes"
	case SectorCOFEPRIS04:
		return "Sustancias tóxicas"
	case SectorCOFEPRIS05:
		return "Plaguicidas y fertilizantes"
	}
	return ""
}

// FormaFarmaceutica Expresar la forma farmacéutica o mezcla del medicamento, precursor, químico de uso dual, psicotrópico o estupefaciente que presenta ciertas características físicas para su adecuada dosificación, conservación y administración.
type FormaFarmaceutica string

const (
	FormaFarmaceutica01 FormaFarmaceutica = "01"
	FormaFarmaceutica02 FormaFarmaceutica = "02"
	FormaFarmaceutica03 FormaFarmaceutica = "03"
	FormaFarmaceutica04 FormaFarmaceutica = "04"
	FormaFarmaceutica05 FormaFarmaceutica = "05"
	FormaFarmaceutica06 FormaFarmaceutica = "06"
	FormaFarmaceutica07 FormaFarmaceutica = "07"
	FormaFarmaceutica08 FormaFarmaceutica = "08"
	FormaFarmaceutica09 FormaFarmaceutica = "09"
	FormaFarmaceutica10 FormaFarmaceutica = "10"
	FormaFarmaceutica11 FormaFarmaceutica = "11"
	FormaFarmaceutica12 FormaFarmaceutica = "12"
	FormaFarmaceutica13 FormaFarmaceutica = "13"
	FormaFarmaceutica14 FormaFarmaceutica = "14"
	FormaFarmaceutica15 FormaFarmaceutica = "15"
	FormaFarmaceutica16 FormaFarmaceutica = "16"
	FormaFarmaceutica17 FormaFarmaceutica = "17"
	FormaFarmaceutica18 FormaFarmaceutica = "18"
	FormaFarmaceutica19 FormaFarmaceutica = "19"
	FormaFarmaceutica20 FormaFarmaceutica = "20"
)

// Desc regresa la descripción de la forma farmacéutica.
func (f FormaFarmaceutica) Desc() string {
	switch f {
	case FormaFarmaceutica01:
		return "Tableta"
	case FormaFarmaceutica02:
		return "Capsulas"
	case FormaFarmaceutica03:
		return "Comprimidos"
	case FormaFarmaceutica04:
		return "Grageas"
	case FormaFarmaceutica05:
		return "Suspensión"
	case FormaFarmaceutica06:
		return "Solución"
	case FormaFarmaceutica07:
		return "Emulsión"
	case FormaFarmaceutica08:
		return "Jarabe"
	case FormaFarmaceutica09:
		return "Inyectable"
	case FormaFarmaceutica10:
		return "Crema"
	case FormaFarmaceutica11:
		return "Ungüento"
	case FormaFarmaceutica12:
		return "Aerosol"
	case FormaFarmaceutica13:
		return "Gas medicinal"
	case FormaFarmaceutica14:
		return "Gel"
	case FormaFarmaceutica15:
		return "Implante"
	case FormaFarmaceutica16:
		return "Óvulo"
	case FormaFarmaceutica17:
		return "Parche"
	case FormaFarmaceutica18:
		return "Pasta"
	case FormaFarmaceutica19:
		return "Polvo"
	case FormaFarmaceutica20:
		return "Supositorio"
	}
	return ""
}

// CondicionesEspTransp Expresa la condición en la cual es necesario mantener el medicamento, precursor, químico de uso dual, psicotrópicos o estupefacientes durante el traslado y almacenamiento.
type CondicionesEspTransp string

const (
	CondicionesEspTransp01 CondicionesEspTransp = "01"
	CondicionesEspTransp02 CondicionesEspTransp = "02"
	CondicionesEspTransp03 CondicionesEspTransp = "03"
	CondicionesEspTransp04 CondicionesEspTransp = "04"
)

// Desc regresa la descripción de las condiciones especiales de transporte.
func (c CondicionesEspTransp) Desc() string {
	switch c {
	case CondicionesEspTransp01:
		return "Congelados"
	case CondicionesEspTransp02:
		return "Refrigerados"
	case CondicionesEspTransp03:
		return "Temperatura controlada"
	case CondicionesEspTransp04:
		return "Temperatura ambiente"
	}
	return ""
}

// RegimenAduanero represents the type of movement for the document.
type RegimenAduanero string

const (
	RegimenAduaneroIMD RegimenAduanero = "IMD"
	RegimenAduaneroEXD RegimenAduanero = "EXD"
	RegimenAduaneroITR RegimenAduanero = "ITR"
	RegimenAduaneroITE RegimenAduanero = "ITE"
	RegimenAduaneroETR RegimenAduanero = "ETR"
	RegimenAduaneroETE RegimenAduanero = "ETE"
	RegimenAduaneroDFI RegimenAduanero = "DFI"
	RegimenAduaneroRFE RegimenAduanero = "RFE"
	RegimenAduaneroRFS RegimenAduanero = "RFS"
	RegimenAduaneroTRA RegimenAduanero = "TRA"
)

// Desc returns the description of the movement type.
func (t RegimenAduanero) Desc() string {
	switch t {
	case RegimenAduaneroIMD:
		return "Definitivo de importación"
	case RegimenAduaneroEXD:
		return "Definitivo de exportación"
	case RegimenAduaneroITR:
		return "Temporales de importación para retomar al extranjero en el mismo estado"
	case RegimenAduaneroITE:
		return "Temporales de importación para elaboración, transformación o reparación para empresas con programa IMMEX"
	case RegimenAduaneroETR:
		return "Temporales de exportación para retornar al país en el mismo estado"
	case RegimenAduaneroETE:
		return "Temporales de exportación para elaboración, transformación o reparación"
	case RegimenAduaneroDFI:
		return "Depósito Fiscal"
	case RegimenAduaneroRFE:
		return "Elaboración, transformación o reparación en recinto fiscalizado"
	case RegimenAduaneroRFS:
		return "Recinto fiscalizado estratégico"
	case RegimenAduaneroTRA:
		return "Tránsitos"
	}
	return ""
}
