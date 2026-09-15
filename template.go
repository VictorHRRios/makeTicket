package main

import (
	"fmt"
	"os"
)

func (c Config) fillTemplatePM(file *os.File) error {
	PMTemplate := fmt.Sprintf(`%s**PM**
PROGRAMAS:
- NOMBRE DE PROGRAMA | VERSION

DEPENDENCIAS:
- NOMBRE DE DEPENDECNIAS | VERSION | WP DE CUANDO FUE ENTREGADO

MODIFICACIONES:
Se implementa ...

REPOSITORIOS:
%s
%s

ENTREGABLE:
%s/%s.zip

QUERIES:
<pre> QUERIES </pre>

PETICIONES:
%s.xml

NOTAS:
notas aqui

CONFIGURACIÓN/INSTRUCCIONES:
1. Pegar los archivos en la estructura de Inssist correspondiente.

EVIDENCIA:
evSVN.png <- Commit al SVN del proyecto.
evEntregable.png <- Commit al SVN de entregables.
%s%%2F%s%%2F%s <- Vídeo con prueba unitaria
`, c.TicketNumber,
		c.CurrentProject.Name,
		c.CurrentVersion.SvnURL,
		c.CurrentVersion.EvidenceURL,
		c.TicketNumber,
		c.TicketNumber,
		c.CloudURL,
		c.DateStr,
		c.TicketNumber)
	_, err := file.WriteString(PMTemplate)
	if err != nil {
		return err
	}
	return nil
}
