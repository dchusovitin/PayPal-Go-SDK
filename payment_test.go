package paypalsdk

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateDirectPaypalPayment(t *testing.T) {
	c, _ := NewClient(testClientID, testSecret, APIBaseSandBox)
	c.SetLog(os.Stdout)
	c.GetAccessToken()

	amount := Amount{
		Total:    "15.11",
		Currency: "USD",
	}

	p, err := c.CreateDirectPaypalPayment(amount, "http://example.com", "http://example.com", "test payment")

	if err != nil || p.ID == "" {
		t.Errorf("Test paypal payment is not created")
	}
}

func TestGetPayment(t *testing.T) {
	c, _ := NewClient(testClientID, testSecret, APIBaseSandBox)
	c.GetAccessToken()

	_, err := c.GetPayment(testPaymentID)

	if err == nil {
		t.Errorf("404 for this payment ID")
	} else {
		fmt.Println(err.Error())
	}
}

func TestGetPayments(t *testing.T) {
	c, _ := NewClient(testClientID, testSecret, APIBaseSandBox)
	c.GetAccessToken()

	_, err := c.GetPayments()

	if err != nil {
		t.Errorf("Nil error expected")
	}
}

func TestExecuteApprovedPayment(t *testing.T) {
	c, _ := NewClient(testClientID, testSecret, APIBaseSandBox)
	c.GetAccessToken()

	_, err := c.ExecuteApprovedPayment(testPaymentID, testPayerID)

	if err == nil {
		t.Errorf("404 for this payment ID")
	} else {
		fmt.Println(err.Error())
	}
}
