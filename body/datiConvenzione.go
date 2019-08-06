package body

import "fmt"

//DatiConvenzione (dati relativi alla convenzione collegata alla fattura)
type datiConvenzione struct {
	//	RiferimentoNumeroLinea: numero della linea o delle linee di
	//dettaglio della fattura alle quali si riferisce la convenzione così
	//come identificata dai tre elementi successivi (IdDocumento,
	//Data, NumItem); nel caso in cui la convenzione si riferisce
	//all’intera fattura, questo elemento non deve essere
	//valorizzato.
	RiferimentoNumeroLinea numeroLinea `xml:"RiferimentoNumeroLinea" json:"RiferimentoNumeroLinea"`

	//IdDocumento: numero della convenzione associata alla
	//fattura o alla linea/linee di fattura indicate nell’elemento
	//RiferimentoNumeroLinea.
	IDDocumento idDocumento `xml:"IdDocumento" json:"IdDocumento"`

	//Data: data della convenzione associata alla fattura o alla
	//linea/linee di fattura indicate  nell’elemento RiferimentoNumeroLinea.
	Data data `xml:"Data" json:"Data"`

	//NumItem: identificativo della singola voce (linea di
	//convenzione) all'interno della convenzione associata alla
	//fattura o alla linea/linee di fattura indicate nell’elemento RiferimentoNumeroLinea.
	NumItem numItem `xml:"NumItem" json:"NumItem"`

	//CodiceCommessaConvenzione: codice della commessa o
	//della convenzione collegata alla fattura.
	CodiceCommessaConvenzione codiceCommessaConvenzione `xml:"CodiceCommessaConvenzione" json:"CodiceCommessaConvenzione"`

	//CodiceCUP: codice gestito dal CIPE che caratterizza ogni
	//progetto di investimento pubblico (Codice Unitario Progetto).
	CodiceCUP codiceCup `xml:"CodiceCUP" json:"CodiceCUP"`

	//CodiceCIG: Codice Identificativo della Gara.
	CodiceCIG codiceCig `xml:"CodiceCIG" json:"CodiceCIG"`
}

func (c datiConvenzione) Validate() error {
	var err error

	if err = c.Data.Validate(); err != nil {
		return fmt.Errorf("Data %s", err)
	}

	if err = c.CodiceCIG.Validate(); err != nil {
		return fmt.Errorf("CodiceCIG %s", err)
	}

	if err = c.CodiceCUP.Validate(); err != nil {
		return fmt.Errorf("CodiceCUP %s", err)
	}
	if err = c.RiferimentoNumeroLinea.Validate(); err != nil {
		return fmt.Errorf("RiferimentoNumeroLinea %s", err)
	}

	if err = c.NumItem.Validate(); err != nil {
		return fmt.Errorf("NumItem %s", err)
	}

	if err = c.CodiceCommessaConvenzione.Validate(); err != nil {
		return fmt.Errorf("CodiceCommessaConvenzione %s", err)
	}

	if err = c.IDDocumento.Validate(); err != nil {
		return fmt.Errorf("IDDocumento %s", err)
	}

	return err
}
