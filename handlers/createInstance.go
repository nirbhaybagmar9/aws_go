package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

//VM
type VM struct {
	ImageID  string `json:"imageId"`
	Region   string `json:"region"`
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue"`
}

//CreateInstance : http.MethodPost.
func CreateInstance() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		w.Header().Set("Content-Type", "application/json")

		if (*r).Method == "OPTIONS" {
			return
		}
		if r.Method == http.MethodPost {

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Request is missing parameter", http.StatusBadRequest)
			}
			var inst VM
			err = json.Unmarshal(body, &inst)
			if err != nil {
				panic(err)
			}
			sess, err := session.NewSession(&aws.Config{
				Region: aws.String(inst.Region)},
			)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(inst)
			// Create EC2 service client
			svc := ec2.New(sess)

			// Specify the details of the instance that you want to create.
			runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
				// An Amazon Linux AMI ID for t2.micro instances in the us-west-2 region
				ImageId:      aws.String(inst.ImageID),
				InstanceType: aws.String("t2.micro"),
				MinCount:     aws.Int64(1),
				MaxCount:     aws.Int64(1),
			})

			if err != nil {
				http.Error(w, "Could not create instance", http.StatusInternalServerError)
				fmt.Println("Could not create instance", err)
				return
			}

			fmt.Println("Created instance", *runResult.Instances[0].InstanceId)

			// Add tags to the created instance
			_, errtag := svc.CreateTags(&ec2.CreateTagsInput{
				Resources: []*string{runResult.Instances[0].InstanceId},
				Tags: []*ec2.Tag{
					{
						Key:   aws.String(inst.TagKey),
						Value: aws.String(inst.TagValue),
					},
				},
			})
			if errtag != nil {
				http.Error(w, "Could not create tags", http.StatusInternalServerError)
				log.Println("Could not create tags for instance", runResult.Instances[0].InstanceId, errtag)
				return
			}
			w.WriteHeader(http.StatusOK)
			values := map[string]string{"imageId": *runResult.Instances[0].InstanceId}
			jsonValue, _ := json.Marshal(values)
			w.Write(jsonValue)
			fmt.Println("Successfully tagged instance")
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

	})
}
