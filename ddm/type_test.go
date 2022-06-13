package ddm

import (
	"encoding/json"
	"testing"
)

type message struct {
	Name      IDName   `json:"name"`
	Mobile    Mobile   `json:"mobile"`
	IDCard    IDCard   `json:"id_card"`
	PassWord  PassWord `json:"password"`
	Email     Email    `json:"email"`
	BankCard1 BankCard `json:"bank_card_1"`
	BankCard2 BankCard `json:"bank_card_2"`
	BankCard3 BankCard `json:"bank_card_3"`
}

func TestMarshalJSON(t *testing.T) {
	msg := new(message)
	msg.Name = IDName("李鸿章")
	msg.Mobile = Mobile("15288887986")
	msg.IDCard = IDCard("341324199001011234")
	msg.PassWord = PassWord("123456")
	msg.Email = Email("tuonijix@163.com")
	msg.BankCard1 = BankCard("53535353535353535")
	msg.BankCard2 = BankCard("864675364674256565")
	msg.BankCard3 = BankCard("145425624562466224")

	marshal, _ := json.Marshal(msg)
	t.Log(string(marshal))
}
