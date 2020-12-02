package infra

import (
	"cloud.google.com/go/pubsub"
	"context"
	"fmt"
)

type PubsubClient struct {
	*pubsub.Client
	ProjectID string
}

func NewPubsubClient(pID string) *PubsubClient {
	c, err := pubsub.NewClient(context.Background(), pID)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &PubsubClient{ProjectID: pID, Client: c}
}

func (p *PubsubClient) SimplePublish(s string, m []byte) error {
	ctx := context.Background()
	t := p.Topic(s)
	defer t.Stop()
	var results []*pubsub.PublishResult
	r := t.Publish(ctx, &pubsub.Message{
		Data: m,
	})

	results = append(results, r)
	for _, r := range results {
		id, err := r.Get(ctx)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("Published a message with a message ID: %s\n", id)
	}
	return nil
}
