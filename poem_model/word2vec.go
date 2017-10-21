package poem_model

import (
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"os"
//	"errors"
)

const maxSize int = 2000
const N int = 10

type W2VModel struct {
	Words   int
	Size    int
	Vocab   []string
	WordIdx map[string]int
	Vec     [][]float32
}

type WordData struct {
	Distance float64
	Word     string
}

func (m *W2VModel) Load(fn string) {
	file, err := os.Open(fn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fscanf(file, "%d", &m.Words)
	fmt.Fscanf(file, "%d", &m.Size)

   // m.Words = 1000
	var ch string
	m.Vocab = make([]string, m.Words)
	m.Vec = make([][]float32, m.Words)
	m.WordIdx = make(map[string]int)
	for b := 0; b < m.Words; b++ {
		m.Vec[b] = make([]float32, m.Size)
		fmt.Fscanf(file, "%s%c", &m.Vocab[b], &ch)
		m.WordIdx[m.Vocab[b]] = b
		binary.Read(file, binary.LittleEndian, m.Vec[b])

		length := 0.0
		for _, v := range m.Vec[b] {
			length += float64(v * v)
		}
		length = math.Sqrt(length)

		for i, _ := range m.Vec[b] {
			m.Vec[b][i] /= float32(length)
		}
	}
	file.Close()
}

//func (m *W2VModel) Vector(idx int) ([]float32, error) {
//	if idx < 0 || idx >= m.Words {
//		return nil, errors.New("index out of range")
//	}
//	vec := make([]float32, m.Size)
//	copy(vec, m.Vec[idx * m.Size : idx * m.Size + m.Size])
//	return vec, nil
//}


func (m *W2VModel) WordVector(word string) ([]float32, error) {
	idx, ok := m.WordIdx[word]
	if ! ok {
		return nil, fmt.Errorf("[%s] isn't in vocabulary", word)
	}
	return m.Vec[idx], nil
}

func (m *W2VModel) MostSimilar(seedWords []string) []WordData {
	if len(seedWords) == 0 {
		return nil
	}

	inputPosition := make([]int, 10)
	for k, v := range seedWords {
		var b int
		for b = 0; b < m.Words; b++ {
			if m.Vocab[b] == v {
				break
			}
		}
		inputPosition[k] = b
		fmt.Printf("Word %v Position %v \n", v, b)
	}
	vec := make([]float32, maxSize)
	for i, _ := range seedWords {
		for j := 0; j < m.Size; j++ {
			vec[j] += m.Vec[inputPosition[i]][j]
		}
	}

	length := 0.0
	for _, v := range vec {
		length += float64(v * v)
	}
	length = math.Sqrt(length)

	for k, _ := range vec {
		vec[k] /= float32(length)
	}

	bestWords := make([]WordData, N)

	for i := 0; i < m.Words; i++ {
		c := 0
		for _, v := range inputPosition {
			if v == i {
				c = 1
			}
		}
		if c == 1 {
			continue
		}
		dist := 0.0
		for j := 0; j < m.Size; j++ {
			dist += float64(vec[j] * m.Vec[i][j])
		}

		for j := 0; j < N; j++ {
			if dist > bestWords[j].Distance {
				for d := N - 1; d > j; d-- {
					bestWords[d] = bestWords[d-1]
				}
				bestWords[j] = WordData{dist, m.Vocab[i]}
				break
			}
		}
	}
	return bestWords
}
