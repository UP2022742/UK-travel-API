package worker

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

func ReturnCountries(id string) ([]string, error) {
	c := colly.NewCollector()
	var countries []string
	c.OnHTML(id, func(e *colly.HTMLElement) {
		e.DOM.Next().Find("tbody").Children().Each(func(i int, s *goquery.Selection) {
			re := regexp.MustCompile(`([\(\[]).*?([\)\]])`)
			thing := strings.TrimSuffix(s.Find("th").Text(), "\n")

			switch thing {
			case "Congo (Democratic Republic)": // Our issue, ish?
				countries = append(countries, "CD")
				break
			case "Congo": // Our issue
				countries = append(countries, "CG")
				break
			case "South Sudan": // Definietly google issue...
				countries = append(countries, "SS")
				break
			case "Côte d’Ivoire":
				countries = append(countries, "CI")
				break
			case "The Bahamas":
				countries = append(countries, "BS")
				break
			case "  Israel and Jerusalem":
				countries = append(countries, "IL")
				break
			case "Kosovo":
				countries = append(countries, "XK")
				break
			case "North Macedonia":
				countries = append(countries, "MK")
				break
			case "Timor-Leste":
				countries = append(countries, "TL")
				break
			case "Eswatini":
				countries = append(countries, "SZ")
				break
			case "Réunion":
				countries = append(countries, "RE")
				break
			default:
				new := re.ReplaceAllString(thing, "")
				countries = append(countries, new)
				break
			}
		})
	})
	c.Visit("https://www.gov.uk/guidance/red-amber-and-green-list-rules-for-entering-england")

	if len(countries) == 0 {
		return nil, fmt.Errorf("Unable to fetch results.")
	}
	return countries, nil
}

func (route *Worker) SendJson(w http.ResponseWriter, r *http.Request, id string, pageRequest string) {
	countries, err := ReturnCountries(id)
	if err != nil {
		route.logger.Error(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	payload, err := json.Marshal(countries)
	if err != nil {
		route.logger.Debug(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	route.logger.Debug("A '" + pageRequest + "' request was made.")

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
}

func (route *Worker) GreenList(w http.ResponseWriter, r *http.Request) {
	route.SendJson(w, r, "#green-list-of-countries-and-territories---from-17-may", "Green")
}

func (route *Worker) AmberList(w http.ResponseWriter, r *http.Request) {
	route.SendJson(w, r, "#amber-list-of-countries-and-territories", "Amber")
}

func (route *Worker) RedList(w http.ResponseWriter, r *http.Request) {
	route.SendJson(w, r, "#red-list-of-countries-and-territories", "Red")
}
