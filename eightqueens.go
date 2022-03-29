package piscine

import (
	"github.com/01-edu/z01"
)

func Checkqueendirect(queenspos [8]int, queenaxe_x int, queenaxe_y int) bool {
	noqueen := true
	seccheck := true
	for x, y := range queenspos {
		if y != 0 && queenaxe_y == y || y != 0 && queenaxe_x == x+1 {
			noqueen = false
		}
		for indexdiago := 1; indexdiago < queenaxe_x-x; indexdiago++ {
			if queenspos[queenaxe_x-1-indexdiago]+indexdiago == queenaxe_y || queenspos[queenaxe_x-1-indexdiago]-indexdiago == queenaxe_y {
				seccheck = false
			}
		}
	}
	return noqueen && seccheck
}

func EightQueens() {
	for q1 := 1; q1 <= 8; q1++ {
		if Checkqueendirect([8]int{0, 0, 0, 0, 0, 0, 0, 0}, 1, q1) {
			for q2 := 1; q2 <= 8; q2++ {
				if Checkqueendirect([8]int{q1, 0, 0, 0, 0, 0, 0, 0}, 2, q2) {
					for q3 := 1; q3 <= 8; q3++ {
						if Checkqueendirect([8]int{q1, q2, 0, 0, 0, 0, 0, 0}, 3, q3) {
							for q4 := 1; q4 <= 8; q4++ {
								if Checkqueendirect([8]int{q1, q2, q3, 0, 0, 0, 0, 0}, 4, q4) {
									for q5 := 1; q5 <= 8; q5++ {
										if Checkqueendirect([8]int{q1, q2, q3, q4, 0, 0, 0, 0}, 5, q5) {
											for q6 := 1; q6 <= 8; q6++ {
												if Checkqueendirect([8]int{q1, q2, q3, q4, q5, 0, 0, 0}, 6, q6) {
													for q7 := 1; q7 <= 8; q7++ {
														if Checkqueendirect([8]int{q1, q2, q3, q4, q5, q6, 0, 0}, 7, q7) {
															for q8 := 1; q8 <= 8; q8++ {
																if Checkqueendirect([8]int{q1, q2, q3, q4, q5, q6, q7, 0}, 8, q8) {
																	z01.PrintRune(rune(q1) + 48)
																	z01.PrintRune(rune(q2) + 48)
																	z01.PrintRune(rune(q3) + 48)
																	z01.PrintRune(rune(q4) + 48)
																	z01.PrintRune(rune(q5) + 48)
																	z01.PrintRune(rune(q6) + 48)
																	z01.PrintRune(rune(q7) + 48)
																	z01.PrintRune(rune(q8) + 48)
																	z01.PrintRune('\n')
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
