package octane

import (
	"context"
	"net/http"
	"time"
	"github.com/getoctane/octane-go/internal/swagger"
)

type (
	CreateCreditGrantArgs struct {
		// The date at which the grant is effective
		EffectiveAt *time.Time `json:"effective_at,omitempty"`
		// Number of credits to grant
		Amount float64 `json:"amount"`
		// The date at which this grant expires
		ExpiresAt *time.Time `json:"expires_at,omitempty"`
		// Name of the customer receving the grant
		CustomerName string `json:"customer_name"`
		// Optional description. This is only viewable internally
		Description string `json:"description,omitempty"`
		// Total price paid for the credits in cents. Defaults to $1 (100 cents) per credit if not specified
		Price int32 `json:"price,omitempty"`
	}

	creditsAPI struct {
		impl *swagger.APIClient
		ctx  func() context.Context
	}

	CreditGrant struct {
		// Number of credits granted
		Amount float64 `json:"amount,omitempty"`
		// The date at which this grant is effective
		EffectiveAt* time.Time `json:"effective_at,omitempty"`
		// A unique identifier for this grant
		Uuid string `json:"uuid,omitempty"`
		// The date at which this grant expires
		ExpiresAt *time.Time `json:"expires_at,omitempty"`
		// Name of the customer who received the grant
		CustomerName string `json:"customer_name,omitempty"`
		// The source of the grant.
		Source string `json:"source,omitempty"`
		// Optional description. This is only viewable internally
		Description string `json:"description,omitempty"`
		// Total price paid for the credits, in cents
		Price int32 `json:"price,omitempty"`
	}
)

func newCreditsAPI(impl *swagger.APIClient, ctx func() context.Context) *creditsAPI {
	return &creditsAPI{impl: impl, ctx: ctx}
}

// Create creates a new credit grant.
func (api *creditsAPI) Create(body CreateCreditGrantArgs) (CreditGrant, *http.Response, error) {
	implCreateCredits := swagger.CreateCreditGrantArgs{
		Amount:                        body.Amount,
		EffectiveAt:                   body.EffectiveAt,
		ExpiresAt:                	   body.ExpiresAt,
		CustomerName:                  body.CustomerName,
		Description: 			       body.Description,
		Price:                         body.Price,
	}
	implCreateCredit, resp, err := api.impl.CreditsApi.CreditsGrantPost(
		api.ctx(), implCreateCredits)
	credit := implCreditToCredit(&implCreateCredit)
	return credit, resp, err
}


func implCreditToCredit(implCredit *swagger.CreditGrant) CreditGrant {
	externalCreditGrant := CreditGrant{
		Price: implCredit.Price,
		Description: implCredit.Description,
		Source:implCredit.Source, 
		CustomerName: implCredit.CustomerName,
		Uuid: implCredit.Uuid,
		Amount: implCredit.Amount,
	}
	if len(implCredit.EffectiveAt) > 0 {
		tsString := implCredit.EffectiveAt + "Z"
		time_val, _ := time.Parse(time.RFC3339, tsString)
		externalCreditGrant.EffectiveAt = &time_val
	} 

	if len(implCredit.ExpiresAt) > 0  {
		tsString := implCredit.ExpiresAt + "Z"
		time_val, _ := time.Parse(time.RFC3339, tsString)
		externalCreditGrant.ExpiresAt = &time_val
	}

	return externalCreditGrant
}