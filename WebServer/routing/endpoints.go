package routing

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"text/template"
)

func GetGreenList() ([]string, error) {
	resp, err := http.Get("https://api:8080/green")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var post []string
	err = json.Unmarshal(body, &post)
	if err != nil {
		return nil, err
	}
	return post, nil
}
func GetAmberList() ([]string, error) {
	resp, err := http.Get("https://api:8080/amber")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var post []string
	err = json.Unmarshal(body, &post)
	if err != nil {
		return nil, err
	}
	return post, nil
}
func GetRedList() ([]string, error) {
	resp, err := http.Get("https://api:8080/red")
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		return nil, err
	}

	var post []string
	err = json.Unmarshal(body, &post)
	if err != nil {
		return nil, err
	}
	return post, nil
}

// HomePage contains the index.
func (route *Router) HomePage(w http.ResponseWriter, r *http.Request) {
	type MetaHome struct{}
	templates, err := template.ParseFiles(
		"templates/index.html",
	)
	if err != nil {
		route.logger.Error("Unable to parse 'Home' page.", "err", err)
	}
	templates.ExecuteTemplate(w, "index.html", MetaHome{})
	route.logger.Debug("A 'Home' page request was made.")
}

// NotificationPage contains the template event page.
func (route *Router) AboutPage(w http.ResponseWriter, r *http.Request) {
	type MetaEvent struct {
	}
	templates, err := template.ParseFiles(
		"templates/about.html",
	)
	if err != nil {
		route.logger.Error("Unable to parse 'About' page.", "err", err)
	}
	templates.ExecuteTemplate(w, "about.html", MetaEvent{})
	route.logger.Debug("A 'About' page request was made.")
}

// TablesPage contains the template event page.
func (route *Router) ProjectsPage(w http.ResponseWriter, r *http.Request) {
	type MetaTables struct {
	}
	templates, err := template.ParseFiles(
		"templates/project.html",
	)
	if err != nil {
		route.logger.Error("Unable to parse 'Projects' page.", "err", err)
	}
	templates.ExecuteTemplate(w, "project.html", MetaTables{})
	route.logger.Debug("A 'Projects' page request was made.")
}

// TypographyPage contains the template event page.
func (route *Router) SocialPage(w http.ResponseWriter, r *http.Request) {
	type MetaTypo struct {
	}
	templates, err := template.ParseFiles(
		"templates/social.html",
	)
	if err != nil {
		route.logger.Error("Unable to parse 'Social' page.", "err", err)
	}
	templates.ExecuteTemplate(w, "social.html", nil)
	route.logger.Debug("A 'Social' page request was made.")
}

// TypographyPage contains the template event page.
func (route *Router) TravelPage(w http.ResponseWriter, r *http.Request) {
	type MetaTravel struct {
		GreenListCountries []string
		AmberListCountries []string
		RedListCountries   []string
	}
	templates, err := template.ParseFiles(
		"templates/travel.html",
	)
	if err != nil {
		route.logger.Error("Unable to parse 'Travel' page.", "err", err)
	}

	greenList, err := GetGreenList()
	if err != nil {
		route.logger.Error("Unable to get 'Green' list.", "err", err)
	}
	amberList, err := GetAmberList()
	if err != nil {
		route.logger.Error("Unable to get 'Amber' list.", "err", err)
	}
	redList, err := GetRedList()
	if err != nil {
		route.logger.Error("Unable to get 'Red' list.", "err", err)
	}

	templates.ExecuteTemplate(
		w, "travel.html", MetaTravel{
			GreenListCountries: greenList,
			AmberListCountries: amberList,
			RedListCountries:   redList,
		},
	)
	route.logger.Debug("A 'Travel' page request was made.")
}
