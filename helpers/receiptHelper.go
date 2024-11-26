package helpers

import (
	"fmt"
	"math"
	"receipt-processor-module/models"
	"strconv"
	"strings"
)

func GetReceiptById(id string, receipts []models.Receipt) (models.Receipt, string) {
	errorMsg := "Invalid Receipt Id"
	for _, receipt := range receipts { // Looping through all the receipts to get the receipt with give ID
		fmt.Println(receipt)
		if id == receipt.ID {
			return receipt, ""
		}
	}
	var emptyStruct models.Receipt
	return emptyStruct, errorMsg
}

/**
 * This function calculates the number of points for the item in the receipt
 */
func GetItemPoints(item models.Item) int {
	total := 0 // Note: Following "total" variable pattern keeping future expansion of point calculation in mid. (e.g. adding more rules) 
	trimmedDescription := strings.TrimSpace(item.ShortDescription)
	if len(trimmedDescription) % 3 == 0 {
		itemPrice, _ := strconv.ParseFloat(item.Price, 64)
		total += int(math.Ceil(itemPrice * 0.2))

	}
	return total
}

/*
	These rules collectively define how many points should be awarded to a receipt.

	1. One point for every alphanumeric character in the retailer name.
	2. 50 points if the total is a round dollar amount with no cents.
	3. 25 points if the total is a multiple of 0.25.
	4. 5 points for every two items on the receipt.
	5. If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.
	6. 6 points if the day in the purchase date is odd.
	7. 10 points if the time of purchase is after 2:00pm and before 4:00pm.
*/

func CalculateReceiptPoints(receipt models.Receipt) (int, string) {

	// 1. Counting alphanumeric characters
	totalPoints := CountAlphanumeric(receipt.Retailer)

	// 2. Checking if the amount is rounded and adding points accordingly
	if isRoundedDollarAmount, errorString := IsRoundedDollarAmount(receipt.Total); errorString != "" {
		return -1, errorString
	} else if isRoundedDollarAmount {
		totalPoints += 50
	}

	// 3. Adding points for receipts pairs
	if isMultipleOfQuarter, errorString := IsMultipleOfQuarter(receipt.Total); errorString != "" {
		return -1, errorString
	} else if isMultipleOfQuarter {
		totalPoints += 50
	}

	// 4. Adding points for receipts pairs
	totalPoints += 5 * (len(receipt.Items) / 2)

	// 5. Adding points for items
	for _, item := range receipt.Items { 
		totalPoints += GetItemPoints(item)
	}

	// 6. Getting Day from YYYY-MM-DD format
	if day, errorString := GetDayFromDate(receipt.PurchaseDate); errorString != "" {
		return -1, errorString
	} else if day%2 == 1 {
		totalPoints += 6
	}

	// 7. Checking if the time is between 2 and 4 PM
	if isTimeBetween2And4PM, errorString := IsTimeBetween2And4PM(receipt.PurchaseTime); errorString != "" {
		return -1, errorString
	} else if isTimeBetween2And4PM {
		totalPoints += 10
	}

	return totalPoints, ""
}
