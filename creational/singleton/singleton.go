/**
Requirements and acceptance criteria
There are some requirements and acceptance criteria to write the described single counter. They are as follows:

- When no counter has been created before, a new one is created with the value 0
- If a counter has already been created, return this instance that holds the actual count
- If we call the method AddOne, the count must be incremented by 1
*/

package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance() Singleton {
	return nil
}

func (s singleton) AddOne() int {
	return 0
}
