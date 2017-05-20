package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"net/http"
	"fmt"
	"github.com/ademuanthony/h2gh/services"
	"github.com/ademuanthony/h2gh/models"
	"time"
	"github.com/ademuanthony/h2gh/utilities"
)

type PaymentController struct {
	beego.Controller
}


func (this *PaymentController) Confirm() {
	o := orm.NewOrm()
	o.Begin()
	accountService := services.AccountService{O:o, Ctx:this.Ctx}
	peer2peerService := services.Peer2PeerService{O:o}
	currentMember, _ := accountService.GetCurrentMember()

	flash := beego.NewFlash()
	flash.Store(&this.Controller)

	paymentId, err := this.GetInt64("paymentId")
	if err != nil{
		panic(err)
		flash.Error("Invalid payment ID")
		this.Redirect("/dashboard", http.StatusTemporaryRedirect)
		return
	}

	var payment models.Payment
	err = o.QueryTable(new(models.Payment)).Filter("id", paymentId).RelatedSel().One(&payment)
	if err != nil{
		flash.Error("Payment not found")
		this.Redirect("/dashboard", http.StatusTemporaryRedirect)
		return
	}

	if payment.ToMember.Id != currentMember.Id{
		flash.Error("You can only confirm your own payments")
		this.Redirect("/dashboard", http.StatusTemporaryRedirect)
		return
	}

	payment.Status = models.StatusConfirmed
	_, err = o.Update(&payment)
	if err != nil{
		panic(err)
		o.Rollback()
		flash.Error("Something want wrong. Please try again")
		this.Redirect("/dashboard", http.StatusTemporaryRedirect)
		return
	}

	//mark the queue a paid out
	if queue, err := peer2peerService.GetQueueById(payment.QueueId); err == nil{
		queue.Status = models.StatusPaidOut
		queue.UpdatedAt = time.Now()
		o.Update(&queue)
	}

	//if the member have cleared all his payment
	if !o.QueryTable(new(models.Payment)).Filter("from_member_id", payment.FromMember.Id).Filter("status", models.StatusPending).Exclude("id", paymentId).Exist(){
		//queue the user that made this payment to receive payment
		err = peer2peerService.QueueUserForPayment(payment.FromMember.Id, utilities.UnitRebateAmount)
		if err == nil{
			err = peer2peerService.QueueUserForPayment(payment.FromMember.Id, utilities.UnitRebateAmount)
		}
		if err != nil && err.Error() != utilities.ErrorUserAlreadyInQueue{
			panic(err)
			o.Rollback()
			flash.Error("Something want wrong. Please try again")
			this.Redirect("/dashboard", http.StatusTemporaryRedirect)
			return
		}

		// queue his referrer for bonus
		if payment.FromMember.ReferralId > 0{
			err = peer2peerService.QueueUserForPayment(payment.FromMember.ReferralId, utilities.ReferralBonusAmount)
			if err != nil && err.Error() != utilities.ErrorUserAlreadyInQueue{
				panic(err)
				o.Rollback()
				flash.Error("Something want wrong. Please try again")
				this.Redirect("/dashboard", http.StatusTemporaryRedirect)
				return
			}
		}
	}

	//if this is a rebate payment
	if payment.Amount == utilities.UnitRebateAmount{
		//and the user have no pending rebate, create another payment for him
		if !o.QueryTable(new(models.Payment)).Filter("to_member_id", currentMember.Id).Filter("status", models.StatusPending).Exist(){
			peer2peerService.CreatePayment(currentMember.Id)
		}
		// requeue the current member for another rebate
		/*if err := peer2peerService.QueueUserForPayment(currentMember.Id, utilities.UnitRebateAmount); err != nil && err.Error() != utilities.ErrorUserAlreadyInQueue{
			panic(err)
			o.Rollback()
			flash.Error("Something want wrong. Please try again")
			this.Redirect("/dashboard", http.StatusTemporaryRedirect)
			return
		}*/
	}



	o.Commit()
	fmt.Println("Thank you for confirming the payment. You have been requeued for another payment. Please wait while you matched again")

	this.Redirect("/dashboard", http.StatusTemporaryRedirect)
}
