package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	subCmdBin = "bin"
	subCmdOct = "oct"
	subCmdDec = "dec"
	subCmdHex = "hex"
)

const (
	title    = `| Binary  | Oct | Dec | Hex| Abbr/Glyph[-name]                   |`
	subTitle = `|---------|-----|-----|----|-------------------------------------|`
)

var asciiList = []string{
	"| 0000000 | 0   | 0   | 0  | NUL - Null                          |",
	"| 0000001 | 1   | 1   | 1  | SOH - Start of Heading              |",
	"| 0000010 | 2   | 2   | 2  | STX - Start of Text                 |",
	"| 0000011 | 3   | 3   | 3  | ETX - End of Text                   |",
	"| 0000100 | 4   | 4   | 4  | EOT - End of Transmission           |",
	"| 0000101 | 5   | 5   | 5  | ENQ - Enquiry                       |",
	"| 0000110 | 6   | 6   | 6  | ACK - Acknowledgement               |",
	"| 0000111 | 7   | 7   | 7  | BEL - Bell                          |",
	"| 0001000 | 10  | 8   | 8  | BS  - Backspace                     |",
	"| 0001001 | 11  | 9   | 9  | HT  - Horizontal Tab                |",
	"| 0001010 | 12  | 10  | 0A | LF  - Line Feed                     |",
	"| 0001011 | 13  | 11  | 0B | VT  - Vertical Tab                  |",
	"| 0001100 | 14  | 12  | 0C | FF  - Form Feed                     |",
	"| 0001101 | 15  | 13  | 0D | CR  - Carriage Return               |",
	"| 0001110 | 16  | 14  | 0E | SO  - Shift Out                     |",
	"| 0001111 | 17  | 15  | 0F | SI  - Shift In                      |",
	"| 0010000 | 20  | 16  | 10 | DLE - Data Link Escape              |",
	"| 0010001 | 21  | 17  | 11 | DC1 - Device Control 1 (often XON)  |",
	"| 0010010 | 22  | 18  | 12 | DC2 - Device Control 2              |",
	"| 0010011 | 23  | 19  | 13 | DC3 - Device Control 3 (often XOFF) |",
	"| 0010100 | 24  | 20  | 14 | DC4 - Device Control 4              |",
	"| 0010101 | 25  | 21  | 15 | NAK - Negative Acknowledgement      |",
	"| 0010110 | 26  | 22  | 16 | SYN - Synchronous Idle              |",
	"| 0010111 | 27  | 23  | 17 | ETB - End of Transmission Block     |",
	"| 0011000 | 30  | 24  | 18 | CAN - Cancel                        |",
	"| 0011001 | 31  | 25  | 19 | EM  - End of Medium                 |",
	"| 0011010 | 32  | 26  | 1A | SUB - Substitute                    |",
	"| 0011011 | 33  | 27  | 1B | ESC - Escape                        |",
	"| 0011100 | 34  | 28  | 1C | FS  - File Separator                |",
	"| 0011101 | 35  | 29  | 1D | GS  - Group Separator               |",
	"| 0011110 | 36  | 30  | 1E | RS  - Record Separator              |",
	"| 0011111 | 37  | 31  | 1F | US  - Unit Separator                |",
	"| 0100000 | 40  | 32  | 20 | space                               |",
	"| 0100001 | 41  | 33  | 21 | !                                   |",
	"| 0100010 | 42  | 34  | 22 | \"                                   |",
	"| 0100011 | 43  | 35  | 23 | #                                   |",
	"| 0100100 | 44  | 36  | 24 | $                                   |",
	"| 0100101 | 45  | 37  | 25 | %                                   |",
	"| 0100110 | 46  | 38  | 26 | &                                   |",
	"| 0100111 | 47  | 39  | 27 | '                                   |",
	"| 0101000 | 50  | 40  | 28 | (                                   |",
	"| 0101001 | 51  | 41  | 29 | )                                   |",
	"| 0101010 | 52  | 42  | 2A | *                                   |",
	"| 0101011 | 53  | 43  | 2B | +                                   |",
	"| 0101100 | 54  | 44  | 2C | ,                                   |",
	"| 0101101 | 55  | 45  | 2D | -                                   |",
	"| 0101110 | 56  | 46  | 2E | .                                   |",
	"| 0101111 | 57  | 47  | 2F | /                                   |",
	"| 0110000 | 60  | 48  | 30 | 0                                   |",
	"| 0110001 | 61  | 49  | 31 | 1                                   |",
	"| 0110010 | 62  | 50  | 32 | 2                                   |",
	"| 0110011 | 63  | 51  | 33 | 3                                   |",
	"| 0110100 | 64  | 52  | 34 | 4                                   |",
	"| 0110101 | 65  | 53  | 35 | 5                                   |",
	"| 0110110 | 66  | 54  | 36 | 6                                   |",
	"| 0110111 | 67  | 55  | 37 | 7                                   |",
	"| 0111000 | 70  | 56  | 38 | 8                                   |",
	"| 0111001 | 71  | 57  | 39 | 9                                   |",
	"| 0111010 | 72  | 58  | 3A | :                                   |",
	"| 0111011 | 73  | 59  | 3B | ;                                   |",
	"| 0111100 | 74  | 60  | 3C | <                                   |",
	"| 0111101 | 75  | 61  | 3D | =                                   |",
	"| 0111110 | 76  | 62  | 3E | >                                   |",
	"| 0111111 | 77  | 63  | 3F | ?                                   |",
	"| 1000000 | 100 | 64  | 40 | @                                   |",
	"| 1000001 | 101 | 65  | 41 | A                                   |",
	"| 1000010 | 102 | 66  | 42 | B                                   |",
	"| 1000011 | 103 | 67  | 43 | C                                   |",
	"| 1000100 | 104 | 68  | 44 | D                                   |",
	"| 1000101 | 105 | 69  | 45 | E                                   |",
	"| 1000110 | 106 | 70  | 46 | F                                   |",
	"| 1000111 | 107 | 71  | 47 | G                                   |",
	"| 1001000 | 110 | 72  | 48 | H                                   |",
	"| 1001001 | 111 | 73  | 49 | I                                   |",
	"| 1001010 | 112 | 74  | 4A | J                                   |",
	"| 1001011 | 113 | 75  | 4B | K                                   |",
	"| 1001100 | 114 | 76  | 4C | L                                   |",
	"| 1001101 | 115 | 77  | 4D | M                                   |",
	"| 1001110 | 116 | 78  | 4E | N                                   |",
	"| 1001111 | 117 | 79  | 4F | O                                   |",
	"| 1010000 | 120 | 80  | 50 | P                                   |",
	"| 1010001 | 121 | 81  | 51 | Q                                   |",
	"| 1010010 | 122 | 82  | 52 | R                                   |",
	"| 1010011 | 123 | 83  | 53 | S                                   |",
	"| 1010100 | 124 | 84  | 54 | T                                   |",
	"| 1010101 | 125 | 85  | 55 | U                                   |",
	"| 1010110 | 126 | 86  | 56 | V                                   |",
	"| 1010111 | 127 | 87  | 57 | W                                   |",
	"| 1011000 | 130 | 88  | 58 | X                                   |",
	"| 1011001 | 131 | 89  | 59 | Y                                   |",
	"| 1011010 | 132 | 90  | 5A | Z                                   |",
	"| 1011011 | 133 | 91  | 5B | [                                   |",
	"| 1011100 | 134 | 92  | 5C | \\                                   |",
	"| 1011101 | 135 | 93  | 5D | ]                                   |",
	"| 1011110 | 136 | 94  | 5E | ^                                   |",
	"| 1011111 | 137 | 95  | 5F | _                                   |",
	"| 1100000 | 140 | 96  | 60 | `                                   |",
	"| 1100001 | 141 | 97  | 61 | a                                   |",
	"| 1100010 | 142 | 98  | 62 | b                                   |",
	"| 1100011 | 143 | 99  | 63 | c                                   |",
	"| 1100100 | 144 | 100 | 64 | d                                   |",
	"| 1100101 | 145 | 101 | 65 | e                                   |",
	"| 1100110 | 146 | 102 | 66 | f                                   |",
	"| 1100111 | 147 | 103 | 67 | g                                   |",
	"| 1101000 | 150 | 104 | 68 | h                                   |",
	"| 1101001 | 151 | 105 | 69 | i                                   |",
	"| 1101010 | 152 | 106 | 6A | j                                   |",
	"| 1101011 | 153 | 107 | 6B | k                                   |",
	"| 1101100 | 154 | 108 | 6C | l                                   |",
	"| 1101101 | 155 | 109 | 6D | m                                   |",
	"| 1101110 | 156 | 110 | 6E | n                                   |",
	"| 1101111 | 157 | 111 | 6F | o                                   |",
	"| 1110000 | 160 | 112 | 70 | p                                   |",
	"| 1110001 | 161 | 113 | 71 | q                                   |",
	"| 1110010 | 162 | 114 | 72 | r                                   |",
	"| 1110011 | 163 | 115 | 73 | s                                   |",
	"| 1110100 | 164 | 116 | 74 | t                                   |",
	"| 1110101 | 165 | 117 | 75 | u                                   |",
	"| 1110110 | 166 | 118 | 76 | v                                   |",
	"| 1110111 | 167 | 119 | 77 | w                                   |",
	"| 1111000 | 170 | 120 | 78 | x                                   |",
	"| 1111001 | 171 | 121 | 79 | y                                   |",
	"| 1111010 | 172 | 122 | 7A | z                                   |",
	"| 1111011 | 173 | 123 | 7B | {                                   |",
	"| 1111100 | 174 | 124 | 7C | |                                   |",
	"| 1111101 | 175 | 125 | 7D | }                                   |",
	"| 1111110 | 176 | 126 | 7E | ~                                   |",
	"| 1111111 | 177 | 127 | 7F | DEL - Delete                        |",
}

func indexASCII(subCmd, repre string) error {
	var idx int64
	var err error
	switch subCmd {
	case subCmdBin:
		idx, err = strconv.ParseInt(repre, 2, 8)
		if err != nil || idx < 0 || idx > 127 {
			return fmt.Errorf("'bin' subcommand exepect argument(0000000-1111111)")
		}

	case subCmdOct:
		idx, err = strconv.ParseInt(repre, 8, 8)
		if err != nil || idx < 0 || idx > 127 {
			return fmt.Errorf("'oct' subcommand exepect argument(0-177)")
		}

	case subCmdDec:
		idx, err = strconv.ParseInt(repre, 10, 8)
		if err != nil || idx < 0 || idx > 127 {
			return fmt.Errorf("'bin' subcommand exepect argument(0-127)")
		}

	case subCmdHex:
		idx, err = strconv.ParseInt(repre, 16, 8)
		if err != nil || idx < 0 || idx > 127 {
			return fmt.Errorf("'bin' subcommand exepect argument(0-7F)")
		}
	default:
		return fmt.Errorf("expected 'bin', 'oct', 'dec' or 'hex' subcommands")
	}

	fmt.Println(title)
	fmt.Println(subTitle)
	fmt.Println(asciiList[idx])
	return nil
}

func main() {

	if len(os.Args) < 2 {
		exitWithError(fmt.Errorf("expected 'bin', 'oct', 'dec' or 'hex' subcommands"))
	}

	var err error
	switch os.Args[1] {
	case subCmdBin, subCmdOct, subCmdDec, subCmdHex:
		if len(os.Args) < 3 {
			exitWithError(fmt.Errorf("'%s' subcommand need argumen", os.Args[1]))
		}
		err = indexASCII(os.Args[1], os.Args[2])
		if err != nil {
			exitWithError(err)
		}
	default:
		exitWithError(fmt.Errorf("expected 'bin', 'oct', 'dec' or 'hex' subcommands"))
	}

}

func exitWithError(err error) {
	fmt.Println("error with: ", err.Error())
	os.Exit(1)
}
