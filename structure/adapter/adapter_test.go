package adapter

import "testing"

func TestComputer(t *testing.T) {
	o := &OSClient{}
	o.InsertLightningConnectorIntoComputer(&Mac{})

}

func TestAdapter(t *testing.T) {
	adapter := WindowsAdapter{&Windows{}}
	adapter.InsertIntoLightningPort()
}
