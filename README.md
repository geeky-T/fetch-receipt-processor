# Receipt Processor (Golang)
This web application is fully developed in Golang, utilizing the Gin web framework for routing and various utility packages to streamline functionality.

## Project Highlights
- The project structure is designed to be production-ready and scalable for future enhancements.
- The repository includes the following components:
  - **api**: Contains route definitions and corresponding controller functions.
  - **helpers**: Houses utility functions supporting core logic.
  - **models**: Defines structs representing real-world entities like receipts and items.
  - **main.go**: Initializes the server and imports the defined routes.

## Local Setup Instructions
1. Install Golang on your system (tested with version 1.21.1).
2. Navigate to the root directory: `cd receipt-processor`
3. Install dependencies by running `go get .` or `go get -t receipt-processor-module`
3. Start the server by running: `go run .`
4. The server will be hosted locally on port **_4000_**. Ensure this port is not occupied by another application.

## Available Routes
- **GET** `/receipts`: Returns a list of all receipts in formatted JSON.
- **POST** `/receipts/process`: Accepts a JSON body and responds with an `id` field indicating the assigned receipt ID.
- **GET** `/receipts/{id}/process`: Retrieves the points assigned to a specific receipt.

## Key Features
- **Modular Design**: The project follows a clean and organized structure.
- **Error Handling**: Includes robust mechanisms to return appropriate error messages for invalid inputs.
- **Persistent Points**: Points associated with a receipt are stored and not recalculated once processed.
