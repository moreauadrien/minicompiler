// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case r == 9: // ['\t','\t']
			return 1
		case r == 10: // ['\n','\n']
			return 1
		case r == 13: // ['\r','\r']
			return 1
		case r == 32: // [' ',' ']
			return 1
		case r == 33: // ['!','!']
			return 2
		case r == 38: // ['&','&']
			return 3
		case r == 40: // ['(','(']
			return 4
		case r == 41: // [')',')']
			return 5
		case r == 42: // ['*','*']
			return 6
		case r == 43: // ['+','+']
			return 7
		case r == 45: // ['-','-']
			return 8
		case r == 47: // ['/','/']
			return 9
		case r == 48: // ['0','0']
			return 10
		case 49 <= r && r <= 57: // ['1','9']
			return 11
		case r == 59: // [';',';']
			return 12
		case r == 60: // ['<','<']
			return 13
		case r == 61: // ['=','=']
			return 14
		case r == 64: // ['@','@']
			return 15
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 91: // ['[','[']
			return 17
		case r == 93: // [']',']']
			return 18
		case r == 95: // ['_','_']
			return 19
		case 97 <= r && r <= 100: // ['a','d']
			return 16
		case r == 101: // ['e','e']
			return 20
		case 102 <= r && r <= 104: // ['f','h']
			return 16
		case r == 105: // ['i','i']
			return 21
		case 106 <= r && r <= 118: // ['j','v']
			return 16
		case r == 119: // ['w','w']
			return 22
		case 120 <= r && r <= 122: // ['x','z']
			return 16
		case r == 123: // ['{','{']
			return 23
		case r == 125: // ['}','}']
			return 24
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 25
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S4
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 26
		case r == 47: // ['/','/']
			return 27
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 28
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		case r == 61: // ['=','=']
			return 29
		}
		return NoState
	},
	// S15
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 16
		}
		return NoState
	},
	// S17
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S18
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 16
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 19
		case 97 <= r && r <= 107: // ['a','k']
			return 16
		case r == 108: // ['l','l']
			return 31
		case 109 <= r && r <= 122: // ['m','z']
			return 16
		}
		return NoState
	},
	// S21
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 19
		case 97 <= r && r <= 101: // ['a','e']
			return 16
		case r == 102: // ['f','f']
			return 32
		case 103 <= r && r <= 122: // ['g','z']
			return 16
		}
		return NoState
	},
	// S22
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 19
		case r == 97: // ['a','a']
			return 33
		case 98 <= r && r <= 103: // ['b','g']
			return 16
		case r == 104: // ['h','h']
			return 34
		case 105 <= r && r <= 122: // ['i','z']
			return 16
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S25
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 35
		default:
			return 26
		}
	},
	// S27
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 36
		default:
			return 27
		}
	},
	// S28
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 28
		}
		return NoState
	},
	// S29
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S30
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 16
		}
		return NoState
	},
	// S31
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 19
		case 97 <= r && r <= 114: // ['a','r']
			return 16
		case r == 115: // ['s','s']
			return 37
		case 116 <= r && r <= 122: // ['t','z']
			return 16
		}
		return NoState
	},
	// S32
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 16
		}
		return NoState
	},
	// S33
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 19
		case 97 <= r && r <= 104: // ['a','h']
			return 16
		case r == 105: // ['i','i']
			return 38
		case 106 <= r && r <= 122: // ['j','z']
			return 16
		}
		return NoState
	},
	// S34
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 19
		case 97 <= r && r <= 104: // ['a','h']
			return 16
		case r == 105: // ['i','i']
			return 39
		case 106 <= r && r <= 122: // ['j','z']
			return 16
		}
		return NoState
	},
	// S35
	func(r rune) int {
		switch {
		case r == 42: // ['*','*']
			return 35
		case r == 47: // ['/','/']
			return 40
		default:
			return 26
		}
	},
	// S36
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S37
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 19
		case 97 <= r && r <= 100: // ['a','d']
			return 16
		case r == 101: // ['e','e']
			return 41
		case 102 <= r && r <= 122: // ['f','z']
			return 16
		}
		return NoState
	},
	// S38
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 19
		case 97 <= r && r <= 115: // ['a','s']
			return 16
		case r == 116: // ['t','t']
			return 42
		case 117 <= r && r <= 122: // ['u','z']
			return 16
		}
		return NoState
	},
	// S39
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 19
		case 97 <= r && r <= 107: // ['a','k']
			return 16
		case r == 108: // ['l','l']
			return 43
		case 109 <= r && r <= 122: // ['m','z']
			return 16
		}
		return NoState
	},
	// S40
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S41
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 16
		}
		return NoState
	},
	// S42
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 16
		}
		return NoState
	},
	// S43
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 19
		case 97 <= r && r <= 100: // ['a','d']
			return 16
		case r == 101: // ['e','e']
			return 44
		case 102 <= r && r <= 122: // ['f','z']
			return 16
		}
		return NoState
	},
	// S44
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 30
		case 65 <= r && r <= 90: // ['A','Z']
			return 16
		case r == 95: // ['_','_']
			return 19
		case 97 <= r && r <= 122: // ['a','z']
			return 16
		}
		return NoState
	},
}
