package jsq

import "os"
import "log"
import "testing"
import "reflect"
import "encoding/json"
import "io/ioutil"

type TestCase struct {
	Description string        `json:description`
	Input       []interface{} `json:input`
	Output      []interface{} `json:output`
	Query       string        `json:query`
}

func TestJSONStreamQuery(t *testing.T) {

	test := readTestCase("test-cases/01.json")

	log.Println("Runnin testcase ", test.Description)
	jsq := BuildJSONStreamQuerier(test.Query)
	actualOutput := jsq.Query(test.Input)
	if !reflect.DeepEqual(test.Output, actualOutput) {
		t.Error("Actual output ", actualOutput, " is not the expected output ", test.Output)
	}
}

func readTestCase(fileName string) TestCase {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	jsonBytes, _ := ioutil.ReadAll(jsonFile)

	var test TestCase
	json.Unmarshal(jsonBytes, &test)

	return test
}
