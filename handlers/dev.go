package handlers

import (
    "net/http"
    "os"

    "crowdfunding/tools/seed"
    "github.com/gin-gonic/gin"
)

// DevSeedHandler runs the seed.Seed function but only when ENABLE_DEV_ENDPOINTS=true
func DevSeedHandler(c *gin.Context) {
    if os.Getenv("ENABLE_DEV_ENDPOINTS") != "true" {
        c.JSON(http.StatusForbidden, gin.H{"error": "dev endpoints disabled"})
        return
    }
    if err := seed.Seed(); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "seeded"})
}
