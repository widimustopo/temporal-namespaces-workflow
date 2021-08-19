package activities

import (
	"context"
	"fmt"

	"github.com/widimustopo/temporal-namespaces-module/entities"
	"go.temporal.io/sdk/client"

	"gorm.io/gorm"
)

type Activities struct {
	*gorm.DB
	client.Client
}

func HandlerActivities(db *gorm.DB, temporalClient client.Client) *Activities {
	return &Activities{
		db,
		temporalClient,
	}
}

func (c Activities) Register(ctx context.Context, req *entities.TemporalRequest) (interface{}, error) {

	fmt.Println("ini datanya kan : ", req)

	return fmt.Sprintf("ini datanya kan : %v", req), nil
}
