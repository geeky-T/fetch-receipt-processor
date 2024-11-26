package api

import (
	"net/http"
	"receipt-processor-module/helpers"
	"receipt-processor-module/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetAllReceipts(c *gin.Context) { // Controller to get all receipts
	c.IndentedJSON(http.StatusOK, models.Receipts)
}

func AddReceipt(c *gin.Context) { // Controller to add a receipt
	var newReceipt models.Receipt
	if err := c.ShouldBindJSON(&newReceipt); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
	}
	newReceipt.ID = uuid.New().String() // Generates new id for every receipt
	newReceipt.Points = -1
	models.Receipts = append(models.Receipts, newReceipt)
	c.JSON(http.StatusCreated, gin.H{"id": newReceipt.ID})
}


func GetReceiptPointsById(c *gin.Context) { // Controller to get points for a receipt
	var id = c.Param("id")
	receipt, errorString := helpers.GetReceiptById(id, models.Receipts) // Getting the receipt from ID
	if(errorString != ""){
		sendErrorResponse(c, errorString) // Returning error is ID is invalid
		return
	} 
	if(receipt.Points > 0) { // Checking if the points for the receipt are already calculated to save re-calculation
		c.JSON(http.StatusCreated, gin.H{"points": receipt.Points}) 
		return
	}
	
	totalPoints, errorString := helpers.CalculateReceiptPoints(receipt)
	if errorString != "" {
		sendErrorResponse(c, errorString)
		return
	}
	receipt.Points = totalPoints
	c.JSON(http.StatusCreated, gin.H{"points": totalPoints}) // Sending total points in JSON
}

func sendErrorResponse(c *gin.Context, message string){
	c.IndentedJSON(http.StatusBadRequest, gin.H{"error": message}) // Sending error response based on error received
}