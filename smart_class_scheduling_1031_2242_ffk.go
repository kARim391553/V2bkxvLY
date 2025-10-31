// 代码生成时间: 2025-10-31 22:42:50
package main

import (
    "echo"
    "net/http"
    "strconv"
)

// ClassSchedule represents a class schedule
type ClassSchedule struct {
    ClassID int     `json:"class_id"`
    Teacher  string `json:"teacher"`
    Subject  string `json:"subject"`
    TimeSlot string `json:"time_slot"`
}

// ScheduleService provides operations for class scheduling
type ScheduleService struct {
    // This could be a database connection in a real implementation
    schedules map[int]ClassSchedule
}

// NewScheduleService creates a new ScheduleService instance
func NewScheduleService() *ScheduleService {
    return &ScheduleService{
        schedules: make(map[int]ClassSchedule),
    }
}

// AddSchedule adds a new class schedule
func (s *ScheduleService) AddSchedule(classID int, teacher, subject, timeSlot string) error {
    if _, exists := s.schedules[classID]; exists {
        return echo.NewHTTPError(http.StatusBadRequest, "Class ID already exists")
    }
    s.schedules[classID] = ClassSchedule{
        ClassID: classID,
        Teacher:  teacher,
        Subject:  subject,
        TimeSlot: timeSlot,
    }
    return nil
}

// GetSchedule retrieves a class schedule by class ID
func (s *ScheduleService) GetSchedule(classID int) (*ClassSchedule, error) {
    schedule, exists := s.schedules[classID]
    if !exists {
        return nil, echo.NewHTTPError(http.StatusNotFound, "Class schedule not found")
    }
    return &schedule, nil
}

// Start the Echo server and define routes
func main() {
    e := echo.New()
    service := NewScheduleService()

    // POST /schedules to add a new schedule
    e.POST("/schedules", func(c echo.Context) error {
        req := new(ClassSchedule)
        if err := c.Bind(req); err != nil {
            return err
        }
        if err := service.AddSchedule(req.ClassID, req.Teacher, req.Subject, req.TimeSlot); err != nil {
            return err
        }
        return c.JSON(http.StatusCreated, req)
    })

    // GET /schedules/:classID to get a schedule by ID
    e.GET("/schedules/:classID", func(c echo.Context) error {
        classID, err := strconv.Atoi(c.Param("classID"))
        if err != nil {
            return echo.NewHTTPError(http.StatusBadRequest, "Invalid class ID")
        }
        schedule, err := service.GetSchedule(classID)
        if err != nil {
            return err
        }
        return c.JSON(http.StatusOK, schedule)
    })

    e.Logger.Fatal(e.Start(":1323"))
}
