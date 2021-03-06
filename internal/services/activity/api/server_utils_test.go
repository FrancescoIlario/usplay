package api_test

import (
	"context"
	"time"

	"github.com/FrancescoIlario/usplay/internal/services/activity/storage"
	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Activity Test repository
type activityTestRepo struct {
	CreateResult struct {
		ID  *uuid.UUID
		Err error
	}
	DeleteResult struct {
		Err error
	}
	ExistResult struct {
		Exists *bool
		Err    error
	}
	ReadResult struct {
		Activity *storage.Activity
		Err      error
	}
	ListResult struct {
		Activities []storage.Activity
		Err        error
	}
	ListInIntervalResult struct {
		Activities []storage.Activity
		Err        error
	}
	UpdateResult struct {
		Err error
	}
}

func (r *activityTestRepo) Create(context.Context, storage.Activity) (*uuid.UUID, error) {
	return r.CreateResult.ID, r.CreateResult.Err
}

func (r *activityTestRepo) Read(ctx context.Context, id uuid.UUID) (*storage.Activity, error) {
	activity := r.ReadResult.Activity
	activity.ID = id.String()
	return activity, r.ReadResult.Err
}

func (r *activityTestRepo) Exist(ctx context.Context, id uuid.UUID) (*bool, error) {
	return r.ExistResult.Exists, r.ReadResult.Err
}

func (r *activityTestRepo) Update(context.Context, storage.Activity) error {
	return r.UpdateResult.Err
}

func (r *activityTestRepo) Delete(context.Context, uuid.UUID) error {
	return r.DeleteResult.Err
}

func (r *activityTestRepo) List(context.Context, []uuid.UUID) (storage.Activities, error) {
	return r.ListResult.Activities, r.ListResult.Err
}

func (r *activityTestRepo) ListInInterval(context.Context, storage.Interval) (storage.Activities, error) {
	return r.ListInIntervalResult.Activities, r.ListInIntervalResult.Err
}

// ActivityType Test Client
type actTestClient struct {
	WaitTime   time.Duration
	ReadResult struct {
		Err   error
		Reply activitytypecomm.ReadActivityTypeReply
	}
	ListResult struct {
		Err   error
		Reply activitytypecomm.ListActivityTypesReply
	}
	ExistResult struct {
		Err   error
		Reply activitytypecomm.ExistActivityTypeReply
	}
}

func (c *actTestClient) Create(ctx context.Context, in *activitytypecomm.CreateActivityTypeRequest, opts ...grpc.CallOption) (*activitytypecomm.CreateActivityTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "Create method is not implemented")
}

// Exists an activity
func (c *actTestClient) Exist(ctx context.Context, in *activitytypecomm.ExistActivityTypeRequest, opts ...grpc.CallOption) (*activitytypecomm.ExistActivityTypeReply, error) {
	return &c.ExistResult.Reply, c.ExistResult.Err
}

// Reads an ActivityType
func (c *actTestClient) Read(ctx context.Context, in *activitytypecomm.ReadActivityTypeRequest, opts ...grpc.CallOption) (*activitytypecomm.ReadActivityTypeReply, error) {
	time.Sleep(c.WaitTime)

	reply := c.ReadResult.Reply
	if reply.GetActivityType() != nil {
		reply.ActivityType.Id = in.Id
	}
	return &reply, c.ReadResult.Err
}

// Delete an ActivityType
func (c *actTestClient) Delete(ctx context.Context, in *activitytypecomm.DeleteActivityTypeRequest, opts ...grpc.CallOption) (*activitytypecomm.DeleteActivityTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "Delete method is not implemented")
}

// Update an ActivityType
func (c *actTestClient) Update(ctx context.Context, in *activitytypecomm.UpdateActivityTypeRequest, opts ...grpc.CallOption) (*activitytypecomm.UpdateActivityTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "Update method is not implemented")
}

// List an ActivityType
func (c *actTestClient) List(ctx context.Context, in *activitytypecomm.ListActivityTypesRequest, opts ...grpc.CallOption) (*activitytypecomm.ListActivityTypesReply, error) {
	return &c.ListResult.Reply, c.ListResult.Err
}

// Order Test Client
type orderTestClient struct {
	WaitTime   time.Duration
	ReadResult struct {
		Err   error
		Reply ordercomm.ReadOrderReply
	}
	ListResult struct {
		Err   error
		Reply ordercomm.ListOrdersReply
	}
	ExistResult struct {
		Err   error
		Reply ordercomm.ExistOrderReply
	}
}

func (c *orderTestClient) Create(ctx context.Context, in *ordercomm.CreateOrderRequest, opts ...grpc.CallOption) (*ordercomm.CreateOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "Create method is not implemented")
}

// Exists an activity
func (c *orderTestClient) Exist(ctx context.Context, in *ordercomm.ExistOrderRequest, opts ...grpc.CallOption) (*ordercomm.ExistOrderReply, error) {
	return &c.ExistResult.Reply, c.ExistResult.Err
}

// Reads an Order
func (c *orderTestClient) Read(ctx context.Context, in *ordercomm.ReadOrderRequest, opts ...grpc.CallOption) (*ordercomm.ReadOrderReply, error) {
	time.Sleep(c.WaitTime)

	reply := c.ReadResult.Reply
	if reply.GetOrder() != nil {
		reply.Order.Id = in.Id
	}
	return &reply, c.ReadResult.Err
}

// Delete an Order
func (c *orderTestClient) Delete(ctx context.Context, in *ordercomm.DeleteOrderRequest, opts ...grpc.CallOption) (*ordercomm.DeleteOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "Delete method is not implemented")
}

// Update an Order
func (c *orderTestClient) Update(ctx context.Context, in *ordercomm.UpdateOrderRequest, opts ...grpc.CallOption) (*ordercomm.UpdateOrderReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "Update method is not implemented")
}

// List an Order
func (c *orderTestClient) List(ctx context.Context, in *ordercomm.ListOrdersRequest, opts ...grpc.CallOption) (*ordercomm.ListOrdersReply, error) {
	return &c.ListResult.Reply, c.ListResult.Err
}
