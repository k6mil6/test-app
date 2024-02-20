package printer

import (
	"context"
	"fmt"
	"github.com/k6mil6/test-app/internal/model"
	"strconv"
	"strings"
)

type OrderStorage interface {
	Create(ctx context.Context, order model.Order) error
	Select(ctx context.Context, orderID int) (model.Order, error)
}

type ShelfStorage interface {
	SelectNameById(ctx context.Context, itemID int) (string, error)
}

func PrintPage(ctx context.Context, orderStorage OrderStorage, shelfStorage ShelfStorage, orderIDs []string) (string, error) {
	itemsByShelf := make(map[int][]model.Item)
	shelves := make(map[int]string)
	orders := make(map[int]model.Order)

	for _, idStr := range orderIDs {
		orderID, _ := strconv.Atoi(idStr)
		order, err := orderStorage.Select(ctx, orderID)
		if err != nil {
			return "", err
		}

		for _, item := range order.Items {
			orders[item.ID] = order
			itemsByShelf[item.MainShelfID] = append(itemsByShelf[item.MainShelfID], item)
			if _, ok := shelves[item.MainShelfID]; !ok {
				shelfName, err := shelfStorage.SelectNameById(ctx, item.MainShelfID)
				if err != nil {
					return "", err
				}
				shelves[item.MainShelfID] = shelfName
			}
		}
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("=+=+=+=\nСтраница сборки заказов %s\n\n", strings.Join(orderIDs, ",")))

	for shelfID, items := range itemsByShelf {
		shelfName := shelves[shelfID]
		sb.WriteString(fmt.Sprintf("===Стеллаж %s\n", shelfName))
		for _, item := range items {
			sb.WriteString(fmt.Sprintf("%s (id=%d)\nзаказ %d, %d шт\n\n", item.Name, item.ID, orders[item.ID].ID, item.Quantity))
			if item.Shelves != nil {
				for _, shelf := range item.Shelves {
					sb.WriteString(fmt.Sprintf("Доп стеллаж - %s\n", shelf.Name))
				}
			}
		}
	}

	return sb.String(), nil
}
