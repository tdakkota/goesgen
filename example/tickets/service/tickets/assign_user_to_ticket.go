package tickets

import (
	"context"
	"fmt"
	"tickets/auth"
	"tickets/generated"
	"tickets/service/tickets/io"
)

func (s *Service) AssignUserToTicket(
	ctx context.Context,
	tx generated.TransactionReader,
	in io.AssignUserToTicketIn,
) ([]generated.Event, error) {
	client, err := auth.User(ctx)
	if err != nil {
		return nil, err
	}

	if _, ok := tx.(transaction).store.state.tickets[in.Ticket]; !ok {
		err = fmt.Errorf("ticket %s not found", in.Ticket)
	}

	return []generated.Event{
		generated.EventUserAssignedToTicket{
			User:   in.User,
			Ticket: in.Ticket,
			By:     client,
		},
	}, nil
}