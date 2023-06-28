package main

import (
	"dcchargen/gen"
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

func main() {
	p := &gen.Profile{
		W1:     gen.RandWeapon(),
		W2:     gen.RandWeapon(),
		Pri:    "",
		Sec:    "",
		Charm:  1,
		Heart:  1,
		Focus:  1,
		Power:  1,
		Spells: 0,
		Speed:  0,
		HP:     0,
	}
	// primary and secondary stats
	p.Pri = gen.Primary(p)
	p.Sec = gen.Secondary(p)

	// list random quirks & traits
	quirks := gen.Quirks(p)
	traits := gen.Traits(p)

	gen.Stats(6, p)

	ability := gen.Abilities(2, p)
	spells := gen.Spells(p)

	//divider := "\n ------ \n"
	weapons := gen.Green.Render("Main Hand: ") + p.W1 + " | " + gen.Green.Render("Offhand: ") + p.W2
	statPrio := gen.Green.Render("Primary ") + p.Pri + " | " + gen.Green.Render("Secondary: ") + p.Sec
	charm := gen.StylePlainBorder.Render(gen.Green.Render("Charm: ") + strconv.Itoa(p.Charm))
	heart := gen.StylePlainBorder.Render(gen.Green.Render("Heart: ") + strconv.Itoa(p.Heart))
	focus := gen.StylePlainBorder.Render(gen.Green.Render("Focus: ") + strconv.Itoa(p.Focus))
	power := gen.StylePlainBorder.Render(gen.Green.Render("Power: ") + strconv.Itoa(p.Power))

	hp := gen.StylePlainBorder.Render(gen.Green.Render("HP: ") + strconv.Itoa(p.HP))
	speed := gen.StylePlainBorder.Render(gen.Green.Render("Speed: ") + strconv.Itoa(p.Speed))
	spn := gen.StylePlainBorder.Render(gen.Green.Render("Spells: ") + strconv.Itoa(p.Spells))

	tr := gen.Style.Render("Traits\n") + traits
	qu := gen.Style.Render("Quirks\n") + quirks

	ab := gen.Style.Render("Abilities\n") + ability
	spdis := gen.Style.Render("Spells\n") + spells

	if p.Spells != 0 {

	}
	var primInfo = lipgloss.JoinVertical(lipgloss.Center, weapons, statPrio, lipgloss.JoinHorizontal(lipgloss.Center, hp, speed, spn))
	var statBlock = lipgloss.JoinHorizontal(lipgloss.Center, charm, heart, focus, power)
	var moreInfo = lipgloss.JoinVertical(lipgloss.Left, tr, qu, ab, spdis)

	var charSheet = gen.StyleCuteBorder.Render(lipgloss.JoinVertical(lipgloss.Center, lipgloss.JoinVertical(lipgloss.Center, primInfo, statBlock, moreInfo)))
	block := lipgloss.PlaceHorizontal(80, lipgloss.Center, charSheet)

	fmt.Print(block)
}
