package store_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/tatucosmin/hotel-system/fixtures"
	"github.com/tatucosmin/hotel-system/store"
)

func TestTicketStore(t *testing.T) {
	env := fixtures.NewTestEnv(t)

	cleanup := env.SetupDb(t)

	ctx := context.Background()

	t.Cleanup(func() {
		cleanup(t)
	})

	userStore := store.NewUserStore(env.Db)
	ticketStore := store.NewTicketStore(env.Db)

	user, err := userStore.CreateUser(ctx, "test@test.com", "test")
	require.NoError(t, err)

	ticket, err := ticketStore.Create(ctx, "test ticket", "test description", user.Id, store.TicketPriorityUrgent)
	require.NoError(t, err)

	require.NotNil(t, ticket.Id)
	now := time.Now()

	require.Equal(t, "test ticket", ticket.Title)
	require.Equal(t, user.Id, ticket.Creator)
	require.True(t, now.After(ticket.CreatedAt))

	require.NoError(t, err)

	err = ticketStore.Update(ctx, ticket.Id, store.TicketPriorityUrgent, store.TicketStatusCreated)
	require.NoError(t, err)

	ticket, err = ticketStore.ById(ctx, ticket.Id)
	require.NoError(t, err)

	err = ticketStore.Delete(ctx, ticket.Id)
	require.NoError(t, err)

}
