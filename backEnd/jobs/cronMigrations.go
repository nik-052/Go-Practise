package jobs

import (
	"CollegeAdministration/config"
	"CollegeAdministration/daos"
	"log"
	"sync"
	"time"
)

var wg sync.WaitGroup

func RunDailyMigrations() {
	//making go routines synchronous
	wg.Add(2)
	ch := make(chan int)
	log.Println("running migrations")
	go CheckOutDatedTokensSetFalse(ch)
	go DeleteNotValidTokens(ch)
	wg.Wait()
}

func CheckOutDatedTokensSetFalse(ch chan int) {

	daos := daos.New(config.DBInit())
	ch <- 1
	all_tokens, err := daos.GetAllTokens()
	if err != nil {
		log.Panic(err)
	}
	var time_now = time.Now()
	for _, token := range all_tokens {

		if token.ValidTill-time_now.Unix() < 0 {
			daos.SetTokenFalse(token.Token)
		}

	}
	defer wg.Done()
}

func DeleteNotValidTokens(ch chan int) {

	<-ch
	daos := daos.New(config.DBInit())
	err := daos.RunMigrationsForRemovingOutDatedTokens()
	if err != nil {
		log.Panic(err)
	}
	defer wg.Done()
}