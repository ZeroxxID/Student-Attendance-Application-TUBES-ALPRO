package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// ==============================================================================
// CONSTANTS & DATA STRUCTURES
// ==============================================================================

const (
	maxStudent    = 50
	maxAttendance = 18250
)

type Student struct {
	NIS  string
	Name string
}

type Dates struct {
	Day   int
	Month int
	Year  int
}

type Attendance struct {
	Date     Dates
	NIS      string
	Name     string
	Presence bool
}

type Statistics struct {
	NIS           string
	Name          string
	TotalPresence int
	TotalMeetings int
	Percentage    float64
}

type TabStudent [maxStudent]Student
type TabAttendance [maxAttendance]Attendance
type TabStatistics [maxStudent]Statistics

var studentData TabStudent
var attendanceData TabAttendance

// ==============================================================================
// MAIN PROGRAM
// ==============================================================================

/*
Subprogram: main
Description: The main entry point of the program. Initializes counters and calls MainMenu.
Parameters: None.
Initial State (I.S.): Program runs, nStudent and nAttendance counters are initialized to 0.
Final State (F.S.): MainMenu is called using pass-by-reference. The program halts normally when MainMenu returns.
*/
func main() {
	var nStudent int = 0
	var nAttendance int = 0

	SeedDummyData(&nStudent, &nAttendance)

	scanner := bufio.NewScanner(os.Stdin)
	MainMenu(scanner, &nStudent, &nAttendance)
}

// ==============================================================================
// DUMMY DATA SEEDER
// ==============================================================================

/*
Subprogram: SeedDummyData
Description: Fill the initial state of the application with dummy data directly.

	Implement relational mapping to ensure data consistency.
*/
func SeedDummyData(nS *int, nA *int) {
	dummyStudents := []Student{
		{NIS: "103032540004", Name: "Elon Musk"},
		{NIS: "103032540005", Name: "Barack Obama"},
		{NIS: "103032500177", Name: "Vladimir Putin"},
		{NIS: "103032500180", Name: "Xi Jinping"},
		{NIS: "103032540001", Name: "Joe Biden"},
		{NIS: "103032500153", Name: "Cristiano Ronaldo"},
		{NIS: "103032500149", Name: "Lionel Messi"},
		{NIS: "103032540002", Name: "LeBron James"},
		{NIS: "103032500146", Name: "Michael Jordan"},
		{NIS: "103032500154", Name: "Bill Gate"},
	}

	for i, s := range dummyStudents {
		studentData[i] = s
	}
	*nS = len(dummyStudents)

	SortStudentAsc(1, nS)

	type seedAtt struct {
		d, m, y int
		nis     string
		pres    bool
	}

	dummyAttendances := []seedAtt{
		// Tanggal 01 01 2026
		{1, 1, 2026, "103032540004", true}, {1, 1, 2026, "103032540005", true}, {1, 1, 2026, "103032500177", true}, {1, 1, 2026, "103032500180", true}, {1, 1, 2026, "103032540001", false},
		{1, 1, 2026, "103032500153", true}, {1, 1, 2026, "103032500149", false}, {1, 1, 2026, "103032540002", true}, {1, 1, 2026, "103032500146", true}, {1, 1, 2026, "103032500154", false},

		// Tanggal 02 01 2026
		{2, 1, 2026, "103032540004", false}, {2, 1, 2026, "103032540005", true}, {2, 1, 2026, "103032500177", true}, {2, 1, 2026, "103032500180", true}, {2, 1, 2026, "103032540001", false},
		{2, 1, 2026, "103032500153", true}, {2, 1, 2026, "103032500149", true}, {2, 1, 2026, "103032540002", true}, {2, 1, 2026, "103032500146", true}, {2, 1, 2026, "103032500154", true},

		// Tanggal 03 01 2026
		{3, 1, 2026, "103032540004", true}, {3, 1, 2026, "103032540005", true}, {3, 1, 2026, "103032500177", true}, {3, 1, 2026, "103032500180", true}, {3, 1, 2026, "103032540001", false},
		{3, 1, 2026, "103032500153", true}, {3, 1, 2026, "103032500149", true}, {3, 1, 2026, "103032540002", true}, {3, 1, 2026, "103032500146", true}, {3, 1, 2026, "103032500154", false},
	}

	for i, att := range dummyAttendances {
		var studentName string
		for j := 0; j < *nS; j++ {
			if studentData[j].NIS == att.nis {
				studentName = studentData[j].Name
				break
			}
		}

		attendanceData[i] = Attendance{
			Date:     Dates{Day: att.d, Month: att.m, Year: att.y},
			NIS:      att.nis,
			Name:     studentName,
			Presence: att.pres,
		}
	}
	*nA = len(dummyAttendances)
}

/*
Subprogram: MainMenu
Description: Displays the root navigation interface.
Parameters: scanner (*bufio.Scanner), nS (*int), nA (*int)
Initial State (I.S.): The application starts from main().
Final State (F.S.): Executes menu navigation, terminates if user inputs 0 or EOF, returning control to main to exit.
*/
func MainMenu(scanner *bufio.Scanner, nS *int, nA *int) {
	for {
		fmt.Println("\n==============================================")
		fmt.Println("||          STUDENT ATTENDANCE APP          ||")
		fmt.Println("==============================================")
		fmt.Println("[1] Student Menu")
		fmt.Println("[2] Attendance Menu")
		fmt.Println("[3] Statistics Menu")
		fmt.Println("[0] Exit")
		fmt.Println("==============================================")
		fmt.Print("Select option: ")

		if !scanner.Scan() {
			return
		}
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			if StudentMenu(scanner, nS, nA) {
				return
			}
		case "2":
			if AttendanceMenu(scanner, nS, nA) {
				return
			}
		case "3":
			if StatisticsMenu(scanner, nS, nA) {
				return
			}
		case "0":
			fmt.Println("\n[!] Message: The program is finished. See you later!")
			return
		default:
			fmt.Println("\n[!] Error: Invalid input! Please choose available options.")
		}
	}
}

// ==============================================================================
// STUDENT MENU & CRUD
// ==============================================================================

/*
Subprogram: StudentMenu
Description: Navigation for CRUD operations on the Student entity.
Parameters: scanner (*bufio.Scanner), nS (*int), nA (*int)
Initial State (I.S.): Called from MainMenu.
Final State (F.S.): Executes student operations, returns false if 'Back', true if 'Exit'.
*/
func StudentMenu(scanner *bufio.Scanner, nS *int, nA *int) bool {
	inMenu := true
	for inMenu {
		fmt.Println("\n==============================================")
		fmt.Println("||               STUDENT MENU               ||")
		fmt.Println("==============================================")
		fmt.Println("[1] Back to Menu")
		fmt.Println("[2] Add Student")
		fmt.Println("[3] Edit Student")
		fmt.Println("[4] Delete Student")
		fmt.Println("[5] Search Student")
		fmt.Println("[6] Sort Student")
		fmt.Println("[7] Show Student Data")
		fmt.Println("[0] Exit")
		fmt.Println("==============================================")
		fmt.Print("Select option: ")

		if !scanner.Scan() {
			return true
		}
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			return false
		case "2":
			AddStudent(scanner, nS)
		case "3":
			EditStudent(scanner, nS, nA)
		case "4":
			DeleteStudent(scanner, nS, nA)
		case "5":
			SearchStudentMenu(scanner, nS)
		case "6":
			if SortStudentMenu(scanner, nS) {
				return true
			}
		case "7":
			ShowStudentData(nS)
		case "0":
			fmt.Println("\n[!] Message: The program is finished. See you later!")
			return true
		default:
			fmt.Println("\n[!] Error: Invalid input!")
		}
	}
	return false
}

/*
Subprogram: isNumeric
Description: Pure numeric validation.
Parameters: s (string)
Initial State (I.S.): String s is received.
Final State (F.S.): Returns true if s only contains digits 0-9.
*/
func isNumeric(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

/*
Subprogram: nisLess
Description: Compares two NIS numerically to prevent misordering (e.g., "9" vs "10").
Parameters: a (string), b (string)
Initial State (I.S.): a and b are numeric-validated NIS strings.
Final State (F.S.): Returns true if the numeric value of a < b; fallback to string comparison if parsing fails.
*/
func nisLess(a, b string) bool {
	ia, ea := strconv.Atoi(a)
	ib, eb := strconv.Atoi(b)
	if ea == nil && eb == nil {
		return ia < ib
	}
	return a < b
}

/*
Subprogram: AddStudent
Description: Adds a new student to the array.
Parameters: scanner (*bufio.Scanner), nS (*int)
Initial State (I.S.): Array capacity is available.
Final State (F.S.): Data is appended, *nS increments, array sorted (Ascending by NIS).
*/
func AddStudent(scanner *bufio.Scanner, nS *int) {
	if *nS >= maxStudent {
		fmt.Println("\n[!] Error: Student data capacity is full!")
		return
	}

	var inputNIS string
	isValid := false

	for !isValid {
		fmt.Print("Input NIS (Numeric): ")
		if !scanner.Scan() {
			return
		}
		inputNIS = strings.TrimSpace(scanner.Text())

		if isNumeric(inputNIS) {
			SortStudentAsc(1, nS)
			if BinarySearchStudent(inputNIS, 0, *nS-1) != -1 {
				fmt.Println("\n[!] Error: NIS is already registered!")
			} else {
				isValid = true
			}
		} else {
			fmt.Println("\n[!] Error: Invalid NIS format! Must be numbers.")
		}
	}

	fmt.Print("Input Name: ")
	if !scanner.Scan() {
		return
	}
	inputName := strings.TrimSpace(scanner.Text())

	if inputName == "" {
		fmt.Println("\n[!] Error: Name cannot be empty!")
		return
	}

	studentData[*nS].NIS = inputNIS
	studentData[*nS].Name = inputName
	*nS++

	SortStudentAsc(1, nS)
	fmt.Println("\n[+] Message: Successfully added student data!")
}

/*
Subprogram: EditStudent
Description: Updates student attributes, cascades changes of name and NIS to attendance.
Parameters: scanner (*bufio.Scanner), nS (*int), nA (*int)
Initial State (I.S.): *nS > 0.
Final State (F.S.): Data is updated, changes propagate to attendanceData, array remains sorted.
*/
func EditStudent(scanner *bufio.Scanner, nS *int, nA *int) {
	if *nS == 0 {
		fmt.Println("\n[!] Message: No student data available!")
		return
	}

	ShowStudentData(nS)
	fmt.Printf("Input student number to edit (1 - %d): ", *nS)
	if !scanner.Scan() {
		return
	}

	var idx int
	_, err := fmt.Sscanf(strings.TrimSpace(scanner.Text()), "%d", &idx)
	if err != nil || idx < 1 || idx > *nS {
		fmt.Println("\n[!] Error: Invalid student number!")
		return
	}

	realIdx := idx - 1
	targetNIS := studentData[realIdx].NIS
	oldName := studentData[realIdx].Name

	SortStudentAsc(1, nS)
	searchIdx := BinarySearchStudent(targetNIS, 0, *nS-1)
	if searchIdx == -1 {
		fmt.Println("\n[!] Error: Data consistency failure!")
		return
	}

	fmt.Printf("\n[?] Found Data: %s - %s\n", targetNIS, oldName)
	fmt.Print("[?] Continue editing? (y/n): ")
	if !scanner.Scan() {
		return
	}

	if strings.ToLower(strings.TrimSpace(scanner.Text())) != "y" {
		fmt.Println("\n[-] Message: Edit operation cancelled.")
		return
	}

	fmt.Println("\n[Info] Enter '-' (dash) to keep current value.")

	fmt.Print("Input new NIS: ")
	if !scanner.Scan() {
		return
	}
	newNIS := strings.TrimSpace(scanner.Text())

	if newNIS != "-" {
		if !isNumeric(newNIS) {
			fmt.Println("\n[!] Error: Invalid NIS format! Must be numbers.")
			return
		}
		if BinarySearchStudent(newNIS, 0, *nS-1) != -1 {
			fmt.Println("\n[!] Error: NIS is already used by another student!")
			return
		}
	} else {
		newNIS = targetNIS
	}

	fmt.Print("Input new Name: ")
	if !scanner.Scan() {
		return
	}
	newName := strings.TrimSpace(scanner.Text())

	if newName != "-" {
		if newName == "" {
			fmt.Println("\n[!] Error: Name cannot be empty!")
			return
		}
	} else {
		newName = oldName
	}

	studentData[searchIdx].NIS = newNIS
	studentData[searchIdx].Name = newName

	for i := 0; i < *nA; i++ {
		if attendanceData[i].NIS == targetNIS {
			attendanceData[i].NIS = newNIS
			attendanceData[i].Name = newName
		}
	}

	if newNIS != targetNIS {
		SortStudentAsc(1, nS)
	}

	fmt.Println("\n[+] Message: Successfully updated student data!")
}

/*
Subprogram: DeleteStudent
Description: Deletes a student by index, implementing CASCADE DELETE on attendance data.
Parameters: scanner (*bufio.Scanner), nS (*int), nA (*int)
Initial State (I.S.): Student array has at least 1 data.
Final State (F.S.): Student deleted, attendance with related NIS cleared, array shifted left.
*/
func DeleteStudent(scanner *bufio.Scanner, nS *int, nA *int) {
	if *nS == 0 {
		fmt.Println("\n[!] Message: No student data available!")
		return
	}

	ShowStudentData(nS)
	fmt.Printf("Input student number to delete (1 - %d): ", *nS)
	if !scanner.Scan() {
		return
	}

	var idx int
	_, err := fmt.Sscanf(strings.TrimSpace(scanner.Text()), "%d", &idx)
	if err != nil || idx < 1 || idx > *nS {
		fmt.Println("\n[!] Error: Invalid student number!")
		return
	}

	realIdx := idx - 1
	targetNIS := studentData[realIdx].NIS
	targetName := studentData[realIdx].Name

	SortStudentAsc(1, nS)
	searchIdx := BinarySearchStudent(targetNIS, 0, *nS-1)
	if searchIdx == -1 {
		fmt.Println("\n[!] Error: Data consistency failure!")
		return
	}

	fmt.Printf("\n[?] Are you sure you want to delete No.%d | %s - %s? (y/n): ", idx, targetNIS, targetName)
	if !scanner.Scan() {
		return
	}

	if strings.ToLower(strings.TrimSpace(scanner.Text())) != "y" {
		fmt.Println("\n[-] Message: Delete operation cancelled.")
		return
	}

	for i := searchIdx; i < *nS-1; i++ {
		studentData[i] = studentData[i+1]
	}
	studentData[*nS-1] = Student{}
	*nS--

	var newAttCount int = 0
	for i := 0; i < *nA; i++ {
		if attendanceData[i].NIS != targetNIS {
			attendanceData[newAttCount] = attendanceData[i]
			newAttCount++
		}
	}

	for i := newAttCount; i < *nA; i++ {
		attendanceData[i] = Attendance{}
	}
	*nA = newAttCount

	fmt.Println("\n[+] Message: Successfully deleted student data and cascaded attendance!")
}

/*
Subprogram: SearchStudentMenu
Description: Searches for a student utilizing Binary Search (if NIS) and Sequential Search (if Name).
Parameters: scanner (*bufio.Scanner), nS (*int)
Initial State (I.S.): Array is defined.
Final State (F.S.): Related data is printed if exists, taking advantage of Binary Search efficiency when possible.
*/
func SearchStudentMenu(scanner *bufio.Scanner, nS *int) {
	if *nS == 0 {
		fmt.Println("\n[!] Message: No student data available!")
		return
	}

	fmt.Print("Input Keyword (NIS or Name): ")
	if !scanner.Scan() {
		return
	}
	keyword := strings.ToLower(strings.TrimSpace(scanner.Text()))

	if keyword == "" {
		fmt.Println("\n[!] Error: Keyword cannot be empty!")
		return
	}

	if isNumeric(keyword) {
		SortStudentAsc(1, nS)
		idx := BinarySearchStudent(keyword, 0, *nS-1)

		if idx != -1 {
			fmt.Println("\n==============================================")
			fmt.Printf("| %-3s | %-13s | %-20s |\n", "NO", "NIS", "NAME")
			fmt.Println("==============================================")
			fmt.Printf("| %-3d | %-13s | %-20s |\n", 1, studentData[idx].NIS, studentData[idx].Name)
			fmt.Println("==============================================")
			return
		}
	}

	count := 0
	found := false
	for i := 0; i < *nS; i++ {
		currName := strings.ToLower(studentData[i].Name)
		currNIS := strings.ToLower(studentData[i].NIS)

		if strings.Contains(currName, keyword) || strings.Contains(currNIS, keyword) {
			if !found {
				fmt.Println("\n==============================================")
				fmt.Printf("| %-3s | %-13s | %-20s |\n", "NO", "NIS", "NAME")
				fmt.Println("==============================================")
				found = true
			}
			count++
			fmt.Printf("| %-3d | %-13s | %-20s |\n", count, studentData[i].NIS, studentData[i].Name)
		}
	}

	if found {
		fmt.Println("==============================================")
	} else {
		fmt.Println("\n[!] Error: No student matches the keyword!")
	}
}

/*
Subprogram: SortStudentMenu
Description: Menu options for Sorting Students (Asc/Desc by NIS/Name).
Parameters: scanner (*bufio.Scanner), nS (*int)
Initial State (I.S.): Selected from StudentMenu.
Final State (F.S.): Calls sorting algorithms (Selection/Insertion) and prints them. Returns bool for graceful exit.
*/
func SortStudentMenu(scanner *bufio.Scanner, nS *int) bool {
	inMenu := true
	for inMenu {
		fmt.Println("\n==============================================")
		fmt.Println("||            SORT STUDENT MENU             ||")
		fmt.Println("==============================================")
		fmt.Println("[1] Ascending by NIS (Selection Sort)")
		fmt.Println("[2] Descending by NIS (Insertion Sort)")
		fmt.Println("[3] Ascending by Name (Selection Sort)")
		fmt.Println("[4] Descending by Name (Insertion Sort)")
		fmt.Println("[5] Back to Student Menu")
		fmt.Println("[0] Exit")
		fmt.Println("==============================================")
		fmt.Print("Select option: ")

		if !scanner.Scan() {
			return true
		}
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			SortStudentAsc(1, nS)
			fmt.Println("\n[+] Message: Sorted Ascending by NIS!")
			ShowStudentData(nS)
		case "2":
			SortStudentDesc(1, nS)
			fmt.Println("\n[+] Message: Sorted Descending by NIS!")
			ShowStudentData(nS)
		case "3":
			SortStudentAsc(2, nS)
			fmt.Println("\n[+] Message: Sorted Ascending by Name!")
			ShowStudentData(nS)
		case "4":
			SortStudentDesc(2, nS)
			fmt.Println("\n[+] Message: Sorted Descending by Name!")
			ShowStudentData(nS)
		case "5":
			inMenu = false
		case "0":
			fmt.Println("\n[!] Message: The program is finished. See you later!")
			return true
		default:
			fmt.Println("\n[!] Error: Invalid input!")
		}
	}
	return false
}

/*
Subprogram: ShowStudentData
Description: Prints the data in the studentData array to the terminal.
Parameters: nS (*int)
Initial State (I.S.): Data is available.
Final State (F.S.): Data is printed nicely formatted.
*/
func ShowStudentData(nS *int) {
	if *nS == 0 {
		fmt.Println("\n[!] Message: No student data available!")
		return
	}
	fmt.Println("\n==============================================")
	fmt.Printf("| %-3s | %-13s | %-20s |\n", "NO", "NIS", "NAME")
	fmt.Println("==============================================")
	for i := 0; i < *nS; i++ {
		fmt.Printf("| %-3d | %-13s | %-20s |\n", i+1, studentData[i].NIS, studentData[i].Name)
	}
	fmt.Println("==============================================")
}

/*
Subprogram: BinarySearchStudent
Description: Recursive Binary Search (Adding value according to spec k).
Parameters: nis (string), left (int), right (int)
Initial State (I.S.): studentData array sorted ascending by NIS.
Final State (F.S.): Returns index if found, -1 if none.
*/
func BinarySearchStudent(nis string, left int, right int) int {
	if left > right {
		return -1
	}
	mid := (left + right) / 2
	if studentData[mid].NIS == nis {
		return mid
	} else if nisLess(nis, studentData[mid].NIS) {
		return BinarySearchStudent(nis, left, mid-1)
	} else {
		return BinarySearchStudent(nis, mid+1, right)
	}
}

/*
Subprogram: SequentialSearchAttendanceByIdx
Description: Sequential Search to find an attendance entry based on its sequence number (1-based).
Parameters: targetIdx (int), nA (*int)
Initial State (I.S.): attendanceData array is available, targetIdx >= 1.
Final State (F.S.): Returns 0-based index if found, -1 if none.
*/
func SequentialSearchAttendanceByIdx(targetIdx int, nA *int) int {
	found := -1
	i := 0
	for i < *nA && found == -1 {
		if i+1 == targetIdx {
			found = i
		}
		i++
	}
	return found
}

/*
Subprogram: SortStudentAsc
Description: Sorts studentData elements from smallest to largest using Selection Sort.
Parameters: mode (int), nS (*int)
Initial State (I.S.): Data is unordered.
Final State (F.S.): Ordered ascending (case-insensitive for string).
*/
func SortStudentAsc(mode int, nS *int) {
	for i := 0; i < *nS-1; i++ {
		minIdx := i
		for j := i + 1; j < *nS; j++ {
			if mode == 1 {
				if nisLess(studentData[j].NIS, studentData[minIdx].NIS) {
					minIdx = j
				}
			} else {
				if strings.ToLower(studentData[j].Name) < strings.ToLower(studentData[minIdx].Name) {
					minIdx = j
				}
			}
		}
		studentData[i], studentData[minIdx] = studentData[minIdx], studentData[i]
	}
}

/*
Subprogram: SortStudentDesc
Description: Sorts studentData elements from largest to smallest using Insertion Sort.
Parameters: mode (int), nS (*int)
Initial State (I.S.): Data is random.
Final State (F.S.): Ordered descending (case-insensitive).
*/
func SortStudentDesc(mode int, nS *int) {
	for i := 1; i < *nS; i++ {
		key := studentData[i]
		j := i - 1

		if mode == 1 {
			for j >= 0 && nisLess(studentData[j].NIS, key.NIS) {
				studentData[j+1] = studentData[j]
				j--
			}
		} else {
			for j >= 0 && strings.ToLower(studentData[j].Name) < strings.ToLower(key.Name) {
				studentData[j+1] = studentData[j]
				j--
			}
		}
		studentData[j+1] = key
	}
}

// ==============================================================================
// ATTENDANCE MENU & CRUD
// ==============================================================================

/*
Subprogram: AttendanceMenu
Description: Navigation and operations for attendance tracking (Attendance CRUD).
Parameters: scanner (*bufio.Scanner), nS (*int), nA (*int)
Initial State (I.S.): Selected from MainMenu.
Final State (F.S.): Returns false if 'Back', true if exiting the program.
*/
func AttendanceMenu(scanner *bufio.Scanner, nS *int, nA *int) bool {
	inMenu := true
	for inMenu {
		fmt.Println("\n==============================================")
		fmt.Println("||             ATTENDANCE MENU              ||")
		fmt.Println("==============================================")
		fmt.Println("[1] Back to Menu")
		fmt.Println("[2] Add Student Attendance")
		fmt.Println("[3] Edit Student Attendance")
		fmt.Println("[4] Delete Student Attendance")
		fmt.Println("[5] Search Student Attendance")
		fmt.Println("[6] Sort Student Attendance")
		fmt.Println("[7] Show Student Attendance")
		fmt.Println("[0] Exit")
		fmt.Println("==============================================")
		fmt.Print("Select option: ")

		if !scanner.Scan() {
			return true
		}
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			return false
		case "2":
			AddAttendance(scanner, nS, nA)
		case "3":
			EditAttendance(scanner, nA)
		case "4":
			DeleteAttendance(scanner, nA)
		case "5":
			SearchAttendance(scanner, nA)
		case "6":
			if SortAttendanceMenu(scanner, nA) {
				return true
			}
		case "7":
			ShowAttendance(nA)
		case "0":
			fmt.Println("\n[!] Message: The program is finished. See you later!")
			return true
		default:
			fmt.Println("\n[!] Error: Invalid input!")
		}
	}
	return false
}

/*
Subprogram: AddAttendance
Description: Records attendance for ALL registered students simultaneously for a single validated date.
Parameters: scanner (*bufio.Scanner), nS (*int), nA (*int)
Initial State (I.S.): Student data array contains *nS elements.
Final State (F.S.): Prompts for a single date, validates it, then iterates through all *nS students

	to record presence. nA increments by the total number of students processed.
*/
func AddAttendance(scanner *bufio.Scanner, nS *int, nA *int) {
	if *nS == 0 {
		fmt.Println("\n[!] Error: No student data exists. Add student first!")
		return
	}

	if *nA+*nS > maxAttendance {
		fmt.Println("\n[!] Error: Attendance capacity is insufficient for all students!")
		return
	}

	fmt.Print("Input Date (DD MM YYYY separated by space): ")
	if !scanner.Scan() {
		return
	}

	var d, m, y int
	parsedCount, err := fmt.Sscanf(strings.TrimSpace(scanner.Text()), "%d %d %d", &d, &m, &y)
	if err != nil || parsedCount != 3 {
		fmt.Println("\n[!] Error: Invalid date format! Use numbers DD MM YYYY.")
		return
	}

	testDate := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
	if testDate.Year() != y || int(testDate.Month()) != m || testDate.Day() != d {
		fmt.Println("\n[!] Error: Calendar validation failed! The date is invalid.")
		return
	}

	dateExists := false
	for i := 0; i < *nA && !dateExists; i++ {
		if attendanceData[i].Date.Day == d && attendanceData[i].Date.Month == m && attendanceData[i].Date.Year == y {
			dateExists = true
		}
	}
	if dateExists {
		fmt.Printf("\n[!] Error: Attendance for %02d/%02d/%d already exists in the system!\n", d, m, y)
		return
	}

	fmt.Println("\n--- RECORDING ATTENDANCE FOR ALL STUDENTS ---")

	for i := 0; i < *nS; i++ {
		currentStudent := studentData[i]
		isValidInput := false
		var pres bool

		for !isValidInput {
			fmt.Printf("[%d/%d] NIS: %-8s | Name: %-20s -> Present? (y/n): ", i+1, *nS, currentStudent.NIS, currentStudent.Name)
			if !scanner.Scan() {
				return
			}
			presStr := strings.ToLower(strings.TrimSpace(scanner.Text()))

			switch presStr {
			case "y":
				pres = true
				isValidInput = true
			case "n":
				pres = false
				isValidInput = true
			default:
				fmt.Println("[!] Error: Input must be 'y' or 'n'!")
			}
		}

		attendanceData[*nA] = Attendance{
			Date:     Dates{Day: d, Month: m, Year: y},
			NIS:      currentStudent.NIS,
			Name:     currentStudent.Name,
			Presence: pres,
		}
		*nA++
	}

	fmt.Println("\n[+] Message: Successfully recorded attendance for all students on this date!")
}

/*
Subprogram: EditAttendance
Description: Updates the date or presence status of a student in attendanceData.
Parameters: scanner (*bufio.Scanner), nA (*int)
Initial State (I.S.): At least 1 attendance entry exists.
Final State (F.S.): Attributes of the modified element are updated.
*/
func EditAttendance(scanner *bufio.Scanner, nA *int) {
	if *nA == 0 {
		fmt.Println("\n[!] Error: No attendance data available!")
		return
	}

	ShowAttendance(nA)
	fmt.Printf("\nInput attendance number to edit (1 - %d): ", *nA)
	if !scanner.Scan() {
		return
	}

	var idx int
	_, err := fmt.Sscanf(strings.TrimSpace(scanner.Text()), "%d", &idx)
	if err != nil || idx < 1 || idx > *nA {
		fmt.Println("\n[!] Error: Invalid attendance number!")
		return
	}

	realIdx := SequentialSearchAttendanceByIdx(idx, nA)
	if realIdx == -1 {
		fmt.Println("\n[!] Error: Attendance record not found!")
		return
	}
	oldData := attendanceData[realIdx]
	dateStr := fmt.Sprintf("%02d/%02d/%d", oldData.Date.Day, oldData.Date.Month, oldData.Date.Year)
	pres := "ABSENT"
	if oldData.Presence {
		pres = "PRESENT"
	}

	fmt.Printf("\n[?] Found Data: %s | %s - %s | Status: %s\n", dateStr, oldData.NIS, oldData.Name, pres)
	fmt.Println("[Info] Enter '-' (dash) to keep current value.")

	fmt.Print("Input new Date (DD MM YYYY): ")
	if !scanner.Scan() {
		return
	}
	newDateStr := strings.TrimSpace(scanner.Text())
	newD, newM, newY := oldData.Date.Day, oldData.Date.Month, oldData.Date.Year

	if newDateStr != "-" {
		parsedCount, err := fmt.Sscanf(newDateStr, "%d %d %d", &newD, &newM, &newY)
		if err != nil || parsedCount != 3 {
			fmt.Println("\n[!] Error: Invalid date format!")
			return
		}

		testDate := time.Date(newY, time.Month(newM), newD, 0, 0, 0, 0, time.UTC)
		if testDate.Year() != newY || int(testDate.Month()) != newM || testDate.Day() != newD {
			fmt.Println("\n[!] Error: Calendar validation failed! The date is invalid.")
			return
		}

		dupFound := false
		for i := 0; i < *nA && !dupFound; i++ {
			if i != realIdx && attendanceData[i].NIS == oldData.NIS && attendanceData[i].Date.Day == newD && attendanceData[i].Date.Month == newM && attendanceData[i].Date.Year == newY {
				dupFound = true
			}
		}
		if dupFound {
			fmt.Println("\n[!] Error: Student already has attendance for this new date!")
			return
		}
	}

	fmt.Print("Is Present? (y/n): ")
	if !scanner.Scan() {
		return
	}
	newPresStr := strings.ToLower(strings.TrimSpace(scanner.Text()))
	newPres := oldData.Presence

	if newPresStr != "-" {
		switch newPresStr {
		case "y":
			newPres = true
		case "n":
			newPres = false
		default:
			fmt.Println("\n[!] Error: Input must be 'y' or 'n'!")
			return
		}
	}

	attendanceData[realIdx].Date = Dates{Day: newD, Month: newM, Year: newY}
	attendanceData[realIdx].Presence = newPres
	fmt.Println("\n[+] Message: Successfully updated attendance data!")
}

/*
Subprogram: DeleteAttendance
Description: Deletes attendance history with left-shifting array elements.
Parameters: scanner (*bufio.Scanner), nA (*int)
Initial State (I.S.): At least 1 attendance exists.
Final State (F.S.): Target index is deleted, subsequent elements shift left.
*/
func DeleteAttendance(scanner *bufio.Scanner, nA *int) {
	if *nA == 0 {
		fmt.Println("\n[!] Error: No attendance data available!")
		return
	}

	ShowAttendance(nA)
	fmt.Printf("\nInput attendance number to delete (1 - %d): ", *nA)
	if !scanner.Scan() {
		return
	}

	var idx int
	_, err := fmt.Sscanf(strings.TrimSpace(scanner.Text()), "%d", &idx)
	if err != nil || idx < 1 || idx > *nA {
		fmt.Println("\n[!] Error: Invalid input!")
		return
	}

	realIdx := SequentialSearchAttendanceByIdx(idx, nA)
	if realIdx == -1 {
		fmt.Println("\n[!] Error: Attendance record not found!")
		return
	}
	oldData := attendanceData[realIdx]
	dateStr := fmt.Sprintf("%02d/%02d/%d", oldData.Date.Day, oldData.Date.Month, oldData.Date.Year)

	fmt.Printf("\n[?] Are you sure you want to delete No.%d | %s | %s? (y/n): ", idx, dateStr, oldData.NIS)
	if !scanner.Scan() {
		return
	}
	if strings.ToLower(strings.TrimSpace(scanner.Text())) != "y" {
		fmt.Println("\n[-] Message: Delete operation cancelled.")
		return
	}

	for i := realIdx; i < *nA-1; i++ {
		attendanceData[i] = attendanceData[i+1]
	}
	attendanceData[*nA-1] = Attendance{}
	*nA--

	fmt.Println("\n[+] Message: Successfully deleted attendance data!")
}

/*
Subprogram: SearchAttendance
Description: Omnibox Sequential Search for attendance data.
Parameters: scanner (*bufio.Scanner), nA (*int)
Initial State (I.S.): Attendance data exists.
Final State (F.S.): Case-insensitive matching records are printed.
*/
func SearchAttendance(scanner *bufio.Scanner, nA *int) {
	if *nA == 0 {
		fmt.Println("\n[!] Error: Data is empty!")
		return
	}

	fmt.Print("Input Keyword (Date DD/MM/YYYY or NIS or Name): ")
	if !scanner.Scan() {
		return
	}
	keyword := strings.ToLower(strings.TrimSpace(scanner.Text()))

	if keyword == "" {
		fmt.Println("\n[!] Error: Keyword cannot be empty!")
		return
	}

	count := 0
	found := false

	for i := 0; i < *nA; i++ {
		dateStr := fmt.Sprintf("%02d/%02d/%d", attendanceData[i].Date.Day, attendanceData[i].Date.Month, attendanceData[i].Date.Year)
		currNIS := strings.ToLower(attendanceData[i].NIS)
		currName := strings.ToLower(attendanceData[i].Name)

		if strings.Contains(dateStr, keyword) || strings.Contains(currNIS, keyword) || strings.Contains(currName, keyword) {
			if !found {
				fmt.Println("\n=====================================================================")
				fmt.Printf("| %-3s | %-10s | %-13s | %-20s | %-7s |\n", "NO", "DATE", "NIS", "NAME", "STATUS")
				fmt.Println("=====================================================================")
				found = true
			}
			count++
			pres := "ABSENT"
			if attendanceData[i].Presence {
				pres = "PRESENT"
			}
			fmt.Printf("| %-3d | %-10s | %-13s | %-20s | %-7s |\n", count, dateStr, attendanceData[i].NIS, attendanceData[i].Name, pres)
		}
	}

	if found {
		fmt.Println("=====================================================================")
	} else {
		fmt.Println("\n[!] Error: No attendance data matches the keyword!")
	}
}

/*
Subprogram: SortAttendanceMenu
Description: Interface for sorting Attendance list data.
Parameters: scanner (*bufio.Scanner), nA (*int)
Initial State (I.S.): Selected from AttendanceMenu.
Final State (F.S.): Displays attendance data sorted by Date/NIS/Name (Asc/Desc). Returns bool for proper exit tracking.
*/
func SortAttendanceMenu(scanner *bufio.Scanner, nA *int) bool {
	inMenu := true
	for inMenu {
		fmt.Println("\n==============================================")
		fmt.Println("||          SORT ATTENDANCE MENU            ||")
		fmt.Println("==============================================")
		fmt.Println("[1] Ascending by Date (Selection Sort)")
		fmt.Println("[2] Descending by Date (Insertion Sort)")
		fmt.Println("[3] Ascending by NIS (Selection Sort)")
		fmt.Println("[4] Descending by NIS (Insertion Sort)")
		fmt.Println("[5] Ascending by Name (Selection Sort)")
		fmt.Println("[6] Descending by Name (Insertion Sort)")
		fmt.Println("[7] Back to Attendance Menu")
		fmt.Println("[0] Exit")
		fmt.Println("==============================================")
		fmt.Print("Select option: ")

		if !scanner.Scan() {
			return true
		}
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			SortAttendanceAsc(1, nA)
			fmt.Println("\n[+] Sorted Ascending by Date!")
			ShowAttendance(nA)
		case "2":
			SortAttendanceDesc(1, nA)
			fmt.Println("\n[+] Sorted Descending by Date!")
			ShowAttendance(nA)
		case "3":
			SortAttendanceAsc(2, nA)
			fmt.Println("\n[+] Sorted Ascending by NIS!")
			ShowAttendance(nA)
		case "4":
			SortAttendanceDesc(2, nA)
			fmt.Println("\n[+] Sorted Descending by NIS!")
			ShowAttendance(nA)
		case "5":
			SortAttendanceAsc(3, nA)
			fmt.Println("\n[+] Sorted Ascending by Name!")
			ShowAttendance(nA)
		case "6":
			SortAttendanceDesc(3, nA)
			fmt.Println("\n[+] Sorted Descending by Name!")
			ShowAttendance(nA)
		case "7":
			inMenu = false
		case "0":
			fmt.Println("\n[!] Message: The program is finished. See you later!")
			return true
		default:
			fmt.Println("\n[!] Error: Invalid input.")
		}
	}
	return false
}

/*
Subprogram: isDateLess
Description: Evaluates Date1 < Date2 chronologically.
Parameters: d1 (Dates), d2 (Dates)
Initial State (I.S.): Valid dates.
Final State (F.S.): Boolean comparison time.
*/
func isDateLess(d1, d2 Dates) bool {
	if d1.Year != d2.Year {
		return d1.Year < d2.Year
	}
	if d1.Month != d2.Month {
		return d1.Month < d2.Month
	}
	return d1.Day < d2.Day
}

/*
Subprogram: SortAttendanceAsc
Description: Selection Sort for attendance array in ascending order.
Parameters: mode (int), nA (*int)
Initial State (I.S.): Randomly ordered array.
Final State (F.S.): Array is sorted from smallest to largest.
*/
func SortAttendanceAsc(mode int, nA *int) {
	for i := 0; i < *nA-1; i++ {
		minIdx := i
		for j := i + 1; j < *nA; j++ {
			switch mode {
			case 1:
				if isDateLess(attendanceData[j].Date, attendanceData[minIdx].Date) {
					minIdx = j
				}
			case 2:
				if nisLess(attendanceData[j].NIS, attendanceData[minIdx].NIS) {
					minIdx = j
				}
			default:
				if strings.ToLower(attendanceData[j].Name) < strings.ToLower(attendanceData[minIdx].Name) {
					minIdx = j
				}
			}
		}
		attendanceData[i], attendanceData[minIdx] = attendanceData[minIdx], attendanceData[i]
	}
}

/*
Subprogram: SortAttendanceDesc
Description: Insertion Sort for attendance array in descending order.
Parameters: mode (int), nA (*int)
Initial State (I.S.): Randomly ordered array.
Final State (F.S.): Array is sorted from largest to smallest.
*/
func SortAttendanceDesc(mode int, nA *int) {
	for i := 1; i < *nA; i++ {
		key := attendanceData[i]
		j := i - 1

		switch mode {
		case 1:
			for j >= 0 && isDateLess(attendanceData[j].Date, key.Date) {
				attendanceData[j+1] = attendanceData[j]
				j--
			}
		case 2:
			for j >= 0 && attendanceData[j].NIS < key.NIS {
				attendanceData[j+1] = attendanceData[j]
				j--
			}
		default:
			for j >= 0 && strings.ToLower(attendanceData[j].Name) < strings.ToLower(key.Name) {
				attendanceData[j+1] = attendanceData[j]
				j--
			}
		}
		attendanceData[j+1] = key
	}
}

/*
Subprogram: ShowAttendance
Description: Displays the attendanceData list to the terminal.
Parameters: nA (*int)
Initial State (I.S.): Attendance data exists.
Final State (F.S.): Printed to the terminal with aligned borders.
*/
func ShowAttendance(nA *int) {
	if *nA == 0 {
		fmt.Println("\n[!] Message: No attendance data available.")
		return
	}
	fmt.Println("\n=====================================================================")
	fmt.Printf("| %-3s | %-10s | %-13s | %-20s | %-7s |\n", "NO", "DATE", "NIS", "NAME", "STATUS")
	fmt.Println("=====================================================================")
	for i := 0; i < *nA; i++ {
		dateStr := fmt.Sprintf("%02d/%02d/%d", attendanceData[i].Date.Day, attendanceData[i].Date.Month, attendanceData[i].Date.Year)
		pres := "ABSENT"
		if attendanceData[i].Presence {
			pres = "PRESENT"
		}
		fmt.Printf("| %-3d | %-10s | %-13s | %-20s | %-7s |\n", i+1, dateStr, attendanceData[i].NIS, attendanceData[i].Name, pres)
	}
	fmt.Println("=====================================================================")
}

// ==============================================================================
// STATISTICS MENU
// ==============================================================================

/*
Subprogram: StatisticsMenu
Description: Main menu for computing attendance analysis.
Parameters: scanner (*bufio.Scanner), nS (*int), nA (*int)
Initial State (I.S.): Selected from MainMenu.
Final State (F.S.): Executes statistical requests. Returns false if 'Back', true if exiting the program.
*/
func StatisticsMenu(scanner *bufio.Scanner, nS *int, nA *int) bool {
	inMenu := true
	for inMenu {
		fmt.Println("\n==============================================")
		fmt.Println("||             STATISTICS MENU              ||")
		fmt.Println("==============================================")
		fmt.Println("[1] Back to Menu")
		fmt.Println("[2] Show Student Statistics")
		fmt.Println("[3] Search Student Statistic")
		fmt.Println("[4] Sort Student Statistic")
		fmt.Println("[0] Exit")
		fmt.Println("==============================================")
		fmt.Print("Select option: ")

		if !scanner.Scan() {
			return true
		}
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			inMenu = false
		case "2":
			ShowStatsData(nS, nA)
		case "3":
			SearchStats(scanner, nS, nA)
		case "4":
			if SortStatsMenu(scanner, nS, nA) {
				return true
			}
		case "0":
			fmt.Println("\n[!] Message: The program is finished. See you later!")
			return true
		default:
			fmt.Println("\n[!] Error: Invalid input!")
		}
	}
	return false
}

/*
Subprogram: GenerateStats
Description: Generates the student attendance tabulation on-the-fly.
Parameters: nS (*int), nA (*int)
Initial State (I.S.): Data exists in memory.
Final State (F.S.): Returns TabStatistics along with the total data count.
*/
func GenerateStats(nS *int, nA *int) (TabStatistics, int) {
	var stats TabStatistics
	for i := 0; i < *nS; i++ {
		stats[i].NIS = studentData[i].NIS
		stats[i].Name = studentData[i].Name
		stats[i].TotalMeetings = 0
		stats[i].TotalPresence = 0

		for j := 0; j < *nA; j++ {
			if attendanceData[j].NIS == stats[i].NIS {
				stats[i].TotalMeetings++
				if attendanceData[j].Presence {
					stats[i].TotalPresence++
				}
			}
		}

		if stats[i].TotalMeetings > 0 {
			stats[i].Percentage = (float64(stats[i].TotalPresence) / float64(stats[i].TotalMeetings)) * 100
		} else {
			stats[i].Percentage = 0
		}
	}
	return stats, *nS
}

/*
Subprogram: PrintStatsFormat
Description: Prints a single row of TabStatistics.
Parameters: s (Statistics)
Initial State (I.S.): Struct variable is provided.
Final State (F.S.): Displayed using left-aligned formatting.
*/
func PrintStatsFormat(s Statistics) {
	fmt.Printf("| %-13s | %-20s | %-3d/ %-3d | %-6.2f%% |\n", s.NIS, s.Name, s.TotalPresence, s.TotalMeetings, s.Percentage)
}

/*
Subprogram: ShowStatsData
Description: Prints the full summary of all students.
Parameters: nS (*int), nA (*int)
Initial State (I.S.): Called from within MainLoop.
Final State (F.S.): Table display is printed.
*/
func ShowStatsData(nS *int, nA *int) {
	stats, count := GenerateStats(nS, nA)
	if count == 0 {
		fmt.Println("\n[!] Message: No student data available to summarize!")
		return
	}
	SortStatsDesc(&stats, count, 3)
	fmt.Println("\n=============================================================")
	fmt.Printf("| %-13s | %-20s | %-7s | %-8s |\n", "NIS", "NAME", "PRS/MTG", "PERCENT")
	fmt.Println("=============================================================")
	for i := 0; i < count; i++ {
		PrintStatsFormat(stats[i])
	}
	fmt.Println("=============================================================")
}

/*
Subprogram: SearchStats
Description: Sequential Search for statistical data.
Parameters: scanner (*bufio.Scanner), nS (*int), nA (*int)
Initial State (I.S.): nS > 0
Final State (F.S.): Filtered matching data is printed.
*/
func SearchStats(scanner *bufio.Scanner, nS *int, nA *int) {
	if *nS == 0 {
		fmt.Println("\n[!] Error: No data to search!")
		return
	}

	fmt.Print("Input Keyword (NIS or Name): ")
	if !scanner.Scan() {
		return
	}
	keyword := strings.ToLower(strings.TrimSpace(scanner.Text()))
	if keyword == "" {
		fmt.Println("\n[!] Error: Keyword cannot be empty!")
		return
	}

	stats, count := GenerateStats(nS, nA)
	found := false

	for i := 0; i < count; i++ {
		currNIS := strings.ToLower(stats[i].NIS)
		currName := strings.ToLower(stats[i].Name)

		if strings.Contains(currNIS, keyword) || strings.Contains(currName, keyword) {
			if !found {
				fmt.Println("\n=============================================================")
				fmt.Printf("| %-13s | %-20s | %-7s | %-8s |\n", "NIS", "NAME", "PRS/MTG", "PERCENT")
				fmt.Println("=============================================================")
				found = true
			}
			PrintStatsFormat(stats[i])
		}
	}

	if found {
		fmt.Println("=============================================================")
	} else {
		fmt.Println("\n[!] Error: Statistics data not found for keyword!")
	}
}

/*
Subprogram: SortStatsMenu
Description: Menu for dynamic sorting of local arrays based on NIS, Name, and Percentage.
Parameters: scanner (*bufio.Scanner), nS (*int), nA (*int)
Initial State (I.S.): Selected from StatisticsMenu.
Final State (F.S.): Sorting algorithms are applied. Returns bool for graceful exit tracking.
*/
func SortStatsMenu(scanner *bufio.Scanner, nS *int, nA *int) bool {
	inMenu := true
	for inMenu {
		fmt.Println("\n==============================================")
		fmt.Println("||           SORT STATISTICS MENU           ||")
		fmt.Println("==============================================")
		fmt.Println("[1] Ascending by NIS (Selection Sort)")
		fmt.Println("[2] Descending by NIS (Insertion Sort)")
		fmt.Println("[3] Ascending by Name (Selection Sort)")
		fmt.Println("[4] Descending by Name (Insertion Sort)")
		fmt.Println("[5] Ascending by Percentage (Selection Sort)")
		fmt.Println("[6] Descending by Percentage (Insertion Sort)")
		fmt.Println("[7] Back to Statistics Menu")
		fmt.Println("[0] Exit")
		fmt.Println("==============================================")
		fmt.Print("Select option: ")

		if !scanner.Scan() {
			return true
		}
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "7":
			inMenu = false
		case "0":
			fmt.Println("\n[!] Message: The program is finished. See you later!")
			return true
		default:
			stats, count := GenerateStats(nS, nA)
			if count == 0 {
				fmt.Println("\n[!] Error: Data is empty!")
			} else {
				switch input {
				case "1":
					SortStatsAsc(&stats, count, 1)
					fmt.Println("\n[+] Sorted Ascending by NIS!")
					ShowSortedStats(stats, count)
				case "2":
					SortStatsDesc(&stats, count, 1)
					fmt.Println("\n[+] Sorted Descending by NIS!")
					ShowSortedStats(stats, count)
				case "3":
					SortStatsAsc(&stats, count, 2)
					fmt.Println("\n[+] Sorted Ascending by Name!")
					ShowSortedStats(stats, count)
				case "4":
					SortStatsDesc(&stats, count, 2)
					fmt.Println("\n[+] Sorted Descending by Name!")
					ShowSortedStats(stats, count)
				case "5":
					SortStatsAsc(&stats, count, 3)
					fmt.Println("\n[+] Sorted Ascending by Percentage!")
					ShowSortedStats(stats, count)
				case "6":
					SortStatsDesc(&stats, count, 3)
					fmt.Println("\n[+] Sorted Descending by Percentage!")
					ShowSortedStats(stats, count)

				default:
					fmt.Println("\n[!] Error: Invalid input.")
				}
			}
		}
	}
	return false
}

/*
Subprogram: SortStatsAsc
Description: Selection Sort for TabStatistics array in ascending order.
Parameters: stats (*TabStatistics), count (int), mode (int)
Initial State (I.S.): Pointer to stats is provided.
Final State (F.S.): Elements are shifted to be sorted from smallest to largest.
*/
func SortStatsAsc(stats *TabStatistics, count int, mode int) {
	for i := 0; i < count-1; i++ {
		minIdx := i
		for j := i + 1; j < count; j++ {
			switch mode {
			case 1:
				if nisLess(stats[j].NIS, stats[minIdx].NIS) {
					minIdx = j
				}
			case 2:
				if strings.ToLower(stats[j].Name) < strings.ToLower(stats[minIdx].Name) {
					minIdx = j
				}
			default:
				if stats[j].Percentage < stats[minIdx].Percentage {
					minIdx = j
				}
			}
		}
		stats[i], stats[minIdx] = stats[minIdx], stats[i]
	}
}

/*
Subprogram: SortStatsDesc
Description: Insertion Sort for TabStatistics array in descending order.
Parameters: stats (*TabStatistics), count (int), mode (int)
Initial State (I.S.): Local stats array is in random order.
Final State (F.S.): Stats array is sorted from largest to smallest.
*/
func SortStatsDesc(stats *TabStatistics, count int, mode int) {
	for i := 1; i < count; i++ {
		key := stats[i]
		j := i - 1

		switch mode {
		case 1:
			for j >= 0 && nisLess(stats[j].NIS, key.NIS) {
				stats[j+1] = stats[j]
				j--
			}
		case 2:
			for j >= 0 && strings.ToLower(stats[j].Name) < strings.ToLower(key.Name) {
				stats[j+1] = stats[j]
				j--
			}
		default:
			for j >= 0 && stats[j].Percentage < key.Percentage {
				stats[j+1] = stats[j]
				j--
			}
		}
		stats[j+1] = key
	}
}

/*
Subprogram: ShowSortedStats
Description: Prints a specific sorted statistics array.
Parameters: stats (TabStatistics), count (int)
Initial State (I.S.): Parametric array is available.
Final State (F.S.): Output is aligned and printed.
*/
func ShowSortedStats(stats TabStatistics, count int) {
	fmt.Println("\n=============================================================")
	fmt.Printf("| %-13s | %-20s | %-7s | %-8s |\n", "NIS", "NAME", "PRS/MTG", "PERCENT")
	fmt.Println("=============================================================")
	for i := 0; i < count; i++ {
		PrintStatsFormat(stats[i])
	}
	fmt.Println("=============================================================")
}
