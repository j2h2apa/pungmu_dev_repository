package main

import (
	"strconv"
	"time"
)

/*Car : */
type Car struct {
	val string
}

/*Plane : */
type Plane struct {
	val string
}

/*MakeTire : */
func MakeTire(carChan, outChan chan Car, planeChan chan Plane, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "Car Tire, "
			outChan <- car
		case plane := <-planeChan:
			plane.val += "Plane Tire, "
			outPlaneChan <- plane
		}
	}
}

/*MakeEngine : */
func MakeEngine(carChan, outChan chan Car, planeChan, outPlaneChan chan Plane) {
	for {
		select {
		case car := <-carChan:
			car.val += "CarEngine, "
			outChan <- car
		case plane := <-planeChan:
			plane.val += "PlaneEngine, "
			outPlaneChan <- plane
		}
	}
}

/*StartWork :*/
func StartWork(carChan chan Car) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		carChan <- Car{val: "Car " + strconv.Itoa(i) + " "}
		i++
	}
}

/*StartPlaneWork :*/
func StartPlaneWork(chan1 chan Plane) {
	i := 0
	for {
		time.Sleep(1 * time.Second)
		chan1 <- Plane{val: "Plane " + strconv.Itoa(i) + " "}
		i++
	}
}
