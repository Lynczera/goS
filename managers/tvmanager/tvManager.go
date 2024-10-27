package tvmanager

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Program struct {
	ShowType     string `json:"showType"`
	EpisodeTitle string `json:"episodeTitle"`
	ListDateTime string `json:"listDateTime"`
}

// returns fixed value. Will call mediatv api in the future
func GetLineup(zip string) (string, error) {
	return "6063", nil
}

/*
take arguments, call MediaTv api, and return an list of objects.
returning test data at the moment.
*/

func GetListings() {
	var programs []Program

	//need to find a way to change base dir of running program?
	dir, direrr := os.Getwd()
	if direrr != nil {
		fmt.Println("Error getting current directory:", direrr)
		return
	}

	jsonFile, jsonerr := os.Open(dir + "/testdata/listing.json")

	if jsonerr != nil {
		fmt.Println(jsonerr)
		return
	}

	defer jsonFile.Close()

	bytes, _ := io.ReadAll(jsonFile)

	json.Unmarshal(bytes, &programs)

	for i := 0; i < 4; i++ {
		fmt.Printf("showtype:%v / ", programs[i].ShowType)
		fmt.Printf("episodetitle:%v / ", programs[i].EpisodeTitle)
		fmt.Printf("listdatetime:%v\n", programs[i].ListDateTime)
	}

}
