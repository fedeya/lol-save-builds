package scrapper

import (
	"strings"

	"github.com/gocolly/colly/v2"
)

// GetChampSkills get the champ skills from op.gg
func (s *Scrapper) GetChampSkills(e *colly.HTMLElement) {
	e.ForEach("li > span", func(i int, e *colly.HTMLElement) {
		s.Champ.Skills[i] = e.Text
	})
}

// GetChampSkillsOrder get champ skills order from op.gg
func (s *Scrapper) GetChampSkillsOrder(e *colly.HTMLElement) {
	e.ForEach("td", func(i int, e *colly.HTMLElement) {
		s.Champ.SkillsOrder[i] = strings.TrimSpace(e.Text)
	})
}