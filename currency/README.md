go run main.go

go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest

grpcurl --plaintext -d '{"Base":"USD","Destination":"INR"}' localhost:9092 currency.Currency.GetRate