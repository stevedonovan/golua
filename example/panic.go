package main

import lua "github.com/stevedonovan/golua/lua51"
import "fmt"

func test(L *lua.State) int {
	fmt.Println("hello world! from go!");
	return 0;
}

func main() {

	var L *lua.State;

	L = lua.NewState();
	defer L.Close()
	L.OpenLibs();
    
    // just to show we can catch the panic
    defer func() {
		fmt.Println("done") 
		if x := recover(); x != nil {
			fmt.Printf("run time panic: %v", x)
		}
	}()    
    
    // we can still catch the panic without this,
    // but it allows us to do some custom recovery
	currentPanicf := L.AtPanic(nil);
	currentPanicf = L.AtPanic(currentPanicf);
	newPanic := func(L1 *lua.State) int {
		fmt.Println("I AM PANICKING!!!");
		return currentPanicf(L1);
	}
	L.AtPanic(newPanic);    

	//force a panic
	L.PushNil();
	L.Call(0,0);
}
