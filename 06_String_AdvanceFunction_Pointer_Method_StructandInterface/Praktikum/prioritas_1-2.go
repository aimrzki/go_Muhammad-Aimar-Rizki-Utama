package main

import (
	"fmt"
	"math"
)

type Student struct {
	name  string
	score float64
}

type StudentList []Student

func (s StudentList) AverageScore() float64 {
	totalScore := 0.0
	for _, student := range s {
		totalScore += student.score
	}
	return totalScore / float64(len(s))
}

func (s StudentList) MinScoreStudent() Student {
	minScore := math.MaxFloat64
	var minStudent Student

	for _, student := range s {
		if student.score < minScore {
			minScore = student.score
			minStudent = student
		}
	}
	return minStudent
}

func (s StudentList) MaxScoreStudent() Student {
	maxScore := -1.0
	var maxStudent Student

	for _, student := range s {
		if student.score > maxScore {
			maxScore = student.score
			maxStudent = student
		}
	}
	return maxStudent
}

func main() {
	var students StudentList

	for i := 1; i <= 5; i++ {
		var name string
		var score float64

		fmt.Printf("input %d student name: ", i)
		fmt.Scanln(&name)
		fmt.Printf("input %d student score: ", i)
		fmt.Scanln(&score)

		student := Student{name, score}
		students = append(students, student)
	}

	averageScore := students.AverageScore()
	minScoreStudent := students.MinScoreStudent()
	maxScoreStudent := students.MaxScoreStudent()

	fmt.Printf("average score: %.2f\n", averageScore)
	fmt.Printf("min score of student: %s (%.0f)\n", minScoreStudent.name, minScoreStudent.score)
	fmt.Printf("max score of student: %s (%.0f)\n", maxScoreStudent.name, maxScoreStudent.score)
}
