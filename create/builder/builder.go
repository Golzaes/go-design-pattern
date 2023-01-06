package builder

type (
	Director struct{ builder IBuilder }
	IBuilder interface {
		setWindowType()
		setDoorType()
		setNumFloor()
		getHouse() House
	}
	House struct {
		windowType string
		doorType   string
		floor      int
	}
)

type (
	IglooBuilder struct {
		windowType string
		doorType   string
		floor      int
	}

	NormalBuilder struct {
		windowType string
		doorType   string
		floor      int
	}
)

func newDirector(b IBuilder) *Director { return &Director{b} }

func (d *Director) setBuilder(b IBuilder) { d.builder = b }

func (d *Director) buildHouse() House {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}
	if builderType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}

func newIglooBuilder() *IglooBuilder   { return &IglooBuilder{} }
func newNormalBuilder() *NormalBuilder { return &NormalBuilder{} }

func (b *IglooBuilder) setWindowType()  { b.windowType = "Snow Window" }
func (b *NormalBuilder) setWindowType() { b.windowType = "Wooden Window" }

func (b *IglooBuilder) setDoorType()  { b.doorType = "Snow Door" }
func (b *NormalBuilder) setDoorType() { b.doorType = "Wooden Door" }

func (b *IglooBuilder) setNumFloor()  { b.floor = 1 }
func (b *NormalBuilder) setNumFloor() { b.floor = 2 }

func (b *IglooBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

func (b *NormalBuilder) getHouse() House {
	return House{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}
