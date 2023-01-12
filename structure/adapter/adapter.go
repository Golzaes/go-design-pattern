package adapter

import "fmt"

type Computer interface {
	InsertIntoLightningPort()
}

type (
	OSClient       struct{}
	Mac            struct{}
	Windows        struct{}
	WindowsAdapter struct{ *Windows }
)

func (*OSClient) InsertLightningConnectorIntoComputer(com Computer) {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}

func (m *Mac) InsertIntoLightningPort() {
	fmt.Println("Lightning connector is plugged into mac machine.")
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.insertIntoUSBPort()
}

func (*Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged into windows machine.")
}
