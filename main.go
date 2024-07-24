/*
package main

import (
	"fmt"
	"net/http"
)

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/world", world)

	fmt.Println("Server is starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

*/

/*
package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/Hello-World", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World Printed"))
	})

	http.ListenAndServe(":8080", nil)
}
*/
/*
package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Goodbye")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/goodbye", goodbye)

	fmt.Println("Server is starting on port 8080...")
	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	fmt.Println("Failed to start server:", err)
	// }

	http.ListenAndServe(":8080", nil)
}

*/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World")

}

func sum(w http.ResponseWriter, r *http.Request) {
	x, y := 10, 20

	sum_of_values := x + y

	fmt.Fprintf(w, "Sum of 10 and 20 is %d", sum_of_values)
}

func main() {
	fmt.Println("Server is runining on port 8080")

	http.Handle("/",
		http.HandlerFunc(
			index,
		),
	)

	http.HandleFunc("/sum_of_numbers", sum)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
