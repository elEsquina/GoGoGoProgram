package data

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
	"log"
)

func GetOrdersInTimeRange(start, end time.Time, repo *OrderRepository) ([]Order, error) {
    repo.mu.RLock()
    defer repo.mu.RUnlock()

    var filteredOrders []Order
    for _, order := range *repo.orders {
        if order.CreatedAt.After(start) && order.CreatedAt.Before(end) {
            filteredOrders = append(filteredOrders, order)
        }
    }
    return filteredOrders, nil
}

func generateSalesReport(repo *OrderRepository) (SalesReport, error) {
    now := time.Now()
    start := now.Add(-24 * time.Hour)

    orders, err := GetOrdersInTimeRange(start, now, repo)
    if err != nil {
        return SalesReport{}, err
    }

    report := SalesReport{
        Timestamp:    now,
        TotalRevenue: 0,
        TotalOrders:  len(orders),
    }

    bookSalesMap := make(map[int]*BookSales)
    for _, order := range orders {
        report.TotalRevenue += order.TotalPrice
        for _, item := range order.Items {
            if _, exists := bookSalesMap[item.Book.ID]; !exists {
                bookSalesMap[item.Book.ID] = &BookSales{
                    Book:     item.Book,
                    Quantity: 0,
                }
            }
            bookSalesMap[item.Book.ID].Quantity += item.Quantity
        }
    }

    for _, sales := range bookSalesMap {
        report.TopSellingBooks = append(report.TopSellingBooks, *sales)
    }

    return report, nil
}

func saveReport(report SalesReport) error {
    filename := fmt.Sprintf("output-reports/report_%s.json", report.Timestamp.Format("20060102150405"))
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    return json.NewEncoder(file).Encode(report)
}

func StartReportGenerator(store *InMemoryStore) {

	dao, err := GetDAO[Order]("order", store)
	if err != nil {
		log.Fatalf("Failed to get DAO for order: %v", err)
	}
	repo, ok := dao.(*OrderRepository)
	if !ok {
		log.Fatalf("Failed to cast DAO to *OrderRepository")
	}
	
    ticker := time.NewTicker(24 * time.Second)
    defer ticker.Stop()

    for range ticker.C {
        report, err := generateSalesReport(repo)
        if err != nil {
            log.Println("Error generating sales report:", err)
            continue
        }
        if err := saveReport(report); err != nil {
            log.Println("Error saving sales report:", err)
        }
    }
}

