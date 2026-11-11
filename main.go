package main

import (
	"fmt"
	"os"

	"github.com/LazyCode2/chip8-disassembler/disassembler"
)

func main() {
	programBytes := []uint16{0x00E0, 0x00EE, 0x1234, 0x3400}

	out, err := os.Create("out.asm")
	if err != nil {
		panic(err)
	}

	defer out.Close()

	for _, opcode := range programBytes {
		prefix := (opcode & 0xF000) >> 12

		if handler, exists := disassembler.OpcodeTable[uint16(prefix)]; exists {
			instruction := handler(uint16(opcode))
			fmt.Fprint(out, instruction)
		} else {
			fmt.Fprint(out, "-> Error: Unimplemented opcode prefix 0x\n", prefix)
			return
		}
	}
}
