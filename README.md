## vc2csv
vc2csv converts vCards (https://en.wikipedia.org/wiki/VCard) to csv. 
It reads vCards from standard input and prints them in csv-format to standard output.
'-fields' is mandatory argument to specify which vCard fields/properties are to be included in resulting csv.

Resulting csv-format uses semicolon to separate field/property values and comma to separate parameter values inside fields/properties. 

vc2csv is written in Go.

### Installation
Install and update with go: 
`go get -u github.com/jasilven/vc2csv`

### Usage
To convert 'input.vcf' to 'output.csv' so that only "FN" and "TEL" vCard
fields/properties are extracted run: `$GOPATH/bin/vc2csv -fields "FN TEL" < input.vcf > output.csv`

### Dependencies
No dependencies.
