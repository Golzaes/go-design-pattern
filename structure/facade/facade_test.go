package facade

import "fmt"

func ExampleAPI_TestIt() {
	api := NewAPI()
	fmt.Println(api.TestIt())
	//output:
	//A model is tested
	//B model is tested
}
