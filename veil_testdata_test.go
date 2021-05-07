package veil_test

// getInputData prepares the test dataset.
func getInputData() []interface{} {

	var data []interface{}

	// string
	data = append(data, "hello 123")

	// struct
	data = append(data, struct {
		Name  string
		Phone string
		ph    string
	}{"Test", "123", "123"})

	// map
	data = append(data, map[string]int{
		"phone": 123,
	})

	return data
}

func getStructData() []interface{} {

	var data []interface{}

	type User struct {
		User  string
		Pwd   string `veil:"obscure"`
		group float32
	}

	type Data struct {
		Name  string
		Phone string
		ph    int
		User  User
	}

	// input data
	d := Data{
		Name:  "Test",
		Phone: "123",
		ph:    123,
		User: User{
			User:  "User",
			Pwd:   "Password",
			group: 1.01,
		},
	}

	data = append(data, d)

	return data
}
