# ESEP Grade Calculator

This library is a simple utility to calculate your grades.

Runtime requirements:
Golang 1.24+

Check installation:
go version

Clone the repo and navigate to it:
git clone https://github.com/<your-username>/esep-grade-calculator.git
cd esep-grade-calculator/grade_calculator

Run unit tests:
```

go test ./...
```

Run unit tests with code coverage:
```

go test -coverprofile=coverage.txt ./...
```

The Grade Calculator computes a student's final letter grade from three categories, assignments (50%), exams (35%), and essays (15%). The weighted average is then converted to a final letter grade (A-F).

