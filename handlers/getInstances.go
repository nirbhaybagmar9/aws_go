package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

//Instance
type Instance struct {
	InstanceID       string
	ImageID          string
	InstanceType     string
	AvailabilityZone string
	State            string
}

//GetInstances : http.MethodGet with region parameter gives an response of array of all the AWS EC2 instance in that region.
func GetInstances() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.Header().Set("Content-Type", "application/json")
			region := r.URL.Query().Get("region")
			if region == "" {
				http.Error(w, "missing Region parameter", http.StatusBadRequest)
			}
			sess, err := session.NewSession(&aws.Config{
				Region: aws.String(region)},
			)
			// Create new EC2 client
			ec2Svc := ec2.New(sess)

			// Call to get detailed information on each instance
			result, err := ec2Svc.DescribeInstances(nil)
			if err != nil {
				fmt.Println("Error", err)
				w.WriteHeader(http.StatusInternalServerError)

			} else {
				fmt.Println(result)
				var Instances []Instance = make([]Instance, 0)
				for idx := range result.Reservations {
					for _, inst := range result.Reservations[idx].Instances {
						instance := Instance{*inst.InstanceId, *inst.ImageId, *inst.InstanceType, *inst.Placement.AvailabilityZone, *inst.State.Name}
						Instances = append(Instances, instance)

					}
				}
				fmt.Println("Success", Instances)
				val, _ := json.Marshal(Instances)
				w.Write(val)
				w.WriteHeader(http.StatusOK)
			}
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

	})
}
