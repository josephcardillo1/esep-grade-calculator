package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 50, Assignment)
	gradeCalculator.AddGrade("exam 1", 40, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 37, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeBoundaries(t *testing.T) {
	gc := NewGradeCalculator()
	// Exactly 90 → A
	gc.AddGrade("a", 90, Assignment)
	gc.AddGrade("e1", 90, Exam)
	gc.AddGrade("s1", 90, Essay)
	if got := gc.GetFinalGrade(); got != "A" {
		t.Fatalf("want A, got %s", got)
	}

	gc = NewGradeCalculator()
	gc.AddGrade("a", 80, Assignment)
	gc.AddGrade("e1", 80, Exam)
	gc.AddGrade("s1", 80, Essay)
	if got := gc.GetFinalGrade(); got != "B" {
		t.Fatalf("want B, got %s", got)
	}

	gc = NewGradeCalculator()
	gc.AddGrade("a", 70, Assignment)
	gc.AddGrade("e1", 70, Exam)
	gc.AddGrade("s1", 70, Essay)
	if got := gc.GetFinalGrade(); got != "C" {
		t.Fatalf("want C, got %s", got)
	}

	gc = NewGradeCalculator()
	gc.AddGrade("a", 60, Assignment)
	gc.AddGrade("e1", 60, Exam)
	gc.AddGrade("s1", 60, Essay)
	if got := gc.GetFinalGrade(); got != "D" {
		t.Fatalf("want D, got %s", got)
	}

	gc = NewGradeCalculator()
	gc.AddGrade("a", 0, Assignment)
	gc.AddGrade("e1", 0, Exam)
	gc.AddGrade("s1", 0, Essay)
	if got := gc.GetFinalGrade(); got != "F" {
		t.Fatalf("want F, got %s", got)
	}
}

func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Fatalf("Assignment.String() wrong")
	}
	if Exam.String() != "exam" {
		t.Fatalf("Exam.String() wrong")
	}
	if Essay.String() != "essay" {
		t.Fatalf("Essay.String() wrong")
	}
}