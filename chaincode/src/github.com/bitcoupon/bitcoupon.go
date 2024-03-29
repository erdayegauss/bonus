/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")
//	_, args := stub.GetFunctionAndParameters()
	var err error


	err = stub.PutState(strconv.Itoa(2), []byte(strconv.Itoa(100000)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(strconv.Itoa(3), []byte(strconv.Itoa(100000)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(strconv.Itoa(4), []byte(strconv.Itoa(100000)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(strconv.Itoa(5), []byte(strconv.Itoa(100000)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(strconv.Itoa(6), []byte(strconv.Itoa(100000)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(strconv.Itoa(7), []byte(strconv.Itoa(100000)))
	if err != nil {
		return shim.Error(err.Error())
	}

		err = stub.PutState(strconv.Itoa(8), []byte(strconv.Itoa(100000)))
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(strconv.Itoa(9), []byte(strconv.Itoa(100000)))
		if err != nil {
			return shim.Error(err.Error())
		}


			err = stub.PutState(strconv.Itoa(10), []byte(strconv.Itoa(100000)))
			if err != nil {
				return shim.Error(err.Error())
			}

			err = stub.PutState(strconv.Itoa(11), []byte(strconv.Itoa(100000)))
			if err != nil {
				return shim.Error(err.Error())
			}


				err = stub.PutState(strconv.Itoa(12), []byte(strconv.Itoa(100000)))
				if err != nil {
					return shim.Error(err.Error())
				}

				err = stub.PutState(strconv.Itoa(13), []byte(strconv.Itoa(100000)))
				if err != nil {
					return shim.Error(err.Error())
				}


					err = stub.PutState(strconv.Itoa(14), []byte(strconv.Itoa(100000)))
					if err != nil {
						return shim.Error(err.Error())
					}

					err = stub.PutState(strconv.Itoa(15), []byte(strconv.Itoa(100000)))
					if err != nil {
						return shim.Error(err.Error())
					}


						err = stub.PutState(strconv.Itoa(16), []byte(strconv.Itoa(100000)))
						if err != nil {
							return shim.Error(err.Error())
						}

						err = stub.PutState(strconv.Itoa(17), []byte(strconv.Itoa(100000)))
						if err != nil {
							return shim.Error(err.Error())
						}


							err = stub.PutState(strconv.Itoa(18), []byte(strconv.Itoa(100000)))
							if err != nil {
								return shim.Error(err.Error())
							}

							err = stub.PutState(strconv.Itoa(19), []byte(strconv.Itoa(100000)))
							if err != nil {
								return shim.Error(err.Error())
							}




	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "invoke" {
		// Make payment of X units from A to B
		return t.invoke(stub, args)
	} else if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	} else if function == "query" {
		// the old "Query" is now implemtned in invoke
		return t.query(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A, B string    // Entities
	var Aval, Bval int // Asset holdings
	var X, Y int          // Transaction value
	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	A = args[0]
	B = args[1]

	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Avalbytes == nil {
		return shim.Error("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Bvalbytes == nil {
		return shim.Error("Entity not found")
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}
	Y, err = strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}

	Aval = Aval - X
	Bval = Bval + Y
	fmt.Printf("%s = %d, %s = %d\n", A ,Aval,B, Bval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string // Entities
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
	}

	A = args[0]

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
	fmt.Printf("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
