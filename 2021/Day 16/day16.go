// Package day16 implements the 2021 Advent of Code Day 16 assignment.
// See the readme.md for details on this assignment or visit the Advent
// of Code website: https://adventofcode.com/2021/day/16
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"

	"fileprocessing"
)

// main() prints the Part 1 and Part 2 solutions. It receives a single
// argument specifying the name of the data file containing the hexidecimal
// string input, which is converted to a binary string and then used to
// calculate the Part 1 and Part 2 solutions.
func main() {
	inputFile := os.Args[1]

	fileContents, err := fileprocessing.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	numLines := len(fileContents)
	if numLines <= 0 {
		// invalid input
		log.Fatal(fmt.Errorf("invalid input in %s", inputFile))
	}

	binaryString := convertHexStringToBinaryString(fileContents[0])

	p := new(Packet)
	_ = p.New(binaryString)

	fmt.Printf("Part 1 - Sum of the packet versions: %d\n", p.SumVersion())
	fmt.Printf("Part 2 - Calculated result of the packet payload: %d\n", p.CalculateResult())
}

// These are the different type IDs a Packet can specify.
const (
	PacketSum          int64 = 0
	PacketProduct      int64 = 1
	PacketMinimum      int64 = 2
	PacketMaximum      int64 = 3
	PacketLiteralValue int64 = 4
	PacketGreaterThan  int64 = 5
	PacketLessThan     int64 = 6
	PacketEqual        int64 = 7
)

type Packet struct {
	Version    int64
	TypeID     int64
	Val        int64
	SubPackets []Packet
}

// String() prints the Packet struct without any formatting.
func (p *Packet) String() string {
	pJSON, _ := json.Marshal(p)
	return string(pJSON)
}

// StringFormatted() prints the Packet struct with formatting to
// that makes it easier to see the nested Packets contained inside.
func (p *Packet) StringFormatted() string {
	pJSON, _ := json.MarshalIndent(p, "", "  ")
	return string(pJSON)
}

// New() parses an input string into a Packet and its sub-packets. The
// input string is a binary string (e.g. '1010010').
func (p *Packet) New(input string) (index int) {
	if len(input) < 6 {
		// can't extract the version and type
		return
	}

	// parse version and type
	p.Version, _ = convertBinaryStringToInt(input[0:3])
	p.TypeID, _ = convertBinaryStringToInt(input[3:6])

	currentIndex := 6

	// if type = 4 then literal value, parse the value and return
	if p.TypeID == PacketLiteralValue {
		binaryValue := ""

		//   grab the next 5 bits
		for {
			if currentIndex+5 > len(input) {
				return
			}
			// append the following 4 bits after the signal bit into a binary string
			binaryValue += input[currentIndex+1 : currentIndex+5]

			//   if the first bit is 0, then this is the last packet
			if input[currentIndex] == '0' {
				currentIndex += 5
				break
			}

			//     grab the next 5 and repeat
			currentIndex += 5
		}

		p.Val, _ = convertBinaryStringToInt(binaryValue)

	} else {
		// if type != 4, then operator value, parse the length type ID
		lengthTypeID := rune(input[currentIndex])
		currentIndex++

		//   if length type ID = 0, parse the next 15 bits (length in bits of sub-packets)
		if lengthTypeID == '0' {
			numBits, _ := convertBinaryStringToInt(input[currentIndex : currentIndex+15])
			currentIndex += 15

			startIndex := currentIndex

			for numBits > int64(currentIndex)-int64(startIndex) {
				var newPacket *Packet = new(Packet)
				currentIndex += newPacket.New(input[currentIndex:])
				p.SubPackets = append(p.SubPackets, *newPacket)
			}
		}

		//   if length type ID = 1, parse the next 11 bits (number of sub-packets to parse)
		if lengthTypeID == '1' {
			//     create new packets until the number of packets created is equal to the number of sub-packets
			//     in the number of sub-packets to parse
			numPackets, _ := convertBinaryStringToInt(input[currentIndex : currentIndex+11])
			currentIndex += 11

			p.SubPackets = make([]Packet, numPackets)
			for i := 0; i < int(numPackets); i++ {
				var newPacket *Packet = new(Packet)
				currentIndex += newPacket.New(input[currentIndex:])

				p.SubPackets[i] = *newPacket
			}
		}
	}

	return currentIndex
}

// SumVersion() will iterate over a packet and its sub-packets
// and return the sum of the 'Version' values. A packet's 'Version'
// is the value contained in the first three bits of a packet's input
// string.
func (p *Packet) SumVersion() int64 {
	if len(p.SubPackets) == 0 {
		return p.Version
	}

	var sum int64 = 0
	for i := 0; i < len(p.SubPackets); i++ {
		sum += p.SubPackets[i].SumVersion()
	}

	return sum + p.Version
}

// CalculateResult() inspects the Packet type ID and calculates its
// value based on the operation(s) and value(s) of its sub-packets. If
// a Packet  is a literal value, the specified packet value is returned.
func (p *Packet) CalculateResult() int64 {
	var result int64

	switch p.TypeID {
	case PacketSum:
		result = 0
		for i := 0; i < len(p.SubPackets); i++ {
			result += p.SubPackets[i].CalculateResult()
		}
	case PacketProduct:
		result = 1
		for i := 0; i < len(p.SubPackets); i++ {
			result *= p.SubPackets[i].CalculateResult()
		}
	case PacketMinimum:
		result = math.MaxInt64
		for i := 0; i < len(p.SubPackets); i++ {
			calculatedResult := p.SubPackets[i].CalculateResult()
			if result > calculatedResult {
				result = calculatedResult
			}
		}
	case PacketMaximum:
		result = 0
		for i := 0; i < len(p.SubPackets); i++ {
			calculatedResult := p.SubPackets[i].CalculateResult()
			if result < calculatedResult {
				result = calculatedResult
			}
		}
	case PacketLiteralValue:
		return p.Val
	case PacketGreaterThan:
		// has exactly two subpackets, per the instructions
		if len(p.SubPackets) == 2 {
			if p.SubPackets[0].CalculateResult() > p.SubPackets[1].CalculateResult() {
				result = 1
			} else {
				result = 0
			}
		}
	case PacketLessThan:
		// has exactly two subpackets, per the instructions
		if len(p.SubPackets) == 2 {
			if p.SubPackets[0].CalculateResult() < p.SubPackets[1].CalculateResult() {
				result = 1
			} else {
				result = 0
			}
		}
	case PacketEqual:
		// has exactly two subpackets, per the instructions
		if len(p.SubPackets) == 2 {
			if p.SubPackets[0].CalculateResult() == p.SubPackets[1].CalculateResult() {
				result = 1
			} else {
				result = 0
			}
		}
	}

	p.Val = result
	return result
}

// convertHexRuneToBinaryString() takes a specified rune (character) and
// returns a 4-bit binary representation of the rune's hexadecimal value.
func convertHexRuneToBinaryString(h rune) string {
	switch h {
	case '0':
		return "0000"
	case '1':
		return "0001"
	case '2':
		return "0010"
	case '3':
		return "0011"
	case '4':
		return "0100"
	case '5':
		return "0101"
	case '6':
		return "0110"
	case '7':
		return "0111"
	case '8':
		return "1000"
	case '9':
		return "1001"
	case 'A':
		return "1010"
	case 'B':
		return "1011"
	case 'C':
		return "1100"
	case 'D':
		return "1101"
	case 'E':
		return "1110"
	case 'F':
		return "1111"
	}

	return ""
}

// convertBinaryStringToInt() takes a binary string and converts it into a
// 64-bit integer value.
func convertBinaryStringToInt(s string) (int64, error) {
	i, err := strconv.ParseInt(s, 2, 64)

	return i, err
}

// convertHexStringToBinaryString() takes a hexadecimal string and converts
// it to a binary string representation.
func convertHexStringToBinaryString(input string) (output string) {
	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		output += convertHexRuneToBinaryString(runes[i])
	}

	return
}
