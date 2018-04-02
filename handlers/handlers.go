package handlers

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/julienschmidt/httprouter"
	"googlemaps.github.io/maps"
)

var (
	t        = template.Must(template.New("FC").ParseGlob("templates/*.html"))
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

// Index is main page
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t.ExecuteTemplate(w, "index", nil)
}

// FoodRequest is websocket connection, that performs query from user.
// It ...
func FoodRequest(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrading:", err)
		return
	}
	defer ws.Close()

	c, err := maps.NewClient(maps.WithAPIKey("AIzaSyAjlaVCKC9i-xPZRNuKFAHBPNukxOaGE_o"))
	if err != nil {
		log.Println("maps.NewClient: ", err)
	}

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Read message:", err)
			return
		}
		log.Println("Проверка на получение: ", string(msg))

		slc := strings.Split(string(msg), "|")
		lat, _ := strconv.ParseFloat(slc[0], 64)
		lng, _ := strconv.ParseFloat(slc[1], 64)

		req := &maps.TextSearchRequest{
			Query:    slc[2],
			Language: "russian",
			Location: &maps.LatLng{
				Lat: lat,
				Lng: lng,
			},
			Radius: 1500,
		}

		resp, err := c.TextSearch(context.Background(), req)
		if err != nil {
			log.Println("TextSearch: ", err)
		}
		ws.WriteJSON(resp.Results)
	}
}
