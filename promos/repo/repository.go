package repo

import (
	"fmt"
	"log"
	"sync"

	src "gihub.com/dna2/promos/source"
)

var (
	once           sync.Once
	singleInstance *Repository
)

type FnOnFinish func()

type Repository struct {
	src                     *src.Source
	ChedrauiRepository      *chedrauiRepository
	BodegaAurreraRepository *bodegaARepository
	ElektraRepository       *elektraRepository
}

func FetchRepository() (*Repository, error) {
	var err error
	if singleInstance == nil {
		once.Do(
			func() {
				log.Println("Repository: Creating single instance now.")
				singleInstance, err = newInstanceRepository()
			})
	} else {
		log.Println("Repository: Single instance already created.")
	}
	return singleInstance, err
}

func newInstanceRepository() (*Repository, error) {
	source, err := src.FetchSource()
	if err != nil {
		return nil, fmt.Errorf("Repository: error instance src - %v", err)
	}
	chedrauiRepo, err := newInstanceChedrauiRepo(source)
	if err != nil {
		return nil, fmt.Errorf("Repository: error instance chedraui - %v", err)
	}
	bodegaARepository, err := newInstanceBodegaARepo(source)
	if err != nil {
		return nil, fmt.Errorf("Repository: error instance bodegaa - %v", err)
	}
	elektraRepository, err := newInstanceElektraRepo(source)
	if err != nil {
		return nil, fmt.Errorf("Repository: error instance elektra - %v", err)
	}
	return &Repository{src: source, ChedrauiRepository: chedrauiRepo, BodegaAurreraRepository: bodegaARepository, ElektraRepository: elektraRepository}, nil
}

func (repo *Repository) Close() {
	if repo.src != nil {
		repo.src.Close()
	}
}
