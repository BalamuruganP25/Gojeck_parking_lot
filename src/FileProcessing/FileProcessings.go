package FileProcessing

import(

    "log"
    "os"
    "fmt"
    "encoding/json"
)


type Logfiledetails struct {

Logfilepath string
Logfilefoldername string
Logfilename string
Useruploadfilepath string

}
var Logdetails Logfiledetails

func Readconfiguration() (bool) {
   readconfig,err := os.Open("/root/Projects/src/Gojeck_Parking_lot_problem_result/src/Configuration/Configuration.json")
    if err !=nil{
      fmt.Println(err)
      return false
    }
    defer readconfig.Close()
    decoder := json.NewDecoder(readconfig)
    Logdetails = Logfiledetails{}
    errs := decoder.Decode(&Logdetails)
    if errs != nil{
      fmt.Println(errs)
      log.Println(errs)
      return false
    }
    fmt.Println("Configuration ==>",Logdetails)
    return true
}

func Createfolder(){

 folderlocation := Logdetails.Logfilepath
 FolderName := Logdetails.Logfilefoldername
 var appendfile = folderlocation + FolderName
 _, err := os.Stat(appendfile)
   
    if err != nil {     
        if os.IsNotExist(err) {
            
            os.Mkdir(appendfile, 755)
        }

    }else {


    	log.Println("Folder Create Alread")

    }

}


func Checklogfilestatus(){

   var appendfile_file_path = Logdetails.Logfilepath + Logdetails.Logfilefoldername +"/"+ Logdetails.Logfilename
   _, err := os.Stat(appendfile_file_path)
    if err != nil {
        if os.IsNotExist(err) {
            log.Println("File does not exist.")
            Createfile(appendfile_file_path)
        }

    }else {
    	Createfile(appendfile_file_path)
    	log.Println("File Create Done")

    }

}


func Createfile(createfileapth string){
	var _, err = os.Stat(createfileapth)
	if os.IsNotExist(err) {
		var file, err = os.Create(createfileapth)
		if err != nil { return }
		defer file.Close()
	}

	log.Println("==> done creating file", createfileapth)
}