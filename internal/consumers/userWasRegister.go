package consumers

import (
	"context"
	"dysn/notify/internal/model/dto"
	"dysn/notify/pkg/log"
	"encoding/json"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type AuthServiceInterface interface {
	ConfirmRegister(ctx context.Context, code dto.Code) error
}

type ConsumerRegister struct {
	reader  *kafka.Reader
	authSrv AuthServiceInterface
	logger  *log.Logger
}

func NewConsumerRegister(reader *kafka.Reader,
	authSrv AuthServiceInterface, logger *log.Logger) *ConsumerRegister {
	return &ConsumerRegister{reader, authSrv, logger}
}

func (c *ConsumerRegister) ReadRegisters(ctx context.Context) {
	c.logger.InfoLog.Println("start reader register user for consumer")
	for {
		message, err := c.reader.ReadMessage(ctx)
		if err != nil {
			fmt.Println("err read", err)
			break
		}
		fmt.Printf("message at offset %d: %s = %s\n", message.Offset, string(message.Key), string(message.Value))

		var userReg dto.Code
		if err = json.Unmarshal(message.Value, &userReg); err != nil {
			c.logger.ErrorLog.Println("cant unmarshal message:", err)
		} else {
			if err := c.authSrv.ConfirmRegister(ctx, userReg); err != nil {
				c.logger.ErrorLog.Println("cant send message:", err)
			}
			c.logger.InfoLog.Println("message was send:", err)
		}
	}
}
