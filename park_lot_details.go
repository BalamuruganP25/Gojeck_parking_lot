package main

	

import (

 "os"
 "bufio"
 "strings"
 "log"
 "fmt"
 "time"
 FileProcessing "Gojeck_Parking_lot_problem_result/src/FileProcessing"
 Processuserinput "Gojeck_Parking_lot_problem_result/src/Processuserinput"

)



func tooldetails(){

	data_and_time := time.Now()
    fmt.Println("************************************************************************************************************************")
    fmt.Println("Tool Name               			  :  parking-lot")
    fmt.Println("Language                			  :  Golang")
    fmt.Println("Developer              			          :  P.Bala")
    fmt.Println("Appliaction Start time                            :",data_and_time.Format("01-02-2006 15:04:05 Monday"))
    fmt.Println("************************************************************************************************************************")
 
 }


func usermanuaenteroptions(){

	fmt.Println("User Enter Manual option")
	fmt.Println("Note - Below Key words only accept")
	fmt.Println("1.create_parking_lot 6")
	fmt.Println("2.park KA-01-HH-1234 White")
	fmt.Println("3.leave 4")
	fmt.Println("4.status")
	fmt.Println("5.registration_numbers_for_cars_with_colour White")
	fmt.Println("6.slot_numbers_for_cars_with_colour White")
	fmt.Println("7.slot_number_for_registration_number KA-01-HH-3141")
	fmt.Println("8.exit")
	fmt.Println("8.Enter file name example - parking-lot.txt")
	

}



func Process_user_upload_file(filename string){

		path := FileProcessing.Logdetails.Useruploadfilepath + filename 


		filedata,err := os.Open(path)

		//log.Println("filedata ==>",filedata)
		
		if err != nil {
		log.Print(err)
		fmt.Println(err)
		}
		
		scanner := bufio.NewScanner(filedata)
		
		for scanner.Scan() {
		
			Processuserinput.Process_inputfile_data(scanner.Text())
			//log.Println("Scanner ==>",scanner.Text())
		
		}
}



 func Process_user_enter_manual_data(){

 	for {

 	     consoleReader := bufio.NewReader(os.Stdin)
         fmt.Print("Type user option(Before refer user option).... \n")
         input, _ := consoleReader.ReadString('\n')
         replace_junk_character := strings.Replace(input, "\n", "", -1)
         log.Println("Your options : ", replace_junk_character)
         fmt.Println("Your options : ", replace_junk_character)
         if strings.Contains(replace_junk_character, ".txt"){

         	Process_user_upload_file(replace_junk_character)

         }else{

         	  Processuserinput.Process_inputfile_data(replace_junk_character)

         }
   
   }


 }


func main(){

	tooldetails()

	Configuration_file_read_status :=  FileProcessing.Readconfiguration()

	if Configuration_file_read_status {

	FileProcessing.Createfolder()
	FileProcessing.Checklogfilestatus()
	usermanuaenteroptions()

	Open_File := FileProcessing.Logdetails.Logfilepath + FileProcessing.Logdetails.Logfilefoldername +"/"+FileProcessing.Logdetails.Logfilename
	fmt.Println("Open File ==>",Open_File)
	file, err := os.OpenFile(Open_File, os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
	    log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	log.SetOutput(file)
	//log.Println("This is a test log entry")
	
	if len(os.Args) == 1{
		
		
		Process_user_enter_manual_data()

	}else{

		filename := os.Args[1] 

		Process_user_upload_file(filename)
		Process_user_enter_manual_data()
		
		

	}

	
  }else{


  	log.Println("Check Configuration file")


  }

 }