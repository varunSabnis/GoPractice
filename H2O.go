package main

import (
   "fmt"
   "sync"
)

type Hydrogen struct {
  sync.Mutex
  hydrogencount int
  hyd_cond *sync.Cond
}

type Oxygen struct {
  sync.Mutex
  oxygencount int
  oxy_cond *sync.Cond
}

func NewOxygen() *Oxygen {
  o := Oxygen{}
  o.oxy_cond = sync.NewCond(&o)
  o.oxygencount = 0
  return(&o)

}

func NewHydrogen() *Hydrogen {
  h := Hydrogen{}
  h.hyd_cond = sync.NewCond(&h)
  h.hydrogencount = 0
  return(&h)

}

func main() {

  var wg sync.WaitGroup
  mx := sync.Mutex{}
  oxy := NewOxygen()
  hyd := NewHydrogen()
  wg.Add(18)
i := 0
for(i<18) {
  //wg.Add(3)
  go func(){
    defer wg.Done()

    mx.Lock()
    //hyd.Lock()
    hyd.hydrogencount = hyd.hydrogencount + 1
    fmt.Printf("Hydrogen %d created\n", hyd.hydrogencount)
    if(hyd.hydrogencount%2==1){
      mx.Unlock()
      //hyd.Unlock()

      hyd.Lock()
      hyd.hyd_cond.Wait()
      //print("1st Hydrogen Signalled\n")
      hyd.Unlock()
    }else{
      oxy.Lock()
      oxy.oxy_cond.Signal()
      //print("Oxygen can now be created as 2nd hydrogen created\n")
      oxy.Unlock()

      hyd.Lock()
      hyd.hyd_cond.Wait()
      //print("2nd Hydrogen signalled after oxyygen created\n")
      hyd.Unlock()

      mx.Unlock()
    }
    //print(" Hydrogen Goroutine returns\n")
    return
  }()

  go func(){

    defer wg.Done()
    //print("hi oxygen\n")
    oxy.Lock()
    oxy.oxy_cond.Wait()
    //print("hi oxygen\n")
    oxy.oxygencount = oxy.oxygencount + 1
    fmt.Printf("oxygen %d created\n",oxy.oxygencount)
    oxy.Unlock()

    hyd.Lock()
    hyd.hyd_cond.Signal()
    //print("First Hydrogen signalled\n")
    hyd.Unlock()

    hyd.Lock()
    hyd.hyd_cond.Signal()
    //print("Second Hydrogen Signalled\n")
    fmt.Println("H20 prepared")
    hyd.Unlock()
    //print(" Oxygen Goroutine returns\n")

    return
  }()
  i = i + 1
  //wg.Wait()
}
  wg.Wait()
}
