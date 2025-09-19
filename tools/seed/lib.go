package seed

import (
    "fmt"
    "log"
    "time"

    "crowdfunding/database"
    "crowdfunding/models"
    "golang.org/x/crypto/bcrypt"
)

// Seed inserts example users, projects and fundings into the database.
// It is safe to run multiple times; existing records (matched by email/title)
// will be left intact.
func Seed() error {
    database.InitDB()

    users := []struct{
        Email string
        Username string
        Password string
    }{
        {"alice@example.com", "alice", "password1"},
        {"bob@example.com", "bob", "password2"},
        {"carol@example.com", "carol", "password3"},
    }

    var created []models.User
    for _, u := range users {
        var mu models.User
        if err := database.DB.Where("email = ?", u.Email).First(&mu).Error; err == nil {
            fmt.Printf("user exists: %s (id=%d)\n", mu.Email, mu.ID)
            created = append(created, mu)
            continue
        }
        hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
        mu = models.User{Username: u.Username, Email: u.Email, Password: string(hash)}
        if err := database.DB.Create(&mu).Error; err != nil {
            return fmt.Errorf("failed to create user %s: %w", u.Email, err)
        }
        fmt.Printf("created user: %s (id=%d)\n", mu.Email, mu.ID)
        created = append(created, mu)
    }

    today := time.Now()
    projects := []models.Project{
        {Title: "Smart Plant Sensor", Description: "IoT sensor to monitor plant health.", Goal: 2000, Deadline: today.AddDate(0,1,0).Format("2006-01-02"), Raised: 150, Status: "published", UserID: created[0].ID},
        {Title: "Indie Game: Skylight", Description: "A cozy indie game.", Goal: 8000, Deadline: today.AddDate(0,2,0).Format("2006-01-02"), Raised: 3200, Status: "published", UserID: created[1].ID},
        {Title: "Art Zine Vol.1", Description: "A community art zine (draft).", Goal: 1000, Deadline: today.AddDate(0,3,0).Format("2006-01-02"), Raised: 0, Status: "draft", UserID: created[0].ID},
        {Title: "Robotics Toolkit", Description: "Education robotics kit (draft)", Goal: 12000, Deadline: today.AddDate(0,4,0).Format("2006-01-02"), Raised: 0, Status: "draft", UserID: created[1].ID},
    }

    for i := range projects {
        p := &projects[i]
        var existing models.Project
        if err := database.DB.Where("title = ?", p.Title).First(&existing).Error; err == nil {
            fmt.Printf("project exists: %s (id=%d)\n", existing.Title, existing.ID)
            continue
        }
        if err := database.DB.Create(p).Error; err != nil {
            return fmt.Errorf("failed to create project %s: %w", p.Title, err)
        }
        fmt.Printf("created project: %s (id=%d)\n", p.Title, p.ID)
    }

    var pubProjects []models.Project
    database.DB.Where("status = ?", "published").Find(&pubProjects)
    for _, p := range pubProjects {
        var f models.Funding
        if err := database.DB.Where("project_id = ?", p.ID).First(&f).Error; err == nil {
            fmt.Printf("project %d already has funding records\n", p.ID)
            continue
        }
        funds := []models.Funding{
            {ProjectID: p.ID, UserID: created[0].ID, Amount: p.Raised * 0.4},
            {ProjectID: p.ID, UserID: created[1].ID, Amount: p.Raised * 0.6},
        }
        for _, ff := range funds {
            if err := database.DB.Create(&ff).Error; err != nil {
                return fmt.Errorf("failed to create funding for project %d: %w", p.ID, err)
            }
        }
        fmt.Printf("created %d fundings for project %d\n", len(funds), p.ID)
    }

    fmt.Println("seed complete")
    return nil
}
