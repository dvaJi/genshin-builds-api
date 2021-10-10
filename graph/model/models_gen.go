// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Artifact struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	MinRarity int          `json:"min_rarity"`
	MaxRarity int          `json:"max_rarity"`
	Flower    *ArtifactSet `json:"flower"`
	Plume     *ArtifactSet `json:"plume"`
	Sands     *ArtifactSet `json:"sands"`
	Goblet    *ArtifactSet `json:"goblet"`
	Circlet   *ArtifactSet `json:"circlet"`
	OnePc     *string      `json:"one_pc"`
	TwoPc     *string      `json:"two_pc"`
	FourPc    *string      `json:"four_pc"`
}

type ArtifactSet struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type Ascension struct {
	Ascension *int               `json:"ascension"`
	Level     []*int             `json:"level"`
	Cost      *int               `json:"cost"`
	MatOne    *AscensionMaterial `json:"mat_one"`
	MatTwo    *AscensionMaterial `json:"mat_two"`
	MatThree  *AscensionMaterial `json:"mat_three"`
	MatFour   *AscensionMaterial `json:"mat_four"`
}

type AscensionMaterial struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Amount *int   `json:"amount"`
	Rarity *int   `json:"rarity"`
}

type Character struct {
	ID              string            `json:"id"`
	Name            string            `json:"name"`
	Title           string            `json:"title"`
	Description     string            `json:"description"`
	WeaponType      string            `json:"weapon_type"`
	Element         string            `json:"element"`
	Gender          string            `json:"gender"`
	Substat         string            `json:"substat"`
	Affiliation     string            `json:"affiliation"`
	Region          string            `json:"region"`
	Rarity          *int              `json:"rarity"`
	Birthday        []*int            `json:"birthday"`
	Constellation   string            `json:"constellation"`
	Domain          string            `json:"domain"`
	Cv              *CharacterVoice   `json:"cv"`
	Skills          []*Skill          `json:"skills"`
	Passives        []*Passive        `json:"passives"`
	Constellations  []*Constellation  `json:"constellations"`
	Ascension       []*Ascension      `json:"ascension"`
	TalentMaterials []*TalentMaterial `json:"talent_materials"`
}

type CharacterVoice struct {
	English  string `json:"english"`
	Chinese  string `json:"chinese"`
	Japanese string `json:"japanese"`
	Korean   string `json:"korean"`
}

type CommonMaterial struct {
	ID          *string              `json:"id"`
	Name        *string              `json:"name"`
	Description *string              `json:"description"`
	Sources     []*string            `json:"sources"`
	Rarity      *int                 `json:"rarity"`
	Craft       *CommonMaterialCraft `json:"craft"`
}

type CommonMaterialCraft struct {
	Cost  *int                  `json:"cost"`
	Items []*CommonMaterialItem `json:"items"`
}

type CommonMaterialItem struct {
	ID     *string `json:"id"`
	Name   *string `json:"name"`
	Amount *int    `json:"amount"`
}

type Constellation struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       *int   `json:"level"`
}

type Passive struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Level       *int   `json:"level"`
}

type Skill struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Info        string            `json:"info"`
	Attributes  []*SkillAttribute `json:"attributes"`
}

type SkillAttribute struct {
	Label  string   `json:"label"`
	Values []string `json:"values"`
}

type TalentMaterial struct {
	Level *int                 `json:"level"`
	Cost  *int                 `json:"cost"`
	Items []*AscensionMaterial `json:"items"`
}
