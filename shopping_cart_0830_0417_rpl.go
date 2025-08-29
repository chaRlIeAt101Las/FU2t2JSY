// 代码生成时间: 2025-08-30 04:17:13
// shopping_cart.go

package main

import (
    "fmt"
    "log"
    "os"

    "github.com/spf13/cobra"
)

// CartItem 表示购物车中的商品项
type CartItem struct {
    ID     int    `json:"id"`
    Name   string `json:"name"`
    Price  float64 `json:"price"`
    Quantity int `json:"quantity"`
}

// ShoppingCart 表示购物车
type ShoppingCart struct {
    Items map[int]*CartItem
}

// AddItem 向购物车添加商品
func (cart *ShoppingCart) AddItem(item *CartItem) error {
    if item == nil {
        return fmt.Errorf("item cannot be nil")
    }
    if cart.Items == nil {
        cart.Items = make(map[int]*CartItem)
    }
    cart.Items[item.ID] = item
    return nil
}

// RemoveItem 从购物车移除商品
func (cart *ShoppingCart) RemoveItem(itemID int) error {
    if _, exists := cart.Items[itemID]; !exists {
        return fmt.Errorf("item with ID %d does not exist", itemID)
    }
    delete(cart.Items, itemID)
    return nil
}

// CalculateTotal 计算购物车中所有商品的总价
func (cart *ShoppingCart) CalculateTotal() (float64, error) {
    total := 0.0
    for _, item := range cart.Items {
        total += item.Price * float64(item.Quantity)
    }
    return total, nil
}

// NewCart 创建一个新的购物车
func NewCart() *ShoppingCart {
    return &ShoppingCart{
        Items: make(map[int]*CartItem),
    }
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
    Use:   "shopping-cart",
    Short: "A brief description of your application",
    Long: `
        A longer description that spans multiple lines and likely contains
        examples and usage of using your application. For example:

        Usage:
        shopping-cart add --item-id 1 --item-name "Apple" --item-price 0.5 --item-quantity 10
    `,
    Args: cobra.MinimumNArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        // 这里可以添加处理命令行参数的逻辑
    },
}

// initCommands 初始化命令
func initCommands() {
    // 添加子命令
    rootCmd.AddCommand(&cobra.Command{
        Use:   "add",
        Short: "Add an item to the shopping cart",
        Long: `Add an item to the shopping cart`,
        Run: func(cmd *cobra.Command, args []string) {
            var itemID int
            var itemName string
            var itemPrice float64
            var itemQuantity int
            err := cmd.Flags().GetInt("item-id", &itemID)
            if err != nil {
                log.Fatal(err)
            }
            err = cmd.Flags().GetString("item-name", &itemName)
            if err != nil {
                log.Fatal(err)
            }
            err = cmd.Flags().GetFloat64("item-price", &itemPrice)
            if err != nil {
                log.Fatal(err)
            }
            err = cmd.Flags().GetInt("item-quantity", &itemQuantity)
            if err != nil {
                log.Fatal(err)
            }
            cart := NewCart()
            cart.AddItem(&CartItem{ID: itemID, Name: itemName, Price: itemPrice, Quantity: itemQuantity})
            fmt.Printf("Item added to cart: %+v
", cart.Items[itemID])
        },
    })

    // 可以添加更多子命令，如 remove, calculateTotal 等
}

func main() {
    initCommands()
    if err := rootCmd.Execute(); err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}
