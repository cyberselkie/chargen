package gen

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data/abilities.json
var abilityFS embed.FS

var ajson, aerr = abilityFS.ReadFile("data/abilities.json")

func flipWeapon(p *Profile) string {
	var s = []string{p.W1, p.W2}
	n := randNum(2)

	return s[n]
}

func abilitySelect(p *Profile) string {
	var ability string
	w := flipWeapon(p)
	tier := "t1"

	a := unmarshal(ajson, fmt.Sprintf("%s.%s.#", w, tier)).Int()
	fmt.Print(a)
	i := randNum(int(a))
	abilityN := unmarshal(ajson, fmt.Sprintf("%s.%s.%s.name", w, tier, strconv.Itoa(i))).Str
	abilityP := unmarshal(ajson, fmt.Sprintf("%s.%s.%s.prose", w, tier, strconv.Itoa(i))).Str

	ability = fmt.Sprintf("%s: %s\n", Bold.Render(abilityN), abilityP)
	return BasicPadding.Render(ability)
}

// listing the abilities as a string
func Abilities(amt int, p *Profile) string {
	var ability string

	for i := 0; i < amt; i++ {
		add := abilitySelect(p)
		if strings.Contains(ability, add) {
			i--
		} else {
			ability += add
		}
	}
	return ability
}
