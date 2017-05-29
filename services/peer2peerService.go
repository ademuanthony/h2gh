package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/ademuanthony/h2gh/models"
	"errors"
	"time"
	"strconv"
	"github.com/ademuanthony/h2gh/utilities"
	"fmt"
)

func init() {
	ticker := time.NewTicker(time.Hour * 1)
	// Start a goroutine that removes expired payments
	p2pService := Peer2PeerService{O:orm.NewOrm()}
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
			p2pService.RemoveExpiredPayments()
		}
	}()
}

type Peer2PeerService struct {
	O orm.Ormer
}

// Create a payment to be made by the From member ID
func (this Peer2PeerService) CreatePayment(fromMemberId int64) error {
	// Get the next queue for rebate
	var queue models.Queue

	//get bonus else get rebate
	err := this.O.QueryTable(new(models.Queue)).Filter("description", utilities.ReferralBonusDescription).
		Filter("status", models.StatusPending).RelatedSel().OrderBy("sort_order").One(&queue)

	if err != nil{
		err = this.O.QueryTable(new(models.Queue)).Filter("description", utilities.UnitRebateDescription).
			Filter("status", models.StatusPending).RelatedSel().OrderBy("sort_order").One(&queue)
		if err != nil{
			return errors.New("No one is on queue to receive payment")
		}
	}

	var payment models.Payment

	fromMember := new(models.Member)
	fromMember.Id = fromMemberId

	toMember := new(models.Member)
	toMember.Id = queue.Member.Id

	payment.Amount = queue.Amount
	payment.Status = models.StatusPending
	payment.Description = "Membership Rebate"
	payment.FromMember = fromMember
	payment.ToMember = toMember
	payment.CreatedAt = time.Now()
	payment.PenaltyTime = time.Now().AddDate(0, 0, 1)
	payment.QueueId = queue.Id

	this.O.Insert(&payment)
	queue.Status = models.StatusPaired
	this.O.Update(&queue)

	/*
	var bonusQueue models.Queue

	// Get the next queue for an active member
	err = this.O.QueryTable(new(models.Queue)).Filter("amount", utilities.ReferralBonusAmount).Filter("Member__status", models.StatusActive).Filter("status", models.StatusPending).RelatedSel().OrderBy("sort_order").One(&bonusQueue)
	if err != nil{
		return errors.New("Please tell admin about this. There is no bonus payment on the current queue")
	}

	var bonusPayment = models.Payment{
		Amount:bonusQueue.Amount,
		Status:models.StatusPending,
		Description:"Referral Bonus",
		FromMember:fromMember,
		ToMember:bonusQueue.Member,
		CreatedAt: time.Now(),
		PenaltyTime:time.Now().AddDate(0, 0, 1),
		QueueId:bonusQueue.Id,
	}

	this.O.Insert(&bonusPayment)
	bonusQueue.Status = models.StatusPaired
	this.O.Update(&bonusQueue)
*/
	this.O.QueryTable(new(models.Member)).Filter("id", fromMember.Id).RelatedSel().One(fromMember)
	this.O.QueryTable(new(models.Member)).Filter("id", toMember.Id).RelatedSel().One(toMember)

	//send message to the donor and receiver
	messageBody := fmt.Sprintf("Dear %s, <br/>" +
		"<p>You have been paired to pay %s the sum of %f from help2gethelp.com</p>" +
		"<p>Below is the beneficiary details<br/>" +
		"Name: %s %s<br/>" +
		"Phone Number: %s<br/>" +
		"Bank Name: %s" +
		"Account Name: %s" +
		"Account Number: %s</p>" +
		"<p>Please make the payment before 24 hours to avoid account termination</p>",
		fromMember.LastName, toMember.LastName, payment.Amount, toMember.LastName, toMember.FirstName,
		toMember.PhoneNumber, toMember.Bank.Name, toMember.AccountName, toMember.AccountNumber,
	)

	messageService := MessageService{}
	go messageService.SendEmail("info@help2gethelp.com", fromMember.Email, "Account Paired", messageBody, "text/html")

	messageBody = fmt.Sprintf("Dear %s, <br/>" +
		"%s has been asked to pay you the sum of %f from help2gethelp.com. His phone number is %s", toMember.LastName, fromMember.LastName,
		payment.Amount, fromMember.PhoneNumber,
	)
	go messageService.SendEmail("info@help2gethelp.com", toMember.Email, "Account Matched", messageBody, "text/html")

	return nil
}

func (this *Peer2PeerService) RemoveExpiredPayments()  {
	// todo use time ticker to make this hourly
	//var ticker time.Ticker

	var expiredPayments []models.Payment
	minDate := time.Now().AddDate(0, 0, -1) // One day ago

	this.O.QueryTable(new(models.Payment)).Filter("status", models.StatusPending).Filter("created_at__lt", minDate).All(&expiredPayments)

	accountService := AccountService{O:this.O}

	for _, payment := range expiredPayments{
		member, err := accountService.GetMemberById(payment.FromMember.Id)
		if err != nil{
			member.Status = models.StatusInActive
			this.O.Update(&member)
		}
		queue, err := this.GetQueueById(payment.QueueId)
		if err != nil{
			queue.Status = models.StatusPending
			this.O.Update(&queue)
		}
		//delete the payment
		this.O.Delete(&payment)
	}

	// Loop continuously
	//this.RemoveExpiredPayments()
}

func (this Peer2PeerService) QueueUserForPayment(memberId int64, amount float64, description string) error {

	if ok := this.O.QueryTable(new(models.Member)).Filter("id", memberId).Exist(); !ok{
		return errors.New("Invalid Member Id")
	}

	qCount, _ := this.O.QueryTable(new(models.Queue)).Filter("member_id", memberId).Filter("amount", amount).
		Filter("status", models.StatusPending).Filter("description", description).Count()

	if description != utilities.ReferralBonusDescription && qCount >= 2{
		return errors.New(utilities.ErrorUserAlreadyInQueue)
	}

	var params []orm.Params
	this.O.Raw("SELECT MAX(sort_order) AS sort_order FROM queue WHERE description = ?", description).Values(&params)

	maxSortOrder, _ := strconv.ParseInt((params[0]["sort_order"]).(string), 10, 64)

	member := new(models.Member)
	member.Id = memberId


	var queue = models.Queue{
		Amount:amount, Member:member,
		Description:description, Status:models.StatusPending,
		SortOrder:maxSortOrder+1,
		CreatedAt:time.Now(),
	}

	_, err := this.O.Insert(&queue)

	return err

}

func (this Peer2PeerService) GetQueueById(id int64) (models.Queue, error) {
	var queue models.Queue
	err := this.O.QueryTable(new(models.Queue)).Filter("id", id).One(&queue)
	return queue, err
}