I am thinking of going all out and including all the necessary inputs that can make this app work.

Let's look at a loan first. A loan will contain the following fields:
- Loan Number
- Principal
- Interest Rate
- Interest
- Duration
- Terrible

---

I have finished with a minimal example of the home page of the Loan Management System. It works fine using Go. I just used an in memory 2D slice as a data store. Great for demo. I implemented a minimal CSS for visual purposes. I will outline ideas with regards to the design as I go. Next up is the loan page.

For the loan page. I will start by spilling out the detaila of the loan. Loan Number, Name, Amount. Then comes a list of repayments. Each repayment must match with the remaining amount to be paid, that is the Balance. Amount, Balance, Paid. Paid refers to the total of repayments. I will use an in memory 2d slice to hold the repayments.

- Loans
- Repayments
- Applications
- Funds
- Tax (Tax is based on monthly interest)
- Borrowers
- Disbursements
- Loan Officer
- Fund Manager
- Penalties
- Terms
- Contract
- Loan Life cycle (Application > Review(Underwriting) > Approval > Agreement > Disbursement > Repayment > Closed)

Storage??? (Copies of letters, per client we are talking 50MB plus database) Assuming a portfolio of 1000 loans, this means you need storage of about 200GB and then 2 cloud backups, 1 offline back up hard drive
