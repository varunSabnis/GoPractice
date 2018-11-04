package main

import (
  "fmt"
  "encoding/csv"
  "bufio"
   "flag"
   "os"
   "strings"
   "time"
)

func main() {
  
   filename := flag.String("file","problems.csv","csv file containing question and answers")
   timelimit := flag.Int("timelimit",30,"time limit for the quiz")
   flag.Parse()

   f,_ := os.Open(*filename)
   r := csv.NewReader(bufio.NewReader(f))

   correct_answer := 0
   var ans string

   records,err := r.ReadAll()
   if(err!=nil){
     fmt.Println("Error")
   }

   total_questions := len(records)

   timer := time.NewTimer(time.Duration(*timelimit) * time.Second)


  problemloop:
   for i, record := range records{

     answer_ch := make(chan string)
     fmt.Printf("question %d : %s\n",i+1,record[0])
     go func(){
       fmt.Scanf("%s\n",&ans)
       answer_ch <- ans
     }()
     select {
     case <- timer.C:
              fmt.Println("Time up!!")
              //fmt.Printf("Thanking you for taking the quiz \n Your score is %d out of %d\n",correct_answer,total_questions)
              break problemloop

     case <- answer_ch:
              strings.Trim(ans," ")
              if(ans == record[1]){
                correct_answer = correct_answer + 1
          }
    }
}

fmt.Printf("Thanking you for taking the quiz \n Your score is %d out of %d",correct_answer,total_questions)
}
