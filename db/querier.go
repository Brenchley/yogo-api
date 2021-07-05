package yogo

import (
	"context"
)

type Querier interface {
	CreateInterest(ctx context.Context, arg CreateInterestParams) (Interest, error)
	CreatePlace(ctx context.Context, arg CreatePlaceParams) (Place, error)
	CreateTrip(ctx context.Context, arg CreateTripParams) (Trip, error)
	CreateTripMember(ctx context.Context, arg CreateTripMemberParams) (TripMember, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetInterest(ctx context.Context, id int32) (Interest, error)
	GetTrip(ctx context.Context, id int32) (Trip, error)
	GetUser(ctx context.Context, id int64) (User, error)
	ListInterests(ctx context.Context) ([]Interest, error)
	ListTripMembers(ctx context.Context, tripID int32) ([]TripMember, error)
	ListTrips(ctx context.Context) ([]Trip, error)
	ListUsers(ctx context.Context) ([]User, error)
}

var _ Querier = (*Queries)(nil)
