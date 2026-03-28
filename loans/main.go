package main

import "fmt"
import "net/http"

func main() {
	var loan = []string{"LN0000", "Andrew", "500"}
	var repayments = []string{"100", "100"}
	var amount int = 500
	var repaymentsInt = []int{100, 100}
	var paid int = 0
	var i int
	for i = 0;i < len(repaymentsInt);i++ {
		paid = paid + repaymentsInt[i]
	}
	var balance int = amount - paid

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
<head>
        <meta name="viewport" content="width=device-widt
h,initial-scale=1.0"/>
</head>
<body style="font-family: Arial">
        <h1>%s</h1>
	<p>Name: %s</p>
	<p>Amount: $%s</p>
	<p>Paid: $%d</p>
	<p>Balance: $%d</p>
	<h2>Repayments</h2>
	<p>$%s</p>
	<p>$%s</p>
</body>
`, loan[0], loan[1], loan[2], paid, balance, repayments[0], repayments[1])
})
	fmt.Println("Running on localhost:8080")
	http.ListenAndServe(":8080", nil)
		
}
