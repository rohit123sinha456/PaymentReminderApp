package main

import (
    "fmt"
    "time"
    "gorm.io/gorm"
	. "github.com/rohit123sinha456/payredapp/models"
)

func DropTablesIfExists(db *gorm.DB) {
    // Drop each table using raw SQL
    db.Exec("DROP TABLE IF EXISTS payments;")
    db.Exec("DROP TABLE IF EXISTS users;")
    db.Exec("DROP TABLE IF EXISTS quotations;")
    db.Exec("DROP TABLE IF EXISTS service_reports;")
    db.Exec("DROP TABLE IF EXISTS nocs;")
    db.Exec("DROP TABLE IF EXISTS clients;")
}

func SeedData(db *gorm.DB) {
    // Seed 20 values for each model (sample data)
    for i := 1; i <= 20; i++ {
        // Seed Clients
        client := Client{
            Name:     fmt.Sprintf("Client%d", i),
            SiteName: fmt.Sprintf("Site %d", i),
            Address:  fmt.Sprintf("Address %d", i),
        }
        db.Create(&client)
        // Seed Payments
        payment := Payment{
            InvoiceDate:   time.Now(),
            InvoiceNumber: fmt.Sprintf("INV-%d", i),
            InvoiceAmount: float64(i * 100),
            SentDate:            time.Now(),
            InvoiceFor:    fmt.Sprintf("Invoice for #%d", i),
            ClientID:      client.ID, // Assuming 5 clients exist
        }
        db.Create(&payment)

        // Seed Users
        user := User{
            Name:        fmt.Sprintf("User%d", i),
            Email:       fmt.Sprintf("user%d@example.com", i),
            PhoneNumber: fmt.Sprintf("1234567%d", i),
            Password:    "password123",
        }
        db.Create(&user)



        // Seed Quotations
        quotation := Quotation{
            QuotationNumber:    fmt.Sprintf("QUO-%d", i),
            QuotationDate:      time.Now(),
            QuotationTotalValue: float64(i * 500),
            QuotationValidTill: time.Now(),
            SentDate:            time.Now(),
            QuotationReason:    fmt.Sprintf("Reason #%d", i),
            ClientID:           client.ID, // Assuming 5 clients exist
        }
        db.Create(&quotation)

        // Seed ServiceReports
        serviceReport := ServiceReport{
            ServiceReportNumber: fmt.Sprintf("SR-%d", i),
            ServiceReportDate:   time.Now(),
            SentDate:            time.Now(),
            ServiceReportNotes:  fmt.Sprintf("Notes #%d", i),
            ClientID:            client.ID, // Assuming 5 clients exist
        }
        db.Create(&serviceReport)

        // Seed NOCs
        noc := NOC{
            StatusNote:   fmt.Sprintf("Status Note #%d", i),
            DocumentLink: fmt.Sprintf("https://example.com/documents/noc%d.pdf", i),
            ClientID:     client.ID, // Assuming 5 clients exist
        }
        db.Create(&noc)

        // Seed Categories
        // category := Category{
        //     Name:         fmt.Sprintf("Category #%d", i),
        //     CategoryType: "payment",
        // }
        // db.Create(&category)
    }
}

func Migrate(db *gorm.DB){
    
    //Drop Table If Exists
	DropTablesIfExists(db)
    // Auto Migrate all models
    db.AutoMigrate(&Payment{}, &User{}, &Client{}, &Quotation{}, &ServiceReport{}, &NOC{})

    // Seed data insertion
    SeedData(db)

    // Close the database connection
    // db.Close()
}