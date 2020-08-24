package einvoice

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

func Validator() *validator.Validate {
	var validate *validator.Validate
	validate = validator.New()
	validate.RegisterValidation("regimeValidate", regimeFiscaleValidator)
	validate.RegisterValidation("isInteger", isInteger)
	validate.RegisterValidation("isDate", isDate)
	validate.RegisterValidation("isPrice", isPrice)
	validate.RegisterValidation("isTypeDocument", isTypeDocument)
	validate.RegisterValidation("isntSDIPec", isntSDIPec)
	validate.RegisterValidation("isNatura", isNatura)
	validate.RegisterValidation("isDateTime", isDateTime)
	validate.RegisterValidation("isMP", isMP)
	validate.RegisterStructValidation(datiTrasmissioneValidate, datiTrasmissione{})
	validate.RegisterStructValidation(cessionarioCommittenteValidate, CessionarioCommittente{})
	return validate
}

func regimeFiscaleValidator(rf validator.FieldLevel) bool {
	RF := rf.Field().String()

	regimeFiscale := make(map[string]string)
	for i := 1; i < 20; i++ {
		regimeFiscale[fmt.Sprintf("RF%02d", i)] = "true"
	}
	_, exists := regimeFiscale[string(RF)]
	//println(exists, RF)
	return exists == true
}

// control id Field is Integer
func isInteger(t validator.FieldLevel) bool {
	_, err := strconv.Atoi(t.Field().String())
	if err != nil {
		return false
	}
	return true
}

//Check the data hold block, which must be filled in case of DatiCassaPrevidenziale.Ritenuta == "SI"
func checkDatiRitenuta(d validator.StructLevel) {
	data := d.Current().Interface().(datiGeneraliDocumento)
	if data.DatiCassaPrevidenziale.Ritenuta == "SI" {
		if *data.DatiRitenuta == (datiRitenuta{}) {
			d.ReportError(data.DatiRitenuta, "struct", "all", "required", "")
		} else {
			validate := Validator()
			if err := validate.Struct(data.DatiRitenuta); err != nil {
				d.ReportError("DatiRitenuta", "", fmt.Sprintf("%s", err), "", "")
			}

		}
	}

}

// Struct Level control
// This control at the structure level is essential, check that if the
// target code is "0000000" the pec is not empty
func datiTrasmissioneValidate(d validator.StructLevel) {
	data := d.Current().Interface().(datiTrasmissione)
	if data.CodiceDestinatario == "0000000" {
		if data.PECDestinatario == "" {
			d.ReportError(data.PECDestinatario, "PECDestinatario", "", "required", "")
		}
	}

}

// if the person making the invoice is a foreigner, the Italian office must be indicated.
// This validator also checks whether the Stabile Organization is in the
// Italian territory so the Nation value is "IT"
func cessionarioCommittenteValidate(d validator.StructLevel) {
	data := d.Current().Interface().(CessionarioCommittente)
	if data.Sede.Nazione != "IT" {
		if *data.StabileOrganizzazione == (indirizzoType{}) {
			d.ReportError(data.StabileOrganizzazione, "StabileOrganizzazione", "", "required", "")
		} else {
			if data.StabileOrganizzazione.Nazione != "IT" {
				d.ReportError(data.StabileOrganizzazione, "StabileOrganizzazione", "", "eq", "")
			}
		}

	}
}

// control is is a price number 0.00
func isPrice(d validator.FieldLevel) bool {
	s := d.Field().String()
	val, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return false
	}
	stringVersion := fmt.Sprintf("%.2f", val)
	if s != stringVersion {
		return false
	}
	return true
}

//is Data format: YYYY-MM-DD.
func isDate(field validator.FieldLevel) bool {
	data := field.Field().String()
	_, err := time.Parse(`2006-01-02`, data)
	if err != nil {
		return false
	}
	return true
}

// validate format: YYYY-MM-DDTHH:MM:SS.
func isDateTime(field validator.FieldLevel) bool {

	return true
}

func isNatura(field validator.FieldLevel) bool {
	c := field.Field().String()

	for key, _ := range NatureWithDescription {
		if key == c {
			return true
		}
	}
	return false
}

func isMP(field validator.FieldLevel) bool {
	c := field.Field().String()

	for key, _ := range MethodsPayments {
		if key == c {
			return true
		}
	}
	return false
}

func isTypeDocument(field validator.FieldLevel) bool {
	c := field.Field().String()

	for key, _ := range TypeDocument {
		if key == c {
			return true
		}
	}
	return false
}

func isntSDIPec(field validator.FieldLevel) bool {
	c := field.Field().String()
	matched, _ := regexp.MatchString(`sdi\d\d@pec\.fatturapa\.it`, c)
	return matched
	//sdixx@pec.fatturapa.it
}
