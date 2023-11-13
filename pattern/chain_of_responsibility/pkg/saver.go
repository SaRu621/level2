package pkg

import "fmt"

type DataService struct {
	Next Service
}

func (upd *DataService) Execute(data *Data) {
	if !data.UpdateSource {
		fmt.Printf("Data not update\n")

		return
	}

	fmt.Println("Data save.")

}

func (upd *DataService) SetNext(svc Service) {
	upd.Next = svc

}
