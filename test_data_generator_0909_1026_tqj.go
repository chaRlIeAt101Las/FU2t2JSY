// 代码生成时间: 2025-09-09 10:26:25
package main

import (
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// define constants for the number of records to generate
const (
	defaultRecordsCount = 10
)

// RandomDataGenerator holds options for the generator
type RandomDataGenerator struct {
	RecordsCount int
}

// NewRandomDataGenerator creates a new instance of RandomDataGenerator with default values
func NewRandomDataGenerator() *RandomDataGenerator {
	return &RandomDataGenerator{
		RecordsCount: defaultRecordsCount,
	}
}

// GenerateData generates random test data
func (g *RandomDataGenerator) GenerateData() ([]map[string]string, error) {
	data := make([]map[string]string, g.RecordsCount)
	for i := 0; i < g.RecordsCount; i++ {
		data[i] = map[string]string{
			"ID":            generateRandomID(8),
			"Name":          generateRandomName(10),
			"Email":         generateRandomEmail(20),
			"PhoneNumber":   generateRandomPhoneNumber(12),
			"RegistrationDate": generateRandomDate(),
		}
	}
	return data, nil
}

// generateRandomID generates a random ID
func generateRandomID(length int) string {
	id, err := generateRandomString(length)
	if err != nil {
		panic(err)
	}
	return id
}

// generateRandomName generates a random name
func generateRandomName(length int) string {
	name, err := generateRandomString(length)
	if err != nil {
		panic(err)
	}
	return name
}

// generateRandomEmail generates a random email address
func generateRandomEmail(length int) string {
	email, err := generateRandomString(length)
	if err != nil {
		panic(err)
	}
	return email + "@example.com"
}

// generateRandomPhoneNumber generates a random phone number
func generateRandomPhoneNumber(length int) string {
	number, err := generateRandomString(length)
	if err != nil {
		panic(err)
	}
	return "+1-" + number
}

// generateRandomString generates a random string of a given length
func generateRandomString(length int) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	bytes := make([]byte, length)
	for i := range bytes {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		bytes[i] = letters[num.Int64()]
	}
	return string(bytes), nil
}

// generateRandomDate generates a random date string
func generateRandomDate() string {
	date := time.Now().AddDate(-rand.Int63n(3650), 0, 0)
	return date.Format("2006-01-02")
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "test-data-generator",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your application. For example:

  Cobra is a CLI library for Go that empowers applications.
  This application is a tool to generate test data.`,
	Run: func(cmd *cobra.Command, args []string) {
		generator := NewRandomDataGenerator()
		data, err := generator.GenerateData()
		if err != nil {
			fmt.Println("Error generating test data: ", err)
			os.Exit(1)
		}
		for _, record := range data {
			fmt.Println(record)
		}
	},
}

// init configures the root command
func init() {
	rootCmd.PersistentFlags().IntVarP(&randomDataGenerator.RecordsCount, "count", "c", defaultRecordsCount, "Number of records to generate")
}

func main() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
