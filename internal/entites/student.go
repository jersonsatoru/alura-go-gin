package entities

type Student struct {
	Name string `json:"name"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}

var Students []Student = []Student{
	{Name: "Jerson", CPF: "00000000000", RG: "11111111"},
	{Name: "Satoru", CPF: "12345678900", RG: "32165487"},
}
