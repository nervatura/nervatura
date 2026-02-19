package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	cu "github.com/nervatura/component/pkg/util"
	"github.com/nervatura/nervatura/v6/example/go/service"

	//pb "github.com/nervatura/nervatura/v6/pkg/service/grpc/proto"
	pb "github.com/nervatura/nervatura/v6/protos/go"
)

// loadEnvFile reads a .env file and returns a map of key/value pairs
func loadEnvFile(filename string) (cu.SM, error) {
	envMap := make(cu.SM)

	data, err := os.ReadFile(filename)
	if err != nil {
		return envMap, err
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		parts := strings.SplitN(line, "=", 2)
		if line == "" || strings.HasPrefix(line, "#") || len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Remove quotes if present
		value = strings.Trim(value, `"'`)

		envMap[key] = value
	}

	return envMap, nil
}

func main() {
	var config cu.SM
	var err error
	if config, err = loadEnvFile(".env.example"); err != nil {
		fmt.Printf("Error loading environment variables: %v\n", err)
		return
	}

	cli := service.CliClient{Config: config}
	rest := service.HttpClient{Config: config}
	rpc := service.RpcClient{Config: config}

	fnMap := map[string][]func() (string, any, error){
		"cli": {
			func() (string, any, error) {
				result, err := cli.Database(cu.IM{
					"alias": "demo", "demo": true,
				})
				return "Database", result, err
			},
			func() (string, any, error) {
				result, err := cli.Function(cu.IM{
					"name": "test", "values": cu.IM{},
				})
				return "Function", result, err
			},
			func() (string, any, error) {
				result, err := cli.ResetPassword(cu.IM{
					"alias": "demo", "code": "USR0000000000N1",
				})
				return "ResetPassword", result, err
			},
			func() (string, any, error) {
				result, err := cli.Create("customer",
					cu.IM{
						"alias": "demo",
					},
					cu.IM{
						"code": "CLI0000000000N1", "customer_name": "Go Test",
					})
				return "Create", result, err
			},
			func() (string, any, error) {
				result, err := cli.Update("customer",
					cu.IM{
						"alias": "demo", "code": "CLI0000000000N1",
					},
					cu.IM{
						"customer_meta": cu.IM{"account": "1234567890"},
					})
				return "Update", result, err
			},
			func() (string, any, error) {
				result, err := cli.Get("customer",
					cu.IM{
						"alias": "demo", "code": "CLI0000000000N1",
					})
				return "Get", result, err
			},
			func() (string, any, error) {
				result, err := cli.Query("customer",
					cu.IM{
						"alias": "demo", "customer_type": "CUSTOMER_COMPANY",
					})
				return "Query", result, err
			},
			func() (string, any, error) {
				result, err := cli.View(cu.IM{
					"alias": "demo", "name": "VIEW_CUSTOMER_EVENTS", "filters": []any{
						map[string]any{"field": "like_subject", "value": "visit"},
						map[string]any{"field": "place", "value": "City1"},
					}, "limit": 10,
				})
				return "View", result, err
			},
			func() (string, any, error) {
				result, err := cli.Delete("customer",
					cu.IM{
						"alias": "demo", "code": "CLI0000000000N1",
					})
				return "Delete", result, err
			},
		},
		"rest": {
			func() (string, any, error) {
				result, err := rest.Get("", "customer", cu.IM{
					"customer_type": "CUSTOMER_COMPANY",
				})
				return "Get", result, err
			},
			func() (string, any, error) {
				result, err := rest.Post("", "customer", cu.IM{
					"code": "REST0000000000N1", "customer_name": "Go Test",
				})
				return "Post", result, err
			},
			func() (string, any, error) {
				result, err := rest.Put("", "customer/REST0000000000N1", cu.IM{
					"customer_name": "Test Customer",
				})
				return "Put", result, err
			},
			func() (string, any, error) {
				result, err := rest.Delete("", "customer/REST0000000000N1")
				return "Delete", result, err
			},
		},
		"rpc": {
			/*
				func() (string, any, error) {
					result, err := rpc.Database(cu.IM{
						"alias": "demo", "demo": true,
					})
					return "Database", result, err
				},
			*/
			func() (string, any, error) {
				result, err := rpc.CustomerUpdate("", &pb.Customer{
					Code: "RPC0000000000N1", CustomerName: "Go Test",
				})
				return "CustomerUpdate", result, err
			},
			func() (string, any, error) {
				result, err := rpc.CustomerGet("", &pb.RequestGet{
					Code: "RPC0000000000N1",
				})
				return "CustomerGet", result, err
			},
			func() (string, any, error) {
				result, err := rpc.CustomerQuery("", &pb.RequestQuery{
					Filters: []*pb.RequestQueryFilter{
						{
							Field: "customer_type",
							Value: "CUSTOMER_COMPANY",
						},
					},
					Limit: 10,
				})
				return "CustomerQuery", result, err
			},
			func() (string, any, error) {
				result, err := rpc.Delete("", &pb.RequestDelete{
					Code: "RPC0000000000N1", Model: pb.Model_CUSTOMER,
				})
				return "Delete", result, err
			},
			func() (string, any, error) {
				result, err := rpc.Function("", &pb.RequestFunction{
					Function: "product_price", Args: cu.SM{
						"product_code": "PRD0000000000N1", "currency_code": "EUR", "price_type": "PRICE_CUSTOMER",
					},
				})
				return "Function", result, err
			},
			func() (string, any, error) {
				result, err := rpc.View("", &pb.RequestView{
					Name: pb.ViewName_VIEW_CUSTOMER_EVENTS,
					// TODO: change to filters
					Filter: "subject like '%visit%' and place='City1'",
					Limit:  10,
				})
				return "View", result, err
			},
		},
	}

	for _, fnName := range []string{"cli", "rest", "rpc"} {
		startTime := time.Now()
		fns := fnMap[fnName]
		for _, fn := range fns {
			fnName, _, err := fn()
			if err != nil {
				fmt.Println(fnName + " error: " + err.Error())
			} else {
				fmt.Println(fnName + " OK")
			}
		}
		endTime := time.Now()
		fmt.Println("--------------------")
		fmt.Println(fnName+" time ", int64(endTime.Sub(startTime).Seconds()*1000))
		fmt.Println("--------------------")
	}

}
