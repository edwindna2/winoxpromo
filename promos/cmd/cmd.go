package cmd

import (
	"log"
	"sync"

	config "gihub.com/dna2/promos/config"
	repo "gihub.com/dna2/promos/repo"
)

func Start() {
	//Init Config
	config.Setup()

	//Repository
	repo, err := repo.FetchRepository()
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer repo.Close()

	wg := sync.WaitGroup{}
	wg.Add(2)
	go repo.ChedrauiRepository.StartScraping(func() {
		log.Println("Chedraui finished")
		wg.Done()
	})

	go repo.BodegaAurreraRepository.StartScraping(func() {
		log.Println("BodegaAurrera finished")
		wg.Done()
	})

	/*go repo.ElektraRepository.StartScraping(func() {
		log.Println("Elektra finished")
		wg.Done()
	})*/

	wg.Wait()
	repo.Close()
}
