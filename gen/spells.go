package gen

import (
	"embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed data/spells.json
var spellsFS embed.FS

var sjson, serr = spellsFS.ReadFile("data/spells.json")

func spellSelect(p *Profile) string {
	var s string
	tier := "t1"

	a := unmarshal(sjson, fmt.Sprintf("%s.#", tier)).Int()
	i := randNum(int(a))
	sN := unmarshal(sjson, fmt.Sprintf("%s.%s.name", tier, strconv.Itoa(i))).Str
	sP := unmarshal(sjson, fmt.Sprintf("%s.%s.prose", tier, strconv.Itoa(i))).Str

	s = fmt.Sprintf("%s: %s\n", Bold.Render(sN), sP)
	return BasicPadding.Render(s)
}

// listing the abilities as a string
func Spells(p *Profile) string {
	var spell string

	for i := 0; i < p.Spells; i++ {
		add := spellSelect(p)
		if strings.Contains(spell, add) {
			i--
		} else {
			spell += add
		}
	}
	return spell
}
