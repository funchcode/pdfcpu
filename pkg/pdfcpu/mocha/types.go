package mocha

import "github.com/pkg/errors"

type Form struct {
	Common 		map[string]Common	`json:"common"`
	Users		map[string]Users	`json:"users"`
}

type Common struct {
	Type			string	`json:"type"`
	Page			int		`json:"page"`
	Value			string	`json:"value"`
	Left			float64	`json:"left"`
	Top				float64 `json:"top"`
	FontFamily 		string	`json:"fontFamily"`
	FontSize 		int		`json:"fontSize"`
}

type Users struct {
	Data	Data `json:"data"`
}

type Data struct {
	SignatureGroup		 map[string]SignatureGroup		`json:"signatureGroup"`
	SignTextGroup		 map[string]SignTextGroup		`json:"signTextGroup"`
	SignCheckboxGroup	 map[string]SignCheckboxGroup	`json:"signCheckboxGroup"`
}

type SignatureGroup struct {
	IsRequired		 string		`json:"isRequired"`
	Page 			 int		`json:"page"`
	Left		     float64	`json:"left"`
	Top 			 float64	`json:"top"`
	Value 			 string		`json:"value"`
}

type SignTextGroup struct {
	IsRequired 		string		`json:"isRequired"`
	Page 			int			`json:"page"`
	Left 			float64		`json:"left"`
	Top 			float64		`json:"top"`
	FontFamily 		string		`json:"fontFamily"`
	FontSize 		int			`json:"fontSize"`
	Value 			string		`json:"value"`
}

type SignCheckboxGroup struct {
	IsRequired 		string					`json:"isRequired"`
	Page 			int						`json:"page"`
	Value 			string					`json:"value"`
	Options 		Options					`json:"options"`
	Checkboxes 		map[string]Checkbox		`json:"checkboxes"`
}

type Options struct {
	Min int		`json:"min"`
	Max int		`json:"max"`
}

type Checkbox struct {
	Id 		string		`json:"id"`
	Left 	float64		`json:"left"`
	Top		float64		`json:"top"`
}

func (form *Form) Validate() error {
	if err := form.validateCommon(); err != nil {
		return err
	}
	if err := form.validateSignText(); err != nil {
		return err
	}
	if err := form.validateSignature(); err != nil {
		return err
	}
	if err := form.validateSignCheckbox(); err != nil {
		return err
	}
	return nil
}

func (form *Form) validateCommon() error {
	for _, common := range form.Common {
		switch common.Type {
		case "text":
			break
		case "image":
			if len(common.Value) <= 0 {
				return errors.Errorf("mocha - common value policy violation")
			}
			break
		}
	}
	return nil
}

func (form *Form) validateSignText() error {
	for _, users := range form.Users {
		for _, signText := range users.Data.SignTextGroup {
			if len(signText.Value) <= 0 &&
				signText.IsRequired == "1" {
				return errors.Errorf("mocha - signTextField value policy violation")
			}
		}
	}
	return nil
}

func (form *Form) validateSignature() error {
	for _, users := range form.Users {
		for _, signature := range users.Data.SignatureGroup {
			if len(signature.Value) <= 0 &&
				signature.IsRequired == "1" {
				return errors.Errorf("mocha - signatureField value policy violation")
			}
		}
	}
	return nil
}

func (form *Form) validateSignCheckbox() error {
	for _, users := range form.Users {
		for _, checkbox := range users.Data.SignCheckboxGroup {
			if len(checkbox.Value) <= 0 &&
				checkbox.IsRequired == "1" {
				return errors.Errorf("mocha - signCheckboxField value policy violation")
			} else if checkbox.IsRequired == "1" {
				// todo options 체크
			}
		}
	}
	return nil
}

