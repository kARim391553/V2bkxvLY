// 代码生成时间: 2025-10-05 02:49:24
package main

import (
    "log"
    "net/http"
    "math"

    "github.com/labstack/echo"
)

// OptionPricingService 结构体封装期权定价相关的数据和方法
type OptionPricingService struct {
    // 可以添加更多字段，如市场数据接口等
}

// NewOptionPricingService 创建一个OptionPricingService实例
func NewOptionPricingService() *OptionPricingService {
    return &OptionPricingService{}
}

// BlackScholes 计算欧式期权的Black-Scholes定价模型
func (s *OptionPricingService) BlackScholes(S float64, K float64, T float64, r float64, sigma float64) (float64, error) {
    // 计算d1和d2
    d1 := (math.Log(S/K) + (r + sigma*sigma/2)*T) / (sigma * math.Sqrt(T))
    d2 := d1 - sigma * math.Sqrt(T)

    // 计算欧式看涨期权的Black-Scholes价格
    call := S * math.Exp(-r * T) * normCDF(d1) - K * math.Exp(-r * T) * normCDF(d2)

    // 计算欧式看跌期权的Black-Scholes价格
    put := K * math.Exp(-r * T) * normCDF(-d2) - S * math.Exp(-r * T) * normCDF(-d1)

    // 这里返回看涨期权价格作为示例，可以根据需要返回看跌期权价格或两者都返回
    return call, nil
}

// normCDF 计算标准正态分布的累积分布函数值
func normCDF(x float64) float64 {
    return 0.5 * (1 + math.Erf(x / (math.Sqrt2)))
}

// handlerOptionPricing 创建期权定价的HTTP处理函数
func handlerOptionPricing(service *OptionPricingService) func(c echo.Context) error {
    return func(c echo.Context) error {
        // 从请求中解析参数
        S, err := c.GetFloat("S")
        if err != nil {
            return err
        }
        K, err := c.GetFloat("K")
        if err != nil {
            return err
        }
        T, err := c.GetFloat("T")
        if err != nil {
            return err
        }
        r, err := c.GetFloat("r")
        if err != nil {
            return err
        }
        sigma, err := c.GetFloat("sigma")
        if err != nil {
            return err
        }

        // 计算期权价格
        price, err := service.BlackScholes(S, K, T, r, sigma)
        if err != nil {
            return err
        }

        // 返回计算结果
        return c.JSON(http.StatusOK, map[string]float64{"price": price})
    }
}

func main() {
    e := echo.New()
    service := NewOptionPricingService()

    // 定义路由
    e.GET("/option-pricing", handlerOptionPricing(service))

    // 启动服务器
    log.Fatal(e.Start(":8080"))
}
