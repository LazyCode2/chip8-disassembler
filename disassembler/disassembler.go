package disassembler

import (
	"fmt"
)

type Instruction func(opcode uint16) string

var OpcodeTable = map[uint16]Instruction{
	0x0: ExecFamily0,
	0x1: ExecJP,
	0x2: ExecCALL,
	0x3: ExecSE,
	0x4: ExecSNE,
	0x5: ExecSE_REG,
	0x6: ExecLD,
	0x7: ExecADD,
	0x8: ExecFamily8,
	0x9: ExecSNE_REG,
	0xA: ExecLD_I,
	0xB: ExecJP_ADDR,
	0xC: ExecRND,
	0xD: ExecDRW,
	0xE: ExecFamilyE,
	0xF: ExecFamilyF,
}

func GetX(opcode uint16) uint16 {
	return (opcode & 0x0F00) >> 8
}

func GetY(opcode uint16) uint16 {
	return (opcode & 0x00F0) >> 4
}

func GetNNN(opcode uint16) uint16 {
	return (opcode & 0x0FFF)
}

func GetNN(opcode uint16) uint16 {
	return (opcode & 0x00FF)
}

func ExecFamily0(opcode uint16) string {
	switch opcode {
	case 0x00E0:
		return fmt.Sprintln("CLS")
	case 0x00EE:
		return fmt.Sprintln("RET")
	default:
		return fmt.Sprint()
	}
}

func ExecJP(opcode uint16) string {
	return fmt.Sprintf("JP %#x\n", GetNNN(opcode))
}

func ExecCALL(opcode uint16) string {
	return fmt.Sprintf("CALL %#x\n", GetNNN(opcode))
}

func ExecSE(opcode uint16) string {
	return fmt.Sprintf("SE V%d, %#x\n", GetX(opcode), GetNN(opcode))
}

func ExecSNE(opcode uint16) string {
	return fmt.Sprintf("SNE V%d, %#x\n", GetX(opcode), GetNN(opcode))
}

func ExecSE_REG(opcode uint16) string {
	return fmt.Sprintf("SE V%d, V%d\n", GetX(opcode), GetY(opcode))
}

func ExecLD(opcode uint16) string {
	return fmt.Sprintf("LD V%d, %#x\n", GetX(opcode), GetNN(opcode))
}

func ExecADD(opcode uint16) string {
	return fmt.Sprintf("ADD V%d, %#x\n", GetX(opcode), GetNN(opcode))
}

func ExecFamily8(opcode uint16) string {
	switch opcode & 0x000F {
	case 0x0:
		return fmt.Sprintf("LD V%d, V%d\n", GetX(opcode), GetY(opcode))
	case 0x1:
		return fmt.Sprintf("OR V%d, V%d\n", GetX(opcode), GetY(opcode))
	case 0x2:
		return fmt.Sprintf("AND V%d, V%d\n", GetX(opcode), GetY(opcode))
	case 0x3:
		return fmt.Sprintf("XOR V%d, V%d\n", GetX(opcode), GetY(opcode))
	case 0x4:
		return fmt.Sprintf("ADD V%d, V%d\n", GetX(opcode), GetY(opcode))
	case 0x5:
		return fmt.Sprintf("SUB V%d, V%d\n", GetX(opcode), GetY(opcode))
	case 0x6:
		return fmt.Sprintf("SHR V%d, V%d\n", GetX(opcode), GetY(opcode))
	case 0x7:
		return fmt.Sprintf("SUBN V%d, V%d\n", GetX(opcode), GetY(opcode))
	case 0xE:
		return fmt.Sprintf("SHL V%d, V%d\n", GetX(opcode), GetY(opcode))

	default:
		return fmt.Sprint()
	}
}

func ExecSNE_REG(opcode uint16) string {
	return fmt.Sprintf("SNE V%d, V%d\n", GetX(opcode), GetY(opcode))
}

func ExecLD_I(opcode uint16) string {
	return fmt.Sprintf("LD I, %#x\n", GetNNN(opcode))
}

func ExecJP_ADDR(opcode uint16) string {
	return fmt.Sprintf("JP V0, %#x\n", GetNNN(opcode))
}

func ExecRND(opcode uint16) string {
	return fmt.Sprintf("RND V%d, %#x\n", GetX(opcode), GetNN(opcode))
}

func ExecDRW(opcode uint16) string {
	return fmt.Sprintf("DRW V%d, V%d, %#x\n", GetX(opcode), GetY(opcode), (opcode & 0x000F))
}

func ExecFamilyE(opcode uint16) string {
	switch opcode & 0x00FF {
	case 0x9E:
		return fmt.Sprintf("SKP V%d\n", GetX(opcode))
	case 0xA1:
		return fmt.Sprintf("SKNP V%d\n", GetX(opcode))
	default:
		return fmt.Sprint()
	}
}

func ExecFamilyF(opcode uint16) string {
	switch opcode & 0x00FF {
	case 0x07:
		return fmt.Sprintf("LD V%d, DT\n", GetX(opcode))
	case 0x0A:
		return fmt.Sprintf("LD V%d, K\n", GetX(opcode))
	case 0x15:
		return fmt.Sprintf("LD DT, V%d\n", GetX(opcode))
	case 0x18:
		return fmt.Sprintf("LD ST, V%d\n", GetX(opcode))
	case 0x1E:
		return fmt.Sprintf("ADD I, V%d\n", GetX(opcode))
	case 0x29:
		return fmt.Sprintf("LD F, V%d\n", GetX(opcode))
	case 0x33:
		return fmt.Sprintf("LD B, V%d\n", GetX(opcode))
	case 0x55:
		return fmt.Sprintf("LD [I], V%d\n", GetX(opcode))
	case 0x65:
		return fmt.Sprintf("LD V%d, [I]\n", GetX(opcode))

	default:
		return fmt.Sprint()
	}
}
