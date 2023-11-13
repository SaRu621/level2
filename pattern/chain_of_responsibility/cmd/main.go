package main

import "chain_of_responsibility/pkg"

func main() {
	device := &pkg.Device{Name: "Device-1"}

	updateSvc := &pkg.UpdateDataService{Name: "Device-1"}

	dataSvc := &pkg.DataService{}

	device.SetNext(updateSvc)
	updateSvc.SetNext(dataSvc)

	data := &pkg.Data{}

	device.Execute(data)
}
