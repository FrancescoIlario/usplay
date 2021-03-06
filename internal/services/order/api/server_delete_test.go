package api_test

import (
	"context"
	"testing"

	"github.com/FrancescoIlario/usplay/internal/services/order/api"
	"github.com/FrancescoIlario/usplay/pkg/services/ordercomm"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Test_DeleteHappyPath(t *testing.T) {
	// arrange
	store := &orderTestRepo{
		DeleteResult: struct{ Err error }{Err: nil},
	}
	svr := api.NewOrderServer(store)
	ctx := context.Background()

	// act
	_, err := svr.Delete(ctx, &ordercomm.DeleteOrderRequest{Id: uuid.New().String()})

	// assert
	if err != nil {
		t.Fatalf("error invoking delete: %v", err)
	}
}

func Test_DeleteInvalidId(t *testing.T) {
	// arrange
	store := &orderTestRepo{
		DeleteResult: struct{ Err error }{Err: nil},
	}
	svr := api.NewOrderServer(store)
	ctx := context.Background()

	// act
	_, err := svr.Delete(ctx, &ordercomm.DeleteOrderRequest{Id: "88888128312319273190"})

	// assert
	if err == nil {
		t.Fatalf("expected error invoking delete with empty id, but none received")
	}

	sts := status.Convert(err)
	if sts == nil {
		t.Fatalf("error is not a status.Status error")
	}

	if expected := codes.InvalidArgument; sts.Code() != expected {
		t.Errorf("expected status code %s, provided code %s", expected, sts.Code())
	}
}

func Test_DeleteEmptyId(t *testing.T) {
	// arrange
	store := &orderTestRepo{
		DeleteResult: struct{ Err error }{Err: nil},
	}
	svr := api.NewOrderServer(store)
	ctx := context.Background()

	// act
	_, err := svr.Delete(ctx, &ordercomm.DeleteOrderRequest{Id: ""})

	// assert
	if err == nil {
		t.Fatalf("expected error invoking delete with empty id, but none received")
	}

	sts := status.Convert(err)
	if sts == nil {
		t.Fatalf("error is not a status.Status error")
	}

	if expected := codes.InvalidArgument; sts.Code() != expected {
		t.Errorf("expected status code %s, provided code %s", expected, sts.Code())
	}
}

func Test_DeleteNilId(t *testing.T) {
	// arrange
	store := &orderTestRepo{
		DeleteResult: struct{ Err error }{Err: nil},
	}
	svr := api.NewOrderServer(store)
	ctx := context.Background()

	// act
	_, err := svr.Delete(ctx, &ordercomm.DeleteOrderRequest{Id: uuid.Nil.String()})

	// assert
	if err == nil {
		t.Fatalf("expected error invoking delete with nil id, but none received")
	}

	sts := status.Convert(err)
	if sts == nil {
		t.Fatalf("error is not a status.Status error")
	}

	if expected := codes.InvalidArgument; sts.Code() != expected {
		t.Errorf("expected status code %s, provided code %s", expected, sts.Code())
	}
}
