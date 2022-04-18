ECHO #######################################################################
ECHO Intall dependencies first go get aqwari.net/xml/...

ECHO #######################################################################
md "./gen/contabilidad13/types" 2>NUL

xsdgen -pkg types -o ./gen/contabilidad13/types/_CatalogosParaEsqContE.go ^
./contabilidad13/CatalogosParaEsqContE.xsd

ECHO #######################################################################
md "./gen/contabilidad13/catcuentas" 2>NUL

xsdgen -pkg catcuentas -o ./gen/contabilidad13/catcuentas/_CatalogoCuentas_1_3.go ^
./contabilidad13/CatalogosParaEsqContE.xsd ^
./contabilidad13/CatalogoCuentas_1_3.xsd


ECHO #######################################################################
md "./gen/contabilidad13/balanza" 2>NUL

xsdgen -pkg balanza -o ./gen/contabilidad13/balanza/_BalanzaComprobacion_1_3.go ^
./contabilidad13/BalanzaComprobacion_1_3.xsd

ECHO #######################################################################
md "./gen/contabilidad13/polizasperiodo" 2>NUL

xsdgen -pkg catcuentas -o ./gen/contabilidad13/polizasperiodo/_polizasperiodo_1_3.go
./contabilidad13/CatalogosParaEsqContE.xsd
./contabilidad13/polizasperiodo_1_3.xsd