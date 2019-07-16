package Utils
import(


	"strconv"
)

func Convertstringtoint(recevienumber string) (int){

  convert_int_value, err := strconv.Atoi(recevienumber)

  if err == nil{
  }else{
    panic(err)
 }

return convert_int_value

}