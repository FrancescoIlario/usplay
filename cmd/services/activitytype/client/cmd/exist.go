package cmd

import (
	"context"
	"log"

	"github.com/FrancescoIlario/usplay/pkg/services/activitytypecomm"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	cmdExist = &cobra.Command{
		Use:   "exist",
		Short: "Exists an ActivityType",
		Run: func(cmd *cobra.Command, args []string) {
			conn, err := grpc.Dial(target, grpc.WithInsecure())
			if err != nil {
				log.Fatalf("cannot connect to %s: %v", target, err)
			}
			defer conn.Close()

			cli := activitytypecomm.NewActivityTypeSvcClient(conn)
			resp, err := cli.Exist(context.TODO(), &activitytypecomm.ExistActivityTypeRequest{Id: id})
			if err != nil {
				log.Fatalf("error calling exist: %v", err)
			}

			if resp.Exists {
				log.Printf("ActivityType with id %v exist", id)
			} else {
				log.Printf("ActivityType with id %v does not exists", id)
			}
		},
	}
)

func init() {
	cmdExist.Flags().StringVarP(&id, "id", "i", "", "ActivityType's id")
}
