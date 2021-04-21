package verifier

import (
	"errors"

	emailverifier "github.com/AfterShip/email-verifier"
	"github.com/markkuit/mailcheck/internal/commons"
)

var emailVerifier *emailverifier.Verifier

func init() {
	emailVerifier = emailverifier.NewVerifier().EnableSMTPCheck()
	emailVerifier.HelloName(commons.HelloName)
	emailVerifier.FromEmail(commons.FromEmail)
}

func Check(address string, c chan<- CheckResult) {
	addressParsed := emailVerifier.ParseAddress(address)
	if addressParsed.Valid {
		if res, err := emailVerifier.CheckSMTP(addressParsed.Domain, addressParsed.Username); err == nil {
			c <- CheckResult{
				Address:     address,
				Deliverable: res.Deliverable,
				FullInbox:   res.FullInbox,
				Disabled:    res.Disabled,
				Error:       nil,
			}
		} else {
			c <- CheckResult{
				Address: address,
				Error:   err,
			}
		}
	} else {
		c <- CheckResult{
			Address: address,
			Error:   errors.New("invalid address"),
		}
	}
}
