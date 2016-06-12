package main

import (
	"github.com/drone/routes"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"io"
)
type Food         struct {
		DrinkAlcohol string `json:"drink_alcohol"`
		Types         string `json:"type"`
	}
type Movie	     struct {
		Movies  []string `json:"movies"`
		TvShows []string `json:"tv_shows"`
	}
type Music struct {
		SpotifyUserId string `json:"spotify_user_id"`
	}
type Flight	struct {
			Seat string `json:"seat"`
		}
type Travel		struct {
		Fl Flight  `json:"flight"`
	}
type Profile struct {
	Country       string `json:"country"`
	Email         string `json:"email"`
	FavoriteColor string `json:"favorite_color"`
	FavoriteSport string `json:"favorite_sport"`
	Fo Food `json:"food"`
	IsSmoking string `json:"is_smoking"`
	Mov Movie`json:"movie"`
	Mus Music `json:"music"`
	Profession string `json:"profession"`
	T   Travel   `json:"travel"`
	Zip string `json:"zip"`
}

type Profiles [5]Profile
//var prof
var P *Profile
var v = []Profile{}

func main() {
	mux := routes.New()
	mux.Get("/profile/:email", GetProfile)
	mux.Post("/profile",CreateProfile)
	mux.Put("/profile/:email",UpdateProfile)
	mux.Del("/profile/:email",DeleteProfile)
	
	http.Handle("/", mux)
	log.Println("Listening...")
	//log.Println("printing P-->"+P.Email)
	http.ListenAndServe(":3000", nil)
}
func DeleteProfile(w http.ResponseWriter, r *http.Request){
	log.Println("in Delete")
	res2B, _ := json.Marshal(v)
	log.Println("before delete json->"+string(res2B))
	params := r.URL.Query()
	email := params.Get(":email")
	
	i:=0
	for i:=0;i<len(v);i++{
		if(v[i].Email==email){
			log.Println("found")
			break
		}
	}
	
	v = append(v[:i], v[i+1:]...)
	res2B, _ = json.Marshal(v)
	log.Println("after delete json->"+string(res2B))
	w.WriteHeader(http.StatusNoContent)
	

}
func UpdateProfile(w http.ResponseWriter, r *http.Request){
	log.Println("in Update")
	res2B, _ := json.Marshal(v)
	log.Println(string(res2B));
	params := r.URL.Query()
	email := params.Get(":email")
	var prof Profile
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &prof); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }
    log.Println("printing Travel-->"+prof.T.Fl.Seat)
    
    for i:=0;i<len(v);i++{
		if(v[i].Email==email){
			log.Println("found")
			log.Println("email in put-->"+v[i].Email)
			if(prof.T.Fl.Seat!=""){
			log.Println("seat before put"+v[i].T.Fl.Seat)
				v[i].T=prof.T
				log.Println("seat after put"+v[i].T.Fl.Seat)
			}
			if(prof.Email!=""){
				v[i].Email=prof.Email
			}
			if(prof.Zip!=""){
				v[i].Zip=prof.Zip
			}
			if(prof.Country!=""){
				v[i].Country=prof.Country
			}
			if(prof.FavoriteColor!=""){
				v[i].FavoriteColor=prof.FavoriteColor
			}
			if(prof.IsSmoking!=""){
				v[i].IsSmoking=prof.IsSmoking
			}
			if(prof.FavoriteSport!=""){
				v[i].FavoriteSport=prof.FavoriteSport
			}
			if(prof.Fo.DrinkAlcohol!="" || prof.Fo.Types!=""){
				v[i].Fo=prof.Fo
			}
			if(prof.Mus.SpotifyUserId!=""){
				v[i].Mus=prof.Mus
			}
			if(len(prof.Mov.Movies)>0){
				v[i].Mov=prof.Mov
			}
			break
		}
	}
	//json.NewEncoder(w).Encode(v)
	log.Println("seat after put"+v[0].T.Fl.Seat)
	res2B, _ = json.Marshal(v[0])
	log.Println(string(res2B));
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusNoContent)
	
}
func CreateProfile(w http.ResponseWriter, r *http.Request){
	var prof Profile
	res2B, _ := json.Marshal(v)
	log.Println("p in json->"+string(res2B))
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    if err != nil {
        panic(err)
    }
    if err := r.Body.Close(); err != nil {
        panic(err)
    }
    if err := json.Unmarshal(body, &prof); err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
            panic(err)
        }
    }
    v=append(v,prof)
	log.Println("printing P-->"+prof.Email)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    
}
func GetProfile(w http.ResponseWriter, r *http.Request) {
	log.Println("In get")
	res2B, _ := json.Marshal(v)
	log.Println("p in json->"+string(res2B))
	params := r.URL.Query()
	email := params.Get(":email")
	log.Println("printing request email-->"+email)
	f:=0
	for i:=0;i<len(v);i++{
		log.Println("item-->"+v[i].Email)
		if(v[i].Email==email){
		log.Println("found")
		//json.NewEncoder(w).Encode(v[i])
		res2B, _ := json.Marshal(v[i])
		log.Println("p in json->"+string(res2B))
		w.Write([]byte(res2B))
		f=1
		break
		}
	}
	if(f==0){
	log.Println("not found")
		//w.Write([]byte("No Such Profile Exists"))
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(404)
	}
}