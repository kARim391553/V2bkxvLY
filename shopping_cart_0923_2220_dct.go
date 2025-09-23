// 代码生成时间: 2025-09-23 22:20:48
package main

import (
    "net/http"
    "github.com/labstack/echo"
)

// CartItem 表示购物车中的一个商品项
type CartItem struct {
    ID          string  `json:"id"`         // 商品ID
    Name        string  `json:"name"`       // 商品名称
    Price       float64 `json:"price"`      // 商品价格
    Quantity    int     `json:"quantity"`    // 商品数量
}

// Cart 表示购物车，包含多个商品项
type Cart struct {
    Items map[string]CartItem
}

// NewCart 创建一个新的购物车实例
func NewCart() *Cart {
    return &Cart{
        Items: make(map[string]CartItem),
    }
}

// AddItem 向购物车添加商品
func (c *Cart) AddItem(item CartItem) error {
    if _, exists := c.Items[item.ID]; exists {
        return echo.NewHTTPError(http.StatusBadRequest, "Item already exists in cart")
    }
    c.Items[item.ID] = item
    return nil
}

// RemoveItem 从购物车移除商品
func (c *Cart) RemoveItem(itemID string) error {
    if _, exists := c.Items[itemID]; !exists {
        return echo.NewHTTPError(http.StatusNotFound, "Item not found in cart")
    }
    delete(c.Items, itemID)
    return nil
}

// UpdateQuantity 更新购物车中商品的数量
func (c *Cart) UpdateQuantity(itemID string, quantity int) error {
    if _, exists := c.Items[itemID]; !exists {
        return echo.NewHTTPError(http.StatusNotFound, "Item not found in cart")
    }
    if quantity < 0 {
        return echo.NewHTTPError(http.StatusBadRequest, "Quantity cannot be negative")
    }
    c.Items[itemID].Quantity = quantity
    return nil
}

// GetCart 获取购物车内容
func (c *Cart) GetCart() map[string]CartItem {
    return c.Items
}

func main() {
    e := echo.New()

    // 创建购物车实例
    cart := NewCart()

    // 添加商品到购物车
    e.POST("/add", func(c echo.Context) error {
        item := CartItem{
            ID:    c.QueryParam("id"),
            Name:  c.QueryParam("name"),
            Price: c.QueryParam("price"),
            Quantity: 1,
        }
        if err := cart.AddItem(item); err != nil {
            return err
        }
        return c.JSON(http.StatusOK, cart.GetCart())
    })

    // 从购物车移除商品
    e.POST("/remove", func(c echo.Context) error {
        itemID := c.QueryParam("id")
        if err := cart.RemoveItem(itemID); err != nil {
            return err
        }
        return c.JSON(http.StatusOK, cart.GetCart())
    })

    // 更新购物车中商品的数量
    e.POST("/update", func(c echo.Context) error {
        itemID := c.QueryParam("id\)
        quantity, err := strconv.Atoi(c.QueryParam("quantity"))
        if err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid quantity")
        }
        if err := cart.UpdateQuantity(itemID, quantity); err != nil {
            return err
        }
        return c.JSON(http.StatusOK, cart.GetCart())
    })

    // 获取购物车内容
    e.GET("/cart", func(c echo.Context) error {
        return c.JSON(http.StatusOK, cart.GetCart())
    })

    // 启动Echo服务器
    e.Logger.Fatal(e.Start(":8080"))
}
