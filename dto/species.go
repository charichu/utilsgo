package dto

type Specie struct {
	Id                    string                 `json:"id"`
	Name                  string                 `json:"name"`
	ApValue               int                    `json:"ap_value"`
	HealthBase            int                    `json:"health_base"`
	SpiritBase            int                    `json:"spirit_base"`
	ToughnessBase         int                    `json:"toughness_base"`
	MovementBase          int                    `json:"movement_base"`
	AttributeAdjustments  []AttributeAdjustments `json:"attribute_adjustments"`
	CommonCultures        []string               `json:"common_cultures"`
	AutoAdvantages        []string               `json:"auto_advantages"`
	CommonAdvantages      []string               `json:"common_advantages"`
	UncommonAdvantages    []string               `json:"uncommon_advantages"`
	AutoDisadvantages     []string               `json:"auto_disadvantages"`
	CommonDisadvantages   []string               `json:"common_disadvantages"`
	UncommonDisadvantages []string               `json:"uncommon_disadvantages"`
	Notes                 string                 `json:"notes"`
}
