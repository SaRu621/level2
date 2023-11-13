package pkg

import "fmt"

type UpdateDataService struct {
	Name string
	Next Service
}

func (upd *UpdateDataService) Execute(data *Data) {
	if data.UpdateSource {
		fmt.Printf("Data from device [%s] already get.\n", upd.Name)

		upd.Next.Execute(data)

		return
	}

	fmt.Printf("Update data from device [%s].\n", upd.Name)

	data.GetSource = true

	upd.Next.Execute(data)

}

func (upd *UpdateDataService) SetNext(svc Service) {
	upd.Next = svc

}
