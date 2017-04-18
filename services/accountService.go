package services

import (
	"github.com/astaxie/beego/orm"
	"github.com/ademuanthony/h2gh/models"
	"github.com/astaxie/beego/context"
	"errors"
)

type AccountService struct {
	O orm.Ormer
	Ctx *context.Context
}

func (this *AccountService) GetMemberById(id int64) (models.Member, error) {
	var member models.Member
	err := this.O.QueryTable(new(models.Member)).RelatedSel().Filter("id", id).One(&member)
	return member, err
}

func (this *AccountService) GetCurrentMemberId() (id int64, err error) {
	id, ok := (this.Ctx.Input.Session("uid")).(int64)
	if !ok{
		err = errors.New("Cannot fetch corrent user id")
	}
	return id, err
}

func (this *AccountService) GetCurrentMember() (models.Member, error) {
	id, err := this.GetCurrentMemberId()
	var member models.Member
	if err != nil{
		return member, errors.New("Cannot fetch current member ID")
	}
	return this.GetMemberById(id)
}
