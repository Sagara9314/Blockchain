/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright statusship.  The ASF licenses this file
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
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

// Define the Passport structure, with 4 properties.  Structure tags are used by encoding/json library
type Passport struct {
	Fname   string `json:"fname"`
	Gender  string `json:"gender"`
	Dob string `json:"dob"`
	Status  string `json:"status"`
	Place  string `json:"place"`
}

/*
 * The Init method is called when the Smart Contract "fabPassport" is instantiated by the blockchain network
 * Best practice is to have any Ledger initialization in separate function -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method is called as a result of an application request to run the Smart Contract "fabPassport"
 * The calling application program has also specified the particular smart contract function to be called, with arguments
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "queryPassport" {
		return s.queryPassport(APIstub, args)
	} else if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "createPassport" {
		return s.createPassport(APIstub, args)
	} else if function == "queryAllPassports" {
		return s.queryAllPassports(APIstub)
	} else if function == "changePassportStatus" {
		return s.changePassportStatus(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) queryPassport(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	passportAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(passportAsBytes)
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	passports := []Passport{
		Passport{Fname: "Sagar", Gender: "Male", Dob: "9-10-1996", Status: "Applied", Place: "Bangalore"},
		Passport{Fname: "Lokesh", Gender: "Male", Dob: "11-5-2002", Status: "Issued", Place: "Mysore"},
		Passport{Fname: "Akash", Gender: "Male", Dob: "14-11-1997", Status: "Issued", Place: "Mandya"},
		Passport{Fname: "Tanu", Gender: "Female", Dob: "11-9-1999", Status: "Applied", Place: "Mangalore"},
		Passport{Fname: "Abhi", Gender: "Male", Dob: "10-12-1998", Status: "Inprocess", Place: "Mumbai"},
		Passport{Fname: "Karna", Gender: "Male", Dob: "5-5-1996", Status: "Issued", Place: "Shivmoga"},
		Passport{Fname: "Sunil", Gender: "Male", Dob: "27-11-1996", Status: "Applied", Place: "Chennai"},
		Passport{Fname: "Raavan", Gender: "Male", Dob: "18-9-1990", Status: "Inprocess", Place: "Delhi"},
		Passport{Fname: "Madhurya", Gender: "Female", Dob: "10-9-2000", Status: "Applied", Place: "Bangalore"},
		Passport{Fname: "Arjun", Gender: "Male", Dob: "15-6-1997", Status: "Issued", Place: "Mumbai"},
	}

	i := 0
	for i < len(passports) {
		fmt.Println("i is ", i)
		passportAsBytes, _ := json.Marshal(passports[i])
		APIstub.PutState("PASSPORT"+strconv.Itoa(i), passportAsBytes)
		fmt.Println("Added", passports[i])
		i = i + 1
	}

	return shim.Success(nil)
}

func (s *SmartContract) createPassport(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}

	var passport = Passport{Fname: args[1], Gender: args[2], Dob: args[3], Status: args[4], Place: args[5]}

	passportAsBytes, _ := json.Marshal(passport)
	APIstub.PutState(args[0], passportAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryAllPassports(APIstub shim.ChaincodeStubInterface) sc.Response {

	startKey := "PASSPORT0"
	endKey := "PASSPORT999"

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

	fmt.Printf("- queryAllPassports:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}

func (s *SmartContract) changePassportStatus(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	passportAsBytes, _ := APIstub.GetState(args[0])
	passport := Passport{}

	json.Unmarshal(passportAsBytes, &passport)
	passport.Status = args[1]

	passportAsBytes, _ = json.Marshal(passport)
	APIstub.PutState(args[0], passportAsBytes)

	return shim.Success(nil)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
