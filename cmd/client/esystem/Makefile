local:
	go run main.go -address 127.0.0.1:52898
	go run main.go -serverAddress localhost:52898
online:
	go run main.go -address 34.143.228.170:8080
url:
	go run main.go -address esystemci.fortress-asya.com:8080
send:
    scp /Users/rhickmercado/Documents/Programming/go/src/simplebank/cmd/server/E3-InActiveCID.csv rhickmercado@34.143.228.170:~/.
.PHONY: local online send


scp Loan\ Listing\ Conso\ March\ 2023_1of2.csv rhickmercado@34.143.228.170:~/.

scp /Users/rhickmercado/Documents/Programming/go/src/simplebank/cmd/client/esystem/csv/LnBeneficiary.csv rhickmercado@34.143.228.170:~/.
scp /Users/rhickmercado/Documents/Programming/go/src/simplebank/cmd/client/esystem/csv/ReactivateWriteoff.csv rhickmercado@34.143.228.170:~/.

scp /Users/rhickmercado/Documents/Programming/go/src/simplebank/cmd/client/esystem/csv/SaTrnMaster.csv rhickmercado@34.143.228.170:~/.


scp -P 2222 /Users/rhickmercado/Documents/Programming/go/src/simplebank/cmd/client/esystem/csv/SaTrnMaster.csv rhickmercado@34.142.152.56:/home/rhickmercado


docker cp LnBeneficiary.csv 57792c1eb13f:/var/lib/postgresql/.
docker cp ReactivateWriteoff.csv 57792c1eb13f:/var/lib/postgresql/.

head -1 List1.csv