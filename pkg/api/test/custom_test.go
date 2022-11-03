package test

import (
	"encoding/json"
	"fmt"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/mocha"
	"testing"
)

func TestFormValidate(t *testing.T) {
	data := `{
    "common": {},
    "current": null,
    "users": {
        "f6ba2006044544c088ae310d7a22fd78": {
            "id": "f6ba2006044544c088ae310d7a22fd78",
            "name": "dev@intosoft.kr",
            "color": "#EDA080",
            "custom": null,
            "data": {
                "signatureGroup": null,
                "signTextGroup": {
                    "1fcb5901e286424d99be1f97f8e67db2": {
                        "type": "signText",
                        "id": "1fcb5901e286424d99be1f97f8e67db2",
                        "userId": "f6ba2006044544c088ae310d7a22fd78",
                        "customId": "",
                        "customDesc": "",
                        "page": 1,
                        "isRequired": "0",
                        "value": "asdf",
                        "label": "",
                        "comment": "",
                        "left": 382,
                        "top": 92,
                        "width": 80,
                        "height": 16,
                        "name": null,
                        "fontFamily": "'NotoSans', sans-serif",
                        "fontName": "본고딕",
                        "fontSize": 10,
                        "fontColor": "#191919",
                        "valueFontSize": 0,
                        "valueHeight": 0
                    }
                },
                "signCheckboxGroup": {
                    "5929ae52cfd04bf9b82dc8be9b61f59d": {
                        "type": "signCheckbox",
                        "id": "5929ae52cfd04bf9b82dc8be9b61f59d",
                        "userId": "f6ba2006044544c088ae310d7a22fd78",
                        "customId": "",
                        "customDesc": "",
                        "page": 1,
                        "isRequired": "0",
                        "value": "",
                        "label": "",
                        "comment": "",
                        "left": 0,
                        "top": 0,
                        "width": 20,
                        "height": 20,
                        "name": null,
                        "options": {
                            "min": 2,
                            "max": 4
                        },
                        "checkboxes": {
                            "ad609fb8b60b4b8aac97a92979d693c3": {
                                "id": "ad609fb8b60b4b8aac97a92979d693c3",
                                "groupId": "5929ae52cfd04bf9b82dc8be9b61f59d",
                                "option": "한국",
                                "left": 260.20000000000005,
                                "top": 259.01875
                            },
                            "4bb426ba1dad499b89aba55d1d9f8e19": {
                                "id": "4bb426ba1dad499b89aba55d1d9f8e19",
                                "groupId": "5929ae52cfd04bf9b82dc8be9b61f59d",
                                "option": "미국",
                                "left": 286.20000000000005,
                                "top": 259.01875
                            },
                            "89a8c71b5a2640608d33bae7e9df8111": {
                                "id": "89a8c71b5a2640608d33bae7e9df8111",
                                "groupId": "5929ae52cfd04bf9b82dc8be9b61f59d",
                                "option": "일본",
                                "left": 312.20000000000005,
                                "top": 259.01875
                            },
                            "6810c78ffcc1415d8dae1a6471f9b367": {
                                "id": "6810c78ffcc1415d8dae1a6471f9b367",
                                "groupId": "5929ae52cfd04bf9b82dc8be9b61f59d",
                                "option": "중국",
                                "left": 338.20000000000005,
                                "top": 259.01875
                            },
                            "a0f72aca0332486883282e01d3c3c183": {
                                "id": "a0f72aca0332486883282e01d3c3c183",
                                "groupId": "5929ae52cfd04bf9b82dc8be9b61f59d",
                                "option": "태국",
                                "left": 364.20000000000005,
                                "top": 259.01875
                            }
                        }
                    }
                },
                "signRadioGroup": null
            },
            "nameId": null
        }
    }
}`
	form := mocha.Form{}
	err := json.Unmarshal([]byte(data), &form)
	if err != nil {
		t.Fatalf("%v\n", err)
	}
	err2 := form.Validate()
	if err2 != nil {
		t.Fatalf("%v\n", err)
	}
}

func TestFormMarshaing(t *testing.T) {
	data := `{
    "common": {},
    "current": null,
    "users": {
        "f6ba2006044544c088ae310d7a22fd78": {
            "id": "f6ba2006044544c088ae310d7a22fd78",
            "name": "dev@intosoft.kr",
            "color": "#EDA080",
            "custom": null,
            "data": {
                "signatureGroup": null,
                "signTextGroup": {
                    "1fcb5901e286424d99be1f97f8e67db2": {
                        "type": "signText",
                        "id": "1fcb5901e286424d99be1f97f8e67db2",
                        "userId": "f6ba2006044544c088ae310d7a22fd78",
                        "customId": "",
                        "customDesc": "",
                        "page": 1,
                        "isRequired": "0",
                        "value": "asdf",
                        "label": "",
                        "comment": "",
                        "left": 382,
                        "top": 92,
                        "width": 80,
                        "height": 16,
                        "name": null,
                        "fontFamily": "'NotoSans', sans-serif",
                        "fontName": "본고딕",
                        "fontSize": 10,
                        "fontColor": "#191919",
                        "valueFontSize": 0,
                        "valueHeight": 0
                    }
                },
                "signCheckboxGroup": {
                    "5929ae52cfd04bf9b82dc8be9b61f59d": {
                        "type": "signCheckbox",
                        "id": "5929ae52cfd04bf9b82dc8be9b61f59d",
                        "userId": "f6ba2006044544c088ae310d7a22fd78",
                        "customId": "",
                        "customDesc": "",
                        "page": 1,
                        "isRequired": "0",
                        "value": "ad609fb8b60b4b8aac97a92979d693c3,4bb426ba1dad499b89aba55d1d9f8e19,89a8c71b5a2640608d33bae7e9df8111",
                        "label": "",
                        "comment": "",
                        "left": 0,
                        "top": 0,
                        "width": 20,
                        "height": 20,
                        "name": null,
                        "options": {
                            "min": 2,
                            "max": 4
                        },
                        "checkboxes": {
                            "ad609fb8b60b4b8aac97a92979d693c3": {
                                "id": "ad609fb8b60b4b8aac97a92979d693c3",
                                "groupId": "5929ae52cfd04bf9b82dc8be9b61f59d",
                                "option": "한국",
                                "left": 260.20000000000005,
                                "top": 259.01875
                            },
                            "4bb426ba1dad499b89aba55d1d9f8e19": {
                                "id": "4bb426ba1dad499b89aba55d1d9f8e19",
                                "groupId": "5929ae52cfd04bf9b82dc8be9b61f59d",
                                "option": "미국",
                                "left": 286.20000000000005,
                                "top": 259.01875
                            },
                            "89a8c71b5a2640608d33bae7e9df8111": {
                                "id": "89a8c71b5a2640608d33bae7e9df8111",
                                "groupId": "5929ae52cfd04bf9b82dc8be9b61f59d",
                                "option": "일본",
                                "left": 312.20000000000005,
                                "top": 259.01875
                            },
                            "6810c78ffcc1415d8dae1a6471f9b367": {
                                "id": "6810c78ffcc1415d8dae1a6471f9b367",
                                "groupId": "5929ae52cfd04bf9b82dc8be9b61f59d",
                                "option": "중국",
                                "left": 338.20000000000005,
                                "top": 259.01875
                            },
                            "a0f72aca0332486883282e01d3c3c183": {
                                "id": "a0f72aca0332486883282e01d3c3c183",
                                "groupId": "5929ae52cfd04bf9b82dc8be9b61f59d",
                                "option": "태국",
                                "left": 364.20000000000005,
                                "top": 259.01875
                            }
                        }
                    }
                },
                "signRadioGroup": null
            },
            "nameId": null
        }
    }
}`
	form := mocha.Form{}
	err := json.Unmarshal([]byte(data), &form)
	fmt.Printf("%#v\n", form)
	fmt.Println()
	for _, v := range form.Users {
		for _, vv := range v.Data.SignatureGroup {
			fmt.Println("ddd")
			fmt.Println(vv)
		}
	}

	json, _ := json.Marshal(form)
	fmt.Println(string(json))
	if err != nil {
		t.Fatalf("%v\n", err)
	}
}
