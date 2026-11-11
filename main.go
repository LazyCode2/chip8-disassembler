package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/LazyCode2/chip8-disassembler/disassembler"
)

func main() {
	romFile := flag.String("in", " ", "input rom filepath")
	outFile := flag.String("out", " ", "output asm filepath")

	flag.Parse()

	file, error := os.Open(*romFile)
	if error != nil {
		panic(error)
	}

	defer file.Close()

	data, _ := io.ReadAll(file)

	out, err := os.Create(*outFile)
	if err != nil {
		panic(err)
	}

	defer out.Close()

	for i := 0; i+1 < len(data); i += 2 {
		opcode := uint16(data[i])<<8 | uint16(data[i+1])

		prefix := (opcode & 0xF000) >> 12

		if handler, exists := disassembler.OpcodeTable[uint16(prefix)]; exists {
			instruction := handler(uint16(opcode))
			fmt.Fprint(out, instruction)
		} else {
			return
		}
	}
}
