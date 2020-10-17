package package1

import "strconv"

//zad 1
func Suma(a, b int) int {
	return a + b
}

//zad 2
func Concatenate(s1, s2 []int) []int {
	return append(s1, s2...)
}

//zad 3
func Is18(age int) bool {
	if age >= 18 {
		return true
	} else {
		return false
	}
}

//zad 4
func PrintDividedBy4(n int) []int {
	m := n
	s1 := []int{}
	for m >= 4 {
		if m%4 == 0 {
			s1 = append(s1, m)
			m -= 4
		} else {
			m -= 1
		}
	}
	return s1
}

//zad 5
func FilterEven(age []int) []int {
	m := []int{}
	for _, v := range age {
		if age[v-1]%2 == 1 {
			m = append(m, age[v-1])
		} else {
			continue
		}
	}
	return m
}

//zad 6
type User struct {
	name     string
	lastName string
	age      int
	isAdult  bool
}

func NewUser(_name, _lastName string, _age int) *User {
	_User := new(User)
	_isAdult := Is18(_age)

	_User.name = _name
	_User.lastName = _lastName
	_User.age = _age
	_User.isAdult = _isAdult

	return _User
}

//zad 7
func (u *User) ToString() string {
	str := "This user has following attributes: \n"
	str += "Name: " + u.name + "\n"
	str += "Last Name: " + u.lastName + "\n"
	str += "Age: " + strconv.Itoa(u.age) + "\n"
	str += "Is an Adult: " + strconv.FormatBool(u.isAdult)
	return str
}
func (u *User) ChangeAge(newAge int) {
	u.age = newAge
	u.isAdult = Is18(newAge)
}
