package main

import (
        "os"
        "os/exec"
        "fmt"
)

func simpleMessageStdout() {
        cmd := exec.Command("echo", "'What the heck is up'")

        cmd.Stdout = os.Stdout

        err := cmd.Start()
        if err != nil {
                panic(err)
        }
        cmd.Wait()
}

func commandMessage(cmd string, args []string) {
        //args := []string{"get", "namespaces"}
        out, err := exec.Command(cmd, args...).Output()
        //out, err := exec.Command("kubectl", "get", "namespaces").Output()
        if err != nil {
                panic(err)
        }
        fmt.Printf("%s", out)

}

func messageFile() {
        cmd := exec.Command("echo", "'What the heck is up'")

        //open the file for writing
        outfile, err := os.Create("./out.txt")
        if err != nil {
                panic(err)
        }
        defer outfile.Close()

        cmd.Stdout = outfile

        err = cmd.Start()
        if err != nil {
                panic(err)
        }
        cmd.Wait()
}

func main() {
        messageFile()
        simpleMessageStdout()
        args := []string{"get", "namespaces"}
        commandMessage("kubectl", args)
}
