package hdy

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

var waitGroup sync.WaitGroup
type A struct {
	monsterIdCounter int32
	Monsters sync.Map
}

type Monster struct {
	 Id int32
	 Name string
}

func TestAtomic(t *testing.T) {
	testA := &A{}

	goCount := 100
	waitGroup.Add(goCount)
	for i := 0; i < goCount; i++ {
		go func() {
			atomic.AddInt32(&testA.monsterIdCounter, 1)
			monster := &Monster{
				Id:   testA.monsterIdCounter,
				Name: fmt.Sprintf("name-%d", testA.monsterIdCounter),
			}
			_, isExist := testA.Monsters.LoadOrStore(monster.Id, monster)
			if isExist {
				fmt.Println(fmt.Sprintf("重复的id:%d", monster.Id))
			}


			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	fmt.Println(fmt.Sprintf("任务量:%d, 最后结果:%d", goCount, testA.monsterIdCounter))

	nameMap := make(map[string]bool, 0)
	testA.Monsters.Range(func(key, value interface{}) bool {
		//id := key.(int32)
		monster := value.(*Monster)

		if _, ok := nameMap[monster.Name]; ok {
			fmt.Println(fmt.Sprintf("怪物名字重复:%s", monster.Name))
		} else {
			nameMap[monster.Name] = true
		}

		return false
	})
}

func TestAtomic2(t *testing.T) {
	testA := &A{}

	goCount := 10000
	waitGroup.Add(goCount)
	for i := 0; i < goCount; i++ {
		go func() {
			eachId := atomic.AddInt32(&testA.monsterIdCounter, 1)
			monster := &Monster{
				Id:   eachId,
				Name: fmt.Sprintf("name-%d", eachId),
			}
			_, isExist := testA.Monsters.LoadOrStore(eachId, monster)
			if isExist {
				fmt.Println(fmt.Sprintf("重复的id:%d", eachId))
			}


			waitGroup.Done()
		}()
	}

	waitGroup.Wait()

	fmt.Println(fmt.Sprintf("任务量:%d, 最后结果:%d", goCount, testA.monsterIdCounter))

	nameMap := make(map[string]bool, 0)
	testA.Monsters.Range(func(key, value interface{}) bool {
		//id := key.(int32)
		monster := value.(*Monster)

		if _, ok := nameMap[monster.Name]; ok {
			fmt.Println(fmt.Sprintf("怪物名字重复:%s", monster.Name))
		} else {
			nameMap[monster.Name] = true
		}

		return false
	})
}