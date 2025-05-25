package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Mahasiswa struct {
    Nama   string `json:"nama"`
    Kelas    string `json:"kelas"`
    NPM    string `json:"npm"`
    Jurusan string `json:"jurusan"`
}

func main() {
    r := gin.Default()

    r.GET("/datamahasiswa", func(c *gin.Context) {
        data := Mahasiswa{
            Nama:   "Resqi Aulia Gita Utami",
            Kelas:  "2A",
            NPM:    "714230003",
            Jurusan: "D4 Teknik Informatika",
        }
        c.JSON(http.StatusOK, data) // Ini yang dimaksud "return c.JSON"
    })

    r.Run(":8080") // Jalankan server di localhost:8080
}
