package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

//State Structure for request body to change the state on AWS EC2 instance.
type State struct {
	Region     string `json:"region"`
	InstanceID string `json:"instanceId"`
	State      string `json:"state"`
}

//ChangeState : http.MethodPost.
func ChangeState() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if (*r).Method == "OPTIONS" {
			return
		}
		if r.Method == http.MethodPost {

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				http.Error(w, "Request is missing parameter", http.StatusBadRequest)
			}
			var s State
			err = json.Unmarshal(body, &s)
			if err != nil {
				panic(err)
			}
			sess, err := session.NewSession(&aws.Config{
				Region: aws.String(s.Region)},
			)
			if err != nil {
				fmt.Println(err)
			}
			// Create new EC2 client
			svc := ec2.New(sess)
			if s.State == "START" {
				input := &ec2.StartInstancesInput{
					InstanceIds: []*string{
						aws.String(s.InstanceID),
					},
					DryRun: aws.Bool(true),
				}
				result, err := svc.StartInstances(input)
				awsErr, ok := err.(awserr.Error)

				if ok && awsErr.Code() == "DryRunOperation" {
					input.DryRun = aws.Bool(false)
					result, err = svc.StartInstances(input)
					if err != nil {
						http.Error(w, "Error processing request", http.StatusInternalServerError)
						fmt.Println("Error", err)
					} else {
						w.WriteHeader(http.StatusOK)
						fmt.Println("Success", result.StartingInstances)
					}
				} else {
					http.Error(w, "Error due to permissions", http.StatusInternalServerError)
					fmt.Println("Error due to permissions", err)
				}
			} else {
				input := &ec2.StopInstancesInput{
					InstanceIds: []*string{
						aws.String(s.InstanceID),
					},
					DryRun: aws.Bool(true),
				}
				result, err := svc.StopInstances(input)
				fmt.Println(result)
				awsErr, ok := err.(awserr.Error)
				if ok && awsErr.Code() == "DryRunOperation" {
					input.DryRun = aws.Bool(false)
					result, err = svc.StopInstances(input)
					fmt.Println(result)
					if err != nil {
						http.Error(w, "Error processing request", http.StatusInternalServerError)
						fmt.Println("Error", err)
					} else {
						w.WriteHeader(http.StatusOK)
						fmt.Println("Success", result.StoppingInstances)
					}
				} else {
					http.Error(w, "Error due to permissions", http.StatusInternalServerError)
					fmt.Println("Error", err)
				}
			}

		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

	})
}
