package roman

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func RomanToInt(w http.ResponseWriter, r *http.Request) {
	romanLowerCase := r.URL.Query().Get("query")
	roman := strings.ToUpper(romanLowerCase)
	result := 0
	var romanArr []string
	romanToInt := map[string]int{"M": 1000, "D": 500, "C": 100, "L": 50, "X": 10, "V": 5, "I": 1}
	for _, letter := range roman {
		romanArr = append(romanArr, string(letter))
	}
	for i := 0; i < len(romanArr)-1; i++ {
		curr := romanToInt[romanArr[i]]
		next := romanToInt[romanArr[i+1]]
		if curr < next {
			result -= curr
		} else {
			result += curr
		}
	}
	result += romanToInt[romanArr[len(romanArr)-1]]

	jsonData, err := json.Marshal(result)
	if err != nil {
		log.Println(err.Error())
		status := http.StatusInternalServerError
		http.Error(w, http.StatusText(status), status)
		return
	}
	w.Write(jsonData)

}
