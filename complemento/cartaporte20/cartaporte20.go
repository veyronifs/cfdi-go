package cartaporte20

import (
	"encoding/xml"

	"github.com/shopspring/decimal"
	"github.com/veyronifs/cfdi-go/types"
)

type CartaPorte20 struct {
	Ubicaciones       Ubicaciones       `xml:"Ubicaciones"`                //
	Mercancias        *Mercancias       `xml:"Mercancias"`                 //
	FiguraTransporte  *FiguraTransporte `xml:"FiguraTransporte,omitempty"` //
	Version           string            `xml:"Version,attr"`
	TranspInternac    string            `xml:"TranspInternac,attr"`
	EntradaSalidaMerc string            `xml:"EntradaSalidaMerc,attr,omitempty"`
	PaisOrigenDestino types.Pais        `xml:"PaisOrigenDestino,attr,omitempty"`
	ViaEntradaSalida  string            `xml:"ViaEntradaSalida,attr,omitempty"`
	TotalDistRec      decimal.Decimal   `xml:"TotalDistRec,attr,omitempty"`
}

type Ubicaciones []*Ubicacion

func (u *Ubicaciones) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var ubics struct {
		Slice []*Ubicacion `xml:"Ubicacion"` //
	}

	if err := d.DecodeElement(&ubics, &start); err != nil {
		return err
	}
	*u = ubics.Slice
	return nil
}

type Ubicacion struct {
	Domicilio                   *Domicilio      `xml:"Domicilio,omitempty"` //
	TipoUbicacion               string          `xml:"TipoUbicacion,attr"`
	IDUbicacion                 string          `xml:"IDUbicacion,attr,omitempty"`
	RFCRemitenteDestinatario    string          `xml:"RFCRemitenteDestinatario,attr"`
	NombreRemitenteDestinatario string          `xml:"NombreRemitenteDestinatario,attr,omitempty"`
	NumRegIdTrib                string          `xml:"NumRegIdTrib,attr,omitempty"`
	ResidenciaFiscal            types.Pais      `xml:"ResidenciaFiscal,attr,omitempty"`
	NumEstacion                 string          `xml:"NumEstacion,attr,omitempty"`
	NombreEstacion              string          `xml:"NombreEstacion,attr,omitempty"`
	NavegacionTrafico           string          `xml:"NavegacionTrafico,attr,omitempty"`
	FechaHoraSalidaLlegada      types.TFechaH   `xml:"FechaHoraSalidaLlegada,attr"`
	TipoEstacion                string          `xml:"TipoEstacion,attr,omitempty"`
	DistanciaRecorrida          decimal.Decimal `xml:"DistanciaRecorrida,attr,omitempty"`
}

type Mercancias struct {
	Mercancia             []*Mercancia           `xml:"Mercancia"`                       //
	Autotransporte        *Autotransporte        `xml:"Autotransporte,omitempty"`        //
	TransporteMaritimo    *TransporteMaritimo    `xml:"TransporteMaritimo,omitempty"`    //
	TransporteAereo       *TransporteAereo       `xml:"TransporteAereo,omitempty"`       //
	TransporteFerroviario *TransporteFerroviario `xml:"TransporteFerroviario,omitempty"` //
	PesoBrutoTotal        decimal.Decimal        `xml:"PesoBrutoTotal,attr"`
	UnidadPeso            string                 `xml:"UnidadPeso,attr"`
	PesoNetoTotal         decimal.Decimal        `xml:"PesoNetoTotal,attr,omitempty"`
	NumTotalMercancias    int                    `xml:"NumTotalMercancias,attr"`
	CargoPorTasacion      decimal.Decimal        `xml:"CargoPorTasacion,attr,omitempty"`
}

type Mercancia struct {
	Pedimentos           []*Pedimentos          `xml:"Pedimentos,omitempty"`          //
	GuiasIdentificacion  []*GuiasIdentificacion `xml:"GuiasIdentificacion,omitempty"` //
	CantidadTransporta   []*CantidadTransporta  `xml:"CantidadTransporta,omitempty"`  //
	DetalleMercancia     *DetalleMercancia      `xml:"DetalleMercancia,omitempty"`    //
	BienesTransp         string                 `xml:"BienesTransp,attr"`
	ClaveSTCC            string                 `xml:"ClaveSTCC,attr,omitempty"`
	Descripcion          string                 `xml:"Descripcion,attr"`
	Cantidad             decimal.Decimal        `xml:"Cantidad,attr"`
	ClaveUnidad          string                 `xml:"ClaveUnidad,attr"`
	Unidad               string                 `xml:"Unidad,attr,omitempty"`
	Dimensiones          string                 `xml:"Dimensiones,attr,omitempty"`
	MaterialPeligroso    string                 `xml:"MaterialPeligroso,attr,omitempty"`
	CveMaterialPeligroso string                 `xml:"CveMaterialPeligroso,attr,omitempty"`
	Embalaje             string                 `xml:"Embalaje,attr,omitempty"`
	DescripEmbalaje      string                 `xml:"DescripEmbalaje,attr,omitempty"`
	PesoEnKg             decimal.Decimal        `xml:"PesoEnKg,attr"`
	ValorMercancia       decimal.Decimal        `xml:"ValorMercancia,attr,omitempty"`
	Moneda               types.Moneda           `xml:"Moneda,attr,omitempty"`
	FraccionArancelaria  string                 `xml:"FraccionArancelaria,attr,omitempty"`
	UUIDComercioExt      string                 `xml:"UUIDComercioExt,attr,omitempty"`
}

type Autotransporte struct {
	IdentificacionVehicular *IdentificacionVehicular `xml:"IdentificacionVehicular"` //
	Seguros                 *Seguros                 `xml:"Seguros"`                 //
	Remolques               Remolques                `xml:"Remolques,omitempty"`     //
	PermSCT                 string                   `xml:"PermSCT,attr"`
	NumPermisoSCT           string                   `xml:"NumPermisoSCT,attr"`
}

type ContenedorMaritimo struct {
	MatriculaContenedor string `xml:"MatriculaContenedor,attr"`
	TipoContenedor      string `xml:"TipoContenedor,attr"`
	NumPrecinto         string `xml:"NumPrecinto,attr,omitempty"`
}

type DerechosDePaso struct {
	TipoDerechoDePaso string          `xml:"TipoDerechoDePaso,attr"`
	KilometrajePagado decimal.Decimal `xml:"KilometrajePagado,attr"`
}

type DetalleMercancia struct {
	UnidadPesoMerc string          `xml:"UnidadPesoMerc,attr"`
	PesoBruto      decimal.Decimal `xml:"PesoBruto,attr"`
	PesoNeto       decimal.Decimal `xml:"PesoNeto,attr"`
	PesoTara       decimal.Decimal `xml:"PesoTara,attr"`
	NumPiezas      int             `xml:"NumPiezas,attr,omitempty"`
}

type Domicilio struct {
	Calle          string     `xml:"Calle,attr,omitempty"`
	NumeroExterior string     `xml:"NumeroExterior,attr,omitempty"`
	NumeroInterior string     `xml:"NumeroInterior,attr,omitempty"`
	Colonia        string     `xml:"Colonia,attr,omitempty"`
	Localidad      string     `xml:"Localidad,attr,omitempty"`
	Referencia     string     `xml:"Referencia,attr,omitempty"`
	Municipio      string     `xml:"Municipio,attr,omitempty"`
	Estado         string     `xml:"Estado,attr"`
	Pais           types.Pais `xml:"Pais,attr"`
	CodigoPostal   string     `xml:"CodigoPostal,attr"`
}

type FiguraTransporte struct {
	TiposFigura []*TiposFigura `xml:"TiposFigura"` //
}

type GuiasIdentificacion struct {
	NumeroGuiaIdentificacion  string          `xml:"NumeroGuiaIdentificacion,attr"`
	DescripGuiaIdentificacion string          `xml:"DescripGuiaIdentificacion,attr"`
	PesoGuiaIdentificacion    decimal.Decimal `xml:"PesoGuiaIdentificacion,attr"`
}

type IdentificacionVehicular struct {
	ConfigVehicular string `xml:"ConfigVehicular,attr"`
	PlacaVM         string `xml:"PlacaVM,attr"`
	AnioModeloVM    string `xml:"AnioModeloVM,attr"`
}

type PartesTransporte struct {
	ParteTransporte string `xml:"ParteTransporte,attr"`
}

type Pedimentos struct {
	Pedimento string `xml:"Pedimento,attr"`
}

type Remolque struct {
	SubTipoRem string `xml:"SubTipoRem,attr"`
	Placa      string `xml:"Placa,attr"`
}

type Remolques []*Remolque

func (rem *Remolques) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var rem2 struct {
		Remolque []*Remolque `xml:"Remolque"` //
	}
	if err := d.DecodeElement(&rem2, &start); err != nil {
		return err
	}
	*rem = rem2.Remolque
	return nil
}

type Seguros struct {
	AseguraRespCivil   string          `xml:"AseguraRespCivil,attr"`
	PolizaRespCivil    string          `xml:"PolizaRespCivil,attr"`
	AseguraMedAmbiente string          `xml:"AseguraMedAmbiente,attr,omitempty"`
	PolizaMedAmbiente  string          `xml:"PolizaMedAmbiente,attr,omitempty"`
	AseguraCarga       string          `xml:"AseguraCarga,attr,omitempty"`
	PolizaCarga        string          `xml:"PolizaCarga,attr,omitempty"`
	PrimaSeguro        decimal.Decimal `xml:"PrimaSeguro,attr,omitempty"`
}

type TiposFigura struct {
	PartesTransporte       []*PartesTransporte `xml:"PartesTransporte,omitempty"` //
	Domicilio              *Domicilio          `xml:"Domicilio,omitempty"`        //
	TipoFigura             string              `xml:"TipoFigura,attr"`
	RFCFigura              string              `xml:"RFCFigura,attr,omitempty"`
	NumLicencia            string              `xml:"NumLicencia,attr,omitempty"`
	NombreFigura           string              `xml:"NombreFigura,attr,omitempty"`
	NumRegIdTribFigura     string              `xml:"NumRegIdTribFigura,attr,omitempty"`
	ResidenciaFiscalFigura types.Pais          `xml:"ResidenciaFiscalFigura,attr,omitempty"`
}

type TransporteAereo struct {
	PermSCT                string     `xml:"PermSCT,attr"`
	NumPermisoSCT          string     `xml:"NumPermisoSCT,attr"`
	MatriculaAeronave      string     `xml:"MatriculaAeronave,attr,omitempty"`
	NombreAseg             string     `xml:"NombreAseg,attr,omitempty"`
	NumPolizaSeguro        string     `xml:"NumPolizaSeguro,attr,omitempty"`
	NumeroGuia             string     `xml:"NumeroGuia,attr"`
	LugarContrato          string     `xml:"LugarContrato,attr,omitempty"`
	CodigoTransportista    string     `xml:"CodigoTransportista,attr"`
	RFCEmbarcador          string     `xml:"RFCEmbarcador,attr,omitempty"`
	NumRegIdTribEmbarc     string     `xml:"NumRegIdTribEmbarc,attr,omitempty"`
	ResidenciaFiscalEmbarc types.Pais `xml:"ResidenciaFiscalEmbarc,attr,omitempty"`
	NombreEmbarcador       string     `xml:"NombreEmbarcador,attr,omitempty"`
}

type TransporteFerroviario struct {
	DerechosDePaso  []DerechosDePaso             `xml:"DerechosDePaso,omitempty"` //
	Carro           []TransporteFerroviarioCarro `xml:"Carro"`                    //
	TipoDeServicio  string                       `xml:"TipoDeServicio,attr"`
	TipoDeTrafico   string                       `xml:"TipoDeTrafico,attr"`
	NombreAseg      string                       `xml:"NombreAseg,attr,omitempty"`
	NumPolizaSeguro string                       `xml:"NumPolizaSeguro,attr,omitempty"`
}

type TransporteFerroviarioCarro struct {
	TipoCarro           string          `xml:"TipoCarro,attr"`
	MatriculaCarro      string          `xml:"MatriculaCarro,attr"`
	GuiaCarro           string          `xml:"GuiaCarro,attr"`
	ToneladasNetasCarro decimal.Decimal `xml:"ToneladasNetasCarro,attr"`
}

type TransporteFerroviarioCarroContenedor struct {
	TipoContenedor      string          `xml:"TipoContenedor,attr"`
	PesoContenedorVacio decimal.Decimal `xml:"PesoContenedorVacio,attr"`
	PesoNetoMercancia   decimal.Decimal `xml:"PesoNetoMercancia,attr"`
}

type TransporteMaritimo struct {
	Contenedor             []ContenedorMaritimo `xml:"Contenedor"` //
	PermSCT                string               `xml:"PermSCT,attr,omitempty"`
	NumPermisoSCT          string               `xml:"NumPermisoSCT,attr,omitempty"`
	NombreAseg             string               `xml:"NombreAseg,attr,omitempty"`
	NumPolizaSeguro        string               `xml:"NumPolizaSeguro,attr,omitempty"`
	TipoEmbarcacion        string               `xml:"TipoEmbarcacion,attr"`
	Matricula              string               `xml:"Matricula,attr"`
	NumeroOMI              string               `xml:"NumeroOMI,attr"`
	AnioEmbarcacion        int                  `xml:"AnioEmbarcacion,attr,omitempty"`
	NombreEmbarc           string               `xml:"NombreEmbarc,attr,omitempty"`
	NacionalidadEmbarc     types.Pais           `xml:"NacionalidadEmbarc,attr"`
	UnidadesDeArqBruto     decimal.Decimal      `xml:"UnidadesDeArqBruto,attr"`
	TipoCarga              string               `xml:"TipoCarga,attr"`
	NumCertITC             string               `xml:"NumCertITC,attr"`
	Eslora                 decimal.Decimal      `xml:"Eslora,attr,omitempty"`
	Manga                  decimal.Decimal      `xml:"Manga,attr,omitempty"`
	Calado                 decimal.Decimal      `xml:"Calado,attr,omitempty"`
	LineaNaviera           string               `xml:"LineaNaviera,attr,omitempty"`
	NombreAgenteNaviero    string               `xml:"NombreAgenteNaviero,attr"`
	NumAutorizacionNaviero string               `xml:"NumAutorizacionNaviero,attr"`
	NumViaje               string               `xml:"NumViaje,attr,omitempty"`
	NumConocEmbarc         string               `xml:"NumConocEmbarc,attr,omitempty"`
}

type CantidadTransporta struct {
	Cantidad       decimal.Decimal `xml:"Cantidad,attr"`
	IDOrigen       string          `xml:"IDOrigen,attr"`
	IDDestino      string          `xml:"IDDestino,attr"`
	CvesTransporte string          `xml:"CvesTransporte,attr,omitempty"`
}
