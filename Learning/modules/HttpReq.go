package modules

import (
	"fmt"
	"net/http"
	"io"
	"time"
)

func HttpReq() {
	
	// res, err := http.Get("https://jsonplaceholder.typicode.com/posts/1");

	// if err!=nil {
	// 	fmt.Println("Error in fetching data", err);
	// 	return;
	// }
	// defer res.Body.Close();
	// data, err := io.ReadAll(res.Body);

	// if err!=nil {
	// 	fmt.Println("Error in reading data", err);
	// 	return;
	// }
	// fmt.Println(string(data))
	client:= &http.Client{
		Timeout: 50 * time.Millisecond,	
	}
	resp, err := client.Get("https://jsonplaceholder.typicode.com/posts");

	if err != nil {
		fmt.Println("Error in fetching data", err);
		return;
	}
	defer resp.Body.Close();

	data, err := io.ReadAll(resp.Body);

	if err != nil {
		fmt.Println("Error in reading data", err);
		return;
	}
	fmt.Println("Response is received", string(data));
}