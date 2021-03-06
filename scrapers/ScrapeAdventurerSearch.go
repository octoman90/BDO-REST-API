package scrapers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gocolly/colly/v2"

	"bdo-rest-api/models"
	"bdo-rest-api/translators"
)

func ScrapeAdventurerSearch(region string, query string, searchType uint8, page uint16) (profiles []models.Profile, status int) {
	c := collyFactory()
	closetime := false

	c.OnHTML(closetimeSelector, func(e *colly.HTMLElement) {
		closetime = true
	})

	c.OnHTML(`.box_list_area li:not(.no_result)`, func(e *colly.HTMLElement) {
		profile := models.Profile{
			Region:        region,
			FamilyName:    e.ChildText(".title a"),
			ProfileTarget: extractProfileTarget(e.ChildAttr(".title a", "href")),
			Characters:    make([]models.Character, 1),
		}

		if region != "SA" {
			profile.Region = e.ChildText(".region_info")
		}

		if len(e.ChildAttr(".state a", "href")) > 0 {
			profile.Guild = &models.GuildProfile{
				Name: e.ChildText(".state a"),
			}
		}

		profile.Characters[0].Class = e.ChildText(".name")
		profile.Characters[0].Name = e.ChildText(".text")

		// Site displays the main character when searching by family name
		// And the searched character when searching by character name
		if 2 == searchType {
			profile.Characters[0].Main = true
		}

		if region == "SA" {
			translators.TranslateClassName(&profile.Characters[0].Class)
		}

		if level, err := strconv.Atoi(e.ChildText(".level")[3:]); err == nil {
			profile.Characters[0].Level = uint8(level)
		}

		profiles = append(profiles, profile)
	})

	c.Visit(fmt.Sprintf("%v/Adventure?region=%v&searchType=%v&searchKeyword=%v&Page=%v", getSiteRoot(region), region, searchType, query, page))

	if closetime {
		status = http.StatusServiceUnavailable
	} else if len(profiles) < 1 {
		status = http.StatusNotFound
	} else {
		status = http.StatusOK
	}

	return
}
