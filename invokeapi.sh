echo "POST invoke chaincode on peers of Org1 and Org2"
echo
curl -s -X POST \
  http://localhost:4000/channels/mychannel/chaincodes/mycc \
  -H "authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjQwOTQyODksInVzZXJuYW1lIjoiSmltIiwib3JnTmFtZSI6Ik9yZzEiLCJpYXQiOjE1NjQwNTgyODl9.4pJIQD5NaK6I4lLzxJsA4lYQ_BLVWcqAjyvvu-fxv-E" \
  -H "content-type: application/json" \
  -d "{
	\"peers\": [\"peer0.org1.example.com\",\"peer0.org2.example.com\"],
	\"fcn\":\"certify\",
	\"args\":[\"a\",\"b\",\"10\",\"a\",\"b\",\"10\",\"a\",\"b\",\"10\",\"a\",\"b\",\"10\"]
}"
echo
echo


