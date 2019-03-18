/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

package main

/* Imports
 * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
 * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
 */
import (
//	"bytes"
	"encoding/json"
	"fmt"
//	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the car structure, with 4 properties.  Structure tags are used by encoding/json library
type Car struct {
	Chassisnumber   string `json:"chassisnumber"`
	Enginenumber  string `json:"enginenumber"`
	Colour string `json:"colour"`
	Model  string `json:"model"`
	ManufacturingYear string `json:"manufacturingdate"`
}

type Customer struct {
	Name   string `json:"name"`
	Age  int8 `json:"age"`
	Phonenum int16 `json:"phonenum"`
	Drivinglicencenum  string `json:"drnum"`
}

type InsurancePolicy struct {
	CarId   string `json:"carid"`
	Instype  int8 `json:"instype"`
	Status bool `json:"status"`
}

type Loan struct {
	CarId   string `json:"carid"`
	CustomerId   string `json:"customerid"`
	Amount string `json:"amount"`
}

type RTADetails struct {
	CarId   string `json:"carid"`
	RTA_Number  int8 `json:"rtanumber"`
	CustomerId   string `json:"customerid"`
	InsuranceId   string `json:"insid"`
	HypotheticationId string `json:"hypotheticationid"`
	SaleTransactionId string `json:"saletransactionid"`
	RegistrationEnddate string `json:"registrationenddate"`
}

type Saletransaction struct {
	CarId   string `json:"carid"`
	CustomerId   string `json:"customerid"`
}

type InsuranceClaim struct {
	InsId   string `json:"insid"`
	CustomerId   string `json:"customerid"`
	ClaimDesc string `json:"claimdesc"`
}

type ServiceDetails struct {
	CarId   string `json:"carid"`
	CustomerId   string `json:"customerid"`
	Servicedetails string `json:"servicedetails"`
}
/*
 * The Init method is called when the Smart Contract "fabcar" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabcar"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "produceCar" {
		return s.produceCar(APIstub, args)
//	} else if function == "queryCar" {
//		return s.queryCar(APIstub, args)
//	} else if function == "transertoDealer" {
//		return s.transertoDealer(APIstub, args)
//	} else if function == "queryAllCars" {
//		return s.queryAllCars(APIstub)
//	} else if function == "addInsurace" {
//		return s.addInsurace(APIstub, args)
//	} else if function == "addRTANumber" {
//		return s.addRTANumber(APIstub, args)
//	} else if function == "changeCarOwner"{
//		return s.changeCarOwner(APIstub, args)
//	} else if function == "addService" {
//		return s.addService(APIstub, args)
//	} else if function == "addnewClaim" {
//		return s.addnewClaim(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

/*func (s *SmartContract) queryCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(carAsBytes)
}*/

func (s *SmartContract) produceCar(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	var car = Car{Chassisnumber: args[1], Enginenumber: args[2], Model: args[3], ManufacturingYear: args[4], Colour: args[5]}

	carAsBytes, _ := json.Marshal(car)
	APIstub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}

/*func (s *SmartContract) queryAllCars(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "CAR0"
	endKey := "CAR999"

	resultsIterator, err := APIstub.GetStateByRange(startKey, endKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllCars:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}*/

/*func (s *SmartContract) changeCarOwner(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	carAsBytes, _ := APIstub.GetState(args[0])
	car := Car{}

	json.Unmarshal(carAsBytes, &car)
	car.Owner = args[1]

	carAsBytes, _ = json.Marshal(car)
	APIstub.PutState(args[0], carAsBytes)

	return shim.Success(nil)
}*/

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

