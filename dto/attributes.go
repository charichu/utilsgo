package dto

type AttributeAdjustments struct {
	Attribute  Attribute `json:"attribute"`
	Adjustment int       `json:"adjustment"`
}

type Attribute struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
