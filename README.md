MATH-SKILLS

DESCRIPTION
This program calculates statistical measures from a file containing numerical data.
It reads numbers from a text file and computes Average, Median, Variance, and Standard Deviation.

USAGE
go run main.go <filename>

Example:
go run main.go data.txt
go run main.go mydata.txt

HOW TO CREATE YOUR OWN DATA FILE
Create a text file with one number per line:
1. Open any text editor
2. Write one number per line
3. Save as .txt file
4. Run the program with your filename

INPUT FORMAT
The input file should contain one number per line:
189
113
121
114
145
110

OUTPUT FORMAT
The program displays results as rounded integers:
Average: 132
Median: 118
Variance: 785
Standard Deviation: 28

FUNCTIONS
ReadData - Reads numbers from file and handles errors
Average - Calculates arithmetic mean of all numbers
Median - Finds middle value using bubble sort algorithm
Variance - Measures spread of data from the mean
StandardDeviation - Square root of variance

ERROR HANDLING
- Checks for missing filename argument
- Validates file exists and can be opened
- Handles invalid number formats in file
- Detects empty files
- Ignores empty lines in input

FEATURES
- Supports negative numbers
- Handles decimal numbers
- Proper error messages for invalid input
- Robust file reading with scanner error checking
- Memory efficient data processing

AUTHOR
Giorgos Salaounis
Zone01 Athens - math-skills project