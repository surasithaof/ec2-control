package main

type Config struct {
	Region       string `envconfig:"AWS_REGION" required:"true"`
	AccessKeyID  string `envconfig:"AWS_ACCESS_KEY_ID" required:"true"`
	SeceretKey   string `envconfig:"AWS_SECRET_KEY" required:"true"`
	SessionToken string `envconfig:"AWS_SESSION_TOKEN" required:"false"`
}
