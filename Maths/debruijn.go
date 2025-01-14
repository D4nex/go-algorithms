package main

import(
    "fmt"
    "math")
 
// Global variables   
var seen map[string]bool //Tracks combinations already visited to avoid duplicates.
var edges []int //Stores the order of the characters used to reconstruct the final sequence.


func Dfs(node string, k int, A []rune){
    
    //We implemented a modified depth-first search (DFS) to generate the sequence nodes.
    //Each node represents a prefix of length , and we generate new combinations by adding characters from the alphabet.
    //It is ensured that no combination is repeated by the "seen" set.
    //The indices of the characters used to build the sequence are stored in "edges".

    
    for i := 0; i < k; i++ {
        
        str := node + string(A[i])
        if !seen[str] {
            seen[str] = true
            Dfs(str[1:], k, A)
            edges = append(edges, i)
        }
    }
}

func DeBruijn(n, k int, A []rune) string {
	// This is the main function for constructing the sequence.
	//It initializes the search with a "startingNode", which is a repetition of the first character of the alphabet times.
	//It calls dfs to generate all possible nodes.
	//Uses the indexes stored in edges to build the final sequence.
	
	seen = make(map[string]bool)
	edges = []int{}

	startingNode := ""
	for i := 0; i < n-1; i++ {
		startingNode += string(A[0])
	}

	Dfs(startingNode, k, A)
	
	S := ""
	l := int(math.Pow(float64(k), float64(n)))
	for i := 0; i < l; i++ {
		S += string(A[edges[i]])
	}
	S += startingNode

	return S
}

func main() {
	n := 6 // length
	k := 3 // alphabet size
	A := []rune{'a', 's', 'd'} // alphabet

	fmt.Println(DeBruijn(n, k, A))
}