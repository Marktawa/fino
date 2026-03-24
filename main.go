package main

import "fmt"
import "net/http"
import "strconv"

func main() {
	var loans = [][]string{
		{"Andrew", "$500"},
		{"Jane", "$1000"},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
<head>
	<meta name="viewport" content="width=device-width,initial-scale=1.0"/>
</head>
<body style="font-family: Arial">
	<h1>Loans</h1>
`)
		for i := 0; i < len(loans); i++ {
			fmt.Fprintf(w, `
	<p>%s - %s
		<a href="/delete?%d">Delete</a>
		<a href="/edit?%d">Edit</a>
	</p>
`, loans[i][0], loans[i][1], i, i)
		}
		fmt.Fprintf(w, `
	<form method="GET" action="/add">
		<label>Name:</label>
		<input type="text" name="q" />
		<label>Amount:</label>
		<input type="text" name="s" />
		<input type="submit" value="Add Loan" />
	</form>
</body>
`)
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		var q string = r.FormValue("q")
		var s string = r.FormValue("s")
		loans = append(loans, []string{q, s})
		fmt.Fprintf(w, `
<head>
	<meta name="viewport" content="width=device-width,initial-scale=1.0"/>
</head>
<body>
	<h1>Loan added successfully</h1>
	<p>%s - %s</p>
	<a href="/">Home</a>
</body>
`, q, s)
	})

	http.HandleFunc("/delete", func(w http.ResponseWriter, r *http.Request) {
		for key := range r.URL.Query() {
			index, _ := strconv.Atoi(key)
			loans = append(loans[:index], loans[index+1:]...)
		}
		fmt.Fprintf(w, `
<head>
	<meta name="viewport" content="width=device-width,initial-scale=1.0"/>
</head>
<body>
	<h1>Loan deleted successfully</h1>
	<a href="/">Home</a>
</body>
`)
	})

	http.HandleFunc("/edit", func(w http.ResponseWriter, r *http.Request) {
		for key := range r.URL.Query() {
			index, _ := strconv.Atoi(key)
			loan := loans[index]
			fmt.Fprintf(w, `
<head>
	<meta name="viewport" content="width=device-width,initial-scale=1.0"/>
</head>
<body style="font-family: Arial">
	<h1>Edit Loan</h1>
	<form method="GET" action="/update">
		<input type="hidden" name="index" value="%d"/>
		<label>Name:</label>
		<input type="text" name="q" value="%s"/>
		<label>Amount:</label>
		<input type="text" name="s" value="%s"/>
		<input type="submit" value="Update Loan"/>
	</form>
	<a href="/">Home</a>
</body>
`, index, loan[0], loan[1])
		}
	})

	http.HandleFunc("/update", func(w http.ResponseWriter, r *http.Request) {
		index, _ := strconv.Atoi(r.FormValue("index"))
		loans[index][0] = r.FormValue("q")
		loans[index][1] = r.FormValue("s")
		fmt.Fprintf(w, `
<head>
	<meta name="viewport" content="width=device-width,initial-scale=1.0"/>
</head>
<body>
	<h1>Loan updated successfully</h1>
	<p>%s - %s</p>
	<a href="/">Home</a>
</body>
`, loans[index][0], loans[index][1])
	})

	fmt.Println("Server running on localhost:8000")
	http.ListenAndServe(":8000", nil)
}
