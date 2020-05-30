package interfaces

import (
	"encoding/json"
	"golang-songs/domain"
	"golang-songs/usecases"
	"net/http"
	"strconv"
)

type SongController struct {
	SongInteractor usecases.SongInteractor
	Logger         usecases.Logger
}

func NewSongController(sqlHandler SQLHandler, logger usecases.Logger) *SongController {
	return &SongController{
		SongInteractor: usecases.SongInteractor{
			SongRepository: &SongRepository{
				SQLHandler: sqlHandler,
			},
		},
		Logger: logger,
	}
}

// Index is display a listing of the resource.
func (pc *SongController) Index(w http.ResponseWriter, r *http.Request) {
	pc.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	posts, err := pc.SongInteractor.Index()
	if err != nil {
		pc.Logger.LogError("%s", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}

// Store is stora a newly created resource in storage.
func (pc *SongController) Store(w http.ResponseWriter, r *http.Request) {
	pc.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	p := domain.Post{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		pc.Logger.LogError("%s", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}

	post, err := pc.SongInteractor.Store(p)
	if err != nil {
		pc.Logger.LogError("%s", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

// Destroy is remove the specified resource from storage.
func (pc *SongController) Destroy(w http.ResponseWriter, r *http.Request) {
	pc.Logger.LogAccess("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	postID, _ := strconv.Atoi(r.URL.Query().Get("id"))

	err := pc.SongInteractor.Destroy(postID)
	if err != nil {
		pc.Logger.LogError("%s", err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(nil)
}
