package disassembler

import (
	"fmt"
)

type Instruction func(opcode uint16)

var OpcodeTable = map[uint16]Instruction{
	0x0: ExecFamily0,
	0x1: ExecJP,
	0x2: ExecCALL,
	0x3: ExecSE_Vx,
}

func GetX(opcode uint16) uint16 {
	return (opcode & 0x0F00) >> 8
}

func GetNNN(opcode uint16) uint16 {
	return (opcode & 0x0FFF)
}

func GetNN(opcode uint16) uint16 {
	return (opcode & 0x00FF)
}

func ExecFamily0(opcode uint16) {
	switch opcode {
	case 0x00E0:
		fmt.Println("CLS")
	case 0x00EE:
		fmt.Println("RET")
	}
}

func ExecJP(opcode uint16) {
	fmt.Printf("JP %#x\n", GetNNN(opcode))
}

func ExecCALL(opcode uint16) {
	fmt.Printf("CALL %#x\n", GetNNN(opcode))
}

func ExecSE_Vx(opcode uint16) {
	fmt.Printf("SE V%d, %#x\n", GetX(opcode), GetNN(opcode))
}
