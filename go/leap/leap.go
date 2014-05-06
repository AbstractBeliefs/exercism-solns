package leap

func IsLeapYear(year int) bool {
    // Check base case
    if year % 4 == 0 {

        // Handle centuries
        if (year % 100 == 0) && !(year % 400 == 0) {
            return false
        }

        return true
    }

    return false
}
