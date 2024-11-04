package main

import (
	"fmt"
	"log"
	"os"

	"github.com/shopspring/decimal"

	"github.com/yangnei/enclave-go/enclave"
	"github.com/yangnei/enclave-go/enclave/api"
	"github.com/yangnei/enclave-go/enclave/client"
	"github.com/yangnei/enclave-go/enclave/model"
	"github.com/yangnei/enclave-go/enclave/util"
)

func main() {
	enclaveClient := client.NewClient(
		os.Getenv("enclave_key"),
		os.Getenv("enclave_secret"),
		enclave.SandboxApiUrl,
	)

	spot := func() {
		tradingPair := util.NewTradingPair("AVAX", "USDC")

		// Create a limit Order
		order, err := enclaveClient.AddOrder(&api.AddOrderRequest{
			Market: tradingPair,
			Side:   model.OrderSideBuy,
			Size:   decimal.NewFromInt(1),
			Price:  decimal.NewFromInt(1),
		})
		if err != nil {
			log.Fatalf("Error creating order: %v", err)
		}
		fmt.Println("Order Created:", util.MustMarshalIndent(order))

		// Get Order
		order, err = enclaveClient.GetOrder(&api.GetOrderRequest{
			OrderID: order.OrderID,
		})
		if err != nil {
			log.Fatalf("Error getting order: %v", err)
		}
		fmt.Println("Order Retrieved:", util.MustMarshalIndent(order))

		// Cancel Order
		order, err = enclaveClient.CancelOrder(&api.CancelOrderRequest{
			OrderID: order.OrderID,
		})
		if err != nil {
			log.Fatalf("Error canceling order: %v", err)
		}
		fmt.Println("Order Canceled:", util.MustMarshalIndent(order))

		// Create a market Order
		order, err = enclaveClient.AddOrder(&api.AddOrderRequest{
			Market: tradingPair,
			Side:   model.OrderSideSell,
			Size:   decimal.NewFromInt(1),
			Type:   model.OrderTypeMarket,
		})
		if err != nil {
			log.Fatalf("Error creating order: %v", err)
		}
		fmt.Println("Order Created:", util.MustMarshalIndent(order))

		// Create another market Order
		order, err = enclaveClient.AddOrder(&api.AddOrderRequest{
			Market: tradingPair,
			Side:   model.OrderSideBuy,
			Size:   decimal.NewFromInt(1),
			Type:   model.OrderTypeMarket,
		})
		if err != nil {
			log.Fatalf("Error creating order: %v", err)
		}
		fmt.Println("Order Created:", util.MustMarshalIndent(order))

		// Create another market Order
		order, err = enclaveClient.AddOrder(&api.AddOrderRequest{
			Market: tradingPair,
			Side:   model.OrderSideSell,
			Size:   decimal.NewFromInt(1),
			Type:   model.OrderTypeMarket,
		})
		if err != nil {
			log.Fatalf("Error creating order: %v", err)
		}
		fmt.Println("Order Created:", util.MustMarshalIndent(order))

		// Get Fills
		fills, err := enclaveClient.GetFills(&api.GetFillsRequest{
			Market: tradingPair,
		})
		if err != nil {
			log.Fatalf("Error getting fills: %v", err)
		}
		fmt.Println("Fills Retrieved:", util.MustMarshalIndent(fills))

		// Get Depth
		depth, err := enclaveClient.GetDepth(&api.GetDepthRequest{
			Market: tradingPair,
		})
		if err != nil {
			log.Fatalf("Error getting order book: %v", err)
		}
		fmt.Println("Depth Retrieved:", util.MustMarshalIndent(depth))
	}

	perps := func() {
		positions, err := enclaveClient.GetPositions()
		if err != nil {
			log.Fatalf("Error getting positions: %v", err)
		}
		fmt.Println("Positions Retrieved:", util.MustMarshalIndent(positions))

		balances, err := enclaveClient.GetBalance()
		if err != nil {
			log.Fatalf("Error getting balance: %v", err)
		}
		fmt.Println("Balance Retrieved:", util.MustMarshalIndent(balances))

		transfer, err := enclaveClient.Transfer(&api.TransferRequest{
			Amount: decimal.NewFromInt(1),
			Symbol: "usdc",
		})
		if err != nil {
			log.Fatalf("Error transferring: %v", err)
		}
		fmt.Println("Transfer Created:", util.MustMarshalIndent(transfer))

		transfer, err = enclaveClient.Transfer(&api.TransferRequest{
			Amount: decimal.NewFromInt(-1),
			Symbol: "usdc",
		})
		if err != nil {
			log.Fatalf("Error transferring: %v", err)
		}
		fmt.Println("Transfer Created:", util.MustMarshalIndent(transfer))
	}

	spot()
	perps()
}
