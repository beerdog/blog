package models

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Date represents a date backed by a time.Time
type Date struct {
	time.Time
}

func NewDate(time time.Time) Date {
	return Date{time}
}

func ParseDate(dateStr string) (*Date, error) {
	t, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, fmt.Errorf("failed to parse date %v: %w", dateStr, err)
	}
	d := NewDate(t)
	return &d, nil
}

// MarshalJSON marshals the date in the following format: 2006-01-02
func (d Date) MarshalJSON() ([]byte, error) {
	return []byte("\"" + d.Format("2006-01-02") + "\""), nil
}

// UnmarshalJSON unmarshals the date from the following format: 2006-01-02
func (d *Date) UnmarshalJSON(value []byte) error {
	t, err := time.Parse("\"2006-01-02\"", string(value))
	if err != nil {
		return err
	}
	*d = Date{t}
	return nil
}

func (d *Date) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{
		Value: d.Format("2006-01-02"),
	}, nil
}

func (d Date) String() string {
	return d.Format("2006-01-02")
}

func (d *Date) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	avS, ok := av.(*types.AttributeValueMemberS)
	if !ok {
		return nil
	}

	t, err := time.Parse("2006-01-02", avS.Value)
	if err != nil {
		return err
	}
	*d = Date{t}
	return nil
}

type BlogPost struct {
	Content string `json:"content"`
	Metadata
}

type Metadata struct {
	Title       string   `json:"title"`
	Preamble    string   `json:"preamble"`
	PublishDate Date     `json:"publishDate"`
	EditDate    Date     `json:"editDate"`
	Category    string   `json:"category"`
	Tags        []string `json:"tags"`
	Key         string   `json:"key"`
}
