package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Certificate struct {
	CertificateID					string	`json:"CertificateID"`
	InstituteID						string	`json:"InstituteID"`
	AffInstituteID					string	`json:"AffInstituteID"`
	CourseID						string	`json:"CourseID"`
	BatchID							string	`json:"BatchID"`
	StudentID						string	`json:"StudentID"`
	Specialization					string	`json:"Specialization"`
	ScoreEarned						string	`json:"ScoreEarned"`
	TotalScore						string	`json:"TotalScore"`
	CGPA							string	`json:"CGPA"`
	Credits							string	`json:"Credits"`
	CompletionDate					string	`json:"CompletionDate"`
	Status							string	`json:"Status"`
}

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Chaincode_custom Invoke")
	function, args := stub.GetFunctionAndParameters()

	if function == "certify" {
		return t.certify(stub, args)
	}
	if function == "query" {
		return t.query(stub, args)
	}

	return shim.Error("Invalid invoke function name.")
}

func (t *SimpleChaincode) certify(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("certify")
	var certificate Certificate
	var err error

	certificate.CertificateID=args[0]
	certificate.InstituteID=args[1]
	certificate.AffInstituteID=args[2]
	certificate.CourseID=args[3]
	certificate.BatchID=args[4]
	certificate.StudentID=args[5]
	certificate.Specialization=args[6]
	certificate.ScoreEarned=args[7]
	certificate.TotalScore=args[8]
	certificate.CGPA=args[9]
	certificate.Credits=args[10]
	certificate.CompletionDate=args[11]
	certificate.Status="Certified"

	certificateAsBytes, _ := json.Marshal(certificate)
	err = stub.PutState(certificate.CertificateID, certificateAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var CertificateID string
	var err error
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting Certificate no. to query")
	}

	CertificateID = args[0]

	Avalbytes, err := stub.GetState(CertificateID)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + CertificateID + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + CertificateID + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success(Avalbytes)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
