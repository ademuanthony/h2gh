package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/ademuanthony/h2gh/models"
	"errors"
	"time"
	"strconv"
	"github.com/ademuanthony/h2gh/utilities"
)

type Peer2PeerService struct {
	O orm.Ormer
}

// Create a payment to be made by the From member ID
func (this Peer2PeerService) CreatePayment(fromMemberId int64) error {
	// Get the next queue for rebate
	var rebateQueue models.Queue
	err := this.O.QueryTable(new(models.Queue)).Filter("amount", 15000).Filter("status", models.StatusPending).RelatedSel().OrderBy("sort_order").One(&rebateQueue)
	if err != nil{
		return errors.New("No one is on queue to receive payment")
	}
	var payment models.Payment

	fromMember := new(models.Member)
	fromMember.Id = fromMemberId

	toMember := new(models.Member)
	toMember.Id = rebateQueue.Member.Id

	payment.Amount = rebateQueue.Amount
	payment.Status = models.StatusPending
	payment.Description = "Membership Rebate"
	payment.FromMember = fromMember
	payment.ToMember = toMember
	payment.CreatedAt = time.Now()
	payment.PenaltyTime = time.Now().AddDate(0, 0, 1)
	payment.QueueId = rebateQueue.Id

	this.O.Insert(&payment)
	rebateQueue.Status = models.StatusPaired
	this.O.Update(&rebateQueue)

	var bonusQueue models.Queue
	err = this.O.QueryTable(new(models.Queue)).Filter("amount", 5000).Filter("status", models.StatusPending).RelatedSel().OrderBy("sort_order").One(&bonusQueue)
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

	return nil
}

func (this Peer2PeerService) QueueUserForPayment(memberId int64, amount float64) error {

	if ok := this.O.QueryTable(new(models.Member)).Filter("id", memberId).Exist(); !ok{
		return errors.New("Invalid Member Id")
	}

	if amount != 5000 &&
		this.O.QueryTable(new(models.Queue)).Filter("member_id", memberId).Filter("amount", amount).Filter("status", models.StatusPending).Exist(){
		return errors.New(utilities.ErrorUserAlreadyInQueue)
	}

	var params []orm.Params
	this.O.Raw("SELECT MAX(sort_order) AS sort_order FROM queue WHERE amount = ?", amount).Values(&params)

	maxSortOrder, _ := strconv.ParseInt((params[0]["sort_order"]).(string), 10, 64)

	member := new(models.Member)
	member.Id = memberId

	var description string
	if amount == 5000{
		description = "Referall Bonus"
	}else{
		description = "Membership Rebate"
	}

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