package cloud

import (
	"context"

	billing "cloud.google.com/go/billing/budgets/apiv1"
	billingpb "google.golang.org/genproto/googleapis/cloud/billing/budgets/v1"
)

func GetGCPBillingData() (*billingpb.Budget, error) {
	ctx := context.Background()
	client, err := billing.NewBudgetClient(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	req := &billingpb.GetBudgetRequest{Name: "projects/your-project/budgets/your-budget-id"}
	resp, err := client.GetBudget(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
