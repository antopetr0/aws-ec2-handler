package ec2

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const AWS_ACCESS_KEY_ID = "AKIAREUC7U5FRBHKPUXJ"
const AWS_SECRET_ACCESS_KEY = "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"

type EC2Manager struct {
	svc *ec2.EC2
}

func NewEC2Manager() (*EC2Manager, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_REGION")),
		Credentials: credentials.NewEnvCredentials(),
	})
	if err != nil {
		return nil, err
	}

	return &EC2Manager{
		svc: ec2.New(sess),
	}, nil
}

func (m *EC2Manager) ListInstances() ([]*ec2.Instance, error) {
	input := &ec2.DescribeInstancesInput{}

	result, err := m.svc.DescribeInstances(input)
	if err != nil {
		return nil, err
	}

	instances := []*ec2.Instance{}
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			instances = append(instances, instance)
		}
	}

	return instances, nil
}

func (m *EC2Manager) StartInstance(instanceID string) error {
	input := &ec2.StartInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	}

	_, err := m.svc.StartInstances(input)
	return err
}

func (m *EC2Manager) StopInstance(instanceID string) error {
	input := &ec2.StopInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	}

	_, err := m.svc.StopInstances(input)
	return err
}
