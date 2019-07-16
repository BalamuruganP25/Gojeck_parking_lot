package Processuserinput


import (
 "fmt"
 "log" 
 "strings"
 "strconv"
 Utils "Gojeck_Parking_lot_problem_result/src/Utils"
 "os"
)

const total_size int = 100
type parkingdetails struct{

Parking_lot_count int

Cardetails [total_size] Parking_car_details

}

type Parking_car_details  struct{

Slot_number int
Carnumber string
CarColor string


}
var lastupdated_park_lot_count int = 1
var store_parking_details parkingdetails
var total_slot int = 0


func Process_inputfile_data(inputfile_data string){
	

	
if strings.Contains(inputfile_data, "create_parking_lot") {
	
	if total_slot == 0{
	
	  split_parkinglot :=  strings.Split(inputfile_data, " ")	 
	
	if len(split_parkinglot) == 2{

		store_parking_details.Parking_lot_count = Utils.Convertstringtoint(split_parkinglot[1])
		total_slot = Utils.Convertstringtoint(split_parkinglot[1])
		log.Println("Created a parking lot with",store_parking_details.Parking_lot_count,"slots")
		fmt.Println("Created a parking lot with",store_parking_details.Parking_lot_count,"slots")
		
	}else{
	
		log.Println("Plz enter valid data")
		fmt.Println("Plz enter valid data")
	}
	
	}else{
	
	
		log.Println("Slot alread create")
		fmt.Println("Slot alread create")
	
	}
  }
	
if strings.Contains(inputfile_data, "park ") {


	if total_slot == 0{

		fmt.Println("please create parking slot")
		return

	}
	
	split_parkingcar_details :=  strings.Split(inputfile_data, " ")

	if len(split_parkingcar_details) == 3{

	log.Println("Allocated slot number :",lastupdated_park_lot_count)
	fmt.Println("Allocated slot number :",lastupdated_park_lot_count)

	if lastupdated_park_lot_count <= store_parking_details.Parking_lot_count {
	
	for i:= 1;i<=store_parking_details.Parking_lot_count;i++{
	
	if store_parking_details.Cardetails[i].Slot_number == 0{ 

		store_parking_details.Cardetails[i].Slot_number = i
		store_parking_details.Cardetails[i].Carnumber = split_parkingcar_details[1]
		store_parking_details.Cardetails[i].CarColor = split_parkingcar_details[2]
		lastupdated_park_lot_count ++
		break
	
		  }
	  }

	}else{
	
		log.Println("Sorry, parking lot is full")
		fmt.Println("Sorry, parking lot is full")
	}
  
   }else{
	
  	log.Println("Plz enter valid data")
  	fmt.Println("Plz enter valid data")
 	
   }
 
 }
	
if strings.Contains(inputfile_data, "leave") {

	//fmt.Println("inputfile_data ==>",inputfile_data)


		if total_slot == 0{

		fmt.Println("please create parking slot")
		return

	   }

	split_leave_value :=  strings.Split(inputfile_data, " ")
	log.Println("split_leave_value ==>",split_leave_value)
	
	if len(split_leave_value) == 2{
	
		solt_number := Utils.Convertstringtoint(split_leave_value[1])
		store_parking_details.Cardetails[solt_number].Slot_number = 0
		store_parking_details.Cardetails[solt_number].Carnumber = ""
		store_parking_details.Cardetails[solt_number].CarColor = ""
		lastupdated_park_lot_count =  lastupdated_park_lot_count -1
		log.Println("Slot number ",solt_number," is free") 
	
	}else{
	
	log.Println("Enter valid data")
	fmt.Println("Enter valid data")
	
	
  }
	
}
	
if strings.Contains(inputfile_data, "status") {

 fmt.Println("Your status ==>",inputfile_data)

	if total_slot == 0{

		fmt.Println("please create parking slot")
		return

	}

	if lastupdated_park_lot_count == 1{


		fmt.Println("please update data")

	}

	log.Println("status -",inputfile_data)

	
	
	log.Println("Slot No.","Registration No","Colour")
	
	var Empty_slot string = ""
	
	for i:= 1 ; i< total_slot;i++{
	
		if store_parking_details.Cardetails[i].Slot_number == 0{
		
		Empty_slot = Empty_slot + strconv.Itoa(i) + " "
		
		}else{
		log.Println(store_parking_details.Cardetails[i].Slot_number ," ",store_parking_details.Cardetails[i].Carnumber,"  ",store_parking_details.Cardetails[i].CarColor)
		fmt.Println(store_parking_details.Cardetails[i].Slot_number ," ",store_parking_details.Cardetails[i].Carnumber,"  ",store_parking_details.Cardetails[i].CarColor)
		}
	
	}
	
	if(Empty_slot != ""){
	
	log.Println("Empty_slot ==>",Empty_slot)
	
	}
	
}
	
if strings.Contains(inputfile_data, "registration_numbers_for_cars_with_colour") {
	
	//log.Println("registration_numbers_for_cars_with_colour -",inputfile_data)
	
	if total_slot == 0{

		fmt.Println("please create parking slot")
		return

	}
	split_car_colour := strings.Split(inputfile_data, " ")
	
	searching_data := ""
	if len(split_car_colour) == 2{
	
	log.Println("registration_numbers_for_cars_with_colour")
	//log.Println()
	for i:=1;i<=lastupdated_park_lot_count;i++{
	
		if store_parking_details.Cardetails[i].CarColor == split_car_colour[1]{
		
		log.Print(store_parking_details.Cardetails[i].Carnumber," ")
		searching_data = store_parking_details.Cardetails[i].Carnumber
	}
	
	}
	
		if searching_data ==""{
		
			log.Println("Your searching data not found")
			fmt.Println("Your searching data not found")
		
	     }
	
	log.Println()
	}else{
	
		log.Println("Enter valid data")
		fmt.Println("Enter valid data")
	
  }
	
}
	
if strings.Contains(inputfile_data, "slot_numbers_for_cars_with_colour") {
	

	if total_slot == 0{

		fmt.Println("please create parking slot")
		return

	}
	
	split_car_colour := strings.Split(inputfile_data, " ")
	
	searching_data := 0
	
	if len(split_car_colour) == 2{
	
	log.Println("slot_numbers_for_cars_with_colour White")
	
	for i:=1;i<=lastupdated_park_lot_count;i++{
	
		if store_parking_details.Cardetails[i].CarColor == split_car_colour[1]{
		log.Print(store_parking_details.Cardetails[i].Slot_number," ")
		searching_data = store_parking_details.Cardetails[i].Slot_number 
		
	 }
	
    }
	
	    if searching_data <= 0{
		
		    log.Println("Your searching data not found")
		    fmt.Println("Your searching data not found")
	 }
	log.Println()
	
	}else{
	
	 log.Println("Enter valid data")
	 fmt.Println("Enter valid data")
	
  }
	
}
	
if strings.Contains(inputfile_data, "slot_number_for_registration_number") {

	
	if total_slot == 0{

		fmt.Println("please create parking slot")
		return

	}

	
	split_car_number := strings.Split(inputfile_data, " ")
	
	searching_data := 0
	
	if len(split_car_number) == 2{
	
	log.Println("slot_number_for_registration_number")
	
	for i:=1;i<=lastupdated_park_lot_count;i++{
	
		if store_parking_details.Cardetails[i].Carnumber == split_car_number[1]{
		
		log.Println(store_parking_details.Cardetails[i].Slot_number)
		searching_data  = store_parking_details.Cardetails[i].Slot_number
		
		}
	
	}
	
	if searching_data <= 0{
	
	 log.Println("Your searching data not found")
	  fmt.Println("Your searching data not found")

	}
	log.Println()
	
	}else{
	
	  log.Println("Enter valid data")
	  fmt.Println("Enter valid data")
	
	}
	
}


if inputfile_data == "exit"{

		os.Exit(3)
	}

if inputfile_data == ""{

		fmt.Println("please enter valid commend")
}








}
