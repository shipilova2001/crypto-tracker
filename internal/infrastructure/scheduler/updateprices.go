package scheduler

import (
	"context"
	"fmt"
	"time"
)
type PriceRefresher interface {
	RefreshAll()
}

// Цены обновляются каждые 30 секунд в фоновом потоке
func UpdatePrices(ctx context.Context, refresher PriceRefresher) {
	
	for {
		fmt.Println("upd 1")
		select {
		case <-ctx.Done():
			return
		default:
			refresher.RefreshAll()
			fmt.Println("upd")
		}
		time.Sleep(time.Duration(30) * time.Second)
	}
}
