package handlers

import (
	"context"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"googlemaps.github.io/maps"
)

var t = template.Must(template.New("FC").ParseGlob("templates/*.html"))

// Index is main page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	apiKey, err := ioutil.ReadFile("api-key")
	if err != nil {
		log.Fatal("ioutil.ReadFile: ", err)
	}

	c, err := maps.NewClient(maps.WithAPIKey(string(apiKey)))
	if err != nil {
		log.Fatal("maps.newClient: ", err)
	}

	gr := &maps.GeocodingRequest{
		LatLng: &maps.LatLng{Lat: 53.347457, Lng: 83.775711},
	}

	gres, err := c.Geocode(context.Background(), gr)
	if err != nil {
		log.Fatal("c.Geocode: ", err)
	}

	t.ExecuteTemplate(w, "index", gres)
}
