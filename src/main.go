package main

import (
	"fmt"
)

func main() {
	/*
		var (
			cfg aws.Config
			err error
		)

		if cfg, err = config.LoadDefaultConfig(context.TODO(),
			config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("key", "secret", "")),
			config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(_, _ string, _ ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{
					URL: "https://endpoint",
				}, nil
			})),
		); err != nil {
			log.Fatalln(err)
		}

		client := s3.NewFromConfig(cfg)

		fmt.Println(client)
	*/

	fmt.Println("Hello, world!")
}
