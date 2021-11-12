package main

import "fmt"

func main() {
	//[["A","B","C","E"],["S","F","E","S"],["A","D","E","E"]]
	board := [][]byte{
		{'A','B','C','E'},
		{'S','F','E','S'},
		{'A','D','E','E'},
	}
	word := "ABCESEEEFS"
	result := exist(board, word)
	fmt.Println("result =",result)
}

var (
	Board [][]byte
	Word  string
	H int
	W int
	Result bool
)

func exist(board [][]byte, word string) bool {
	//初始化全局变量（避免leetcode报错）
	Board = [][]byte{}
	Word = ""
	H = 0
	W = 0
	Result = false

	Board = board
	Word = word
	H = len(board)		//高
	W = len(board[0])	//宽

	fmt.Printf("board = %#v\n",board)
	fmt.Printf("word = %#v\n",word)

	fmt.Printf("Board = %#v\n",Board)
	fmt.Printf("Word = %#v\n",Word)
	fmt.Printf("H = %#v\n",H)
	fmt.Printf("W = %#v\n",W)

	fmt.Println("Board[0][0] == Word[0]",Board[0][0] == Word[0])

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			//fmt.Println("Board[i][j] == Word[0]",Board[i][j] == Word[0])
			if Board[i][j] == Word[0] {
				dfs(i,j,0)
				if Result {
					return true
				}
			}
		}
	}
	return false
}

func dfs(i,j,n int)  {
	//如果越界或者word中的指针越界或者board[i][j]被访问过
	if i < 0 ||
		j < 0 ||
		i == H ||
		j == W ||
		n == len(Word) ||
		Board[i][j] == '0' {
		return
	}
	//如果i,j位置的字符与word中下标指向的元素匹配
	if Board[i][j] == Word[n] {
		//fmt.Printf("[%v,%v]\n",i,j)
		//并且word中的下标位置已经到达最后一个，则代表该单词存在于网格中
		if n == len(Word)-1 {
			Result = true
			return
		}
		temp := Board[i][i]
		//fmt.Println("temp =",temp)
		Board[i][j] = '0'	//访问过就置为'0'
		dfs(i+1,j,n+1)
		dfs(i,j+1,n+1)
		dfs(i-1,j,n+1)
		dfs(i,j-1,n+1)
		//fmt.Println("temp =",temp)
		Board[i][j] = temp	//复原访问
	}
}
