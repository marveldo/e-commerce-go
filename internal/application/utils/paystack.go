package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/marveldo/gogin/internal/config"
)

func CallPaystackUrl(g *gin.Context, email string, price float64, callback_url string, ord_id uint) (*http.Response, error) {
	data := map[string]interface{}{
		"email":        email,
		"amount":       price,
		"callback_url": callback_url,
		"reference":    fmt.Sprintf("ORD-%v", ord_id),
	}
	json_data, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(
		g.Request.Context(),
		"POST",
		"https://api.paystack.co/transaction/initialize",
		bytes.NewBuffer(json_data),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+config.LoadConfig().Paystack_Secret_Key)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Request Retrieved an not 200 code")
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}
