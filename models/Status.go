package models

const (
	StatusActive = "Active"
	StatusInActive = "In Active"
	StatusRemoved = "Removed"
	StatusInVisible = "In Visible"
	StatusPending = "Pending" //Your package have not been paired by the system
	StatusPaired = "Paired" //Your package have been paired but you have not made payment
	StatusAwaitingAlert = "Awaiting Alert" //You told your beneficiary that you have paid but he said he is waiting for alert
	StatusConfirm = "Confirm" //Your beneficiary have confirmed your payment
	StatusRejected = "Rejected" //Your payment claim your rejected

	StatusTicketsCreated = "Tickets Created"
	StatusAwaitingPayment = "Awaiting Payment" //the user have been paired, not all have paid
	StatusPaidOut = "Paid Out" //you have been paired and paid
	StatusUrgentPairingNeeded = "Urgent Pairing Needed"
)

/*

 */