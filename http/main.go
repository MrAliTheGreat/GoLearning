package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://www.google.com/")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Copy will take a writer and reader and would write any data that it gets from reader src to writer dst
	io.Copy(os.Stdout, resp.Body)

}

// When we put an interface as the type of a field in a struct it means that the field can be of any type as long as that type satisfies the mentioned interface
/*
	In Go we can make an interface of number of other interfaces meaning that fields in an interface can also be an interface
	What this will do is that if one wants to satisfy the main interface it must satisfy the each of the sub-interfaces as well
*/
/*
	What reader interface does is basically an abstraction.
	Since there are many sources of input with different types, reader interface will convert each of them to []byte so we can use it easily
	So, each type will implement the func in reader interface in a way that converts the raw data to []byte. Abstraction baby! Cool ngl!
	The read func in reader interface will get a []byte as input. This []byte is exactly the output of read like the func will convert the raw data to []byte and put it in this slice that we pass to it
	It is exactly like sending a pointer and saying that whatever you do put the final output in this address
	The read func will return an int and an error. The int is for showing how many bytes were put in the provided []byte
	The important thing for us is the []byte since the data we want is in it after calling the func with it

	The read func in reader interface will write in the []byte until it is full
	So if we pass a element with 0 empty elements the func will think this []byte is full so I'm not writing in it
	So we have to initialize the []byte with enough empty elements so that the whole data can fit into it
	The read func can not change the size of the slice!
*/
// Whenever we see a slice of bytes we can think of slice of strings so the convertion is easy just wrap string() around the slice of bytes
