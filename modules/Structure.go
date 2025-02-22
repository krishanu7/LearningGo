package modules

import "fmt"
type Country struct {
    Name string
    Code string
}

type Address struct {
    City string
    State string
    Zipcode int
    Country Country
}

type Contact struct {
    Email string
    Phone string
}

type Person struct {
    Name string
    Age int
    Address Address
    Contact Contact
}

func Structure() {
    // p := Person{
    //     Name: "Krishanu",
    //     Age:  22,
    //     Address: Address{
    //         City:    "Kolkata",
    //         State:   "West Bengal",
    //         Zipcode: 700001,
    //         Country: Country{
    //             Name: "India",
    //             Code: "IN",
    //         },
    //     },
    // }
    person := Person{
        Name: "Krishanu",
        Age: 22,
        Address: Address{
            City: "Kolkata",
            State: "West Bengal",
            Zipcode: 700001,
            Country: Country{
                Name: "India",
                Code: "IN",
            },
        },
        Contact: Contact{
            Email: "krishanu@gmail.com",
            Phone: "1234567890",
        },
    }

    fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
    fmt.Println("City:", person.Address.City)
    fmt.Println("Country:", person.Address.Country.Name)
    fmt.Println("Country Code:", person.Address.Country.Code)
    fmt.Println("Email:", person.Contact.Email)
    fmt.Println("Phone:", person.Contact.Phone)
}
