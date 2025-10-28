package database

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/learning_new_techs/go/structs"
)

// SeedMovies lê o arquivo movies.json e insere os registros na tabela movies
func SeedMovies(jsonPath string) {
	data, err := ioutil.ReadFile(jsonPath)
	if err != nil {
		log.Fatalf("Erro ao ler arquivo de seed: %v", err)
	}

	var movies []structs.Movie
	if err := json.Unmarshal(data, &movies); err != nil {
		log.Fatalf("Erro ao fazer parse do JSON: %v", err)
	}

	log.Printf("Seed: criando %d registros na tabela movies...", len(movies))
	for _, movie := range movies {
		DB.Create(&movie)
		log.Printf("Seed: criando ro egistro na tabela movies: %s", movie.Name)
	}
	log.Printf("Seed realizado: %d filmes inseridos.", len(movies))
}
