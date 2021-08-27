//generim parol s spec simvolami i simvolami raznoi visoti
//generiruem parol na osnove vvedenogo ranee slova

package main

import (
	"fmt"
  "os"
  "os/exec"
	"math/rand"
	"time"
	"github.com/eiannone/keyboard"

)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func Cls() {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()
}

func rndint(min int,max int) int {
	a := min + rand.Intn(max-min+1)
	return a
}

func main() {


	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()
  Cls()
  state := 4
	var old,newpass string

  for {
		if state == 4 {
			Cls()

			fmt.Println("Generate new strong password press 1 ,return menu press 4")
			fmt.Println("Make your old password stronger press 2")
		  fmt.Println("Press Esc to quit")
		} else if state == 2 {
			Cls()
      //OBRABOTKa VVEDENOGO PAROLJA I EGO USLOZNENIE
			fmt.Println("Enter text and press Enter : ")
			keyboard.Close()
			fmt.Scanf("%s\n", &old)
			keyboard.Open()
			//proverka dlinni i esli mense 8 znakov napisat sto paro sliwkom mal,menjaem vtoruju i tretju bukvu!
			if len(old) > 4 {
				old = old[:2] + string(rndint(65,90)) + old[3:]
			} else {
				fmt.Println("Your password to short need not lower 8 symbols ")
				}

			old = old + string(rndint(33,47))
			fmt.Println("You Enter text: ",old)
			state = 4
		} else if state == 0 {
			break
		} else if state == 1 {
      Cls()
      //TUT GENERATOR PAROLEI ZELATELNO
			//na tretjem etape sdelat generator iz normalnih slov!


			for i := 1; i < 11; i++{
        newpass = newpass + string(rndint(33,126)) //33 + rand.Intn(126-33+1))
      }
      fmt.Println("Your new pass 10 symbols long,enter to change : ",newpass)
		  newpass=""
    }

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		if key == keyboard.KeyEsc { state = 0
		} else if char == 50  { state = 2
		} else if char == 49 { state = 1
    } else if char == 52 { state = 4 }


	}

}
