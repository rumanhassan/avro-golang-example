/**
 * Autogenerated by Avro
 * 
 * DO NOT EDIT DIRECTLY
 */
package test;

import (
	"github.com/BBN-D/avro-golang-example/src"
)
import "errors"

// TestRecordSchema is an AvroSchema object
var TestRecordSchema = src.AvroRegisterSchema(TestRecordFromRecord, "{\"type\":\"record\",\"name\":\"TestRecord\",\"namespace\":\"org.apache.avro.test\",\"doc\":\"A TestRecord.\",\"fields\":[{\"name\":\"name\",\"type\":{\"type\":\"string\",\"avro.java.string\":\"String\"},\"default\":\"foo\",\"order\":\"ignore\"},{\"name\":\"kind\",\"type\":{\"type\":\"enum\",\"name\":\"Kind\",\"doc\":\"A kind of record.\",\"symbols\":[\"FOO\",\"BAR\",\"BAZ\"],\"aliases\":[\"org.foo.KindOf\"]},\"doc\":\"The kind of record.\",\"order\":\"descending\"},{\"name\":\"hash\",\"type\":{\"type\":\"fixed\",\"name\":\"MD5\",\"doc\":\"An MD5 hash.\",\"size\":16,\"foo\":\"bar\"},\"default\":\"0000000000000000\"},{\"name\":\"nullableHash\",\"type\":[\"null\",\"MD5\"],\"default\":null,\"aliases\":[\"hash\",\"hsh\"]},{\"name\":\"value\",\"type\":\"double\",\"default\":\"NaN\"},{\"name\":\"average\",\"type\":\"float\",\"default\":\"-Infinity\"}],\"my-property\":{\"key\":3}}")

/** A TestRecord. */
type TestRecord struct {
	Name string 
	/** The kind of record. */
	Kind Kind 
	Hash *MD5 
	NullableHash interface{} /* union types: null, org.apache.avro.test.MD5 */
	Value float64 
	Average float32 
}

// ToAvroContainer converts TestRecord to a byte array
func (this *TestRecord) ToAvroContainer() ([]byte, error) {
	return src.ToAvroContainer([]src.AvroObject{this})
}

// ToJSON converts TestRecord to a JSON string
func (this *TestRecord) ToJSON(prettyPrint bool) (string, error) {
	record, err := this.ToRecord()
	if err != nil {
		return "", errors.New("ToJSON: " + err.Error())
	}
	return record.ToJSON(prettyPrint)
}

// TestRecordFromAvroContainer returns a slice
// of TestRecord objects
func TestRecordFromAvroContainer(bytes []byte) ([]src.AvroObject, error) {
	return src.FromAvroContainer(bytes, TestRecordSchema)
}

//
// AvroObject interface; should not need to use these directly
//

// GetSchema returns the schema object associated with this type
func (this *TestRecord) GetSchema() *src.AvroSchema {
	return TestRecordSchema
}


func TestRecordFromRecord(record *src.AvroRecord) (src.AvroObject, error) {
	var err error
	obj := new(TestRecord)

	obj.Name, err = record.Get("name").GetString()
	if err != nil {
		return nil, errors.New("get[Name]: " + err.Error())
	}
	obj.Kind.Value, err = record.Get("kind").GetEnum()
	if err != nil {
		return nil, errors.New("get[Kind]: " + err.Error())
	}
	obj_Hash, err := MD5FromRecord(record.Get("hash").ToRecord())
	if err != nil {
		return nil, errors.New("get[Hash]: " + err.Error())
	}
	obj.Hash = obj_Hash.(*MD5) 
	obj.NullableHash, err = record.Get("nullableHash").GetUnion() 
	if err != nil {
		return nil, errors.New("get[NullableHash]: " + err.Error())
	}
	obj.Value, err = record.Get("value").GetDouble()
	if err != nil {
		return nil, errors.New("get[Value]: " + err.Error())
	}
	obj.Average, err = record.Get("average").GetFloat()
	if err != nil {
		return nil, errors.New("get[Average]: " + err.Error())
	}
	
	return obj, nil
}



func (this *TestRecord) ToRecord() (*src.AvroRecord, error) {
	record := this.GetSchema().CreateRecord()
	

	obj_Name, err := src.DatumFromString(this.Name)

	if err != nil {
		return nil, errors.New("set[name]: " + err.Error())
	}
	if err = record.Set("name", obj_Name); err != nil {
		return nil, err
	}
	

	obj_Kind, err := src.DatumFromEnum(KindSchema, this.Kind.Value)

	if err != nil {
		return nil, errors.New("set[kind]: " + err.Error())
	}
	if err = record.Set("kind", obj_Kind); err != nil {
		return nil, err
	}
	


	obj_Hash_rec, err := this.Hash.ToRecord()
	if err != nil {
		return nil, errors.New("set[hash]: " + err.Error())
	}
	obj_Hash, err := src.DatumFromRecord(obj_Hash_rec)
	

	if err != nil {
		return nil, errors.New("set[hash]: " + err.Error())
	}
	if err = record.Set("hash", obj_Hash); err != nil {
		return nil, err
	}
	

	obj_NullableHash, err := src.DatumFromUnion(this.GetSchema().GetFieldSchemaUnsafe("nullableHash"), this.NullableHash)

	if err != nil {
		return nil, errors.New("set[nullableHash]: " + err.Error())
	}
	if err = record.Set("nullableHash", obj_NullableHash); err != nil {
		return nil, err
	}
	

	obj_Value, err := src.DatumFromDouble(this.Value)

	if err != nil {
		return nil, errors.New("set[value]: " + err.Error())
	}
	if err = record.Set("value", obj_Value); err != nil {
		return nil, err
	}
	

	obj_Average, err := src.DatumFromFloat(this.Average)

	if err != nil {
		return nil, errors.New("set[average]: " + err.Error())
	}
	if err = record.Set("average", obj_Average); err != nil {
		return nil, err
	}
	
	return record, nil
}
