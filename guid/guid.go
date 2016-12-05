package guid

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"io"
	"os"
	"sync/atomic"
	"time"
)

// objectIDCounter is atomically incremented when generating a new ObjectID
// using NewObjectID() function. It's used as a counter part of an id.
// objectIDCounter object id counter
var objectIDCounter uint32

// machineID stores machine id generated once and used in subsequent calls
// to NewObjectID function.
var machineID = readMachineID()

// ObjectID is a unique ID identifying a BSON value. It must be exactly 12 bytes
// long. MongoDB objects by default have such a property set in their "_id"
// property.
//
// http://www.mongodb.org/display/DOCS/Object+IDs
type ObjectID string

// readMachineID generates machine id and puts it into the machineID global
// variable. If this function fails to get the hostname, it will cause
// a runtime error.
func readMachineID() []byte {
	var sum [3]byte
	id := sum[:]
	hostname, err1 := os.Hostname()
	if err1 != nil {
		_, err2 := io.ReadFull(rand.Reader, id)
		if err2 != nil {
			// panic(fmt.Errorf("cannot get hostname: %v; %v", err1, err2))
		}
		return id
	}
	hw := md5.New()
	hw.Write([]byte(hostname))
	copy(id, hw.Sum(nil))
	// fmt.Println("readMachineID:" + string(id))
	return id
}

// NewObjectID returns a new unique ObjectID.
// 4byte 时间，
// 3byte 机器ID
// 2byte pid
// 3byte 自增ID
func NewObjectID() ObjectID {
	var b [12]byte
	// Timestamp, 4 bytes, big endian
	binary.BigEndian.PutUint32(b[:], uint32(time.Now().Unix()))
	// Machine, first 3 bytes of md5(hostname)
	b[4] = machineID[0]
	b[5] = machineID[1]
	b[6] = machineID[2]
	// Pid, 2 bytes, specs don't specify endianness, but we use big endian.
	pid := os.Getpid()
	b[7] = byte(pid >> 8)
	b[8] = byte(pid)
	// Increment, 3 bytes, big endian
	i := atomic.AddUint32(&objectIDCounter, 1)
	b[9] = byte(i >> 16)
	b[10] = byte(i >> 8)
	b[11] = byte(i)
	return ObjectID(b[:])
}

// Hex returns a hex representation of the ObjectID.
// 返回16进制对应的字符串
func (id ObjectID) Hex() string {
	return hex.EncodeToString([]byte(id))
}

//e.g
// func main() {
// fmt.Println(NewObjectID().Hex())
// }
