// 代码生成时间: 2025-10-02 02:05:25
package main

import (
    "fmt"
    "github.com/labstack/echo"
    "os"
    "os/signal"
    "syscall"
    "time"
)

// SignalHandler 结构体，用于封装 Echo 实例和信号处理逻辑
type SignalHandler struct {
    echoInstance *echo.Echo
    signalChan   chan os.Signal
}

// NewSignalHandler 创建 SignalHandler 实例
func NewSignalHandler(echoInstance *echo.Echo) *SignalHandler {
    return &SignalHandler{
        echoInstance: echoInstance,
        signalChan:   make(chan os.Signal, 1),
    }
}

// HandleSignals 设置信号处理器
func (sh *SignalHandler) HandleSignals() {
    // 监听 SIGINT 和 SIGTERM 信号
    signal.Notify(sh.signalChan, syscall.SIGINT, syscall.SIGTERM)

    // 阻塞等待信号
    <-sh.signalChan

    // 优雅地关闭 Echo 实例
    fmt.Println("Received shutdown signal, starting graceful shutdown...")
    sh.echoInstance.Shutdown(ctx)
    fmt.Println("Graceful shutdown complete.")
}

func main() {
    echoInstance := echo.New()
    defer echoInstance.Close()

    // 启动 Echo 服务
    go func() {
        if err := echoInstance.Start(":8080"); err != nil && err != echo.ErrServerClosed {
            fmt.Printf("Echo server failed to start: %v", err)
        return
        }
    }()

    // 创建信号处理器
    signalHandler := NewSignalHandler(echoInstance)
    // 处理信号
    signalHandler.HandleSignals()
}

// 上下文对象
var ctx = context.Background()
