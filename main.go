package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/bugsnag/bugsnag-go/v2"
)

func main() {
	fmt.Println("Start of main!")
	bugsnag.Configure(bugsnag.Configuration{})
	region := "eu-central-1"
	my_session, err := session.NewSession(&aws.Config{Region: &region})
	if err != nil {
		fmt.Println(err)
		return
	}
	my_sts := sts.New(my_session, aws.NewConfig().WithRegion(region))
	fmt.Println("About to call getCallerIdentity. Together with Bugsnag.Configure, this causes main to be called for a second time? But this second invocation doesn't really start, only the first few lines are executed? The second invocation nevers gets to printing this exact line??")
	_, err = my_sts.GetCallerIdentity(nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("End of main!")
}
