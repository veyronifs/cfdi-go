ECHO #######################################################################
ECHO Intall dependencies first go get aqwari.net/xml/...

ECHO #######################################################################
md "./gen/cfdi40/td" 2>NUL

xsdgen -pkg cfdi40 -o ./gen/cfdi40/td/_tdCFDI.go ^
./cfdi/tdCFDI.xsd

ECHO #######################################################################
md "./gen/cfdi40/cat" 2>NUL

xsdgen -pkg cfdi40 -o ./gen/cfdi40/cat/_catCFDI.go ^
./cfdi/catCFDI.xsd

ECHO #######################################################################
md "./gen/cfdi40/cfdi40" 2>NUL

xsdgen -pkg cfdi40 -o ./gen/cfdi40/cfdi40/_cfdi40.go ^
./cfdi/catCFDI.xsd ^
./cfdi/tdCFDI.xsd ^
./cfdi/cfv40.xsd

ECHO #######################################################################
md "./gen/cfdi40/pagos20" 2>NUL

xsdgen -pkg pagos20 -o ./gen/cfdi40/pagos20/_pagos20.go ^
./cfdi/catCFDI.xsd ^
./cfdi/tdCFDI.xsd ^
./cfdi/catPagos.xsd ^
./cfdi/Pagos20.xsd

ECHO #######################################################################
md "./gen/cfdi40/comext11" 2>NUL

xsdgen -pkg comext11 -o ./gen/cfdi40/comext11/_comext11.go ^
./cfdi/catCFDI.xsd ^
./cfdi/tdCFDI.xsd ^
./cfdi/catComExt.xsd ^
./cfdi/ComercioExterior11.xsd